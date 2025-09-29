package encode

import (
    "unicode"
    "strings"
    "strconv"
)

func RunLengthEncode(input string) string {
    if len(input) == 0 {
        return ""
    }
    
    var res string
    count := 1
    var currentVal rune = rune(input[0])
    
    for i := 1; i < len(input); i++ {
        if currentVal != rune(input[i]) {
            if count == 1 {
                res += string(currentVal)
            } else {
                res += strconv.Itoa(count)
                res += string(currentVal)
            }
            currentVal = rune(input[i])
            count = 1
        } else {
            count++
        }
    }
    
    // Handle the last character
    if count == 1 {
        res += string(currentVal)
    } else {
        res += strconv.Itoa(count)
        res += string(currentVal)
    }
    
    return res
}

func RunLengthDecode(input string) string {
    result := ""
    count := ""
    
    for i := 0; i < len(input); i++ {
        if unicode.IsDigit(rune(input[i])) {
            count += string(input[i])
            continue
        }
        
        if count == "" {
            result += string(input[i])
        } else {
            num, _ := strconv.Atoi(count)
            result += strings.Repeat(string(input[i]), num)
            count = ""
        }
    }
    
    return result
}