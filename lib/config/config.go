package config

import (
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/config/globals"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/provider"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/rule"
)

// Config is configuration of the dns proxy
type Config struct {
	Global    globals.CoreConfiguration `yaml:"global"`
	providers []provider.Provider       `yaml:"providers"`
	rules     []rule.Rule               `yaml:"rules"`
}

// Validate will check current configuration (rules/providers/...)
func (config Config) Validate() bool {
	for _, provider := range config.providers {
		if !provider.Validate() {
			panic("validation failed for providers")
		}
	}
	for _, rule := range config.rules {
		if !rule.Validate() {
			panic("validation failed for rules")
		}
	}
	if !config.Global.Validate() {
		panic("validation failed for rules")
	}
	if config.getDefaultProvider() == nil {
		panic("default provider was not found")
	}
	return true
}

func (config Config) getDefaultProvider() *provider.Provider {
	return config.findProvider(config.Global.DefaultProvider)
}
func (config Config) findProvider(name string) *provider.Provider {
	for _, p := range config.providers {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

func (config Config) findRuleFor(address string) *rule.Rule {
	for _, r := range config.rules {
		if r.Match(address) {
			return &r
		}
	}
	return nil //, fmt.Errorf("no rule was found for address: `%s`", address)
}
