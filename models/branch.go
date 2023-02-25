package models

type Branch struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BranchPrimaryKey struct {
	Id string `json:"id"`
}

type BranchReq struct {
	Name string `json:"name"`
}

type GetBranchListRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetBranchListResponse struct {
	Count    int      `json:"count"`
	Branches []Branch `json:"users"`
}
