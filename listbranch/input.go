package listbranch

type BranchByAreaInput struct {
	AreaID string `json:"areaID" binding"required"`
}
type BranchByPosInput struct {
	PosID string `json:"posID" binding"required"`
}
