package utils

import (
	"fmt"
	"time"
)

func PrintProgressBar(completed, total int) {
	percentage := (completed * 100) / total
	barLength := 50
	completedBars := (completed * barLength) / total
	remainingBars := barLength - completedBars

	fmt.Printf("\r[%s%s] %d%% Complete",
		string(Replicate('=', completedBars)),
		string(Replicate(' ', remainingBars)),
		percentage)

	if completed == total {
		fmt.Println()
	}
}

func Replicate(char rune, count int) []rune {
	result := make([]rune, count)
	for i := range result {
		result[i] = char
	}
	return result
}

func StartSpinner(done chan bool) {
	chars := []rune{'|', '/', '-', '\\'}
	for {
		select {
		case <-done:
			return
		default:
			for _, char := range chars {
				fmt.Printf("\r%c Loading...", char)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
}
