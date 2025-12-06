package dto

type ProblemCreate struct {
	Category    string
	Difficult   string
	Title       string
	Content     string
	MemoryLimit int
	TimeLimit   int
	TestCase    []TestCaseData
}
