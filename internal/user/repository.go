package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	Get(id string) (*User, error)
	Delete(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (repo *repo) Create(user *User) error {
	user.ID = uuid.New().String()

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	}
	repo.log.Println("user fue creado con ID ", user.ID)
	repo.log.Println(user)
	return nil
}

func (repo *repo) GetAll() ([]User, error) {
	var users []User
	result := repo.db.Model(&users).Order("first_name desc").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo *repo) Get(id string) (*User, error) {
	user := User{ID: id}
	result := repo.db.First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (repo *repo) Delete(id string) error {
	user := User{ID: id}
	result := repo.db.Delete(&user)

	if result.Error != nil {
		return result.Error
	}
	log.Println("valor deleted")
	log.Println(user.Deleted.Value())
	return nil
}

func (repo *repo) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {
	values := make(map[string]interface{})
	if firstName != nil {
		values["first_name"] = *firstName
	}
	if lastName != nil {
		values["last_name"] = *lastName
	}
	if email != nil {
		values["email"] = *email
	}
	if phone != nil {
		values["phone"] = *phone
	}

	if result := repo.db.Model(&User{}).Where("id = ?", id).Updates(values); result.Error != nil {
		return result.Error
	}

	return nil
}
