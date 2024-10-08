package unittest

import (
	"fmt"
	"testing"
)

// func TestSuma(t *testing.T) {
// 	total := Suma(2, 2)
// 	if total != 4 {
// 		t.Errorf("Expected 4, got %d", total)
// 	}
// }

// func TestSuma(t *testing.T) {
// 	tabla := []struct {
// 		a int
// 		b int
// 		c int
// 	}{
// 		{1, 2, 3},
// 		{2, 2, 4},
// 		{3, 3, 6},
// 	}

// 	for _, item := range tabla {
// 		total := Suma(item.a, item.b)

// 		if total != item.c {
// 			t.Errorf("Expected %d, got %d", item.c, total)
// 		} else {
// 			t.Logf("Correct sum %d", total)
// 		}
// 	}
// }

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 2},
		{7, 2, 7},
		{3, 10, 10},
	}
	for _, item := range tabla {
		max := GetMax(item.a, item.b)
		if max != item.c {
			t.Errorf("Expected %d, got %d", item.c, max)
		} else {
			t.Logf("Correct max %d", max)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tabla := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 7778742049},
	}

	for _, item := range tabla {
		max := Fibonacci(item.n)
		fmt.Println(max)
		if max != item.r {
			t.Errorf("Expected %d, got %d", item.r, max)
		}
	}
}
