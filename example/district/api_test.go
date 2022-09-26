package district

import "testing"

func TestAdd(t *testing.T) {
	record := DistrictKV{
		Title:    "区域",
		NodeID:   "1003",
		ParentID: "1002",
		Label:    "area",
	}
	err := Add(record)
	if err != nil {
		panic(err)
	}
}
