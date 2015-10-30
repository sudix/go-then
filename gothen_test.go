package gothen

import "testing"

func TestThen(t *testing.T) {
	c1 := Cmd{"cmd1", "param1"}
	c2 := Cmd{"cmd2", "param2"}
	c3 := Cmd{"cmd3"}

	cmds := c1.Then(c2).Then(c3)

	expected := []string{
		"cmd1", "param1", separetor,
		"cmd2", "param2", separetor,
		"cmd3",
	}

	if isDifferentSlice(expected, []string(cmds)) {
		t.Errorf("command concat error. expected:%q, got:%q", expected, cmds)
	}

}

func TestSplit(t *testing.T) {
	cmd := Cmd{"cmd1", "param1-1", "param1-2", separetor,
		"cmd2", separetor, "cmd3", "param3-1"}

	splitted := cmd.split()

	if len(splitted) != 3 {
		t.Errorf("command size error. expected:%d, got:%d", 3, len(splitted))
	}

	expcmds := [][]string{
		[]string{"cmd1", "param1-1", "param1-2"},
		[]string{"cmd2"},
		[]string{"cmd3", "param3-1"},
	}

	for i, expcmd := range expcmds {
		actual := splitted[i]
		if isDifferentSlice(expcmd, actual) {
			t.Errorf("cmd%d not correct. expected: %q, got %q", i+1, expcmd, actual)
		}
	}
}

func isDifferentSlice(expected, actual []string) bool {
	if len(expected) != len(actual) {
		return true
	}
	for i, e := range expected {
		if e != actual[i] {
			return true
		}
	}
	return false
}
