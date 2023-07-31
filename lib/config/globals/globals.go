package globals

// CoreConfiguration is holds information of dns server listen port,AllowTransfer and default provider
type CoreConfiguration struct {
	Address         string   `yaml:"address"`
	AllowTransfer   []string `yaml:"allowTransfer"`
	DefaultProvider string   `yaml:"defaultProvider"`
}

// Validate will check core configurations and verify it
func (r *CoreConfiguration) Validate() bool {
	//TODO add validation checks
	return true
}
