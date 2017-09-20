package matching

import (
	"encoding/json"
	"fmt"
	"github.com/nbutton23/zxcvbn-go/match"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

//DateSepMatch("1991-09-11jibjab11.9.1991")
//[{date 16 25  . 9 11 1991} {date 0 10  - 9 11 1991}]

func TestDateSepMatch(t *testing.T) {
	matches := dateSepMatchHelper("1991-09-11jibjab11.9.1991")

	assert.Len(t, matches, 2, "Length should be 2")

	for _, match := range matches {
		if match.Separator == "." {
			assert.Equal(t, 16, match.I)
			assert.Equal(t, 25, match.J)
			assert.Equal(t, int64(9), match.Day)
			assert.Equal(t, int64(11), match.Month)
			assert.Equal(t, int64(1991), match.Year)
		} else {
			assert.Equal(t, 0, match.I)
			assert.Equal(t, 10, match.J)
			assert.Equal(t, int64(9), match.Day)
			assert.Equal(t, int64(11), match.Month)
			assert.Equal(t, int64(1991), match.Year)
		}
	}

}

func TestRepeatMatch(t *testing.T) {
	//aaaBbBb
	matches := repeatMatch("aaabBbB")

	assert.Len(t, matches, 2, "Lenght should be 2")

	for _, match := range matches {
		if strings.ToLower(match.DictionaryName) == "b" {
			assert.Equal(t, 3, match.I)
			assert.Equal(t, 6, match.J)
			assert.Equal(t, "bBbB", match.Token)
			assert.NotZero(t, match.Entropy, "Entropy should be set")
		} else {
			assert.Equal(t, 0, match.I)
			assert.Equal(t, 2, match.J)
			assert.Equal(t, "aaa", match.Token)
			assert.NotZero(t, match.Entropy, "Entropy should be set")

		}
	}
}

func TestSequenceMatch(t *testing.T) {
	//abcdjibjacLMNOPjibjac1234  => abcd LMNOP 1234

	matches := sequenceMatch("abcdjibjacLMNOPjibjac1234")
	assert.Len(t, matches, 3, "Lenght should be 2")

	for _, match := range matches {
		if match.DictionaryName == "lower" {
			assert.Equal(t, 0, match.I)
			assert.Equal(t, 3, match.J)
			assert.Equal(t, "abcd", match.Token)
			assert.NotZero(t, match.Entropy, "Entropy should be set")
		} else if match.DictionaryName == "upper" {
			assert.Equal(t, 10, match.I)
			assert.Equal(t, 14, match.J)
			assert.Equal(t, "LMNOP", match.Token)
			assert.NotZero(t, match.Entropy, "Entropy should be set")
		} else if match.DictionaryName == "digits" {
			assert.Equal(t, 21, match.I)
			assert.Equal(t, 24, match.J)
			assert.Equal(t, "1234", match.Token)
			assert.NotZero(t, match.Entropy, "Entropy should be set")
		} else {
			assert.True(t, false, "Unknow dictionary")
		}
	}
}

func TestSpatialMatchQwerty(t *testing.T) {
	matches := spatialMatch("qwerty")
	assert.Len(t, matches, 1, "Lenght should be 1")
	assert.NotZero(t, matches[0].Entropy, "Entropy should be set")

	matches = spatialMatch("asdf")
	assert.Len(t, matches, 1, "Lenght should be 1")
	assert.NotZero(t, matches[0].Entropy, "Entropy should be set")

}

func TestSpatialMatchDvorak(t *testing.T) {
	matches := spatialMatch("aoeuidhtns")
	assert.Len(t, matches, 1, "Lenght should be 1")
	assert.NotZero(t, matches[0].Entropy, "Entropy should be set")

}

func TestDictionaryMatch(t *testing.T) {
	var matches []match.Match
	for _, dicMatcher := range DICTIONARY_MATCHERS {
		matchesTemp := dicMatcher("first")
		matches = append(matches, matchesTemp...)
	}

	assert.Len(t, matches, 4, "Lenght should be 4")
	for _, match := range matches {
		assert.NotZero(t, match.Entropy, "Entropy should be set")

	}

}

func TestDateWithoutSepMatch(t *testing.T) {
	matches := dateWithoutSepMatch("11091991")
	assert.Len(t, matches, 1, "Lenght should be 1")

	matches = dateWithoutSepMatch("20010911")
	assert.Len(t, matches, 1, "Lenght should be 1")
	log.Println(matches)

	//matches := dateWithoutSepMatch("110991")
	//assert.Len(t, matches, 21, "Lenght should be blarg")
}

//l33t
func TestLeetSubTable(t *testing.T) {
	subs := relevantL33tSubtable("password")
	assert.Len(t, subs, 0, "password should produce no leet subs")

	subs = relevantL33tSubtable("p4ssw0rd")
	assert.Len(t, subs, 2, "p4ssw0rd should produce 2 subs")

	subs = relevantL33tSubtable("1eet")
	assert.Len(t, subs, 2, "1eet should produce 2 subs")
	assert.Equal(t, subs["i"][0], "1")
	assert.Equal(t, subs["l"][0], "1")

	subs = relevantL33tSubtable("4pple@pple")
	assert.Len(t, subs, 1, "4pple@pple should produce 1 subs")
	assert.Len(t, subs["a"], 2)

}

func TestPermutationsOfLeetSubstitutions(t *testing.T) {
	password := "p4ssw0rd" //[passw0rd, password, p4ssword]
	possibleSubs := relevantL33tSubtable(password)

	permutations := getAllPermutationsOfLeetSubstitutions(password, possibleSubs)

	assert.Len(t, permutations, 3, "There should be 3 permutations for "+password)

	password = "p4$sw0rd" //[pa$sw0rd, passw0rd, password, pa$sword, p4ssw0rd, p4ssword, p4$sword]
	possibleSubs = relevantL33tSubtable(password)

	permutations = getAllPermutationsOfLeetSubstitutions(password, possibleSubs)
	assert.Len(t, permutations, 7, "There should be 7 (? check my math) permutations for "+password)

	password = "p4$$w0rd" //[pa$sw0rd, passw0rd, password, pa$sword, p4ssw0rd, p4ssword, p4$sword]
	possibleSubs = relevantL33tSubtable(password)

	permutations = getAllPermutationsOfLeetSubstitutions(password, possibleSubs)
	assert.Len(t, permutations, 15, "Check my math 2*2*2*2 - 1 "+password)

	password = "1337"
	possibleSubs = relevantL33tSubtable(password)
	permutations = getAllPermutationsOfLeetSubstitutions(password, possibleSubs)
	assert.Len(t, permutations, 35, "check my math 3*2*2*3 -1 ")
}

func TestLeet(t *testing.T) {
	password := "1337"
	matches := l33tMatch(password)
	bytes, _ := json.Marshal(matches)
	fmt.Println(string(bytes))

	fmt.Println(matches[0].J)
}
