package entrypoint

import (
	"fmt"
	"strings"

	"github.com/FMotalleb/cord-locator/lib/validator"
)

type DnsQueryCallback func()

// EntryPoint of dns server
// indicates where to listen for dns queries
type EntryPoint interface {
	GetName() string
	GetPort() int
	GetType() Type
	String() string
	validator.Validatable
}

// Data of entry point contains its (optional) name, port and type
type Data struct {
	Name *string `yaml:"name,omitempty"`
	Port *int    `yaml:"port,omitempty"`
	Type *string `yaml:"type,omitempty"`
	EntryPoint
}

func (receiver Data) GetName() string {
	if receiver.Name == nil {
		return ""
	}
	return *receiver.Name
}
func (receiver Data) GetPort() int {
	if receiver.Port == nil {
		return 53
	}
	return *receiver.Port
}

func (receiver Data) getTypeRaw() string {
	if receiver.Type == nil {
		return "UDP"
	}
	return *receiver.Type
}

func (receiver Data) GetType() Type {
	if receiver.Type == nil {
		return UDP
	}
	current := parseType(receiver.Type)
	return current
}

func (receiver Data) String() string {
	buffer := strings.Builder{}
	buffer.WriteString("EntryPoint(")
	if len(receiver.GetName()) > 0 {
		buffer.WriteString(fmt.Sprintf("Name: %s, ", receiver.GetName()))
	} else {
		buffer.WriteString("Name<Empty>: '', ")
	}
	buffer.WriteString(fmt.Sprintf("Port: %d, ", receiver.GetPort()))
	buffer.WriteString(fmt.Sprintf("Type: %v", receiver.GetType().String()))
	buffer.WriteString(")")
	return buffer.String()
}
func (receiver Data) Validate() error {
	if receiver.GetPort() < 1 || receiver.GetPort() > 65535 {
		return validator.NewValidationError(
			"a valid entry_points.*.port value (1-65535)",
			fmt.Sprintf("%d", receiver.GetPort()),
			fmt.Sprintf("port value in entrypoint: %s", receiver.GetName()))
	}

	switch parseType(receiver.Type) {
	case Undefined:
		return validator.NewValidationError(
			"one of (udp | tcp | tls | https) in entry_points.*.type",
			fmt.Sprintf("given type (%s) does not match expected values", receiver.getTypeRaw()),
			fmt.Sprintf("type value in entrypoint: %s", receiver.GetName()))
	default:
		break
	}
	return nil
}
