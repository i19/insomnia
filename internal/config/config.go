package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

type c struct {
	Port                  int
	EtcdHosts             []string
	EtcdLintingRulePrefix string

	SessionAddress string
	ProjectAddress string

	SessionCatch struct {
		On  bool
		TTL int64
	}
}

var (
	Config c
	doOnce sync.Once
)

func Init(filePath string) {
	doOnce.Do(func() {
		content, err := os.ReadFile(filePath)
		if err != nil {
			panic(fmt.Sprintf("failed to load config file: %s", err.Error()))
		}

		if err = yaml.Unmarshal(content, &Config); err != nil {
			panic(fmt.Sprintf("failed to generate config: %s", err.Error()))
		}

		runningConfig, _ := yaml.Marshal(&Config)
		fmt.Printf("--- configuration ---\n\n%s\n\n--- configuration ---", string(runningConfig))
		time.Sleep(time.Second)
	})
}
