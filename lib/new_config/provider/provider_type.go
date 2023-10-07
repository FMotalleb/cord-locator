package provider

// Type can be of type Raw,Lua,Tls,HTTPS...
type Type int

const (
	// Raw dns type which is ipv4:53 udp dns server
	Raw Type = iota
	// LUA
	// TLS
	// HTTPS
	// MYSQL
	// SQLLITE
	// CSV

	// Undefined dns type is an unacceptable dns type and provider of this type will be unused
	Undefined = -1
)

func parseType(typeName *string) Type {
	if typeName == nil {
		return Undefined
	}
	switch *typeName {
	case "raw":
		return Raw
	// case "tls":
	// 	return TLS
	// case "https":
	// 	return HTTPS
	default:
		return Undefined
	}
}
func (receiver Type) String() string {
	switch receiver {
	case 0:
		return "Raw"
	// case 1:
	// 	return "TLS"
	// case 2:
	// 	return "HTTPS"
	default:
		return "Undefined"
	}
}
