package provider

import (
	"fmt"
	"strings"

	"github.com/FMotalleb/cord-locator/lib/validator"
	"github.com/miekg/dns"
)

// Provider is (in normal cases) an external recourse the is able to resolve dns questions
type Provider interface {
	GetName() string
	GetType() Type
	Resolve([]dns.Question) (answer []dns.RR, err error)
	String() string
}

// Data of provider configuration in yaml file
type Data struct {
	Name      *string   `yaml:"name,omitempty"`
	Type      *string   `yaml:"type,omitempty"`
	Addresses *[]string `yaml:"addresses,omitempty"`
	Provider
	validator.Validatable
}

// GetName of this dns provider (used in rules to point to this provider)
func (receiver Data) GetName() string {
	if receiver.Name == nil {
		panic("provider name is missing")
	}
	return *receiver.Name
}

// GetType of this dns provider
func (receiver Data) GetType() Type {
	if receiver.Type == nil {
		return UDP
	}
	current := parseType(receiver.Type)
	return current
}

// Resolve a dns request and returns its answer or error
func (receiver Data) Resolve([]dns.Question) (answer []dns.RR, err error) {
	// TODO: Implement Resolve method
	panic("UnImplemented")
}

func (receiver Data) String() string {
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

// Validate this instance of provider data to make sure it will work correctly
func (receiver Data) Validate() error {
	if receiver.Name == nil {
		return validator.NewValidationError(
			"a name for every provider -> providers.*.name",
			"no name",
			"missing name on a provider in yaml config")
	}
	switch receiver.GetType() {
	case Undefined:
		return validator.NewValidationError(
			"one of (udp,tcp,lua,https,tls...) in providers.*.type",
			fmt.Sprintf("`%s`", *receiver.Type),
			fmt.Sprintf("the type in provider: `%s`", receiver.GetName()))
	default:
		break
	}
	return nil
}
