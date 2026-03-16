package power

import (
	"bytes"
	"strconv"
	// "fmt"
)

func (m *Manager) PowerState() (PowerState, error) {

	args := []string{
		"enumerate",
		"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_AssociatedPowerManagementService",
		"-h", m.Host,
		"-P", m.Port,
		"-u", m.Username,
		"-p", m.Password,
		"--auth=digest",
	}

	out, err := m.run(args...)
	if err != nil {
		return PowerUnknown, err
	}

	if bytes.Contains(out, []byte("<p:PowerState>2</p:PowerState>")) {
		return PowerOn, nil
	}

	if bytes.Contains(out, []byte("<p:PowerState>8</p:PowerState>")) {
		return PowerOff, nil
	}

	if bytes.Contains(out, []byte("<p:PowerState>10</p:PowerState>")) {
		return PowerReset, nil
	}

	return PowerUnknown, nil
}

func (m *Manager) changePowerState(state int) error {
	args := []string{
		"invoke",
		"-a", "RequestPowerStateChange",
		"http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/CIM_PowerManagementService",
		"-h", m.Host,
		"-P", m.Port,
		"-u", m.Username,
		"-p", m.Password,
		"--auth=digest",
		"-k", "Name=ManagedSystem:0",
		"-k", "RequestedPowerState=" + strconv.Itoa(state),
	}

	_, err := m.run(args...)
	return err
}

func (m *Manager) PowerOn() error {

	state, err := m.PowerState()
	if err != nil {
		return err
	}

	if state == PowerOn {
		return nil
	}

	return m.changePowerState(2)
}

func (m *Manager) PowerOff() error {

	state, err := m.PowerState()
	if err != nil {
		return err
	}

	if state == PowerOff {
		return nil
	}

	return m.changePowerState(8)
}

func (m *Manager) PowerCycle() error {

	err := m.changePowerState(8)
	if err != nil {
		return err
	}

	return m.changePowerState(10)
}