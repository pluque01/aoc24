package printer

import "sort"

type Printer struct {
	Rules      Rules
	UpdateList []Update
}

func NewPrinter(content []string) *Printer {
	rules := NewRules()
	updateList := make([]Update, 0)
	var i int
	for i = 0; i < len(content); i++ {
		if content[i] == "" {
			break
		}
		rules.AddRule(content[i])
	}
	for j := i + 1; j < len(content); j++ {
		updateList = append(updateList, *NewUpdate(content[j]))
	}

	return &Printer{
		Rules:      *rules,
		UpdateList: updateList,
	}
}

func (p *Printer) IsRequisite(page int, requisite int) bool {
	return p.Rules.isRequisite(page, requisite)
}

func (p *Printer) GetGoodAndBadUpdates() (goodUpdates []Update, badUpdates []Update) {
	goodUpdates = make([]Update, 0)
	badUpdates = make([]Update, 0)

	for _, update := range p.UpdateList {
		checkedPages := make([]int, 0)
		isUpdateCorrect := true
		for _, u := range update.Update {
			for _, page := range checkedPages {
				if p.IsRequisite(page, u) {
					isUpdateCorrect = false
					break
				}
			}
			if !isUpdateCorrect {
				badUpdates = append(badUpdates, update)
				break
			}
			checkedPages = append(checkedPages, u)
		}
		if len(checkedPages) == len(update.Update) {
			goodUpdates = append(goodUpdates, update)
		}
	}

	// sum := 0
	// for _, update := range goodUpdates {
	// 	sum += update.GetMiddleElement()
	// }

	return goodUpdates, badUpdates
}

func (p *Printer) FixBadUpdates(badUpdates []Update) []Update {
	fixedUpdates := make([]Update, 0)
	for _, update := range badUpdates {
		s := update.GetSlice()
		newUpdate := make([]int, len(s))
		copy(newUpdate, s)
		sort.Slice(newUpdate, func(i, j int) bool {
			return !p.IsRequisite(newUpdate[i], newUpdate[j])
		})
		fixedUpdates = append(fixedUpdates, Update{Update: newUpdate})
	}
	return fixedUpdates
}
