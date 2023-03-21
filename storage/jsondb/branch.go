package jsondb

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

func (b *branchRepo) CreateBranch(w http.ResponseWriter, r *http.Request) {
	// readFile
	branches := make([]models.Branch, 0)
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	// create branch
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	var branch models.Branch

	err = json.Unmarshal(body, &branch)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	branch.Id = uuid.New().String()

	body, err = json.Marshal(branch)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	branches = append(branches, branch)
	// write into json file
	body, err = json.MarshalIndent(branches, "", "   ")
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	jsonBranch, err := json.MarshalIndent(branch, "", "   ")
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(jsonBranch)
}

// ==============================================================================================================================================
func (b *branchRepo) GetAll(w http.ResponseWriter, r *http.Request) {
	// readFile
	branches := make([]models.Branch, 0)
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	// json stringify
	body, err := json.Marshal(branches)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(200)
	w.Write(body)
}

// ==============================================================================================================================================
func (b *branchRepo) GetBranchById(w http.ResponseWriter, r *http.Request) {
	// readFile
	branches := make([]models.Branch, 0)
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	id := r.URL.Path[len("/branch/"):]

	var (
		branch models.Branch
	)

	for ind, _ := range branches {
		if id == branches[ind].Id {
			branch = branches[ind]
			break
		}
	}

	body, err := json.Marshal(branch)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(200)
	w.Write(body)
}

// ==============================================================================================================================================
func (b *branchRepo) UpdateBranch(w http.ResponseWriter, r *http.Request) {
	// readFile
	branches := make([]models.Branch, 0)
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	var branch models.Branch

	err = json.Unmarshal(body, &branch)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	flag := false
	for i, _ := range branches {
		if branch.Id == branches[i].Id {
			branches[i].Name = branch.Name
			flag = true
		}
	}

	if !flag {
		res := "branch with id" + " " + branch.Id + " " + "not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}

	// write into json file
	body, err = json.MarshalIndent(branches, "", "   ")

	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Updated"))
}

// ==============================================================================================================================================
func (b *branchRepo) DeleteBranch(w http.ResponseWriter, r *http.Request) {
	// readFile
	branches := make([]models.Branch, 0)
	data, err := ioutil.ReadFile(b.fileName)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &branches)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	var branchId models.BranchPrimaryKey

	err = json.Unmarshal(body, &branchId)
	if err != nil {
		log.Println("Post err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	flag := false
	for i, v := range branches {
		if v.Id == branchId.Id {
			branches = append(branches[:i], branches[i+1:]...)
			flag = true
		}
	}

	if !flag {
		res := "branch with id" + " " + branchId.Id + " " + "not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}

	// write into json file
	body, err = json.MarshalIndent(branches, "", "   ")

	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	err = ioutil.WriteFile(b.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Deleted"))
}
