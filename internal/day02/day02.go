package day02

import (
	"log"
	"pluque01/aoc24/pkg/math"
	"pluque01/aoc24/pkg/reader"
)

func isReportSafe(report []int) bool {
	isSafe := true
	lastInc := 0
	for i := 0; i < len(report)-1; i++ {
		incr := report[i] - report[i+1]
		if math.Abs(incr) == 0 || math.Abs(incr) > 3 {
			// Too much or none increment
			// log.Printf("Report: %v, Increment too big or too small: %d", report, incr)
			isSafe = false
			break
		}
		if !(math.EqualSign(incr, lastInc)) {
			// Change from increment or decrement
			// log.Printf("Report: %v, Change from increment or decrement: %d", report, incr)
			isSafe = false
			break
		}
		lastInc = incr

	}

	return isSafe
}

func Solution1() int {
	content, err := reader.ReadIntByLineFile("./inputs/day02.txt")
	log.Printf("content: %v", content)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	countSafeReports := 0
	for _, report := range content {
		if isReportSafe(report) {
			countSafeReports++
		}
	}
	return countSafeReports
}

func isReportSafeWithDampener(report []int) bool {
	// Create a copy of the slice without the element at index i
	reportCopy := make([]int, len(report))
	copy(reportCopy, report)
	for i := 0; i < len(report); i++ {
		tryReport := append(reportCopy[:i], reportCopy[i+1:]...)
		log.Printf("TryReport: %v", tryReport)
		if isReportSafe(tryReport) {
			return true
		}
		copy(reportCopy, report)
	}
	return false
}

func Solution2() int {
	content, err := reader.ReadIntByLineFile("./inputs/day02.txt")
	// log.Printf("content: %v", content)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	countSafeReports := 0
	for _, report := range content {
		if isReportSafeWithDampener(report) {
			countSafeReports++
		}
	}
	return countSafeReports
}
