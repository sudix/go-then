package gothen

import "github.com/mattn/go-pipeline"

const (
	separetor = "!//-gothen-sep-//!"
)

// Cmd represents a command
type Cmd []string

// Then concatenates commands and return a new Cmd
func (c Cmd) Then(next Cmd) Cmd {
	newCmd := make(Cmd, 0, len(c)+len(next)+1)
	for _, e := range c {
		newCmd = append(newCmd, e)
	}
	newCmd = append(newCmd, separetor)
	for _, e := range next {
		newCmd = append(newCmd, e)
	}
	return newCmd
}

// Exec executes a Cmd
func (c Cmd) Exec() ([]byte, error) {
	cmds := c.split()
	out, err := pipeline.Output(cmds...)
	return out, err
}

func (c Cmd) split() [][]string {
	cmds := make([][]string, 1, 1)
	cmds[0] = make([]string, 0, 1)
	i := 0
	for _, e := range c {
		if e == separetor {
			cmds = append(cmds, make([]string, 0, 1))
			i++
			continue
		}
		cmds[i] = append(cmds[i], e)
	}
	return cmds
}
