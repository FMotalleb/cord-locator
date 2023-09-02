package rule

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"
)

// Rule set of rules to find resolver of each request
type Rule struct {
	Name           *string            `yaml:"name"`
	Matcher        string             `yaml:"matcher"`
	MatcherParams  []string           `yaml:"matcherParams"`
	Resolvers      []string           `yaml:"resolvers,alias:resolvers"`
	ResolverParams *string            `yaml:"resolverParams"`
	Raw            *map[string]string `yaml:"raw"`
	IsBlocked      bool               `yaml:"isBlocked,alias:blocked,default:false"`
}

func (r *Rule) String() string {
	return fmt.Sprintf("rule(Name: %v,Matcher: %s,MatcherParams: %v,Resolver: %v,Raw: %v, IsBlocked: %v)", r.Name, r.Matcher, r.MatcherParams, r.Resolvers, r.IsBlocked, r.IsBlocked)
}

// GetRaw will try to find raw response in the rule
func (r *Rule) GetRaw(qType string) *string {
	if r.Raw == nil {
		return nil
	}
	for key, value := range *r.Raw {
		if strings.ToLower(qType) == strings.ToLower(key) {
			return &value
		}
	}
	return nil
}

// Match returns true if given address matches this rule
func (r *Rule) Match(address string) bool {
	switch r.Matcher {
	case "regex":
		for _, pattern := range r.MatcherParams {
			log.Trace().Msgf("compiling regex: %s", pattern)
			matcher := regexp.MustCompile(pattern)
			if matcher != nil {
				result := matcher.FindIndex([]byte(address))
				if result != nil {
					if result[0] == 0 {
						log.Trace().Msgf("address: %s matched regex pattern: %s", address, pattern)
						return true
					}
				}
				log.Trace().Msgf("address: %s could not match regex pattern: %s", address, pattern)
			}
		}
	case "exact":
		for _, pattern := range r.MatcherParams {
			if address == pattern {
				log.Trace().Msgf("address: %s matched exactly: %s", address, pattern)
				return true
			}
		}
	}
	return false
}

func (r *Rule) validateResolveMethod() bool {
	if r.IsBlocked {
		return true
	}
	if (len(r.Resolvers) == 0) && (r.Raw == nil) {
		log.Error().Msgf("no resolver or raw response found for rule: %s", r)
		return false
	}
	return true
}
func (r *Rule) validateMatcher() bool {
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
		return true
	case "exact":
		if len(r.MatcherParams) == 0 {
			log.Debug().Msgf("failed to validate rule:%s, received exact matcher with no params", r)
			return false
		}
		return true
	}
	log.Debug().Msgf("failed to validate rule:%s", r)
	return false
}

// Validate this rule is correctly configured
func (r *Rule) Validate() bool {
	return r.validateMatcher() && r.validateResolveMethod()
}
