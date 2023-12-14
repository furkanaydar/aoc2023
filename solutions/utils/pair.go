package utils

type Pair struct {
	E1 int
	E2 int
}

type PairArray []Pair

func (p PairArray) Less(i, j int) bool {
	if p[i].E1 == p[j].E1 {
		return p[i].E2 < p[j].E2
	}

	return p[i].E1 < p[j].E1
}

func (p PairArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairArray) Len() int {
	return len(p)
}
