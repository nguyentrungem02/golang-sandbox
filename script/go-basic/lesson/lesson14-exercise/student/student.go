package student

import "fmt"

type Student struct {
	Id     int
	Name   string
	Class  string
	Scores []float64
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("Id: %d | Ho ten: %s | Lop: %s | Diem TB: %.2f", s.Id, s.Name, s.Class, s.CalculateAverageScore())
}

func (s Student) CalculateAverageScore() float64 {
	if len(s.Scores) == 0 {
		return 0.0
	}

	total := 0.0
	for _, score := range s.Scores {
		total += score
	}

	return total / float64(len(s.Scores))
}

func (s Student) GetId() int {
	return s.Id
}
