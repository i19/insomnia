package validate

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"insomnia/pkg/rule_engine"
)

var Validate = &cli.Command{
	Name:  "validate",
	Usage: "yaml to json",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "r",
			Usage: "path to rule file as yaml",
		},
		&cli.StringFlag{
			Name:  "i",
			Usage: "path to input file as json",
		},
	},
	Action: func(cCtx *cli.Context) error {
		rule, err := os.ReadFile(cCtx.String("r"))
		if err != nil {
			panic(err)
		}
		input, err := os.ReadFile(cCtx.String("i"))
		if err != nil {
			panic(err)
		}
		schema, err := rule_engine.GenerateSchemaFromYaml(rule)
		if err != nil {
			panic(err)
		}
		result, err := rule_engine.RenderRaw(schema, input)
		if err != nil {
			panic(err)
		}
		if result.IsValid {
			fmt.Print("input applied with no error")
		} else {
			fmt.Printf("input applied with errors as below: \n")
			fmt.Printf(strings.Join(result.Errors, "\n"))
		}

		return nil
	},
}
