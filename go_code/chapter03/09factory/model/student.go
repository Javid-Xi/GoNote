package model

type student struct {
	name  string
	score float64
}

// student是私有的  因此只能在吗model包使用
// 可通过工厂模式解决
func NewStudent(name string, score float64) *student {
	return &student{name, score}
}

func (s *student) GetName() string {
	return s.name
}

func (s *student) GetScore() float64 {
	return s.score
}
