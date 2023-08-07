package config

import (
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/config/globals"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/provider"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/rule"
)

// Config is configuration of the dns proxy
type Config struct {
	Global          globals.CoreConfiguration `yaml:"global"`
	Providers       []provider.Provider       `yaml:"providers"`
	Rules           []rule.Rule               `yaml:"rules"`
	defaultProvider []provider.Provider
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
	if c.GetDefaultProvider() == nil {
		panic("default provider was not found")
	}
	return true
}

// GetDefaultProvider set in global config
func (c *Config) GetDefaultProvider() []provider.Provider {

	if c.defaultProvider == nil {
		c.defaultProvider = c.FindProviders(c.Global.DefaultProvider)
	}

	return c.defaultProvider
}

// FindProvider with given name
func (c *Config) FindProvider(name string) *provider.Provider {
	for _, p := range c.Providers {
		if p.Name == name {
			return &p
		}
	}
	return nil
}

// FindProviders with given names
func (c *Config) FindProviders(names []string) []provider.Provider {
	providers := make([]provider.Provider, 0)
	for _, name := range names {
		for _, p := range c.Providers {
			if p.Name == name {
				providers = append(providers, p)
			}
		}
	}

	return providers
}

// FindRuleFor given address, this will only find first rule that matches given address
func (c *Config) FindRuleFor(address string) *rule.Rule {
	for _, r := range c.Rules {
		if r.Match(address) {
			return &r
		}
	}
	return nil
}
