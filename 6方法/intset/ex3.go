package intset

// Set s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Set s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	temp := s.Copy()
	temp.IntersectWith(t)
	for i := range s.words {
		s.words[i] ^= temp.words[i]
	}
}

// Set s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
