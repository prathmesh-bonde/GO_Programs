package main

import (
	"fmt"
	"log"
)

// exercise 4

type sqrtErr struct {
	lat  string
	long string
	err  error
}

func (sq sqrtErr) Error() string {
	return fmt.Sprintf("math error: %v %v %v", sq.lat, sq.long, sq.err)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("more coffee needed - value was %v", f)
		return 0, sqrtErr{"50.2289 N", "99.4656 W", e}
	}

	return 42, nil
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}
