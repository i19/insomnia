package server

import (
	"github.com/urfave/cli/v2"
	"insomnia/internal/api/router"
	"insomnia/internal/platform/etcd"
	"insomnia/pkg/rule_engine"

	"insomnia/internal/config"
	"insomnia/internal/service/project_access"
	"insomnia/internal/service/session"
)

var Server = &cli.Command{
	Name:  "server",
	Usage: "start insomnia server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "conf, c",
			Usage: "path to config file",
			Value: "./config.yaml",
		},
	},
	Action: func(cCtx *cli.Context) error {
		config.Init(cCtx.String("conf"))
		etcd.Init(config.Config.EtcdHosts)
		rule_engine.Init(config.Config.EtcdLintingRulePrefix)
		project_access.Init(config.Config.ProjectAddress)
		session.Init(config.Config.SessionAddress)
		router.Run(config.Config.Port)
		return nil
	},
}
