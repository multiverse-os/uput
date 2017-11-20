## Developer Notes



TODO: Try to squash the number of validation needed in valid/string/is, have
more in valid/string by just returning !IsEmpty for NotEmpty(). This
will give the valid/string more readability while reducing the code
footprint of valid/string/is

TODO: Add HasXCharacters (for example, graphic characters, mark
characters, control characters, etc) that allows passing a number, so
one can check if a string has 2 punctuation symbols for example or two
symbol characters. 


var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// isAlphaNum reports whether the byte is an ASCII letter, number, or underscore
func isAlphaNum(c uint8) bool {
	return c == '_' || '0' <= c && c <= '9' || 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'
}


## Add Validations: 
strings.HasPrefix(substring)
strings.HasSuffix(substring)
path.Dir and path.Clean can be used to validate a path - can use stat to
check if it exists
os.ISExist and IsNotExist can be used to find if a file exists yet or
directory
Permission is checked with IsPermission
Executable is IsExecutatable under os
path.Filepath has more, has own Clean(path string) IsAbs will be useuful
too

Could validate groups, usernames and such using os stdlib

Can validate against current hostname using os.Hostname()

## Add Transforms:
strings.ToUpper(string, special bool)
strings.ToLower(string, special bool)
strings.ToTitle(string, special, bool)
strings.TrimSpace(string)
strings.TrimPrefix(substring)
strings.TrimSuffix(substring)
strings.Replace(string, old, new, n int) => present as Replace(Old, New)

for the special path type of string wil probably want:
os.MkDir can create
os.MkDirAll will create all necesssary parents of a path
os.Pipe(fileRead, fileWrite, err)
os.Rename(old, new)
os.TmpDir() uses whatever is set in env

==================
look at text

https://github.com/golang/go/tree/master/src/vendor/golang_org/x/text/secure


https://github.com/golang/go/blob/master/src/vendor/golang_org/x/net/http2/hpack/hpack.go
https://github.com/golang/go/blob/master/src/vendor/golang_org/x/net/idna/idna.go

// ToASCII is a wrapper for Punycode.ToASCII.
func ToASCII(s string) (string, error) {
	return Punycode.process(s, true)
}

// ToUnicode is a wrapper for Punycode.ToUnicode.
func ToUnicode(s string) (string, error) {
	return Punycode.process(s, false)
}\\\


// Transitional sets a Profile to use the Transitional mapping as defined in UTS
// #46. This will cause, for example, "ß" to be mapped to "ss". Using the
// transitional mapping provides a compromise between IDNA2003 and IDNA2008
// compatibility. It is used by most browsers when resolving domain names. This
// option is only meaningful if combined with MapForLookup.
func Transitional(transitional bool) Option {
	return func(o *options) { o.transitional = true }
}

// VerifyDNSLength sets whether a Profile should fail if any of the IDN parts
// are longer than allowed by the RFC.
func VerifyDNSLength(verify bool) Option {
	return func(o *options) { o.verifyDNSLength = verify }
}

// RemoveLeadingDots removes leading label separators. Leading runes that map to
// dots, such as U+3002 IDEOGRAPHIC FULL STOP, are removed as well.
//
// This is the behavior suggested by the UTS #46 and is adopted by some
// browsers.
func RemoveLeadingDots(remove bool) Option {
	return func(o *options) { o.removeLeadingDots = remove }
}

// ValidateLabels sets whether to check the mandatory label validation criteria
// as defined in Section 5.4 of RFC 5891. This includes testing for correct use
// of hyphens ('-'), normalization, validity of runes, and the context rules.
func ValidateLabels(enable bool) Option {
	return func(o *options) {
		// Don't override existing mappings, but set one that at least checks
		// normalization if it is not set.
		if o.mapping == nil && enable {
			o.mapping = normalize
		}
		o.trie = trie
		o.validateLabels = enable
		o.fromPuny = validateFromPunycode
	}
}

