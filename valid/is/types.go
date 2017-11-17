package is

// TODO: Trying to create a reasonable list of datatypes for most user input usecases

// Used by IsFilePath func
const (
	// Unknown is unresolved OS type
	Unknown = iota
	// Win is Windows type
	Windows
	// Unix is *nix OS types
	Posix
)

// Used by IsZero

var zeros = []interface{}{
	int(0),
	int8(0),
	int16(0),
	int32(0),
	int64(0),
	uint(0),
	uint8(0),
	uint16(0),
	uint32(0),
	uint64(0),
	float32(0),
	float64(0),
}

// Data Types
const (
	Slice    = "slice"
	Between  = "between"
	Boolean  = "boolean"
	Date     = "date"
	Email    = "email"
	Exists   = "exists"
	In       = "in"
	Int      = "int"
	IP       = "ip"
	Max      = "max"
	Min      = "min"
	NotIn    = "not_in"
	Numeric  = "numeric"
	Regex    = "regex"
	Required = "required"
	Size     = "size"
	String   = "string"
	Unique   = "unique"
	URL      = "url"
)

// Type Regex
const (
	numericRegex = `^(-|\+)?(([1-9]\d*(.\d*)?)|0(.\d*)?)$`
	phoneRegex   = `^((\+86)|(86))?(13\d|15[^4\D]|17[13678]|18\d)\d{8}|170[^346\D]\d{7}$`
	emailRegex   = `^([\w-_]+(?:\.[\w-_]+)*)@((?:[a-z0-9]+(?:-[a-zA-Z0-9]+)*)+\.[a-z]{2,6})$`
	ipv4Regex    = `^(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)\.(25[0-5]|2[0-4]\d|[0-1]\d{2}|[1-9]?\d)$`
	// TODO: Add Bitcoin address
)

const (
	emailRegex3 = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
)

const (
	ipv4Regex2 = regexp.MustCompile(`/^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}$/`)
)

const (
	//regex for email structure;
	emailRegex4 = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	// FROM THIS PARSER: https://github.com/nickbruun/goinput/blob/master/email.go
	// E-mail address user part pattern.
	emailUserPartPattern = regexp.MustCompile(`^[-!#$%&'*+/=?^_` + "`" + `{}|~0-9A-Za-z]+(\.[-!#$%&'*+/=?^_` + "`" + `{}|~0-9A-Za-z]+)*$`)
	// E-mail address domain part pattern.
	emailDomainPartPattern = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+([a-zA-Z0-9\-]{2,63})$`)
)

const (
	ipRegex1     = regexp.MustCompile(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	domainRegex1 = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`)
)

const (
	uuidRegex2 = regexp.MustCompile("^[a-z0-9]{8}-[a-z0-9]{4}-[1-5][a-z0-9]{3}-[a-z0-9]{4}-[a-z0-9]{12}$")
)

