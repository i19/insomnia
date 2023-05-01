package rule_engine

import (
	"github.com/xeipuuv/gojsonschema"
	"strings"
)

type DescribeFormatChecker struct{}

func (f DescribeFormatChecker) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if !ok {
		return false
	}

	return strings.HasPrefix(asString, "description")
}

// Add it to the library
func init() {
	gojsonschema.FormatCheckers.Add("description", DescribeFormatChecker{})
}
