package valueobject

type PersonStatus int

const (
	ENABLE PersonStatus = iota
	DISABLE
	STATUS_UNKNOWN
)

func NewPersonStatus(personStatus int) PersonStatus {
	switch personStatus {
	case 0:
		return ENABLE
	case 1:
		return DISABLE
	}

	return STATUS_UNKNOWN
}
