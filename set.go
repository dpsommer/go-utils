package goutils

type Set[K comparable] map[K]struct{}

// Add adds a key k to the set
func (s Set[K]) Add(k K) {
	s[k] = struct{}{}
}

// Update returns the set with all keys from others added
func (s Set[K]) Update(others ...Set[K]) Set[K] {
	for _, set := range others {
		for k, v := range set {
			s[k] = v
		}
	}
	return s
}

// Union returns a new set containing all values in the set and others
func (s Set[K]) Union(others ...Set[K]) Set[K] {
	n := s.Copy()
	return n.Update(others...)
}

// Intersection returns a new set containing values that are in both the set
// and all others
func (s Set[K]) Intersection(others ...Set[K]) Set[K] {
	n := Set[K]{}

outer:
	for k, v := range s {
		for _, set := range others {
			if _, ok := set[k]; !ok {
				continue outer
			}
		}
		n[k] = v
	}

	return n
}

// Difference returns a new set containing values that are not in others
func (s Set[K]) Difference(others ...Set[K]) Set[K] {
	n := Set[K]{}

outer:
	for k, v := range s {
		for _, set := range others {
			if _, ok := set[k]; ok {
				continue outer
			}
		}
		n[k] = v
	}

	return n
}

// SymmetricDifference returns a new set containing values that are in the
// original set or other but not both
func (s Set[K]) SymmetricDifference(other Set[K]) Set[K] {
	n := Set[K]{}

	for k, v := range s {
		if _, ok := other[k]; !ok {
			n[k] = v
		}
	}

	for k, v := range other {
		if _, ok := s[k]; !ok {
			n[k] = v
		}
	}

	return n
}

// Copy returns a shallow copy of the set
func (s Set[K]) Copy() Set[K] {
	n := Set[K]{}

	for k, v := range s {
		n[k] = v
	}

	return n
}
