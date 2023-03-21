package jsondb

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

type userRepo struct {
	fileName string
	file     *os.File
}

// Constructor
func NewUserRepo(fileName string, file *os.File) *userRepo {
	return &userRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *userRepo) CreateUser(w http.ResponseWriter, r *http.Request) {
	// read users from file
	var users []*models.User
	err := json.NewDecoder(u.file).Decode(&users)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	//  unmarshal data
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	// add new user
	id := uuid.NewString()
	users = append(users, &models.User{
		Id:      id,
		Name:    user.Name,
		Surname: user.Surname,
		Balance: user.Balance,
	})
	// marshal it to json format
	body, err = json.MarshalIndent(users, "", "   ")
	if err != nil {
		log.Println("Marshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}
	// write it into json file
	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Write file err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created successfully!"))
}

// Get list of Users
func (u *userRepo) GetList(w http.ResponseWriter, r *http.Request) {
	// read users from  file
	users := make([]models.User, 0)
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	// get query params limit offset
	var (
		limit    int
		offset   int
		response *models.GetListResponse
		e        error
	)

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	if limitStr == "" {
		limit = len(users)
	} else {
		limit, e = strconv.Atoi(limitStr)
		if e != nil {
			log.Println("strconv err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
	}
	if offsetStr == "" {
		offset = 0
	} else {
		offset, e = strconv.Atoi(offsetStr)
		if e != nil {
			log.Println("strconv err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
	}

	if limit+offset > len(users) {
		if offset > len(users) {
			response = &models.GetListResponse{
				Count: len(users),
				Users: []models.User{},
			}
		} else {
			response = &models.GetListResponse{
				Count: len(users),
				Users: users[offset:],
			}
		}

	} else {
		response = &models.GetListResponse{
			Count: len(users),
			Users: users[offset : limit+offset],
		}
	}

	body, err := json.Marshal(response)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Write(body)
}

// Get list by id
func (u *userRepo) GetUserById(w http.ResponseWriter, r *http.Request) {
	// read All users from file
	users := make([]models.User, 0)
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	// unmarshal it
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	//  read request body
	id := r.URL.Path[len("/user/"):]

	for _, v := range users {
		if v.Id == id {
			body, err := json.Marshal(v)
			if err != nil {
				log.Println("Unmarshal err:", err)
				w.WriteHeader(500)
				w.Write([]byte("Incorrect data"))
				return
			}
			w.WriteHeader(http.StatusFound)
			w.Write(body)
			return
		}
	}

	res := "user with id" + " " + id + " " + "not existed"
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(res))
}

// Update user by id
func (u *userRepo) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// read file
	users := make([]models.User, 0)
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Println("unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	updatedUser := models.User{}
	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	//  unmarshal data
	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	flag := false
	for i, v := range users {
		if v.Id == updatedUser.Id {
			if updatedUser.Name != "" {
				users[i].Name = updatedUser.Name
			}
			if updatedUser.Surname != "" {
				users[i].Surname = updatedUser.Surname
			}
			if updatedUser.Balance != 0 {
				users[i].Balance = updatedUser.Balance
			}
			flag = true
		}
	}

	if !flag {
		res := "user with id" + " " + updatedUser.Id + " " + "not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}
	// marshal updated  data and write it into  file
	body, err = json.MarshalIndent(users, "", "   ")
	if err != nil {
		log.Println("Marshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}
	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Write file err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User updated successfully!"))
}

// Delete user by id
func (u *userRepo) DeleteUser(w http.ResponseWriter, r *http.Request) {
	users := make([]models.User, 0)
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	//  unmarshal data
	var userId models.UserPrimaryKey
	err = json.Unmarshal(body, &userId)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	flag := false
	for i, v := range users {
		if v.Id == userId.Id {
			users = append(users[:i], users[i+1:]...)
			flag = true

		}
	}

	if !flag {
		res := "user with id" + " " + userId.Id + " " + "not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}

	body, err = json.MarshalIndent(users, "", "   ")
	if err != nil {
		log.Println("Marshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Write file err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted successfully!"))
}
