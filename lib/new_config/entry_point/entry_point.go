package entrypoint

import (
	"fmt"
	"strings"

	"github.com/FMotalleb/cord-locator/lib/validator"
)

type EntryPoint interface {
	GetName() string
	GetPort() int
	GetType() EntryType
	String() string
	validator.Validatable
}
type EntryPointData struct {
	Name *string `yaml:"name,omitempty"`
	Port *int    `yaml:"port,omitempty"`
	Type *string `yaml:"type,omitempty"`
	EntryPoint
}

func (receiver EntryPointData) GetName() string {
	if receiver.Name == nil {
		return ""
	}
	return *receiver.Name
}
func (receiver EntryPointData) GetPort() int {
	if receiver.Port == nil {
		return 53
	}
	return *receiver.Port
}

func (receiver EntryPointData) getTypeRaw() string {
	if receiver.Type == nil {
		return "Raw"
	}
	return *receiver.Type
}

func (receiver EntryPointData) GetType() EntryType {
	if receiver.Type == nil {
		return Raw
	}
	current := parseType(receiver.Type)
	return current
}

func (receiver EntryPointData) String() string {
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
func (receiver EntryPointData) Validate() error {
	if receiver.GetPort() < 1 || receiver.GetPort() > 65535 {
		return validator.NewValidationError(
			"a valid entry_points.*.port value (1-65535)",
			fmt.Sprintf("%d", receiver.GetPort()),
			fmt.Sprintf("port value in entrypoint: %s", receiver.GetName()))
	}

	switch parseType(receiver.Type) {
	case Undefined:
		return validator.NewValidationError(
			"one of (raw | tls | https) in entry_points.*.type",
			fmt.Sprintf("given type (%s) does not match expected values", receiver.getTypeRaw()),
			fmt.Sprintf("type value in entrypoint: %s", receiver.GetName()))
	default:
		break
	}
	return nil
}
