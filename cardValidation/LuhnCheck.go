package cardvalidation

func IsValid(number []int) bool {
	var oddSum, evenSum int

	for i, v := range number {
		if i%2 == 1 {
			oddSum += v
		} else {
			even := v * 2
			if (even) > 9 {
				evenSum = even%10 + even/10
			} else {
				evenSum += even
			}

		}
	}
	println(oddSum + evenSum)
	if (oddSum+evenSum)%10 == 0 {
		return true
	} else {
		return false
	}
}
