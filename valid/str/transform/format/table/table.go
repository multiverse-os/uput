package table

import (
	"fmt"
	"reflect"
)

type bd struct {
	H  rune // BOX DRAWINGS HORIZONTAL
	V  rune // BOX DRAWINGS VERTICAL
	VH rune // BOX DRAWINGS VERTICAL AND HORIZONTAL
	HU rune // BOX DRAWINGS HORIZONTAL AND UP
	HD rune // BOX DRAWINGS HORIZONTAL AND DOWN
	VL rune // BOX DRAWINGS VERTICAL AND LEFT
	VR rune // BOX DRAWINGS VERTICAL AND RIGHT
	DL rune // BOX DRAWINGS DOWN AND LEFT
	DR rune // BOX DRAWINGS DOWN AND RIGHT
	UL rune // BOX DRAWINGS UP AND LEFT
	UR rune // BOX DRAWINGS UP AND RIGHT
}

var m = map[string]bd{
	"ascii":       bd{'-', '|', '+', '+', '+', '+', '+', '+', '+', '+', '+'},
	"box-drawing": bd{'─', '│', '┼', '┴', '┬', '┤', '├', '┐', '┌', '┘', '└'},
}

// Output formats slice of structs data and writes to standard output.(Using box drawing characters)
func Output(slice interface{}) {
	fmt.Println(Table(slice))
}

// OutputA formats slice of structs data and writes to standard output.(Using standard ascii characters)
func OutputA(slice interface{}) {
	fmt.Println(AsciiTable(slice))
}

// Table formats slice of structs data and returns the resulting string.(Using box drawing characters)
func Table(slice interface{}) string {
	coln, colw, rows := parse(slice)
	table := table(coln, colw, rows, m["box-drawing"])
	return table
}

// Table formats slice of structs data and returns the resulting string.(Using standard ascii characters)
func AsciiTable(slice interface{}) string {
	coln, colw, rows := parse(slice)
	table := table(coln, colw, rows, m["ascii"])
	return table
}

func parse(slice interface{}) (
	coln []string, // name of columns
	colw []int, // width of columns
	rows [][]string, // rows of content
) {
	for i, u := range sliceconv(slice) {
		v := reflect.ValueOf(u)
		t := reflect.TypeOf(u)
		if v.Kind() != reflect.Struct {
			panic("Table: items of slice should be on struct value")
		}
		var row []string

		m := 0 // count of unexported field
		for n := 0; n < v.NumField(); n++ {
			if t.Field(n).PkgPath != "" {
				m++
				continue
			}
			cn := t.Field(n).Name
			cv := fmt.Sprintf("%+v", v.FieldByName(cn).Interface())

			if i == 0 {
				coln = append(coln, cn)
				colw = append(colw, len(cn))
			}
			if colw[n-m] < len(cv) {
				colw[n-m] = len(cv)
			}

			row = append(row, cv)
		}
		rows = append(rows, row)
	}
	return coln, colw, rows
}

func table(coln []string, colw []int, rows [][]string, b bd) (table string) {
	head := [][]rune{[]rune{b.DR}, []rune{b.V}, []rune{b.VR}}
	bttm := []rune{b.UR}
	for i, v := range colw {
		head[0] = append(head[0], []rune(repeat(v+2, b.H)+string(b.HD))...)
		head[1] = append(head[1], []rune(" "+coln[i]+repeat(v-len(coln[i])+1, ' ')+string(b.V))...)
		head[2] = append(head[2], []rune(repeat(v+2, b.H)+string(b.VH))...)
		bttm = append(bttm, []rune(repeat(v+2, b.H)+string(b.HU))...)
	}
	head[0][len(head[0])-1] = b.DL
	head[2][len(head[2])-1] = b.VL
	bttm[len(bttm)-1] = b.UL

	var body [][]rune
	for _, r := range rows {
		row := []rune{b.V}
		for i, v := range colw {
			// handle non-ascii character
			lb := len(r[i])
			lr := len([]rune(r[i]))

			row = append(row, []rune(" "+r[i]+repeat(v-lb+(lb-lr)/2+1, ' ')+string(b.V))...)
		}
		body = append(body, row)
	}

	for _, v := range head {
		table += string(v) + "\n"
	}
	for _, v := range body {
		table += string(v) + "\n"
	}
	table += string(bttm)
	return table
}

func sliceconv(slice interface{}) []interface{} {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		panic("sliceconv: param \"slice\" should be on slice value")
	}

	l := v.Len()
	r := make([]interface{}, l)
	for i := 0; i < l; i++ {
		r[i] = v.Index(i).Interface()
	}
	return r
}

func repeat(time int, char rune) string {
	var s = make([]rune, time)
	for i := range s {
		s[i] = char
	}
	return string(s)
}
