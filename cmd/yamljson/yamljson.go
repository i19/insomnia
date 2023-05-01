package yamljson

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"insomnia/pkg/utils"
)

var YamlToJson = &cli.Command{
	Name:  "y2j",
	Usage: "yaml to json",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "f",
			Usage: "path to yaml file",
			Value: "./yaml.yaml",
		},
	},
	Action: func(cCtx *cli.Context) error {
		content, err := os.ReadFile(cCtx.String("f"))
		if err != nil {
			panic(err)
		}

		if result, err := utils.YamlToJson(content); err != nil {
			panic(err)
		} else {
			fmt.Print(string(result))
		}
		return nil
	},
}

var JsonToYaml = &cli.Command{
	Name:  "j2y",
	Usage: "json to yaml",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "f",
			Usage: "path to yaml file",
			Value: "./json.json",
		},
	},
	Action: func(cCtx *cli.Context) error {
		content, err := os.ReadFile(cCtx.String("f"))
		if err != nil {
			panic(err)
		}

		if result, err := utils.JsonToYaml(content); err != nil {
			panic(err)
		} else {
			fmt.Print(string(result))
		}
		return nil
	},
}
