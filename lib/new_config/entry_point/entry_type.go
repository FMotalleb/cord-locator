package entrypoint

// Type indicates the entrypoint connection method
// Typically
//
//	UDP: port 53
//	TCP: random port above 1023
type Type int

const (
	// UDP dns type normally listens on port 53 and relies on UDP packets,
	// this is default dns type supported by almost any device on the planet
	UDP Type = iota
	// TLS
	// HTTPS

	// Undefined dns type is an unacceptable dns type and provider of this type will be unused
	Undefined = -1
)

func parseType(typeName *string) Type {
	if typeName == nil {
		return Undefined
	}
	switch *typeName {
	case "udp":
		return UDP
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
	case UDP:
		return "UDP"
	// case 1:
	// 	return "TLS"
	// case 2:
	// 	return "HTTPS"
	default:
		return "Undefined"
	}
}
