package config

import (
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/config/globals"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/provider"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/rule"
)

// Config is configuration of the dns proxy
type Config struct {
	Global    globals.CoreConfiguration `yaml:"global"`
	Providers []provider.Provider       `yaml:"providers"`
	Rules     []rule.Rule               `yaml:"rules"`
}

// Validate will check current configuration (rules/providers/...)
func (c *Config) Validate() bool {
	for _, p := range c.Providers {
		if !p.Validate() {
			panic("validation failed for providers")
		}
	}
	for _, r := range c.Rules {
		if !r.Validate() {
			panic("validation failed for rules")
		}
	}
	if !c.Global.Validate() {
		panic("validation failed for rules")
	}
	if c.getDefaultProvider() == nil {
		panic("default provider was not found")
	}
	return true
}

func (c *Config) getDefaultProvider() *provider.Provider {
	return c.findProvider(c.Global.DefaultProvider)
}
func (c *Config) findProvider(name string) *provider.Provider {
	for _, p := range c.Providers {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

func (c *Config) findRuleFor(address string) *rule.Rule {
	for _, r := range c.Rules {
		if r.Match(address) {
			return &r
		}
	}
	return nil //, fmt.Errorf("no rule was found for address: `%s`", address)
}
