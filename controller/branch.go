package controller

import "app/models"

func (c *Controller) CreateBranch(req models.BranchReq) (string, error) {
	id, err := c.store.Branch().CreateBranch(&req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Controller) GetBranchList(req *models.GetBranchListRequest) (*models.GetBranchListResponse, error) {
	branches, err := c.store.Branch().GetList(req)
	if err != nil {
		return &models.GetBranchListResponse{}, err
	}

	return branches, nil
}

func (c *Controller) GetBranchByIdController(req *models.BranchPrimaryKey) (models.Branch, error) {
	branch, err := c.store.Branch().GetBranchById(req)
	if err != nil {
		return models.Branch{}, err
	}

	return branch, nil
}

func (c *Controller) UpdateBranchController(req *models.Branch) (models.Branch, error) {
	branch, err := c.store.Branch().UpdateBranch(req)
	if err != nil {
		return models.Branch{}, err
	}

	return branch, nil
}
func (c *Controller) DeleteBranchController(req *models.BranchPrimaryKey) (models.Branch, error) {
	b, err := c.store.Branch().DeleteBranch(req)
	if err != nil {
		return models.Branch{}, err
	}

	return b, nil
}