const (
	// Alpha represents regular expression for alpha chartacters
	Alpha string = "^[a-zA-Z]+$"
	// AlphaDash represents regular expression for alpha chartacters with underscore and ash
	AlphaDash string = "^[a-zA-Z0-9_-]+$"
	// AlphaNumeric represents regular expression for alpha numeric chartacters
	AlphaNumeric string = "^[a-zA-Z0-9]+$"
	// CreditCard represents regular expression for credit cards like (Visa, MasterCard, American Express, Diners Club, Discover, and JCB cards). Ref: https://stackoverflow.com/questions/9315647/regex-credit-card-number-tests
	CreditCard string = "^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
	// Coordinate represents latitude and longitude regular expression
	Coordinate string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?),\\s*[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$" // Ref: https://stackoverflow.com/questions/3518504/regular-expression-for-matching-latitude-longitude-coordinates
	// CSSColor represents css valid color code with hex, rgb, rgba, hsl, hsla etc. Ref: http://www.regexpal.com/97509
	CSSColor string = "^(#([\\da-f]{3}){1,2}|(rgb|hsl)a\\((\\d{1,3}%?,\\s?){3}(1|0?\\.\\d+)\\)|(rgb|hsl)\\(\\d{1,3}%?(,\\s?\\d{1,3}%?){2}\\))$"
	// Date represents regular expression for valid date like: yyyy-mm-dd
	Date string = "^(((19|20)([2468][048]|[13579][26]|0[48])|2000)[/-]02[/-]29|((19|20)[0-9]{2}[/-](0[4678]|1[02])[/-](0[1-9]|[12][0-9]|30)|(19|20)[0-9]{2}[/-](0[1359]|11)[/-](0[1-9]|[12][0-9]|3[01])|(19|20)[0-9]{2}[/-]02[/-](0[1-9]|1[0-9]|2[0-8])))$"
	// DateDDMMYY represents regular expression for valid date of format dd/mm/yyyy , dd-mm-yyyy etc.Ref: http://regexr.com/346hf
	DateDDMMYY string = "^(0?[1-9]|[12][0-9]|3[01])[\\/\\-](0?[1-9]|1[012])[\\/\\-]\\d{4}$"
	// Email represents regular expression for email
	Email string = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$"
	// Float represents regular expression for finding fload number
	Float string = "^[+-]?([0-9]*[.])?[0-9]+$"
	// IP represents regular expression for ip address
	IP string = "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	// IPV4 represents regular expression for ip address version 4
	IPV4 string = "^([0-9]{1,3}\\.){3}[0-9]{1,3}(\\/([0-9]|[1-2][0-9]|3[0-2]))?$"
	// IPV6 represents regular expression for ip address version 6
	IPV6 string = `^s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:)))(%.+)?s*(\/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8]))?$`
	// Latitude represents latitude regular expression
	Latitude string = "^(\\+|-)?(?:90(?:(?:\\.0{1,6})?)|(?:[0-9]|[1-8][0-9])(?:(?:\\.[0-9]{1,6})?))$"
	// Longitude represents longitude regular expression
	Longitude string = "^(\\+|-)?(?:180(?:(?:\\.0{1,6})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\\.[0-9]{1,6})?))$"
	// Numeric represents regular expression for numeric
	Numeric string = "^[0-9]+$"
	// URL represents regular expression for url
	URL string = "^(?:http(s)?:\\/\\/)?[\\w.-]+(?:\\.[\\w\\.-]+)+[\\w\\-\\._~:/?#[\\]@!\\$&'\\(\\)\\*\\+,;=.]+$" // Ref: https://stackoverflow.com/questions/136505/searching-for-uuids-in-text-with-regex
	// UUID represents regular expression for UUID
	UUID string = "^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}$"
	// UUID3 represents regular expression for UUID version 3
	UUID3 string = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	// UUID4 represents regular expression for UUID version 4
	UUID4 string = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	// UUID5 represents regular expression for UUID version 5
	UUID5 string = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
)

// Basic regular expressions for validating strings
const (
	emailRegex2    = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	ISBN10         = `^(?:[0-9]{9}X|[0-9]{10})$`
	CreditCard     = `^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$`
	ISBN13         = `^(?:[0-9]{13})$`
	UUID3          = `^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$`
	UUID4          = `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	UUID5          = `^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`
	UUID           = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`
	Alpha          = `^[a-zA-Z]+$`
	Alphanumeric   = `^[a-zA-Z0-9]+$`
	Numeric        = `^[0-9]+$`
	Int            = `^(?:[-+]?(?:0|[1-9][0-9]*))$`
	Float          = `^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$`
	Hexadecimal    = `^[0-9a-fA-F]+$`
	Hexcolor       = `^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`
	RGBcolor       = `^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$`
	ASCII          = `^[\x00-\x7F]+$`
	Multibyte      = `[^\x00-\x7F]`
	FullWidth      = `[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]`
	HalfWidth      = `[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]`
	Base64         = `^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`
	PrintableASCII = `^[\x20-\x7E]+$`
	DataURI        = `^data:.+\\/(.+);base64$`
	Latitude       = `^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$`
	Longitude      = `^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$`
	DNSName        = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9_-]{1,62})*$`
	IP             = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	URLSchema      = `((ftp|tcp|udp|wss?|https?):\/\/)`
	URLUsername    = `(\S+(:\S*)?@)`
	Hostname       = ``
	URLPath        = `((\/|\?|#)[^\s]*)`
	URLPort        = `(:(\d{1,5}))`
	URLIP          = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	URLSubdomain   = `((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))`
	URL            = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
	SSN            = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	WinPath        = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	UnixPath       = `^(/[^/\x00]*)+/?$`
	Semver         = `^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$`
	tagName        = `valid`
)

func Empty(value interface{}) bool {
	if value == nil {
		return true
	}

	switch v := value.(type) {
	case bool:
		return !v
	case string:
		return v == ""
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		for _, zero := range zeros {
			if v == zero {
				return true
			}
		}
	case time.Time:
		return v.IsZero()
	}

	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			if !Empty(val.FieldByIndex([]int{i}).Interface()) {
				return false
			}
		}
		return true
	case reflect.Map, reflect.Slice, reflect.Chan:
		return (val.Len() == 0)
	case reflect.Ptr:
		return Empty(val.Elem().Interface())
	}
	return false
}

