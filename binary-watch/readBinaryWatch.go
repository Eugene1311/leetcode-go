package binary_watch

import "fmt"

var turnedOnToHours = map[int][]string{
	0: {"0"},
	1: {"1", "2", "4", "8"},
	2: {"3", "5", "9", "6", "10"},
	3: {"7", "11"},
}

var turnedOnToMinutes = map[int][]string{
	0: {"00"},
	1: {"01", "02", "04", "08", "16", "32"},
	2: {"03", "05", "09", "17", "33", "06", "10", "18", "34", "12", "20", "36", "24", "40", "48"},
	3: {"07", "11", "19", "35", "13", "21", "37", "25", "41", "49", "14", "22", "38", "26", "42", "50", "28", "44", "52", "56"},
	4: {"15", "23", "39", "27", "43", "51", "29", "45", "57", "30", "46", "54", "58", "53"},
	5: {"31", "47", "55", "59"},
}

func readBinaryWatch(turnedOn int) []string {
	if turnedOn > 8 {
		return make([]string, 0)
	}

	result := make([]string, 0)

	// turnedOnHours - max 3 LEDs is on
	// minutes - max 5 LEDs is on
	for turnedOnHours := 0; turnedOnHours <= turnedOn; turnedOnHours++ {
		turnedOnMinutes := turnedOn - turnedOnHours
		hours := turnedOnToHours[turnedOnHours]
		minutes := turnedOnToMinutes[turnedOnMinutes]
		for _, hour := range hours {
			for _, minute := range minutes {
				result = append(result, hour+":"+minute)
			}
		}
	}

	return result
}

func main() {
	fmt.Println(readBinaryWatch(8))
}
