package validation

import (
	"testing"

	"github.com/asaskevich/govalidator"
)

type item struct {
	Name string `valid:"required"`
}

type items struct {
	Items []item `valid:"required,arrayStructValidator"`
}

func Test_ArrayValidator(t *testing.T) {
	testcase := items{
		Items: []item{
			{
				Name: "a",
			},
		},
	}

	ok, err := govalidator.ValidateStruct(testcase)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fail()
	}

	testcase = items{
		Items: []item{
			{},
		},
	}

	ok, err = govalidator.ValidateStruct(testcase)
	if err == nil {
		t.Fatal(err)
	}
	if ok {
		t.Fail()
	}
}
