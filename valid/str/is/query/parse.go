package qs

import (
	"fmt"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
	"strconv"
	"strings"
	"time"
)

// ParseError is the error type returned by Parse()
type ParseError struct {
	// Pos is the character position where the error occured
	Pos int
	// Msg is a description of the error
	Msg string
}

func (pe ParseError) Error() string { return fmt.Sprintf("%d: %s", pe.Pos, pe.Msg) }

type OpType int

const (
	OR  OpType = 0
	AND        = 1
)

type Parser struct {
	tokens []token
	pos    int
	// DefaultOp is used when no explict OR or AND is present
	// ie: foo bar => foo OR bar | foo AND bar
	// TODO: not sure AND/OR is the right terminology (but it's what others use)
	// Doesn't actually use OR or AND. AND is treated as an implied '+' prefix
	DefaultOp OpType

	// Loc is the location to use for parsing dates in range queries.
	// If nil, UTC is assumed.
	Loc *time.Location
}

// context is used to hold settings active within a given scope during parsing
// (TODO: maybe this should just be absorbed back into the Parser struct instead?)
type context struct {
	// field is the name of the field currently in scope (or "")
	field    string
	fieldPos int
}

// Parse takes a query string and turns it into a bleve Query.
//
// Returned errors are type ParseError, which includes the position
// of the offending part of the input string.
//
// BNF(ish) query syntax:
//   exprList = expr1*
//   expr1 = expr2 {"OR" expr2}
//   expr2 = expr3 {"AND" expr3}
//   expr3 = {"NOT"} expr4
//   expr4 = {("+"|"-")} expr5
//   expr5 = {field} part {boost}
//   part = lit {"~" number} | range | "(" exprList ")"
//   field = lit ":"
//   range = ("["|"}") {lit} "TO" {lit} ("]"|"}")
//   relational = ("<"|">"|"<="|">=") lit
//   boost = "^" number
//
// (where lit is a string, quoted string or number)
func (p *Parser) Parse(q string) (query.Query, error) {
	p.tokens = lex(q)
	ctx := context{field: ""}
	return p.parseExprList(ctx)
}

// Parse takes a query string and turns it into a bleve Query using
// the default Parser.
// Returned errors are type ParseError, which includes the position
// of the offending part of the input string.
func Parse(q string) (query.Query, error) {
	p := Parser{DefaultOp: OR}
	return p.Parse(q)
}

// peek looks at the next token without consuming it.
// peeks beyond the end of the token stream will return EOF
func (p *Parser) peek() token {
	if p.pos < len(p.tokens) {
		tok := p.tokens[p.pos]
		return tok
	}
	return token{typ: tEOF}
}

// backup steps back one position in the token stream
func (p *Parser) backup() {
	p.pos -= 1
}

// next fetches the next token in the stream
func (p *Parser) next() token {
	if p.pos < len(p.tokens) {
		tok := p.tokens[p.pos]
		p.pos += 1
		return tok
	}
	p.pos += 1 // to make sure backup() works
	return token{typ: tEOF}
}

// starting point
//   exprList = expr1*
func (p *Parser) parseExprList(ctx context) (query.Query, error) {
	// <empty>
	if p.peek().typ == tEOF {
		return bleve.NewMatchNoneQuery(), nil
	}

	must := []query.Query{}
	mustNot := []query.Query{}
	should := []query.Query{}

	for {
		tok := p.peek()
		if tok.typ == tEOF {
			break
		}
		// slightly kludgy...
		if tok.typ == tRPAREN {
			break
		}

		prefix, q, err := p.parseExpr1(ctx)
		if err != nil {
			return nil, err
		}

		switch prefix {
		case tPLUS:
			must = append(must, q)
		case tMINUS:
			mustNot = append(mustNot, q)
		default:
			if p.DefaultOp == AND {
				must = append(must, q)
			} else { // OR
				should = append(should, q)
			}
		}
	}

	// some obvious shortcuts
	total := len(must) + len(mustNot) + len(should)
	if total == 0 {
		return bleve.NewMatchNoneQuery(), nil
	}
	if total == 1 && len(must) == 1 {
		return must[0], nil
	}
	if total == 1 && len(should) == 1 {
		return should[0], nil
	}

	// no shortcuts - go with the full-fat version
	q := bleve.NewBooleanQuery()
	if len(must) > 0 {
		q.AddMust(must...)
	}
	if len(should) > 0 {
		q.AddShould(should...)
	}
	if len(mustNot) > 0 {
		q.AddMustNot(mustNot...)
	}
	return q, nil
}

