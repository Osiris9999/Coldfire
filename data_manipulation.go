package coldfire

import (
	"math/rand"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// RemoveFromSlice removes a string from a list of strings if it exists.
func RemoveFromSlice(slice []string, element string) []string {
	res := []string{}

	for _, e := range slice {
		if e != element {
			res = append(res, e)
		}
	}

	return res
}

// CreateWordList generates possible variations of each word in the wordlist.
func CreateWordlist(words []string) []string {
	wordlist := []string{}
	for _, w := range words {
		word := w
		first_to_upper := strings.ToUpper(string(word[0])) + string(word[1:])
		wordlist = append(wordlist, strings.ToUpper(word))
		wordlist = append(wordlist, Revert(word))
		wordlist = append(wordlist, first_to_upper)
		wordlist = append(wordlist, first_to_upper+"1")
		wordlist = append(wordlist, first_to_upper+"12")
		wordlist = append(wordlist, first_to_upper+"123")
		wordlist = append(wordlist, word+"1")
		wordlist = append(wordlist, word+"12")
		wordlist = append(wordlist, word+"123")
	}

	return wordlist
}

// RemoveStr removes a given string from a list of strings.
func RemoveStr(slice []string, s string) []string {
	final := []string{}
	for _, e := range slice {
		if e != s {
			final = append(final, e)
		}
	}

	return final
}

// RemoveInt removes a given integer from a list of integers.
func RemoveInt(slice []int, s int) []int {
	final := []int{}
	for _, e := range slice {
		if e != s {
			final = append(final, e)
		}
	}

	return final
}

// SplitJoin splits a string then joins them using given delimiters.
func SplitJoin(s, splittBy, joinBy string) string {
	splitted := strings.Split(s, splittBy)
	joined := strings.Join(splitted, joinBy)

	return joined
}

// RevertSlice reverses a slice type agnostically.
func RevertSlice(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)

	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func SplitMultiSep(s string, seps []string) []string {
	f := func(c rune) bool {
		for _, sep := range seps {
			if string(c) == sep {
				return true
			}
		}
		return false
	}
	fields := strings.FieldsFunc(s, f)
	return fields
}

func SplitChunks(s string, chunk int) []string {
	if chunk >= len(s) {
		return []string{s}
	}
	var chunks []string
	c := make([]rune, chunk)
	len := 0
	for _, r := range s {
		c[len] = r
		len++
		if len == chunk {
			chunks = append(chunks, string(c))
			len = 0
		}
	}
	if len > 0 {
		chunks = append(chunks, string(c[:len]))
	}
	return chunks
}

// ExtractIntFromString extracts a list of possible integers from a given string.
func ExtractIntFromString(s string) []int {
	res := []int{}
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	submatchall := re.FindAllString(s, -1)

	for _, element := range submatchall {
		res = append(res, StrToInt(element))
	}

	return res
}

// ShuffleSlice randomly shuffles a list of strings.
func ShuffleSlice(s []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	return s
}

// IpIncrement increments an IP address by 1.
func IpIncrement(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// RemoveDuplicatesStr returns an array of strings that are unique to each other.
func RemoveDuplicatesStr(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// RemoveDuplicatesInt returns an array of integers that are unique to each other.
func RemoveDuplicatesInt(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
