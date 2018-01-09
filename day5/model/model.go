package model

type School struct {
	Name string
	Addr string
}

func NewSchool(name, addr string) *School {
	return &School{name, addr}
}

func (s *School) GetName() string {
	return s.Name
}

func (s *School) GetAddr() string {
	return s.Addr
}
