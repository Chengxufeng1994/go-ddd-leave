package valueobject

type Status string

const (
	APPROVING Status = "APPROVING"
	APPROVED         = "APPROVED"
	REJECTED         = "REJECTED"
	UNKNOWN          = "UNKNOWN"
)

func NewStatus(status string) Status {
	switch status {
	case "APPROVING":
		return APPROVING
	case "APPROVED":
		return APPROVED
	case "REJECTED":
		return REJECTED
	}

	return UNKNOWN
}
