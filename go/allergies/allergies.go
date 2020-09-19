// Exercism exercise allergies.
package allergies

var numToAllergy = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

const maxNum = 128

// Given a `score` determine whether or not the person is allergic to `substance`.
func AllergicTo(score uint, substance string) bool {
	for num, allergy := range numToAllergy {
		if substance == allergy {
			return score&num != 0
		}
	}
	return false
}

// Given a `score`, return the list of allergies of the person.
func Allergies(score uint) []string {
	// The test cases require the allergies to be returned in order of increasing numerical encoded
	// value. This is not specified in the requirements.
	allergies := []string{}
	for num := uint(1); num <= maxNum; num *= 2 {
		if score&num != 0 {
			allergies = append(allergies, numToAllergy[num])
		}
	}
	return allergies
}
