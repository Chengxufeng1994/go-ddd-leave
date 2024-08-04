package po

type Relationship struct {
	ID       string
	PersonID string
	LeaderID string
}

func (r Relationship) TableName() string {
	return "relationships"
}
