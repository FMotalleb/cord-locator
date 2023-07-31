package core_configuration

type CoreConfiguration struct {
	Address         string   `yaml:"address"`
	AllowTransfer   []string `yaml:"allowTransfer"`
	DefaultProvider string   `yaml:"defaultProvider"`
}

func (r *CoreConfiguration) Validate() bool {
	//TODO add validation checks
	return true
}
