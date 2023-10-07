package provider

type ProviderType int

const (
	Raw ProviderType = iota
	// LUA
	// TLS
	// HTTPS
	// MYSQL
	// SQLLITE
	// CSV
	Undefined = -1
)

func parseType(typeName *string) ProviderType {
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
func (receiver ProviderType) String() string {
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
