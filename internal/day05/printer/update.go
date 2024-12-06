package printer

import (
	"log"
	"regexp"
	"strconv"
)

type Update struct {
	Update []int
}

func NewUpdate(s string) *Update {
	rg := regexp.MustCompile(`\d+`)
	matches := rg.FindAllString(s, -1)
	if len(matches) == 0 {
		log.Printf("No numbers found in update: %v", s)
		return nil
	}
	updates := make([]int, len(matches))
	for i, value := range matches {
		v, err := strconv.Atoi(value)
		if err != nil {
			log.Printf("Failed to convert string to int: %v", err)
			return nil
		}
		updates[i] = v
	}
	return &Update{
		Update: updates,
	}
}

func (u *Update) GetMiddleElement() int {
	return u.Update[len(u.Update)/2]
}

func (u *Update) GetSlice() []int {
	return u.Update
}

func SumMiddleElements(updates []Update) int {
	sum := 0
	for _, update := range updates {
		sum += update.GetMiddleElement()
	}
	return sum
}
