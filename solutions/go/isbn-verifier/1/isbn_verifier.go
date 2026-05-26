package isbnverifier

import (
    "regexp"
    "strings"
)

func IsValidISBN(isbn string) bool {

    var re = regexp.MustCompile("^[0-9-]*X?$")
    if (!re.MatchString(isbn)) {
        return false
    }
    var sanitized = strings.ReplaceAll(isbn, "-", "")

    if (len(sanitized) != 10) {
        return false
    }
    var pos = 10
    var sum = 0
    
    for _, t := range sanitized {
        if (t == 'X') {
    		sum += (10 * pos)
        } else {
	        sum += (int(t - '0') * pos)
        }
        pos--
    }

    return sum % 11 == 0
}
