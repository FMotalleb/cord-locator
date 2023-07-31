package rule_test

import (
	"testing"

	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/rule"
)

func TestValidateRegexFail(t *testing.T) {
	name := "test"
	item := rule.Rule{
		Matcher: "regex",
		Name:    &name,
	}
	if item.Validate() {
		t.Error("regex validation failed")
	}
}
func TestValidateRegexPass(t *testing.T) {
	rgp := make([]string, 2)
	rgp = append(rgp, ".*")
	rgp = append(rgp, ".*")
	name := "test"
	item := rule.Rule{
		Matcher:       "regex",
		MatcherParams: rgp,
		Resolver:      &name,
	}
	if !item.Validate() {
		t.Error("regex validation failed")
	}
}
func TestFailValidate(t *testing.T) {
	item := rule.Rule{}
	if item.Validate() {
		t.Error("this test must fail")
	}
}
