package day7

import "AdventOfCode2023/solutions/utils"

type HandTypeEnum int

// Hand: ZZZZZ 1
const (
	HighCard HandTypeEnum = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	Cards    utils.String
	HandType HandTypeEnum
	Bid      int
}

type Hands []Hand

func (h Hands) Len() int {
	return len(h)
}

var cardLabels = utils.StringArray{"A", "K", "Q", "B", "T", "9", "8", "7", "6", "5", "4", "3", "2", "Z"}

func (h Hands) Less(i, j int) bool {
	if h[i].HandType == h[j].HandType {

		for it := 0; it < len(h[i].Cards); it++ {
			scoreI := cardLabels.IndexOf(utils.String(h[i].Cards[it]))
			scoreJ := cardLabels.IndexOf(utils.String(h[j].Cards[it]))

			if scoreI != scoreJ {
				return scoreI > scoreJ
			}
		}
	}

	return h[i].HandType < h[j].HandType
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func NewHand(input utils.String, bid int, handType HandTypeEnum) Hand {
	return Hand{Cards: input, HandType: handType, Bid: bid}
}
