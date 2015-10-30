# go-then

concatenate shell commands and execute it on golang.

Requirements
===================

This library relies on [mattn's go-pipeline](https://github.com/mattn/go-pipeline)

Instllation
===================

```
$ go get github.com/sudix/go-then
```

Usage
===================

```go
func ExampleExec() {
	c1 := gothen.Cmd{"git", "log", "--oneline"}
	c2 := gothen.Cmd{"grep", "first import"}
	c3 := gothen.Cmd{"wc", "-l"}

	out, err := c1.Then(c2).Then(c3).Exec()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
```

License
===================

MIT
