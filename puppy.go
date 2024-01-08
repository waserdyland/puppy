package puppy

import (
	"fmt"

	"github.com/waserdyland/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof Woof Woof!"
}

func BigBark() string {
	return dog.WhenGrowsUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowsUp(Barks())
}

func From11() {
	fmt.Println("Hello from v1.1.0")
}

func From12() {
	fmt.Println("Hello from v1.2.0")
}

func From13() {
	fmt.Println("Hellow from v1.3.0")
}
