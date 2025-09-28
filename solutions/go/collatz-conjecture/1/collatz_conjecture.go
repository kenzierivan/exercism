package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n < 1 {
        return 0, errors.New("number is negative")
    }
	steps := 0
    num := n
    for {
        if num == 1 {
            return steps, nil
        } else if num % 2 == 0 {
            num = num / 2
            steps++
        } else if num % 2 != 0 {
            num = (num*3)+1
            steps++
        }
    }
}
