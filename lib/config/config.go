package config

import (
	"github.com/FMotalleb/cord-locator/lib/config/globals"
	"github.com/FMotalleb/cord-locator/lib/provider"
	"github.com/FMotalleb/cord-locator/lib/rule"
)

// Config is configuration of the dns proxy
type Config struct {
	Global           globals.CoreConfiguration `yaml:"global"`
	Providers        []provider.Provider       `yaml:"providers"`
	Rules            []rule.Rule               `yaml:"rules"`
	defaultProviders []provider.Provider
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
	//TODO: check length of rules and providers
	if len(c.GetDefaultProviders()) == 0 {
		panic("default provider was not found")
	}
	return true
}

// GetDefaultProviders set in global config
func (c *Config) GetDefaultProviders() []provider.Provider {

	if c.defaultProviders == nil {
		c.defaultProviders = c.FindProviders(c.Global.DefaultProviders)
	}

	return c.defaultProviders
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
