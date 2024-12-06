package printer

import (
	"log"
	"regexp"
	"strconv"
)

type Requisites struct {
	Requisites map[int]bool
}

// Rules is a struct that contains the rules for the book
// They are stored as a map where the key is the second page and the values the requisites for that page
type Rules struct {
	rules map[int]Requisites
}

func NewRules() *Rules {
	return &Rules{
		rules: make(map[int]Requisites),
	}
}

func (r *Rules) At(k int) {
}

func (r *Rules) isRequisite(page int, requisite int) bool {
	if _, ok := r.rules[page].Requisites[requisite]; !ok {
		return false
	}
	return true
}

func (r *Rules) addRequisite(page int, requisite int) {
	if _, ok := r.rules[page]; !ok {
		r.rules[page] = Requisites{
			Requisites: make(map[int]bool),
		}
	}
	r.rules[page].Requisites[requisite] = true
}

func (r *Rules) AddRule(s string) {
	rg := regexp.MustCompile(`\d+`)
	matches := rg.FindAllString(s, -1)
	if len(matches) != 2 {
		log.Printf("Incorrect amount of numbers found in rule: %v", s)
		return
	}
	pages := make([]int, 2) // 0: first page, 1: second page
	for i, value := range matches {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("Failed to convert string to int: %v", err)
			return
		}
		pages[i] = v
	}
	r.addRequisite(pages[1], pages[0])
}
