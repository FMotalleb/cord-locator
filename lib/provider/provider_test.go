package provider_test

import (
	"testing"

	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/provider"
)

func TestValidatePass(t *testing.T) {
	ips := make([]string, 0)
	ips = append(ips, "0.0.0.0:53")
	name := "test"
	item := provider.Provider{
		Name: name,
		IP:   ips,
	}
	if !item.Validate() {
		t.Error("Item has valid regex and configuration it must pass")
	}
}
func TestValidateFailMissingIps(t *testing.T) {
	ips := make([]string, 0)
	name := "test"
	item := provider.Provider{
		Name: name,
		IP:   ips,
	}
	if item.Validate() {
		t.Error("Item has valid regex and configuration it must pass")
	}
}
func TestValidateFailIPIssue(t *testing.T) {
	ips := make([]string, 0)
	name := "test"
	ips = append(ips, "0.0.0.0")
	item := provider.Provider{
		Name: name,
		IP:   ips,
	}
	if item.Validate() {
		t.Error("Item has valid regex and configuration it must pass")
	}
}
