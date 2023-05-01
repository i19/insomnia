package rule_engine

import "testing"

func TestFormatDescribe(t *testing.T) {
	rule := `
type: object
properties:
  okDescription:
    type: string
    format: description
  failDescription:
    type: string
    format: description
`

	input := `{"okDescription": "description hello world","failDescription": "hello world"}`

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

	if result.Errors[0] != "(root).failDescription : Does not match format 'description'" {
		t.Errorf("match result error ....")
	}
}
