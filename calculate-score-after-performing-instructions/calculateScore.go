package main

import "fmt"

func calculateScore(instructions []string, values []int) int64 {
	var score int64
	visited := make([]bool, len(instructions))
	i := 0
	for i < len(instructions) && i >= 0 && !visited[i] {
		visited[i] = true
		switch instructions[i] {
		case "add":
			score = score + int64(values[i])
			i += 1
		case "jump":
			i += values[i]
		default:
			return score
		}
	}

	return score
}

func main() {
	fmt.Println(calculateScore([]string{"jump", "add", "add", "jump", "add", "jump"}, []int{2, 1, 3, 1, -2, -3}))
}
