package rule

import (
	"fmt"
	"regexp"

	log "github.com/rs/zerolog/log"
)

type Rule struct {
	Name          *string  `yaml:"name"`
	Matcher       string   `yaml:"matcher"`
	MatcherParams []string `yaml:"matcherParams"`
	Resolver      string   `yaml:"resolver"`
}

func (r *Rule) String() string {
	return fmt.Sprintf("rule(Name: %s,Matcher: %s,MatcherParams: %v,Resolver: %s)", *r.Name, r.Matcher, r.MatcherParams, r.Resolver)
}

func (r *Rule) Match(address string) bool {
	switch r.Matcher {
	case "regex":
		for _, pattern := range r.MatcherParams {
			matcher := regexp.MustCompile(pattern)
			if matcher != nil {
				result := matcher.FindIndex([]byte(address))
				if result != nil {
					return result[0] == 0
				}
			}
		}
	}
	return false
}
func (r *Rule) Validate() bool {
	switch r.Matcher {
	case "regex":
		return true
	}
	log.Fatal().Msgf("failed to validate rule:%s", r)
	return false
}
