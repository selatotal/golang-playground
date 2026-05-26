package collatzconjecture
import (
    "errors";
    "fmt";
)
func CollatzConjecture(n int) (int, error) {
    fmt.Printf("Debug message %d\n", n)

    if (n <= 0) {
        return 0, errors.New("Invalid number")
    }
    if (n == 1) {
        return 0, nil
    }
    if (n % 2 == 0) {
        var totalSteps, err = CollatzConjecture(n/2)
        return totalSteps + 1, err
    } else {
        var totalSteps, err = CollatzConjecture((n * 3) + 1)
        return totalSteps + 1, err
    }
}
