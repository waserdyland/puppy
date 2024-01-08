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
