package students

type Student struct {
	name  string
	age   int
	grade int
}

func (s *Student) Get() string {
	return s.name
}

func (s *Student) Put(name string, age int, grade int) {
	s.name = name
	s.age = age
	s.grade = grade
}
