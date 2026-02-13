package atbash

import (
    "unicode"
    "strings"
)

func Atbash(s string) string {
    cleaned := strings.ToLower(s)
    res := ""
	counter := 0
    for i := 0; i < len(cleaned); i++ {
        if counter == 5 {
            res += " "
            counter = 0
        }
        if unicode.IsLetter(rune(cleaned[i])) {
            newChar := 'a' + 'z' - rune(cleaned[i])	
            res += string(newChar)
            counter++
        } else if unicode.IsDigit(rune(cleaned[i])) {
            res += string(cleaned[i])
            counter++
        }
        
    }
    return strings.Trim(res, " ")
}
