package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
    result := map[string]int {}
	for pointValue, letters := range in {
        for _, letter := range letters {
            result[strings.ToLower(letter)] = pointValue
        }
    }
    return result
}
