package main

type Scrambler struct {
	numList []int64
}

func New(seed int64) *Scrambler {
	numList := Shuffle[int64](CreateNumList(0, 26))
	s := &Scrambler{numList: numList}
	return s
}
