package hugo

import "os/exec"

// Build runs `hugo` and builds the static site located
// in the current working directory.
func Build() error {
	return exec.Command("hugo").Run()
}
