package collectiongroup

type CreateCGInput struct {
	CollectionGroup CollectionGroup `json:"CollectionGroup" binding:"required"`
}
