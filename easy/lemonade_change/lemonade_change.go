package lemonadechange

func lemonadeChange(bills []int) bool {
	fives, tens := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			fives++
		case 10:
			if fives > 0 {
				fives--
				tens++
			} else {
				return false
			}
		default:
			if tens > 0 && fives > 0 {
				tens--
				fives--
			} else if fives >= 3 {
				fives -= 3
			} else {
				return false
			}
		}
	}

	return true
}
