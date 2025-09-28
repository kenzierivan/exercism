package wordcount

import (
    "strings"
    "unicode"
)


type Frequency map[string]int

func WordCount(phrase string) Frequency {
    freq := Frequency{}
	words := strings.FieldsFunc(phrase, func(r rune)bool {
        return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r!='\''
    })
    for _, word := range words {
        cleaned := strings.Trim(strings.ToLower(word), "'")

        if cleaned == "" {
            continue
        }
        freq[cleaned]++
        
    }
    return freq
}
