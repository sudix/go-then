package gothen

import (
	"fmt"
	"log"

	"github.com/sudix/go-then"
)

func ExampleExec() {
	c1 := gothen.Cmd{"git", "log", "--oneline"}
	c2 := gothen.Cmd{"grep", "first import"}
	c3 := gothen.Cmd{"wc", "-l"}

	out, err := c1.Then(c2).Then(c3).Exec()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
	// Output:
	// 1
}
