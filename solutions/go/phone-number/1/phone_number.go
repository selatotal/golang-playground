package phonenumber

import (
    "errors"
    "fmt"
    "regexp"
    "strconv"
    "strings"
)

func getNumbers(phoneNumber string) (string, error) {
    re := regexp.MustCompile(`\D`)
    result := re.ReplaceAllString(phoneNumber, "")
    if len(result) == 11 && result[0] == '1' && result[1] == '1' {
        return "", errors.New("Invalid area code")
    }
    result = strings.TrimPrefix(result, "1")
    if len(result) != 10{
        return "", errors.New("Invalid number")
    }
    if result[3] == '0' || result[3] == '1'{
        return "", errors.New("Invalid exchange code")
    }
    return result, nil
}

func Number(phoneNumber string) (string, error) {
    number, err := getNumbers(phoneNumber)
    if (err != nil) {
        return "", err
    }
	_, errArea := AreaCode(phoneNumber)    
    if (errArea != nil) {
        return "", errArea
    }
    return number, nil
}

func AreaCode(phoneNumber string) (string, error) {
    number, err := getNumbers(phoneNumber)
    if (err != nil) {
        return "", err
    }
    areaCode := number[:3]
    if n, _ := strconv.Atoi(areaCode); n > 1000 || n < 100 {
        return "", errors.New("Invalid area code")
    }
    return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
    number, err := Number(phoneNumber)
    if (err != nil) {
        return "", err
    }
    return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
