package baseballgame

import "strconv"

func calPoints(ops []string) int {
	records := make([]int, 0, len(ops))
	for _, o := range ops {
		l := len(records) - 1
		switch o {
		case "C":
			records = records[:l]
		case "D":
			records = append(records, records[l]*2)
		case "+":
			records = append(records, records[l]+records[l-1])
		default:
			r, _ := strconv.Atoi(o)
			records = append(records, r)
		}
	}

	sum := 0
	for _, r := range records {
		sum += r
	}

	return sum
}
