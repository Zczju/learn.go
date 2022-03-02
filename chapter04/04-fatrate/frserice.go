package main

import (
	"log"
)

type fatRateService struct {
	s *fatRateSuggestions
}

func (frsvc *fatRateService) GiveSuggestionToPerson(person *Person) string {
	if err := person.calcBmi(); err != nil {
		log.Printf("无法给%s计算bmi:%v", person.name, err)
		return "错误"
	}
	person.calcFatRate()
	return frsvc.s.GetSuggestions(person)
}

func (frsvc *fatRateService) GiveSuggestionToPersons(persons ...*Person) map[string]string {
	out := map[string]string{}
	for _, item := range persons {
		out[item.name] = frsvc.GiveSuggestionToPerson(item)
	}
	return out
}
