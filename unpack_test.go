package unpack

import (
	"testing"
)

func TestDo(t *testing.T) {

	t.Run("Should assign all names", func(t *testing.T) {
		names := []string{"Tom", "Jerry", "Sparrow"}
		var tom, jerry, sparrow string
		_ = Do(names, &tom, &jerry, &sparrow)
		if tom != "Tom" || jerry != "Jerry" || sparrow != "Sparrow" {
			t.Fail()
		}
	})

	t.Run("Should assign available variables when lesser source", func(t *testing.T) {
		names := []string{"Tom", "Jerry"}
		var tom, jerry string
		_ = Do(names, &tom, &jerry)
		if tom != "Tom" && jerry != "Jerry" {
			t.Fail()
		}
	})

	t.Run("Should assign available variables when lesser destination", func(t *testing.T) {
		names := []string{"Tom", "Jerry", "Sparrow"}
		var tom, jerry string
		_ = Do(names, &tom, &jerry)
		if tom != "Tom" && jerry != "Jerry" {
			t.Fail()
		}
	})

	t.Run("Should assign custom types", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		users := []User{{Name: "Tom", Age: 30}, {Name: "Jerry", Age: 40}}
		var tom, jerry User
		_ = Do(users, &tom, &jerry)
		if tom.Name != "Tom" && tom.Age != 30 && jerry.Name != "Jerry" && jerry.Age != 40 {
			t.Fail()
		}
	})

	t.Run("Should assign array of custom types", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		users := [2]User{{Name: "Tom", Age: 30}, {Name: "Jerry", Age: 40}}
		var tom, jerry User
		_ = Do(users, &tom, &jerry)
		if tom.Name != "Tom" && tom.Age != 30 && jerry.Name != "Jerry" && jerry.Age != 40 {
			t.Fail()
		}
	})

	t.Run("Should not panic while empty source", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		var names interface{}
		var tom, jerry User
		_ = Do(names, &tom, &jerry)
		if tom.Name != "" && tom.Age != 0 && jerry.Name != "" && jerry.Age != 0 {
			t.Fail()
		}
	})

	t.Run("Should not panic while empty destination", func(t *testing.T) {
		type User struct {
			Name string
			Age  int
		}
		users := []User{{Name: "Tom", Age: 30}, {Name: "Jerry", Age: 40}}
		_ = Do(users)
	})

	t.Run("Should return error while empty destination", func(t *testing.T) {
		names := []string{"Tom", "Jerry"}
		if err := Do(names); err == nil {
			t.Fail()
		}
	})

	t.Run("Should return error while source and destination type mismatch", func(t *testing.T) {
		names := []string{"Tom", "Jerry"}
		var tom int
		if err := Do(names, &tom); err == nil {
			t.Fail()
		}
	})

	t.Run("Should return error while the destination is not a pointer", func(t *testing.T) {
		names := []string{"Tom", "Jerry"}
		var tom string
		if err := Do(names, tom); err == nil {
			t.Fail()
		}
	})
}
