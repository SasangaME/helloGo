/**
* Reflection
* Allows a program to examine its own structure and properties at runtime.
 */

package advanced

import (
	"fmt"
	"reflect"
)

type User struct {
	Name  string
	Email string
}

func Reflection() {
	user := User{
		Name:  "John",
		Email: "john@email.com",
	}

	t := reflect.TypeOf(user)

	fmt.Println("struct name: ", t.Name())
	fmt.Println("user type: ", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i))
	}
}
