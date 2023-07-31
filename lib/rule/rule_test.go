package rule_test

import (
	"testing"

	"github.com/FMotalleb/dns-reverse-proxy-docker/lib/rule"
)

func TestRegexMatcherPass(t *testing.T) {
	rgp := make([]string, 0)
	rgp = append(rgp, ".*")
	name := "test"
	item := rule.Rule{
		Matcher:       "regex",
		MatcherParams: rgp,
		Resolver:      &name,
	}
	if !item.Match("google.com") {
		t.Error("matcher is working incorrectly, expected to match `google.com`")
	}
}
func TestRegexMatcherPassEvenIfFailedBefore(t *testing.T) {
	rgp := make([]string, 0)
	rgp = append(rgp, ".com")
	rgp = append(rgp, ".*")
	name := "test"
	item := rule.Rule{
		Matcher:       "regex",
		MatcherParams: rgp,
		Resolver:      &name,
	}
	if !item.Match("google.com") {
		t.Error("matcher is working incorrectly, expected to match `google.com`")
	}
}
func TestRegexMatcherFail(t *testing.T) {
	rgp := make([]string, 0)
	rgp = append(rgp, "not-google")
	name := "test"
	item := rule.Rule{
		Matcher:       "regex",
		MatcherParams: rgp,
		Resolver:      &name,
	}
	if item.Match("google.com") {
		t.Error("matcher is working incorrectly, expected to fail matching")
	}
}
func TestValidateRegexFail(t *testing.T) {
	name := "test"
	item := rule.Rule{
		Matcher: "regex",
		Name:    &name,
	}
	if item.Validate() {
		t.Error("item has no regex parameter but has regex as matcher which is invalid and must fail")
	}
}

func TestValidateRegexPass(t *testing.T) {
	rgp := make([]string, 1)
	rgp = append(rgp, ".*")
	name := "test"
	item := rule.Rule{
		Matcher:       "regex",
		MatcherParams: rgp,
		Resolver:      &name,
		Name:          &name,
	}
	if !item.Validate() {
		t.Error("Item has valid regex and configuration it must pass")
	}
}
func TestValidateRegexPassWithMissingName(t *testing.T) {
	rgp := make([]string, 1)
	rgp = append(rgp, ".*")
	name := "test"
	item := rule.Rule{
		Matcher:       "regex",
		MatcherParams: rgp,
		Resolver:      &name,
	}
	if !item.Validate() {
		t.Error("Item has valid regex and configuration it must pass, only missing Name")
	}
}
func TestValidateRegexFailOnWrongRegex(t *testing.T) {
	rgp := make([]string, 1)
	rgp = append(rgp, "**")
	name := "test"
	item := rule.Rule{
		Matcher:       "regex",
		MatcherParams: rgp,
		Resolver:      &name,
	}
	if item.Validate() {
		t.Error("Given regex is invalid this item should fail at validation")
	}
}
func TestValidateFailOnMissingResolver(t *testing.T) {
	rgp := make([]string, 1)
	rgp = append(rgp, ".*")
	item := rule.Rule{
		Matcher:       "regex",
		MatcherParams: rgp,
	}
	if item.Validate() {
		t.Error("items must fail at validation if they do not identify their providers")
	}
}
func TestFailValidate(t *testing.T) {
	item := rule.Rule{}
	if item.Validate() {
		t.Error("empty configuration should fail")
	}
}
