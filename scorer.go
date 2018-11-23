package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Running")
}

// Score receives game array and returns a score as int
func Score(score []string) (total int) {

	// var total int
	for i := 0; i < len(score); i++ {
		if i >= 10 {
			return
		}
		val := score[i]

		if len(val) == 2 && string(val[1]) == "/" {
			var nxt int
			if len(score)-(i+1) >= 1 {
				nxt = scr(score[i+1][0])
			}
			total += 10 + nxt
			continue
		}

		if string(val[0]) == "X" {
			total += getFrameScore(val, false)
			rem := len(score) - (i + 1)
			if rem >= 1 {
				var isBonus bool
				if i+2 >= 10 {
					isBonus = true
				}
				if score[i+1] == "X" {
					total += getFrameScore(score[i+1], false)
					total += getFrameScore(score[i+2], isBonus)
				} else {
					total += getFrameScore(score[i+1], false)
				}
			}
			continue
		}

		total += getFrameScore(val, false)
	}
	return
}

func scr(n byte) int {
	if n == '-' {
		return 0
	}
	val, _ := strconv.Atoi(string(n))
	return val
}

func getFrameScore(frame string, isBonus bool) int {

	if frame == "X" {
		return 10
	}

	if len(frame) == 2 && frame[1] == '/' {
		return 10
	}

	first := scr(frame[0])
	var second int
	if len(frame) == 2 && !isBonus {
		second = scr(frame[1])
	}
	return first + second
}
