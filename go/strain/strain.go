package strain

// Ints is a slice of ints.
type Ints []int

// Lists is a slice of slices of ints.
type Lists [][]int

// Strings is a slice of strings.
type Strings []string

// Ints.Keep returns the subset of Ints for which the predicate returns true.
func (ints Ints) Keep(predicate func(int) bool) (filteredInts Ints) {
	if ints == nil {
		return nil
	}
	for _, i := range ints {
		if predicate(i) {
			filteredInts = append(filteredInts, i)
		}
	}
	return
}

// Ints.Discard returns the subset of Ints for which the predicate returns false.
func (ints Ints) Discard(predicate func(int) bool) Ints {
	reversePredicate := func(i int) bool { return !predicate(i) }
	return ints.Keep(reversePredicate)
}

// Lists.Keep returns the subset of Lists for which the predicate returns true.
func (lists Lists) Keep(predicate func([]int) bool) (filteredLists Lists) {
	if lists == nil {
		return nil
	}
	for _, l := range lists {
		if predicate(l) {
			filteredLists = append(filteredLists, l)
		}
	}
	return
}

// Strings.Keep returns the subset of Strings for which the predicate returns true.
func (strings Strings) Keep(predicate func(string) bool) (filteredStrings Strings) {
	if strings == nil {
		return nil
	}
	for _, s := range strings {
		if predicate(s) {
			filteredStrings = append(filteredStrings, s)
		}
	}
	return
}
