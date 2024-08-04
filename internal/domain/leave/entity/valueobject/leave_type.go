package valueobject

type LeaveType string

const (
	INTERNAL           LeaveType = "INTERNAL"
	EXTERNAL                     = "EXTERNAL"
	OFFICIAL                     = "OFFICIAL"
	UNKNOWN_LEAVE_TYPE           = "UNKNOWN"
)

func NewLeaveType(leaveType string) LeaveType {
	switch leaveType {
	case "INTERNAL":
		return INTERNAL
	case "EXTERNAL":
		return EXTERNAL
	case "OFFICIAL":
		return OFFICIAL
	}

	return UNKNOWN_LEAVE_TYPE
}
