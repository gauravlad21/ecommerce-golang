package common

type Unit string

const (
	Grams  Unit = "gms"
	Kilos  Unit = "kgs"
	Liters Unit = "ltr"
	Nos    Unit = "nos"
)

func (u Unit) String() string {
	switch u {
	case Grams:
		return "gms"
	case Kilos:
		return "kgs"
	case Liters:
		return "ltr"
	case Nos:
		return "nos"
	default:
		return "Unknown"
	}
}

func StringToUnit(str string) Unit {
	switch str {
	case "gms":
		return Grams
	case "kgs":
		return Kilos
	case "ltr":
		return Liters
	case "nos":
		return Nos
	default:
		return ""
	}
}

func GetUnits() []Unit {
	return []Unit{
		Grams,
		Kilos,
		Liters,
		Nos,
	}
}
