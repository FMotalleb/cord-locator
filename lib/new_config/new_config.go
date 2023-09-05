package newconfig

import entrypoint "github.com/FMotalleb/dns-reverse-proxy-docker/lib/new_config/entry_point"

type Config struct {
	EntryPoints []entrypoint.EntryPoint `yaml:"entry_points,flow"`
}
