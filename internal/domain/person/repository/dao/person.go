package dao

import (
	"context"

	"github.com/Chengxufeng1994/go-ddd-leave/internal/domain/person/repository/po"
	"gorm.io/gorm"
)

type PersonDao interface {
	Save(*po.Person)
	FindByID(personID string) *po.Person
	FindLeaderByPersonID(personID string) *po.Person
}

type PersonDaoImpl struct {
	db *gorm.DB
}

var _ PersonDao = (*PersonDaoImpl)(nil)

func NewPersonDao(db *gorm.DB) PersonDao {
	return &PersonDaoImpl{
		db: db,
	}
}

// Save implements PersonDao.
func (dao *PersonDaoImpl) Save(person *po.Person) {
	dao.db.
		WithContext(context.Background()).
		Model(&po.Person{}).
		Save(&person)
}

// FindByID implements PersonDao.
func (dao *PersonDaoImpl) FindByID(personID string) *po.Person {
	var person *po.Person
	dao.db.
		WithContext(context.Background()).
		Model(&po.Person{}).Where("id = ?", personID).First(&person)
	return person
}

// FindLeaderByPersonID implements PersonDao.
func (dao *PersonDaoImpl) FindLeaderByPersonID(personID string) *po.Person {
	var person *po.Person

	dao.db.
		WithContext(context.Background()).
		Model(&po.Person{}).
		Joins("left join relationships on relationships.leader_id = person.person_id").
		Where("relationships.person_id = ?", personID).
		First(&person)

	return person
}
