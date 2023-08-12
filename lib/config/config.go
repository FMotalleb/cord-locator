package config

import (
	"strings"

	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/config/globals"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/provider"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/rule"
	"github.com/rs/zerolog/log"
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
			log.Error().Msgf("a provider(%s) was not able to validate itself check the configuration file", p.Name)
			panic("one of providers are invalid")
		}
	}
	for _, r := range c.Rules {
		if !r.Validate() {
			log.Error().Msgf("a rule(%v) was not able to validate itself check the configuration file", r)
			panic("one of rules are invalid")
		}
	}
	if !c.Global.Validate() {
		log.Error().Msgf("global section in config is invalid")
		panic("global section is invalid")
	}
	c.GetDefaultProviders()
	return true
}

// GetDefaultProviders set in global config
func (c *Config) GetDefaultProviders() []provider.Provider {
	if c.defaultProviders == nil {
		log.Debug().Msg("first try to read default provider.")
		c.defaultProviders = c.FindProviders(c.Global.DefaultProviders)
	}
	return c.defaultProviders
}

// FindProviders with given names
func (c *Config) FindProviders(names []string) []provider.Provider {
	log.Debug().Msgf("finding providers: `%s`", strings.Join(names, ","))
	providers := make([]provider.Provider, len(names))
	index := 0
	expected := len(names)
	for _, name := range names {
		for _, p := range c.Providers {
			if p.Name == name {
				providers[index] = p
				index++
			}
		}
	}
	if expected != index {
		log.Error().Msgf("from %d expected providers only %d providers was found", expected, index)
		panic("some of providers was not found")
	}
	return providers
}

// FindRuleFor given address, this will only find first rule that matches given address
func (c *Config) FindRuleFor(address string) *rule.Rule {
	log.Debug().Msgf("finding rule that matches `%s` address", address)
	for _, r := range c.Rules {
		if r.Match(address) {
			log.Debug().Msgf("found rule: `%v` that matches %s", r.String(), address)
			return &r
		}
	}
	log.Debug().Msgf("no rule found for address `%s`", address)
	return nil
}
