package data

import (
	"github.com/FMotalleb/cord-locator/lib/utils/iterable"
	"github.com/miekg/dns"
)

type Question struct {
	Type  string
	Class string
	Name  string
}

func GetQuestions(msg *dns.Msg) []Question {
	if msg == nil || len(msg.Question) == 0 {
		return make([]Question, 0)
	}
	mapper := func(entry dns.Question) Question {
		Qtype := dns.TypeToString[entry.Qtype]
		Qclass := dns.ClassToString[entry.Qclass]
		return Question{
			Type:  Qtype,
			Class: Qclass,
			Name:  entry.Name,
		}
	}
	return iterable.Map[dns.Question, Question](msg.Question, mapper)
}
