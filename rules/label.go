package rules

import (
	"regexp"

	"github.com/maxmellon/onion/types"
)

var labels = []string{"新規", "継続", "完了"}

func contains(s []string, t string) bool {
	for _, v := range s {
		reg := regexp.MustCompile(v)
		if reg.MatchString(t) {
			return true
		}
	}
	return false
}

func LabelStyle(m types.Tml) types.Status {
	if m.Category != "LIST_ITEM" {
		return types.Status{Code: types.S}
	}

	rep := regexp.MustCompile(`○ (\S+) \[(\d+)\]`)
	strList := rep.FindAllString(m.Statement, 1)
	for _, v := range strList {
		if !contains(labels, v) {
			return types.Status{
				Line:     m.Line,
				Column:   3,
				Message:  "[新規], [継続], [完了] のいずれかであるべき",
				RuleName: "LabelStyle",
				Code:     types.C,
			}
		}
	}
	return types.Status{Code: types.S}
}
