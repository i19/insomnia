package rule_engine

import "testing"

func TestFormatVersion(t *testing.T) {
	rule := `
type: object
properties:
  okVersion:
    type: string
    format: version
  failVersion:
    type: string
    format: version
`

	input := `{"okVersion": "1.0.0","failVersion": "1.0"}`

	schema, err := GenerateSchemaFromYaml([]byte(rule))
	if err != nil {
		t.Fatal(err)
	}
	result, err := RenderRaw(schema, []byte(input))
	if err != nil {
		t.Fatal(err)
	}
	if result.IsValid {
		t.Errorf("result.IsValid should be false")
		return
	}
	if len(result.Errors) != 1 {
		t.Errorf("result.Errors size should be 1")
		return
	}

	if result.Errors[0] != "(root).failVersion : Does not match format 'version'" {
		t.Errorf("match result error ....")
	}
}
