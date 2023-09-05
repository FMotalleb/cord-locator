package entrypoint

import (
	"fmt"
	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/validator"
	"strings"
)

type EntryPoint struct {
	Name *string `yaml:"name,omitempty"`
	Port *int    `yaml:"port,omitempty"`
	Type *string `yaml:"type,omitempty"`
}

func (receiver EntryPoint) GetName() string {
	if receiver.Name == nil {
		return ""
	}
	return *receiver.Name
}
func (receiver EntryPoint) GetPort() int {
	if receiver.Port == nil {
		return 53
	}
	return *receiver.Port
}

func (receiver EntryPoint) getTypeRaw() string {
	if receiver.Type == nil {
		return "<unset>"
	}
	return *receiver.Type
}

func (receiver EntryPoint) GetType() EntryType {
	actual := parseType(receiver.Type)
	switch actual {
	case Undefined:
		return Raw
	default:
		return actual
	}
}

func (receiver EntryPoint) String() string {
	buffer := strings.Builder{}
	buffer.WriteString("EntryPoint(")
	if len(receiver.GetName()) > 0 {
		buffer.WriteString(fmt.Sprintf("Name: %s, ", receiver.GetName()))
	} else {
		buffer.WriteString("Name<Empty>: '', ")
	}
	buffer.WriteString(fmt.Sprintf("Port: %d, ", receiver.GetPort()))
	buffer.WriteString(fmt.Sprintf("Type: %v", parseType(receiver.Type).String()))
	buffer.WriteString(")")
	return buffer.String()
}
func (receiver EntryPoint) Validate() error {
	// if len(receiver.GetName()) == 0 {
	// 	return validator.NewValidationError(
	// 		"to receive entry_points.name",
	// 		"no name or empty name received",
	// 		"missing name for an entry point in config file")
	// }
	if receiver.GetPort() < 1 || receiver.GetPort() > 65535 {
		return validator.NewValidationError(
			"to receive a valid entry_points.port value (1-65535)",
			fmt.Sprintf("%d", receiver.GetPort()),
			fmt.Sprintf("port value in entrypoint: %s", receiver.GetName()))
	}

	switch parseType(receiver.Type) {
	case Undefined:
		return validator.NewValidationError(
			"to receive one of (raw | tls | https) in entry_points.type",
			fmt.Sprintf("given type does not match expectee %s", receiver.getTypeRaw()),
			fmt.Sprintf("type value in entrypoint: %s", receiver.GetName()))
	default:
		break
	}
	return nil
}
