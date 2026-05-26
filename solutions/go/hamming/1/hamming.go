package hamming
import "errors"

func Distance(a, b string) (int, error) {
    aLen := len(a)
    bLen := len(b)
    if (aLen != bLen) {
		return 0, errors.New("different lengths")
    }
    hamming := 0
    for pos := 0; pos < aLen; pos++ {
        if a[pos] != b[pos] {
            hamming++
        }
    }
    return hamming, nil
}
