package main

import (
	"math/rand"
	"time"
)

func CreateNumList(init int64, last int64) []int64 {
	var list []int64
	for i := init; i < last; i++ {
		list = append(list, i)
	}
	return list
}

func Shuffle[T any](arr []T) []T {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}
