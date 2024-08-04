package valueobject

type Applicant struct {
	PersonID   string
	PersonName string
	PersonType string
}

func NewApplicant(personID string, personName string, personType string) Applicant {
	return Applicant{
		PersonID:   personID,
		PersonName: personName,
		PersonType: personType,
	}
}
