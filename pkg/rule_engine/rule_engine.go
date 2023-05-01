package rule_engine

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/xeipuuv/gojsonschema"
	"insomnia/internal/platform/etcd"
	"insomnia/pkg/utils"
)

var (
	ruleByProjectID = make(map[string]*gojsonschema.Schema)
	lock            sync.RWMutex
	once            sync.Once
)

type ValidationResult struct {
	IsValid bool
	Errors  []string
}

func GenerateSchemaFromYaml(lrInYaml []byte) (*gojsonschema.Schema, error) {
	jsonByte, err := utils.YamlToJson(lrInYaml)
	if err != nil {
		return nil, fmt.Errorf("failed to json.Marshal lintingRule: %s", err.Error())
	}

	// todo 针对未定义的 format 进行检查
	schema, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(jsonByte))
	if err != nil {
		return nil, fmt.Errorf("failed to compile lintingRule: %s", err.Error())
	}
	return schema, nil
}

func Init(lintingRulePath string) {
	once.Do(
		func() {
			ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
			defer cancel()
			resp, err := etcd.Client.Get(ctx, lintingRulePath)
			if err != nil {
				panic(fmt.Sprintf("Failed to get all lintingRule from etcd: %s", err.Error()))
			}
			for _, kv := range resp.Kvs {
				projectID := string(kv.Key)
				if schema, err := GenerateSchemaFromYaml(kv.Value); err != nil {
					panic(err.Error())
				} else {
					ruleByProjectID[projectID] = schema
				}
			}

			go func() {
				for change := range etcd.Client.Watch(context.Background(), lintingRulePath) {
					for _, event := range change.Events {
						switch event.Type.String() {
						case etcd.ActionPUT:
							projectID := string(event.Kv.Key)
							if schema, err := GenerateSchemaFromYaml(event.Kv.Value); err != nil {
								panic(err.Error())
							} else {
								lock.Lock()
								ruleByProjectID[projectID] = schema
								lock.Unlock()
							}
						case etcd.ActionDELETE:
							projectID := string(event.Kv.Key)
							lock.Lock()
							delete(ruleByProjectID, projectID)
							lock.Unlock()
						}
					}
				}
			}()
		},
	)
}

func InitByRaw(projectID string, ruleYamlByte []byte) {
	schema, err := GenerateSchemaFromYaml(ruleYamlByte)
	if err != nil {
		panic(err)
	}

	ruleByProjectID[projectID] = schema
}

func Render(projectID string, jsonContent []byte) (*ValidationResult, error) {
	lock.RLock()
	schema, ok := ruleByProjectID[projectID]
	lock.RUnlock()

	if !ok {
		return nil, fmt.Errorf("project %s not valid", projectID)
	}

	return RenderRaw(schema, jsonContent)
}

func RenderRaw(schema *gojsonschema.Schema, jsonContent []byte) (*ValidationResult, error) {
	result, err := schema.Validate(gojsonschema.NewBytesLoader(jsonContent))
	if err != nil {
		return nil, err
	}
	if result.Valid() {
		return &ValidationResult{
			IsValid: true,
		}, nil
	}

	r := ValidationResult{IsValid: false, Errors: make([]string, len(result.Errors()))}
	for i, e := range result.Errors() {
		r.Errors[i] = fmt.Sprintf("%s : %s", e.Context().String(), e.Description())
	}

	return &r, nil
}
