package rule_engine

import (
	"regexp"

	"github.com/xeipuuv/gojsonschema"
)

type VersionFormatChecker struct{}

var versionPattern = regexp.MustCompile(`^(\d+\.)(\d+\.)(\d+)$`)

func (f VersionFormatChecker) IsFormat(input interface{}) bool {
	asString, ok := input.(string)
	if !ok {
		return false
	}
	return versionPattern.MatchString(asString)
}

func init() {
	gojsonschema.FormatCheckers.Add("version", VersionFormatChecker{})
}
