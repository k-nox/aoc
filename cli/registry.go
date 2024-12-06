package cli

type Part func(useSample bool) int

type Day struct {
	PartOne Part
	PartTwo Part
}

type Registry map[int]Day
