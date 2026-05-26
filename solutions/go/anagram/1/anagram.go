package anagram

import (
    "sort"
    "strings"
    )

func sortString(source string) string {
    r := []rune(source)
    sort.Slice(r, func(i, j int) bool {
        return r[i] < r[j]
    })
    return string(r)
} 

func Detect(subject string, candidates []string) []string {

    normalizedSubject := strings.ToLower(strings.Trim(subject, ""))
    source := sortString(normalizedSubject)
	var result []string
	
    for _, value := range candidates {
	    normalizedCandidate := strings.ToLower(strings.Trim(value, ""))
        candidate := sortString(normalizedCandidate)
        if (normalizedSubject != normalizedCandidate && candidate == source) {
            result = append(result, value)
        }
    }

    return result
}
