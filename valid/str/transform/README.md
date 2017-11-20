## String Transformations for Pre & Post Validation
Enable a developer to include these type of transofmraitons, to include basic transforms, that support chaining by taking in StringInput data type and returning StringInput data type. 

ToUpper and ToLower are provided by strings stdlib

https://github.com/golang/go/tree/f1966de63f1174f93cdca87b36e8c02b7b95f652/src/strings

Any that can be supplied by stdlibs should be, or at least the code
should be extracted from stdlibs.


// Contains reports whether substr is within s.
Contains()
// ContainsAny reports whether any Unicode code points in chars are within s.
ContainsAny()
// ContainsRune reports whether the Unicode code point r is within s.
ContainsRune()
