package district

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	record := District{
		Title:      "县",
		Code:       "1004",
		ParentCode: "1003",
		Label:      "town",
	}
	err := Add(record)
	if err != nil {
		panic(err)
	}
}

func TestGetByProviceCodeWithChildren(t *testing.T) {
	code := "1001"
	out, err := GetByCodeWithChildren(code)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

func TestGetParent(t *testing.T) {
	code := "1004"
	out, err := GetParent(code)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
