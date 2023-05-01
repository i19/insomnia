package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"insomnia/cmd/server"
	"insomnia/cmd/swagger"
	"insomnia/cmd/validate"
	"insomnia/cmd/yamljson"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Handle requests about custom linting rules  from insomnia"
	app.Commands = []*cli.Command{
		server.Server,
		swagger.Server,
		yamljson.YamlToJson,
		yamljson.JsonToYaml,
		validate.Validate,
	}
	if err := app.Run(os.Args); err != nil {
		panic(err.Error())
	}
}
