package dto

type ProblemResponse struct {
	Id           uint
	Category     string
	Difficult    string
	Title        string
	Content      string
	IsDeleted    bool
	IsDailyToday bool
	PointDaily   int
	MemoryLimit  int
	TimeLimit    int
}
