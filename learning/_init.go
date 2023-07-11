package power

import (
	"fmt"
)

func hello(name string) string {
	message:=fmt.Sprintf("Hi ! %v,welcome "name)
	return message
}