package collectiongroup

import (
	"log"

	"gorm.io/gorm"
)

type Repository interface {
	ListPrivelege() ([]Privelege, error)
	ListCollectionGroup() ([]CollectionGroup, error)
	CreateCollectionGroup(collectionGroup CollectionGroup) (CollectionGroup, error)
	UpdateCollectionGroup(collectionGroup CollectionGroup) (CollectionGroup, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// List Privelege Group
func (r *repository) ListPrivelege() ([]Privelege, error) {

	var listPrivelege []Privelege

	err := r.db.Debug().Table("ms_privelege").Find(&listPrivelege).Error

	if err != nil {
		return listPrivelege, err
	}

	return listPrivelege, nil

}

func (r *repository) ListCollectionGroup() ([]CollectionGroup, error) {

	var listCollection []CollectionGroup

	err := r.db.Debug().Preload("CollectionGroupDetail").Table("ms_collection_group").Find(&listCollection).Error

	if err != nil {
		return listCollection, err
	}

	return listCollection, nil

}

func (r *repository) CreateCollectionGroup(collectionGroup CollectionGroup) (CollectionGroup, error) {

	var result int64

	err2 := r.db.Table("ms_collection_group").Count(&result).Error
	if err2 != nil {
		log.Fatal(err2)
	}

	collectionGroup.Modify(result)
	err := r.db.Debug().Table("ms_collection_group").Create(&collectionGroup).Error

	if err != nil {
		return collectionGroup, err
	}

	return collectionGroup, nil
}

func (r *repository) UpdateCollectionGroup(collectionGroup CollectionGroup) (CollectionGroup, error) {

	err := r.db.Debug().Model(&collectionGroup).Association("CollectionGroupDetail").Replace(&collectionGroup)
	if err != nil {
		return collectionGroup, err
	}

	return collectionGroup, nil
}
