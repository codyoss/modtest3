package dog

import (
	"fmt"

	"github.com/codyoss/modtest3/cat"
)

type Dog struct {
	Name string
}

func (d *Dog) Bark(c cat.Cat) {
	fmt.Printf("%s barks at the cat named %s\n", d.Name, c.Name)
}
