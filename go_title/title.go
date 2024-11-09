package go_title

import (
	"regexp"
	"strings"
)

var (
	floatReg      = regexp.MustCompile(`:(\s*\d*)\.0`)
	numReg        = regexp.MustCompile(`^\d+$`)
	formatReg     = regexp.MustCompile(`[^A-Za-z0-9]`)
	timeReg       = regexp.MustCompile(`\d{4}-\d\d-\d\dT\d\d:\d\d:\d\d(\.\d+)?(\+\d\d:\d\d|Z)`)
	reLiteralUUID = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)
	caseReg1      = regexp.MustCompile(`(^|[^a-zA-Z])([a-z]+)`)
	caseReg2      = regexp.MustCompile(`([A-Z])([a-z]+)`)
	uuidReg       = regexp.MustCompile(`[xy]`)
)

func TitleCase(str string) string {
	if numReg.MatchString(str) {
		str = "Num" + str
	} else if str[0] <= '9' && str[0] >= '0' {
		numbers := map[byte]string{
			'0': "Zero_",
			'1': "One_",
			'2': "Two_",
			'3': "Three_",
			'4': "Four_",
			'5': "Five_",
			'6': "Six_",
			'7': "Seven_",
			'8': "Eight_",
			'9': "Nine_",
		}

		str = numbers[str[0]] + str[1:]
	}

	return formatReg.ReplaceAllString(toProperCase(str), "")
}

// Proper cases a string according to Go conventions
func toProperCase(str string) string {
	// https://github.com/golang/lint/blob/5614ed5bae6fb75893070bdc0996a68765fdd275/lint.gogogo#L771-L810
	commonInitialisms := map[string]struct{}{
		"ACL":   {},
		"API":   {},
		"ASCII": {},
		"CPU":   {},
		"CSS":   {},
		"DNS":   {},
		"EOF":   {},
		"GUID":  {},
		"HTML":  {},
		"HTTP":  {},
		"HTTPS": {},
		"ID":    {},
		"IP":    {},
		"JSON":  {},
		"LHS":   {},
		"QPS":   {},
		"RAM":   {},
		"RHS":   {},
		"RPC":   {},
		"SLA":   {},
		"SMTP":  {},
		"SQL":   {},
		"SSH":   {},
		"TCP":   {},
		"TLS":   {},
		"TTL":   {},
		"UDP":   {},
		"UI":    {},
		"UID":   {},
		"UUID":  {},
		"URI":   {},
		"URL":   {},
		"UTF8":  {},
		"VM":    {},
		"XML":   {},
		"XMPP":  {},
		"XSRF":  {},
		"XSS":   {},
	}

	str = replaceAllStringSubmatchFunc(caseReg1, str, func(groups []string) string {
		sep := groups[1]
		frag := groups[2]

		upFrag := strings.ToUpper(frag)
		_, ok := commonInitialisms[upFrag]
		if ok {
			return sep + upFrag
		} else {
			return sep + string(upFrag[0]) + strings.ToLower(frag[1:])
		}
	})

	str = replaceAllStringSubmatchFunc(caseReg2, str, func(groups []string) string {
		sep := groups[1]
		frag := groups[2]

		upFrag := strings.ToUpper(frag)

		_, ok := commonInitialisms[sep+upFrag]
		if ok {
			return strings.ToUpper(sep + frag)
		} else {
			return sep + frag
		}
	})

	return str
}

func replaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, str[v[i]:v[i+1]])
		}

		//result += repl(groups)
		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}
