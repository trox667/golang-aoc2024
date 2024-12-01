package days

type Day interface {
	ReadSample(dayIndex int8) []string
	ReadInput(dayIndex int8) []string
	Part1(input []string) (result string, err error)
	Part2(input []string) (result string, err error)
}
