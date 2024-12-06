package day05

import (
	"log"
	"pluque01/aoc24/internal/day05/printer"
	"testing"
)

var testData = struct {
	data         []string
	goodUpdates  []printer.Update
	badUpdates   []printer.Update
	fixedUpdates []printer.Update
}{
	[]string{
		"47|53",
		"97|13",
		"97|61",
		"97|47",
		"75|29",
		"61|13",
		"75|53",
		"29|13",
		"97|29",
		"53|29",
		"61|53",
		"97|53",
		"61|29",
		"47|13",
		"75|47",
		"97|75",
		"47|61",
		"75|61",
		"47|29",
		"75|13",
		"53|13",
		"",
		"75,47,61,53,29",
		"97,61,53,29,13",
		"75,29,13",
		"75,97,47,61,53",
		"61,13,29",
		"97,13,75,29,47",
	},
	[]printer.Update{
		{Update: []int{75, 47, 61, 53, 29}},
		{Update: []int{97, 61, 53, 29, 13}},
		{Update: []int{75, 29, 13}},
	},
	[]printer.Update{
		{Update: []int{75, 97, 47, 61, 53}},
		{Update: []int{61, 13, 29}},
		{Update: []int{97, 13, 75, 29, 47}},
	},
	[]printer.Update{
		{Update: []int{97, 75, 47, 61, 53}},
		{Update: []int{61, 29, 13}},
		{Update: []int{97, 75, 47, 29, 13}},
	},
}

func TestGetGoodUpdates(t *testing.T) {
	p := printer.NewPrinter(testData.data)
	g, b := p.GetGoodAndBadUpdates()
	if len(g) != len(testData.goodUpdates) {
		t.Errorf("GetGoodUpdates(%v) = %v; want %v", testData.data, g, testData.goodUpdates)
	}
	if len(b) != len(testData.badUpdates) {
		t.Errorf("GetGoodUpdates(%v) = %v; want %v", testData.data, b, testData.badUpdates)
	}
	for i := range g {
		if len(g[i].Update) != len(testData.goodUpdates[i].Update) {
			t.Errorf("GetGoodUpdates(%v) = %v; want %v", testData.data, g, testData.goodUpdates)
		}
	}
	for i := range b {
		if len(b[i].Update) != len(testData.badUpdates[i].Update) {
			t.Errorf("GetGoodUpdates(%v) = %v; want %v", testData.data, b, testData.badUpdates)
		}
	}
}

func TestFixBadUpdates(t *testing.T) {
	p := printer.NewPrinter(testData.data)
	_, b := p.GetGoodAndBadUpdates()
	fixed := p.FixBadUpdates(b)
	log.Printf("fixed: %v", fixed)

	if len(fixed) != len(testData.fixedUpdates) {
		t.Errorf("FixBadUpdates(%v) = %v; want %v", b, fixed, testData.fixedUpdates)
	}
	for i := range fixed {
		if len(fixed[i].Update) != len(testData.fixedUpdates[i].Update) {
			t.Errorf("FixBadUpdates(%v) = %v; want %v", b, fixed, testData.fixedUpdates)
		}
		for j := range fixed[i].Update {
			if fixed[i].Update[j] != testData.fixedUpdates[i].Update[j] {
				t.Errorf("FixBadUpdates(%v) = %v; want %v", b, fixed, testData.fixedUpdates)
			}
		}
	}
}
