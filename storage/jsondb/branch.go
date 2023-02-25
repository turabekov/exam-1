package jsondb

import (
	"app/models"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/google/uuid"
)

type branchRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewBranchRepo(fileName string, file *os.File) *branchRepo {
	return &branchRepo{
		fileName: fileName,
		file:     file,
	}
}

func (b *branchRepo) CreateBranch(req *models.BranchReq) (id string, err error) {

	var branches []*models.Branch
	err = json.NewDecoder(b.file).Decode(&branches)
	if err != nil {
		return "", err
	}

	id = uuid.NewString()

	branches = append(branches, &models.Branch{
		Id:   id,
		Name: req.Name,
	})

	body, err := json.MarshalIndent(branches, "", "   ")

	if err != nil {
		return "", err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return "", err
	}

	return id, nil
}

// Get list of Users
func (b *branchRepo) GetList(req *models.GetBranchListRequest) (*models.GetBranchListResponse, error) {
	branches := make([]models.Branch, 0)

	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		return nil, err
	}

	if req.Limit+req.Offset > len(branches) {
		if req.Offset > len(branches) {
			return &models.GetBranchListResponse{
				Count:    len(branches),
				Branches: []models.Branch{},
			}, nil
		}

		return &models.GetBranchListResponse{
			Count:    len(branches),
			Branches: branches[req.Offset:],
		}, nil
	}

	response := &models.GetBranchListResponse{
		Count:    len(branches),
		Branches: branches[req.Offset : req.Limit+req.Offset],
	}

	return response, nil
}

// Get branch by id
func (b *branchRepo) GetBranchById(req *models.BranchPrimaryKey) (models.Branch, error) {
	branches := make([]models.Branch, 0)

	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return models.Branch{}, err
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		return models.Branch{}, err
	}

	for _, v := range branches {
		if v.Id == req.Id {
			return v, nil
		}
	}

	return models.Branch{}, errors.New("branch not found")
}

// Update user by id
func (b *branchRepo) UpdateBranch(req *models.Branch) (models.Branch, error) {
	branches := make([]models.Branch, 0)

	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return models.Branch{}, err
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		return models.Branch{}, err
	}

	updatedBranch := models.Branch{}
	for i, v := range branches {
		if v.Id == req.Id {
			branches[i].Name = req.Name

			updatedBranch = branches[i]
		}
	}

	if len(updatedBranch.Name) <= 0 {
		return models.Branch{}, errors.New("user not found")
	}

	body, err := json.MarshalIndent(branches, "", "   ")

	if err != nil {
		return models.Branch{}, err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return models.Branch{}, err
	}

	return updatedBranch, nil
}

// Delete branch by id
func (b *branchRepo) DeleteBranch(req *models.BranchPrimaryKey) (models.Branch, error) {
	branches := make([]models.Branch, 0)

	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		return models.Branch{}, err
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		return models.Branch{}, err
	}

	deletedBranch := models.Branch{}
	for i, v := range branches {
		if v.Id == req.Id {
			deletedBranch = branches[i]
			branches = append(branches[:i], branches[i+1:]...)
		}
	}

	if len(deletedBranch.Name) <= 0 {
		return models.Branch{}, errors.New("user not found")
	}

	body, err := json.MarshalIndent(branches, "", "   ")

	if err != nil {
		return models.Branch{}, err
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		return models.Branch{}, err
	}

	return deletedBranch, nil
}
