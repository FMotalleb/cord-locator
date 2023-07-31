package rule

import (
	"fmt"
	"regexp"

	"github.com/rs/zerolog/log"
)

// Rule set of rules to find resolver of each request
type Rule struct {
	Name          *string  `yaml:"name"`
	Matcher       string   `yaml:"matcher"`
	MatcherParams []string `yaml:"matcherParams"`
	Resolver      *string  `yaml:"resolver"`
}

func (r *Rule) String() string {
	return fmt.Sprintf("rule(Name: %v,Matcher: %s,MatcherParams: %v,Resolver: %v)", r.Name, r.Matcher, r.MatcherParams, r.Resolver)
}

// Match returns true if given address matches this rule
func (r *Rule) Match(address string) bool {
	switch r.Matcher {
	case "regex":
		for _, pattern := range r.MatcherParams {
			matcher := regexp.MustCompile(pattern)
			if matcher != nil {
				result := matcher.FindIndex([]byte(address))
				if result != nil {
					if result[0] == 0 {
						return true
					}
				}
			}
		}
	}
	return false
}

// Validate this rule is correctly configured
func (r *Rule) Validate() bool {
	switch r.Matcher {
	case "regex":
		if len(r.MatcherParams) == 0 {
			log.Debug().Msgf("failed to validate rule:%s, received regex matcher with no params", r)
			return false
		}
		for _, rule := range r.MatcherParams {
			_, err := regexp.Compile(rule)
			if err != nil {
				log.Debug().Msgf("failed to validate regex: %s, msg: %v", rule, err)
				return false
			}
			if r.Name != nil {
				log.Debug().Msgf("validation succeeded for rule: `%s` - regexp: `%s`", *r.Name, rule)
			} else {
				log.Debug().Msgf("validation succeeded for an Unnamed Rule - regexp: `%s`", rule)
			}
		}
		if r.Resolver == nil {
			log.Debug().Msgf("resolver is empty in rule: %s", r)
			return false
		}
		return true
	}
	log.Debug().Msgf("failed to validate rule:%s", r)
	return false
}
