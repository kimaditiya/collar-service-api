package collectiongroup

import "gorm.io/gorm"

type Repository interface {
	ListPrivelege() ([]Privelege, error)
	CreateCollectionGroup(collectionGroup CollectionGroup) (CollectionGroup, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// List Privelege
func (r *repository) ListPrivelege() ([]Privelege, error) {

	var listPrivelege []Privelege

	err := r.db.Debug().Table("ms_privelege").Find(&listPrivelege).Error

	if err != nil {
		return listPrivelege, err
	}

	return listPrivelege, nil

}

func (r *repository) CreateCollectionGroup(collectionGroup CollectionGroup) (CollectionGroup, error) {
	// err := r.db.Debug().Preload("Area").Table("ms_collection_group").Create(&collectionGroup).Error
	err := r.db.Debug().Table("ms_collection_group").Create(&collectionGroup).Error

	if err != nil {
		return collectionGroup, err
	}
	return collectionGroup, nil
}
