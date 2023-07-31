package config

import (
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/config/core_configuration"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/provider"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/rule"
)

type Config struct {
	Global    core_configuration.CoreConfiguration `yaml:"global"`
	Providers []provider.Provider                  `yaml:"providers"`
	Rules     []rule.Rule                          `yaml:"rules"`
}

func (config Config) Validate() bool {
	for _, provider := range config.Providers {
		if !provider.Validate() {
			panic("validation failed for providers")
		}
	}
	for _, rule := range config.Rules {
		if !rule.Validate() {
			panic("validation failed for rules")
		}
	}
	if !config.Global.Validate() {
		panic("validation failed for rules")
	}
	if config.GetDefaultProvider() == nil {
		panic("default provider was not found")
	}
	return true
}
func (config Config) GetDefaultProvider() *provider.Provider {
	return config.FindProvider(config.Global.DefaultProvider)
}
func (config Config) FindProvider(name string) *provider.Provider {
	for _, p := range config.Providers {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

func (config Config) FindRuleFor(address string) *rule.Rule {
	for _, r := range config.Rules {
		if r.Match(address) {
			return &r
		}
	}
	return nil //, fmt.Errorf("no rule was found for address: `%s`", address)
}