func NotEmpty(value interface{}) bool {
	return !Empty(value)
}

func Range(value, min, max int) bool {
	return min <= value && value <= max
}

func StringSize(value string, min, max int) bool {
	l := len([]rune(value))
	return min <= l && l <= max
}

func Regexp(value string, pattern interface{}) bool {
	var r *regexp.Regexp
	if str, ok := pattern.(string); ok {
		r = regexp.MustCompile(str)
	} else if rpattern, ok := pattern.(*regexp.Regexp); ok {
		r = rpattern
	}
	return r.MatchString(value)
}

func Equal(value, expected interface{}) bool {
	if value == nil && expected == nil {
		return true
	}
	if reflect.DeepEqual(value, expected) {
		return true
	}
	return false
}

func Contain(value, expected interface{}) bool {
	v := reflect.ValueOf(expected)
	if v.Kind() == reflect.Slice {
		for i := 0; i < v.Len(); i++ {
			if Equal(value, v.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

func TimeRange(value, from, to time.Time) bool {
	return from.UnixNano() <= value.UnixNano() && value.UnixNano() <= to.UnixNano()
}

var Messages = struct {
	NotEmpty, Range, StringSize, Regexp, Equal, Contain, TimeRange string
}{
	NotEmpty:   "can't be blank",
	Range:      "must be between %v and %v",
	StringSize: "string length must be between %v and %v",
	Regexp:     "must match with pattern\"%v\"",
	Equal:      "must be %v",
	Contain:    "must be one of following values. %v",
	TimeRange:  "must be between %v and %v",
}

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) HasErrors() bool {
	return len(v.Errors) > 0
}

func (v *Validator) AddError(key string, msg ...string) {
	v.Errors[key] = strings.Join(msg, " ")
}

func (v *Validator) SetError(result bool, key string, msg ...string) {
	if !result {
		v.AddError(key, msg...)
	}
}

func (v *Validator) NotEmpty(value interface{}, key string, msg ...string) bool {
	var m string
	if len(msg) == 0 {
		m = fmt.Sprintf(Messages.NotEmpty)
	} else {
		m = fmt.Sprintf(msg[0])
	}
	result := NotEmpty(value)
	v.SetError(result, key, m)
	return result
}

func (v *Validator) Range(value, min, max int, key string, msg ...string) bool {
	var m string
	if len(msg) == 0 {
		m = fmt.Sprintf(Messages.Range, min, max)
	} else {
		m = fmt.Sprintf(msg[0], min, max)
	}
	result := Range(value, min, max)
	v.SetError(result, key, m)
	return result
}

func (v *Validator) StringSize(value string, min, max int, key string, msg ...string) bool {
	var m string
	if len(msg) == 0 {
		m = fmt.Sprintf(Messages.StringSize, min, max)
	} else {
		m = fmt.Sprintf(msg[0], min, max)
	}
	result := StringSize(value, min, max)
	v.SetError(result, key, m)
	return result
}

func (v *Validator) Regexp(value string, pattern interface{}, key string, msg ...string) bool {
	var m string
	if len(msg) == 0 {
		m = fmt.Sprintf(Messages.Regexp, pattern)
	} else {
		m = fmt.Sprintf(msg[0], pattern)
	}
	result := Regexp(value, pattern)
	v.SetError(result, key, m)
	return result
}

func (v *Validator) Equal(value, expected interface{}, key string, msg ...string) bool {
	var m string
	if len(msg) == 0 {
		m = fmt.Sprintf(Messages.Equal, expected)
	} else {
		m = fmt.Sprintf(msg[0], expected)
	}
	result := Equal(value, expected)
	v.SetError(result, key, m)
	return result
}

func (v *Validator) Contain(value, expected interface{}, key string, msg ...string) bool {
	var m string
	if len(msg) == 0 {
		m = fmt.Sprintf(Messages.Contain, expected)
	} else {
		m = fmt.Sprintf(msg[0], expected)
	}
	result := Contain(value, expected)
	v.SetError(result, key, m)
	return result
}

func (v *Validator) TimeRange(value, from, to time.Time, key string, msg ...string) bool {
	var m string
	if len(msg) == 0 {
		m = fmt.Sprintf(Messages.Contain, from, to)
	} else {
		m = fmt.Sprintf(msg[0], from, to)
	}
	result := TimeRange(value, from, to)
	v.SetError(result, key, m)
	return result
}
