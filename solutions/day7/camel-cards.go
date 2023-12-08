package day7

import (
	"AdventOfCode2023/solutions/utils"
	"sort"
)

func CamelCards1() int {
	return Solve("B", func(handStr utils.String) HandTypeEnum {
		return ParseHandType(handStr)
	})
}

func CamelCards2() int {
	return Solve("Z", func(handStr utils.String) HandTypeEnum {
		return ParseHandTypeWithJoker(handStr)
	})
}

func Solve(changeJokersTo string, handTypeParser func(handStr utils.String) HandTypeEnum) int {
	input := utils.NewProblem("solutions/day7/input.txt").InputAsLines()
	hands := make(Hands, input.Len())

	for index, elem := range input {
		handAndBid := elem.Splitter(" ")
		handStr := handAndBid[0].ReplaceAll("J", changeJokersTo)
		hand := NewHand(handStr, handAndBid[1].ToIntOrDefault(-1), handTypeParser(handStr))
		hands[index] = hand
	}

	sort.Sort(hands)
	result := 0

	for index, hand := range hands {
		result += (index + 1) * hand.Bid
	}

	return result
}

func getHandType(distinctCount int, maxi int) HandTypeEnum {
	if distinctCount == 5 {
		return HighCard
	}

	if distinctCount == 1 {
		return FiveOfAKind
	}

	if distinctCount == 2 {
		if maxi == 4 {
			return FourOfAKind
		}

		if maxi == 3 {
			return FullHouse
		}
	}

	if distinctCount == 3 {
		if maxi == 3 {
			return ThreeOfAKind
		}

		return TwoPair
	}

	return OnePair
}

func ParseHandTypeWithJoker(input utils.String) HandTypeEnum {
	charCount := make(map[rune]int)
	maxi := 0
	var maxChar rune
	for _, char := range input {
		charCount[char]++
	}

	distinctCount := len(charCount)

	for val := range charCount {
		if charCount[val] > maxi && val != 'Z' {
			maxi = charCount[val]
			maxChar = val
		}
	}

	if maxi == 0 {
		return FiveOfAKind
	}

	if charCount['Z'] > 0 {
		charCount[maxChar] += charCount['Z']
		maxi = charCount[maxChar]
		distinctCount--
		delete(charCount, 'Z')
	}

	return getHandType(distinctCount, maxi)
}

func ParseHandType(input utils.String) HandTypeEnum {
	charCount := make(map[rune]int)
	maxi := 0
	for _, char := range input {
		charCount[char]++
	}

	distinctCount := len(charCount)
	for val := range charCount {
		if charCount[val] > maxi {
			maxi = charCount[val]
		}
	}

	return getHandType(distinctCount, maxi)
}