// parseExpr1 handles OR expressions
//
//   expr1 = expr2 {"OR" expr2}
func (p *Parser) parseExpr1(ctx context) (tokType, query.Query, error) {

	queries := []query.Query{}
	prefixes := []tokType{}

	for {
		prefix, q, err := p.parseExpr2(ctx)
		if err != nil {
			return tEOF, nil, err
		}

		prefixes = append(prefixes, prefix)
		queries = append(queries, q)

		tok := p.next()
		if tok.typ != tOR {
			p.backup()
			break
		}
	}

	// let single, non-OR expressions bubble upward, prefix intact
	if len(queries) == 1 {
		return prefixes[0], queries[0], nil
	}

	// KLUDGINESS - prefixes on terms in OR expressions
	// we'll ignore "+" and treat "-" as NOT
	// eg:
	// `+alice OR -bob OR chuck`  => `alice OR (NOT bob) OR chuck`
	for i, _ := range queries {
		if prefixes[i] == tMINUS {
			q := bleve.NewBooleanQuery()
			q.AddMustNot(queries[i])
			queries[i] = q
		}
	}

	return tEOF, bleve.NewDisjunctionQuery(queries...), nil
}

// parseExpr2 handles AND expressions
//
//   expr2 = expr3 {"AND" expr3}
func (p *Parser) parseExpr2(ctx context) (tokType, query.Query, error) {

	queries := []query.Query{}
	prefixes := []tokType{}

	for {
		prefix, q, err := p.parseExpr3(ctx)
		if err != nil {
			return tEOF, nil, err
		}

		prefixes = append(prefixes, prefix)
		queries = append(queries, q)

		tok := p.next()
		if tok.typ != tAND {
			p.backup()
			break
		}
	}

	// let single, non-AND expressions bubble upward, prefix intact
	if len(queries) == 1 {
		return prefixes[0], queries[0], nil
	}

	// KLUDGINESS - prefixes on terms in AND expressions
	// we'll ignore "+" and treat "-" as NOT
	// eg:
	// `+alice AND -bob AND chuck`  => `alice AND (NOT bob) AND chuck`
	for i, _ := range queries {
		if prefixes[i] == tMINUS {
			q := bleve.NewBooleanQuery()
			q.AddMustNot(queries[i])
			queries[i] = q
		}
	}

	return tEOF, bleve.NewConjunctionQuery(queries...), nil
}

// parseExpr3 handles NOT expressions
//
//   expr3 = {"NOT"} expr4
func (p *Parser) parseExpr3(ctx context) (tokType, query.Query, error) {

	tok := p.next()
	if tok.typ != tNOT {
		p.backup()
		// just let the lower, non-NOT expression bubble up with its prefix
		return p.parseExpr4(ctx)
	}

	prefix, q, err := p.parseExpr4(ctx)
	if err != nil {
		return tEOF, nil, err
	}

	// KLUDGINESS - prefixes on terms in NOT expressions:
	// `NOT -bob`  => `bob`
	// `NOT +bob`  => `NOT bob`
	if prefix != tMINUS {
		q2 := bleve.NewBooleanQuery()
		q2.AddMustNot(q)
		q = q2
	}
	return tEOF, q, nil
}

// Here's where all the prefix-bubbling-up begins...
//   expr4 = {("+"|"-")} expr5
func (p *Parser) parseExpr4(ctx context) (tokType, query.Query, error) {
	var prefix tokType
	tok := p.next()
	switch tok.typ {
	case tMINUS, tPLUS:
		prefix = tok.typ
	default:
		p.backup()
		prefix = tEOF
	}

	q, err := p.parseExpr5(ctx)
	return prefix, q, err

}

