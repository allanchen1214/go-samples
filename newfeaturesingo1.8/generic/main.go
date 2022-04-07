package main

import (
	"fmt"
)

func SumInts(m map[string]int32) int32 {
	var sum int32
	sum = 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumFloat(m map[string]float32) float32 {
	var sum float32
	sum = 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumNumber1[K comparable, V int32 | float32](m map[K]V) V {
	var sum V
	for _, item := range m {
		sum += item
	}
	return sum
}

type Number interface {
	int32 | float32
}

func SumNumber2[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, item := range m {
		sum += item
	}
	return sum
}

func main() {
	imap := map[string]int32{
		"first":  100,
		"second": 200,
	}
	fmt.Println(SumInts(imap))

	fmap := map[string]float32{
		"first":  300.2,
		"second": 400.2,
	}
	fmt.Println(SumFloat(fmap))

	fmt.Println(SumNumber1(imap))
	fmt.Println(SumNumber1(fmap))

	fmt.Println(SumNumber2(imap))
	fmt.Println(SumNumber2(fmap))
}
