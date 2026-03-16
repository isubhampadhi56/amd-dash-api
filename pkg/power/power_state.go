package power

type PowerState int

const (
	PowerOn        PowerState = 2
	PowerSleep     PowerState = 4
	PowerHibernate PowerState = 7
	PowerOff       PowerState = 8
	PowerReset     PowerState = 10
	PowerUnknown   PowerState = 99
)

func (p PowerState) String() string {

	switch p {

	case PowerOn:
		return "on"

	case PowerOff:
		return "off"

	case PowerSleep:
		return "sleep"

	case PowerHibernate:
		return "hibernate"

	case PowerReset:
		return "reset"

	default:
		return "unknown"
	}
}