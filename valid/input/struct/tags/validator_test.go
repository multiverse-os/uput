package validstruct

import (
	"testing"
)

type struct1 struct {
	Name  string
	Email string
	Age   int
}

type struct2 struct {
	Name  string `validates:"presence"`
	Email string `validates:"presence,format=email"`
	Age   int    `validates:"min=18,max=100"`
}

func Test_Struct_Without_Tags(t *testing.T) {
	strc := struct1{
		"Hugo",
		"test@test.com",
		10,
	}

	valid := Valid(strc)

}

func Test_Struct_Being_Valid(t *testing.T) {
	strc := struct2{
		"Hugo",
		"test@test.com",
		30,
	}

	valid := Valid(strc)

}

func Test_Struct_Not_Valid_Min(t *testing.T) {

	strc := struct2{
		"Hugo",
		"test@test.com",
		10,
	}

	valid := Valid(strc)

}

func Test_Struct_Not_Valid_Max(t *testing.T) {

	strc := struct2{
		"Hugo",
		"test@test.com",
		200,
	}

	valid := Valid(strc)

}

func Test_Struct_Not_Valid_Format(t *testing.T) {

	strc := struct2{
		"Hugo",
		"test",
		30,
	}

	valid := Valid(strc)

}
