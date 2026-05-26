package clock

import (
    "fmt"
)

// Define the Clock type here.
type Clock struct {
    h int
    m int
}

func New(h, m int) Clock {
    min := h * 60 + m
    return Clock{0, 0}.Add(min)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (c Clock) Add(m int) Clock {
    fmt.Println(m)
    c.m = c.m + m % 60
	c.h = c.h + m / 60 % 24
	fmt.Println(c)
    if (c.m < 0) {
        c.m = 60 - c.m * -1
        c.h--
    } else if (c.m >= 60) {
        c.m = 60 - c.m
        c.h++
    }
    fmt.Println(c)
    if (c.h < 0) {
        c.h = 24 - c.h * -1
    } else if (c.h >= 24) {
        c.h = c.h % 24
    }
    fmt.Println(c)
    c.h = abs(c.h)
    c.m = abs(c.m)
    return c
}

func (c Clock) Subtract(m int) Clock {
    return c.Add(m * -1)
}

func (c Clock) String() string {
    return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
