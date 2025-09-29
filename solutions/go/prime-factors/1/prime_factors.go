package prime

func Factors(n int64) []int64 {
    factors := []int64 {}
    num := n
    var divisor int64 = 2

    if n <=1 {
        return factors 
    }
	for num > 1 {
        if num % divisor == 0 {
            factors = append(factors, divisor)
            num = num / divisor
        } else {
            divisor++
        }
    }
    
    return factors
    
}
