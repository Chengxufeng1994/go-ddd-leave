package valueobject

import "github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/entity"

type Approver struct {
	PersonID   string
	PersonName string
	Level      int
}

func NewApprover(personID string, personName string) Approver {
	return Approver{
		PersonID:   personID,
		PersonName: personName,
	}
}

func (approver Approver) IsEmpty() bool {
	return approver.PersonID == "" || approver.PersonName == ""
}

func FromPerson(person *entity.Person) Approver {
	return Approver{
		PersonID:   person.PersonID,
		PersonName: person.PersonName,
		Level:      person.RoleLevel,
	}
}
