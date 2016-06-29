package intset

func (s *IntSet) AddAll(x ...int) {
	for _, y := range x {
		s.Add(y)
	}
}
