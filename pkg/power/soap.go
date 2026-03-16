package power

import (
	"os/exec"
)

func (m *Manager) run(args ...string) ([]byte, error) {

	cmd := exec.Command("wsman", args...)

	return cmd.CombinedOutput()
}