package collectiongroup

type CreateCGInputFormatter struct {
	CollectionGrup CollectionGroup `json:"CollectionGroup"`
}

func FormatCreateCG(collectionGroup CollectionGroup) CreateCGInputFormatter {
	formatter := CreateCGInputFormatter{
		CollectionGrup: collectionGroup,
	}
	return formatter
}
