package util

var exists = struct{}{}

type set struct {
	m map[string]struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *set) List() (list []string) {
	for k := range s.m {
		list = append(list, k)
	}
	return
}

func (s *set) AddAll(a []string) {
	for _, i := range a {
		s.Add(i)
	}
}

func (s *set) Len() int {
	return len(s.m)
}
