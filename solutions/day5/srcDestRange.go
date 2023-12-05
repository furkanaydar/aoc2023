package day5

type SrcDestRange struct {
	DestStart int
	SrcRange  Range
}

type Range struct {
	Start int
	End   int
}

type RangeCompareResult int

const (
	X1andX2BeforeRange RangeCompareResult = iota
	X2InRange
	InRange
	X1InRange
	X1AfterRange
	X1BeforeAndX2AfterRange
)

func (_range Range) compareRange(input Range) RangeCompareResult {
	if input.End < _range.Start {
		return X1andX2BeforeRange
	}

	if input.Start < _range.Start && input.End >= _range.Start {
		return X2InRange
	}

	if input.Start >= _range.Start && input.End <= _range.End {
		return InRange
	}

	if input.Start <= _range.End && input.End > _range.End {
		return X1InRange
	}

	if input.Start > _range.End {
		return X1AfterRange
	}

	return X1BeforeAndX2AfterRange
}

func (_range Range) splitRange(input Range) []Range {
	compareResult := _range.compareRange(input)

	if compareResult == X2InRange {
		return []Range{{input.Start, _range.Start - 1}, {_range.Start, input.End}}
	}

	if compareResult == X1InRange {
		return []Range{{input.Start, _range.End}, {_range.End + 1, input.End}}
	}

	return []Range{input}
}

type SrcDestRangeGroup []SrcDestRange

func (arr SrcDestRangeGroup) Len() int { return len(arr) }

func (arr SrcDestRangeGroup) Less(i1, i2 int) bool {
	if arr[i1].SrcRange.Start == arr[i2].SrcRange.End {
		return arr[i1].DestStart < arr[i2].DestStart
	}

	return arr[i1].SrcRange.Start < arr[i2].SrcRange.Start
}

func (arr SrcDestRangeGroup) Swap(i1, i2 int) {
	arr[i1], arr[i2] = arr[i2], arr[i1]
}

type RangeGroup []Range

func (arr RangeGroup) Swap(i1, i2 int) {
	arr[i1], arr[i2] = arr[i2], arr[i1]
}

func (arr RangeGroup) Len() int { return len(arr) }

func (arr RangeGroup) Less(i1, i2 int) bool {
	if arr[i1].Start == arr[i2].Start {
		return arr[i1].End < arr[i2].End
	}

	return arr[i1].Start < arr[i2].Start
}