//   expr5 = {field} part {boost}
func (p *Parser) parseExpr5(ctx context) (query.Query, error) {

	fldpos := p.peek().pos
	fld, err := p.parseField()
	if err != nil {
		return nil, err
	}
	if fld != "" {
		if ctx.field != "" {
			return nil, ParseError{fldpos, fmt.Sprintf("'%s:' clashes with '%s:'", fld, ctx.field)}
		}
		ctx.field = fld
		ctx.fieldPos = fldpos
	}

	q, err := p.parsePart(ctx)
	if err != nil {
		return nil, err
	}

	// parse (optional) suffix
	boostpos := p.peek().pos
	boost, err := p.parseBoostSuffix()
	if err != nil {
		return nil, err
	}
	if boost > 0 {
		if boostable, ok := q.(query.BoostableQuery); ok {
			boostable.SetBoost(boost)
		} else {
			return nil, ParseError{boostpos, "can't specify a boost value here"}
		}
	}

	return q, nil
}

//   part = lit {"~" number} | range | "(" exprList ")"
func (p *Parser) parsePart(ctx context) (query.Query, error) {

	tok := p.next()

	//   lit
	if tok.typ == tLITERAL {
		var q query.Query
		if strings.ContainsAny(tok.val, "*?") {
			q = bleve.NewWildcardQuery(tok.val)
		} else {
			if p.peek().typ == tFUZZY {
				fuzziness, err := p.parseFuzzySuffix()
				if err != nil {
					return nil, err
				}
				fuzz := bleve.NewFuzzyQuery(tok.val)
				fuzz.SetFuzziness(fuzziness)
				q = fuzz
			} else {
				q = bleve.NewMatchPhraseQuery(tok.val)
			}
		}
		if ctx.field != "" {
			if fieldable, ok := q.(query.FieldableQuery); ok {
				fieldable.SetField(ctx.field)
			} else {
				return nil, ParseError{ctx.fieldPos, "unexpected field"}
			}
		}
		return q, nil
	}
	if tok.typ == tQUOTED {
		// strip quotes (ugh)
		txt := string(tok.val[1 : len(tok.val)-1])
		/*
			if strings.ContainsAny(txt, "*?") {
				return nil, ParseError{tok.pos, "wildcards not supported in phrases"}
			}
		*/
		q := bleve.NewMatchPhraseQuery(txt)
		if ctx.field != "" {
			q.SetField(ctx.field)
		}
		return q, nil
	}

	//   | "(" exprList ")"
	if tok.typ == tLPAREN {
		q, err := p.parseExprList(ctx)
		if err != nil {
			return nil, err
		}
		tok = p.next()
		if tok.typ != tRPAREN {
			return nil, ParseError{tok.pos, "missing )"}
		}
		return q, nil
	}

	//   | range
	if tok.typ == tLSQUARE || tok.typ == tLBRACE {
		p.backup()
		q, err := p.parseRange(ctx)
		if err != nil {
			return nil, err
		}
		return q, nil
	}

	//   | relational
	if tok.typ == tGREATER || tok.typ == tLESS {
		p.backup()
		q, err := p.parseRelational(ctx)
		if err != nil {
			return nil, err
		}
		return q, nil
	}

	if tok.typ == tERROR {
		return nil, ParseError{tok.pos, tok.val}
	}

	return nil, ParseError{tok.pos, fmt.Sprintf("unexpected %s", tok.val)}
}

// returns >0 if there is a value given
func (p *Parser) parseBoostSuffix() (float64, error) {
	tok := p.next()
	if tok.typ != tBOOST {
		p.backup()
		return 0, nil
	}

	v := tok.val[1:]
	if v == "" {
		return 1.0, nil
	}
	boost, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0, ParseError{tok.pos, "bad boost value"}
	}

	return boost, nil
}

//
func (p *Parser) parseFuzzySuffix() (int, error) {
	tok := p.next()
	if tok.typ != tFUZZY {
		return 0, ParseError{tok.pos, "expected ~"}
	}

	v := tok.val[1:]
	if v == "" {
		return 1, nil // Default fuzziness is 1
	}
	fuzz, err := strconv.Atoi(v)
	if err != nil {
		return 0, ParseError{tok.pos, "bad fuzziness value"}
	}

	return fuzz, nil
}

