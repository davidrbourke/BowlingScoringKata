package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Running")
}

// Score receives game array and returns a score as int
func Score(game []string) (total int) {
	for i := 0; i < len(game) && i < 10; i++ {
		frame := game[i]
		if strings.Contains(frame, "/") {
			var nxt int
			if hasNextFrame(game, i) {
				nxt = bowlToInt(game[i+1][0])
			}
			total += 10 + nxt
			continue
		}
		if string(frame[0]) == "X" {
			total += getFrameScore(frame, false)
			if hasNextFrame(game, i) {
				var isBonus bool
				if i+2 >= 10 {
					isBonus = true
				}
				if game[i+1] == "X" {
					total += getFrameScore(game[i+1], false)
					total += getFrameScore(game[i+2], isBonus)
				} else {
					total += getFrameScore(game[i+1], false)
				}
			}
			continue
		}
		total += getFrameScore(frame, false)
	}
	return
}

func bowlToInt(n byte) int {
	if n == '-' {
		return 0
	}
	// val, _ := strconv.Atoi(string(n))
	// return val
	return int(n - '0')
}

func hasNextFrame(game []string, currentFrame int) bool {
	return len(game)-(currentFrame+1) >= 1
}

func getFrameScore(frame string, isBonus bool) int {

	if frame == "X" {
		return 10
	}

	if strings.Contains(frame, "/") {
		return 10
	}

	first := bowlToInt(frame[0])
	var second int
	if len(frame) == 2 && !isBonus {
		second = bowlToInt(frame[1])
	}
	return first + second
}
