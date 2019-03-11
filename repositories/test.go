package repositories

import (
	"emptyApi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

type TestRepository struct {
	db *gorm.DB
}

// FindByID ...
func (repo *TestRepository) FindByID(id int64) (obj *models.Test, err error) {
	obj = &models.Test{ID:id}
	if result := repo.db.Find(obj); result.Error != nil {
		return nil, result.Error
	}

	return obj, nil
}

func (repo *TestRepository) Store(entity *models.Test) error {
	return repo.db.Save(entity).Error
}
func (repo *TestRepository) StoreWithTx(tx *gorm.DB, entity *models.Test) error {
	if err := tx.Save(entity).Error; err != nil {
		return err
	}
	return nil
}
func NewTestRepository(Connect func() *gorm.DB) *TestRepository {
	db := Connect()
	if db == nil {
		log.Fatal("cannot connect to db")
		return nil
	}
	return &TestRepository{db: db}
}
