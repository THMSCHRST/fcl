package render

import "fmt"

/*
layout := []render.Attribute{
	{Index: 0, Size: 3, Offset: 0},
	{Index: 1, Size: 3, Offset: 12},
}*/

func NewLayout(IdxCount int, IdxLengths []int) ([]Attribute, error) {
	if IdxCount != len(IdxLengths) {return nil, fmt.Errorf("Index count and index lengths do not match")}
	layout := make([]Attribute, IdxCount)
	offsetCounter := 0
	for i := range IdxCount {
		if i == 0 {
			layout[i] = Attribute{Index: uint32(i), Size: int32(IdxLengths[i]), Offset: 0}
		} else {
			layout[i] = Attribute{Index: uint32(i), Size: int32(IdxLengths[i]), Offset: offsetCounter*4} //4 = 4 bytes per float in ram
		}
		offsetCounter += IdxLengths[i]
	}
	return layout, nil
}