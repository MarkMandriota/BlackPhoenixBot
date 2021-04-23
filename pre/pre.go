package pre

import (
	"regexp"
	"strings"
)

var (
	r_ident = regexp.MustCompile(`%\w+%`)
	r_value = regexp.MustCompile(`".*"|\S*`)
	r_var   = regexp.MustCompile(`%\w+%(=(".*"|\S*))?`)
)

var lets = make(map[uint64]map[string]string)

type Text struct {
	g  uint64
	in []byte
}

func NewText(group uint64, inner string) *Text {
	return &Text{g: group, in: []byte(inner)}
}

func (t *Text) Parse() string {
	t.in = r_var.ReplaceAllFunc(t.in, func(str []byte) []byte {
		ident := string(r_ident.Find(str))

		var value string
		if i := strings.IndexByte(string(str), '='); i > 0 {
			value = NewText(t.g, strings.Trim(string(r_value.Find(str[i+1:])), "\"")).Parse()
		}

		if len(value) > 0 {
			lets[t.g] = make(map[string]string)
			lets[t.g][ident[1:len(ident)-1]] = value
			return []byte(value)
		}

		return []byte(lets[t.g][string(ident[1:len(ident)-1])])
	})

	return string(t.in)
}