// parse (optional) field specifier
// [ lit ":" ]
// returns field name or "" if not a field
func (p *Parser) parseField() (string, error) {
	tok := p.next()
	if tok.typ != tLITERAL {
		// not a field
		p.backup()
		return "", nil
	}
	field := tok.val

	tok = p.next()
	if tok.typ != tCOLON {
		// oop, not a field after all!
		p.backup()
		p.backup()
		return "", nil
	}

	return field, nil // it's OK
}

//   range = ("["|"}") {lit} "TO" {lit} ("]"|"}")
func (p *Parser) parseRange(ctx context) (query.Query, error) {

	var minVal, maxVal string
	var minInclusive, maxInclusive bool
	openTok := p.next()
	switch openTok.typ {
	case tLSQUARE:
		minInclusive = true
	case tLBRACE:
		minInclusive = false
	default:
		return nil, ParseError{openTok.pos, "expected range"}
	}

	tok := p.next()
	switch tok.typ {
	case tLITERAL:
		minVal = tok.val
	case tQUOTED:
		minVal = tok.val[1 : len(tok.val)-1]
	case tTO:
		p.backup()
		// empty start
	default:
		return nil, ParseError{tok.pos, fmt.Sprintf("unexpected %s", tok.val)}
	}

	tok = p.next()
	if tok.typ != tTO {
		return nil, ParseError{tok.pos, "expected TO"}
	}

	tok = p.next()
	switch tok.typ {
	case tLITERAL:
		maxVal = tok.val
	case tQUOTED:
		maxVal = tok.val[1 : len(tok.val)-1]
	case tRSQUARE:
		p.backup() // empty end value
	case tRBRACE:
		p.backup() // empty end value
	default:
		return nil, ParseError{tok.pos, fmt.Sprintf("unexpected %s", tok.val)}
	}

	closeTok := p.next()
	switch closeTok.typ {
	case tRSQUARE:
		maxInclusive = true
	case tRBRACE:
		maxInclusive = false
	default:
		return nil, ParseError{closeTok.pos, "expected ] or }"}
	}

	rp := newRangeParams(minVal, maxVal, minInclusive, maxInclusive, p.Loc)
	q, err := rp.generate()
	if err != nil {
		return nil, ParseError{openTok.pos, err.Error()}
	}
	if ctx.field != "" {
		if fieldable, ok := q.(query.FieldableQuery); ok {
			fieldable.SetField(ctx.field)
		} else {
			return nil, ParseError{ctx.fieldPos, "unexpected field"}
		}
	}
	return q, nil
}

// parseRelational handles greaterthan/lessthan etc...
// Implemented as a range.
//   relational = ("<"|">"|"<="|">=") lit
func (p *Parser) parseRelational(ctx context) (query.Query, error) {

	var minVal, maxVal string
	var minInclusive, maxInclusive bool

	rel := p.next()
	if rel.typ != tGREATER && rel.typ != tLESS {
		return nil, ParseError{rel.pos, "expected > or <"}
	}

	eq := p.next()
	if eq.typ != tEQUAL {
		p.backup()
	}

	var val string
	tok := p.next()
	switch tok.typ {
	case tLITERAL:
		val = tok.val
	case tQUOTED:
		val = tok.val[1 : len(tok.val)-1]
	default:
		return nil, ParseError{tok.pos, fmt.Sprintf("unexpected %s", tok.val)}
	}

	if rel.typ == tGREATER {
		minVal = val
		minInclusive = (eq.typ == tEQUAL)
	} else { // if rel.typ == tLESS
		maxVal = val
		maxInclusive = (eq.typ == tEQUAL)
	}

	rp := newRangeParams(minVal, maxVal, minInclusive, maxInclusive, p.Loc)
	q, err := rp.generate()
	if err != nil {
		return nil, ParseError{rel.pos, err.Error()}
	}
	if ctx.field != "" {
		if fieldable, ok := q.(query.FieldableQuery); ok {
			fieldable.SetField(ctx.field)
		} else {
			return nil, ParseError{ctx.fieldPos, "unexpected field"}
		}
	}
	return q, nil
}
