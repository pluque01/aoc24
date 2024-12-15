package day14

import "testing"

func TestMoveRobot(t *testing.T) {
	testRobots := []string{
		"p=0,4 v=3,-3",
		"p=6,3 v=-1,-3",
		"p=10,3 v=-1,2",
		"p=2,0 v=2,-1",
		"p=0,0 v=1,3",
		"p=3,0 v=-2,-2",
		"p=7,6 v=-1,-3",
		"p=3,0 v=-1,-2",
		"p=9,3 v=2,3",
		"p=7,3 v=-1,2",
		"p=2,4 v=2,-3",
		"p=9,5 v=-3,-3",
	}
	bm := NewBathroomMap(11, 7)
	for _, tr := range testRobots {
		r := NewRobot(tr)
		bm.MoveRobot(r, 100)
	}
	safetyFactor := bm.CalculateSafetyFactor()
	if safetyFactor != 12 {
		t.Errorf("Expected safety factor 12, got %d", safetyFactor)
	}
}
