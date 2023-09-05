package entrypoint

type EntryType int

const (
	Raw EntryType = iota
	TLS
	HTTPS
	Undefined = -1
)

func parseType(typeName *string) EntryType {
	if typeName == nil {
		return Undefined
	}
	switch *typeName {
	case "raw":
		return Raw
	case "tls":
		return TLS
	case "https":
		return HTTPS
	default:
		return Undefined
	}
}
func (receiver EntryType) String() string {
	switch receiver {
	case 0:
		return "Raw"
	case 1:
		return "TLS"
	case 2:
		return "HTTPS"
	default:
		return "Undefined"
	}
}
