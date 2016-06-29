package intset

// return the number of elements
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= (1 << bit)
	}
}

// Clear removes all elements from the set.
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// Copy returns a copy of the set.
func (s *IntSet) Copy() *IntSet {
	var n IntSet
	n.words = make([]uint64, len(s.words))
	copy(n.words, s.words)
	return &n
}
