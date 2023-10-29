package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/miekg/dns"
)

// Answer of a dns query
type Answer struct {
	TTL     int
	Name    string
	Address string
	Type    string
	Class   string
}

// NewAnswer is generated using given parameters
func NewAnswer(aName string, aType string, aClass string, aAddress string, aTTL int) Answer {
	return Answer{
		TTL:     aTTL,
		Name:    aName,
		Address: aAddress,
		Type:    aType,
		Class:   aClass,
	}
}

func (answer Answer) String() string {
	return fmt.Sprintf("%s\t%d\t%s\t%s\t%s", answer.Name, answer.TTL, answer.Class, answer.Type, answer.Address)
}

func (answer Answer) ToRR() (result dns.RR, err error) {
	return dns.NewRR(answer.String())
}

func NewAnswerFromString(str string) Answer {
	var name string
	var aType string
	var class string
	var address string
	var ttl int
	lastFilled := 0
	for _, v := range strings.Split(str, "\t") {
		if len(v) > 0 {
			switch lastFilled {
			case 0:
				name = v
				break
			case 1:
				ttl, _ = strconv.Atoi(v)
				break
			case 2:
				class = v
				break
			case 3:
				aType = v
				break
			case 5:
				address = v
				break
			default:
				address = address + "\t" + v

			}
			lastFilled++
		}
	}
	if lastFilled != 5 {
		panic("unexpected length")
	}

	ans := NewAnswer(name, aType, class, address, ttl)
	return ans
}
func NewAnswerFromRR(r dns.RR) Answer {
	return NewAnswerFromString(r.String())
}
