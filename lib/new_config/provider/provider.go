package provider

import (
	"fmt"
	"strings"

	"github.com/FMotalleb/cord-locator/lib/validator"
	"github.com/miekg/dns"
)

type Provider interface {
	GetName() string
	GetType() ProviderType
	Resolve([]dns.Question) (answer []dns.RR, err error)
	String() string
	validator.Validatable
}

type ProviderData struct {
	Name      *string   `yaml:"name,omitempty"`
	Type      *string   `yaml:"type,omitempty"`
	Addresses *[]string `yaml:"addresses,omitempty"`
	Provider
}

func (receiver ProviderData) GetName() string {
	if receiver.Name == nil {
		panic("provider name is missing")
	}
	return *receiver.Name
}

func (receiver ProviderData) GetType() ProviderType {
	if receiver.Type == nil {
		return Raw
	}
	current := parseType(receiver.Type)
	return current
}

func (receiver ProviderData) Resolve([]dns.Question) (answer []dns.RR, err error) {
	// TODO: Implement Resolve method
	panic("UnImplemented")
}

func (receiver ProviderData) String() string {
	buffer := strings.Builder{}
	buffer.WriteString("Provider(")
	buffer.WriteString(fmt.Sprintf("Name: %s, ", receiver.GetName()))
	if receiver.Addresses != nil && len(*receiver.Addresses) > 0 {
		addr := *receiver.Addresses
		buffer.WriteString(fmt.Sprintf("Addresses: %s, ", strings.Join(addr, ", ")))
	}

	buffer.WriteString(fmt.Sprintf("Type: %v", receiver.GetType().String()))
	buffer.WriteString(")")
	return buffer.String()
}
func (receiver ProviderData) Validate() error {
	if receiver.Name == nil {
		return validator.NewValidationError(
			"a name for every provider -> providers.*.name",
			"no name",
			"missing name on a provider in yaml config")
	}
	switch receiver.GetType() {
	case Undefined:
		return validator.NewValidationError(
			"one of (raw,lua,https,tls...) in providers.*.type",
			fmt.Sprintf("`%s`", *receiver.Type),
			fmt.Sprintf("the type in provider: `%s`", receiver.GetName()))
	default:
		break
	}
	return nil
}
