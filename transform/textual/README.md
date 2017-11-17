textual
=======

Package textual provides some simple functions for manipulating text.

FUNCTIONS

```go 
func ToCamel(text string, private ...bool) string
```
Convert string from database column names to corresponding struct field names (e.g. field_name to FieldName)

```go 
func ToPlural(text string) (plural string)
```

Provide the plural version of an English word using some simple rules and a table of exceptions.

```go 
func ToSnake(text string) string
```
Convert string from struct field names to corresponding database column names (e.g. FieldName to field_name)

```go 
func Truncate(s string, length int) string
``` 

Truncate the given string to length using … as ellipsis.

```go 
func TruncateWithEllipsis(s string, length int, ellipsis string) string
``` 

Truncate the given string to length using provided ellipsis.

