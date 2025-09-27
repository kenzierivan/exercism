package lsproduct

import "errors"

func LargestSeriesProduct(digits string, span int) (int64, error) {
    if span < 0 {
        return 0, errors.New("span is negative yo")
    }
    if span > len(digits) {
        return 0, errors.New("span longer than digits")
    }
	var largestDigit int64 = 0
	for i := 0; i < len(digits)-(span-1); i++ {
        result := 1
        digit := digits[i:i+span]
        for j := 0; j < span; j++ {
            num := int(digit[j]-48)
            if num < 0 || num > 9 { 
        		return 0, errors.New("invalid character")
    		}
            result *= num
        }
        if int64(result) > largestDigit {
            largestDigit = int64(result)
        }
    }
    

    return largestDigit, nil
}
