package config_test

import (
	"testing"

	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/config"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/config/globals"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/provider"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/rule"
)

func makeArray[T any](args ...T) []T {
	arr := make([]T, 0)
	for _, item := range args {
		arr = append(arr, item)
	}
	return arr
}
func TestValidateFailingOnEmptyConfig(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("config is missing mandatory parts it must fail (missing panic)")
		}
	}()
	DNSConfig := config.Config{}
	isValid := DNSConfig.Validate()
	if isValid {
		t.Error("config is missing mandatory parts it must fail")
	}
}
func TestValidateFailingOnEmptyDefaultProviderConfig(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("config is missing mandatory parts it must fail (missing panic)")
		}
	}()
	DNSConfig := config.Config{
		Global: globals.CoreConfiguration{
			Address:       ":53",
			AllowTransfer: makeArray("0.0.0.0"),
			// DefaultProviders: makeArray("cf"),
		},
	}
	isValid := DNSConfig.Validate()
	if isValid {
		t.Error("config is missing mandatory parts it must fail")
	}
}

func TestValidateFailOnMissingProviderInProviderList(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("config is missing mandatory parts it must fail (missing panic)")
		}
	}()
	DNSConfig := config.Config{
		Global: globals.CoreConfiguration{
			Address:          ":53",
			AllowTransfer:    makeArray("0.0.0.0"),
			DefaultProviders: makeArray("cf"),
		},
	}
	isValid := DNSConfig.Validate()
	if isValid {
		t.Error("config is carrying mandatory parts it must pass")
	}
}

func TestValidatePassOnCompleteConfig(t *testing.T) {
	DNSConfig := config.Config{
		Global: globals.CoreConfiguration{
			Address:          ":53",
			AllowTransfer:    makeArray("0.0.0.0"),
			DefaultProviders: makeArray("cf"),
		},
		Providers: makeArray(provider.Provider{
			Name: "cf",
			IP:   makeArray("1.1.1.1:53"),
		}),
		Rules: makeArray(rule.Rule{
			Matcher:       "regex",
			MatcherParams: makeArray(".*"),
			Resolvers:     makeArray("cf"),
		}),
	}
	isValid := DNSConfig.Validate()
	if !isValid {
		t.Error("config is carrying mandatory parts it must pass")
	}
}
func TestValidateFailOnMissingProviderConfig(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("config is missing mandatory parts it must fail (missing panic)")
		}
	}()
	DNSConfig := config.Config{
		Global: globals.CoreConfiguration{
			Address:          ":53",
			AllowTransfer:    makeArray("0.0.0.0"),
			DefaultProviders: makeArray("cf"),
		},
		Providers: makeArray(provider.Provider{
			Name: "cf",
			IP:   makeArray("1.1.1.1"),
		}),
		Rules: makeArray(rule.Rule{
			Matcher:       "regex",
			MatcherParams: makeArray(".*"),
			Resolvers:     makeArray("cf"),
		}),
	}
	isValid := DNSConfig.Validate()
	if isValid {
		t.Error("config is carrying mandatory parts it must pass")
	}
}

func TestValidateFailOnMissingRuleConfig(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Error("config is missing mandatory parts it must fail (missing panic)")
		}
	}()
	DNSConfig := config.Config{
		Global: globals.CoreConfiguration{
			Address:          ":53",
			AllowTransfer:    makeArray("0.0.0.0"),
			DefaultProviders: makeArray("cf"),
		},
		Providers: makeArray(provider.Provider{
			Name: "cf",
			IP:   makeArray("1.1.1.1:53"),
		}),
		Rules: makeArray(rule.Rule{
			Matcher:       "regex",
			MatcherParams: makeArray("**"),
			Resolvers:     makeArray("cf"),
		}),
	}
	isValid := DNSConfig.Validate()
	if isValid {
		t.Error("config is carrying mandatory parts it must pass")
	}
}
func TestGetRulePass(t *testing.T) {
	DNSConfig := config.Config{
		Global: globals.CoreConfiguration{
			Address:          ":53",
			AllowTransfer:    makeArray("0.0.0.0"),
			DefaultProviders: makeArray("cf"),
		},
		Providers: makeArray(provider.Provider{
			Name: "cf",
			IP:   makeArray("1.1.1.1:53"),
		}),
		Rules: makeArray(rule.Rule{
			Matcher:       "regex",
			MatcherParams: makeArray(".*"),
			Resolvers:     makeArray("cf"),
		}),
	}
	rule := DNSConfig.FindRuleFor("test.com")
	if rule == nil {
		t.Error("config is carrying mandatory parts it must pass")
	}
}
func TestGetRuleFail(t *testing.T) {
	DNSConfig := config.Config{
		Global: globals.CoreConfiguration{
			Address:          ":53",
			AllowTransfer:    makeArray("0.0.0.0"),
			DefaultProviders: makeArray("cf"),
		},
		Providers: makeArray(provider.Provider{
			Name: "cf",
			IP:   makeArray("1.1.1.1:53"),
		}),
		Rules: makeArray(rule.Rule{
			Matcher:       "regex",
			MatcherParams: makeArray(".*\\.org"),
			Resolvers:     makeArray("cf"),
		}),
	}
	rule := DNSConfig.FindRuleFor("test.com")
	if rule != nil {
		t.Error("config is carrying mandatory parts it must pass")
	}
}
