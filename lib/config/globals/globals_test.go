package globals_test

import (
	"testing"

	"github.com/FMotalleb/cord-locator/lib/config/globals"
)

func makeArray[T any](args ...T) []T {
	arr := make([]T, 0)
	for _, item := range args {
		arr = append(arr, item)
	}
	return arr
}
func TestValidateFailingOnEmptyConfig(t *testing.T) {
	DNSConfig := globals.CoreConfiguration{}
	isValid := DNSConfig.Validate()
	if isValid {
		t.Error("config is missing mandatory parts it must fail")
	}
}
func TestValidateFailingOnMissingDefaultProvidersConfig(t *testing.T) {
	DNSConfig := globals.CoreConfiguration{
		AllowTransfer: makeArray("0.0.0.0"),
		Address:       ":53",
	}
	isValid := DNSConfig.Validate()
	if isValid {
		t.Error("config is missing mandatory parts it must fail")
	}
}
func TestValidateFailingOnMissingAddressConfig(t *testing.T) {
	DNSConfig := globals.CoreConfiguration{
		AllowTransfer:    makeArray("0.0.0.0"),
		DefaultProviders: makeArray("cf"),
	}
	isValid := DNSConfig.Validate()
	if isValid {
		t.Error("config is missing mandatory parts it must fail")
	}
}
func TestValidatePass(t *testing.T) {
	DNSConfig := globals.CoreConfiguration{
		AllowTransfer:    makeArray("0.0.0.0"),
		Address:          ":53",
		DefaultProviders: makeArray("cf"),
	}
	isValid := DNSConfig.Validate()
	if !isValid {
		t.Error("config is missing mandatory parts it must fail")
	}
}
