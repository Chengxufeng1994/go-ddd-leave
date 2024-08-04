package valueobject

type PersonType int

const (
	INTERNAL PersonType = iota
	EXTERNAL
	UNKNOWN
)

func NewPersonType(personType int) PersonType {
	switch personType {
	case 0:
		return INTERNAL
	case 1:
		return EXTERNAL
	}

	return UNKNOWN
}
