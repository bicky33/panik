package panik_test

import (
	"fmt"
	"testing"

	"github.com/bicky33/panik"
)

func TestInvalidNIK(t *testing.T) {
	test := []panik.NIK{
		{NIK: "1234567890123456"},
		{NIK: "123456789012345"},
		{NIK: "123456789012345a"},
		{NIK: "loremipsumdolorsitamet"},
		{NIK: "@#$%^&*()_+{}[]|:;?/>.<,~`"},
		{NIK: "0000000000000000"},
		{NIK: "3213213213213213"},
		{NIK: "1234567890123456"},
		{NIK: "3210100401990000"},
		{NIK: "3210100413990001"},
		{NIK: "3210107713990001"},
		{NIK: "9999107713990001"},
		{NIK: "0000007713990001"},
	}
	for _, v := range test {
		check, err := v.IsValid()
		if check {
			t.Error(err)
		}
	}
}

func TestValidNIK(t *testing.T) {
	test := []panik.NIK{
		{NIK: "3210100401990001"},
		{NIK: "3210104401990001"},
		{NIK: "3210107101990001"},
	}

	for _, v := range test {
		check, err := v.IsValid()
		if !check {
			t.Error(err)
		}
	}
}

func TestDataNIK(t *testing.T) {
	test := panik.NIK{NIK: "3210100401990001"}
	data, _ := test.Data()
	fmt.Println(data)
}
