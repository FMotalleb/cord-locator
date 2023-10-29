package data

import (
	"github.com/FMotalleb/cord-locator/lib/utils/iterable"
	"github.com/miekg/dns"
)

// Question of a dns query
type Question struct {
	Type  string
	Class string
	Name  string
	From  *string
}

// NewQuestion built with given parameters
func NewQuestion(qType string, qClass string, qName string) Question {
	return Question{
		Type:  qType,
		Class: qClass,
		Name:  qName,
	}
}

// GetQuestions from a dns message
func GetQuestions(msg *dns.Msg) []Question {

	if msg == nil || len(msg.Question) == 0 {
		return make([]Question, 0)
	}
	mapper := func(entry dns.Question) Question {
		qType := dns.TypeToString[entry.Qtype]
		qClass := dns.ClassToString[entry.Qclass]
		return NewQuestion(qType, qClass, entry.Name)
	}
	return iterable.Map[dns.Question, Question](msg.Question, mapper)
}
