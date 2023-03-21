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

type categoryRepo struct {
	fileName string
	file     *os.File
}

func NewCategoryRepo(fileName string, file *os.File) *categoryRepo {
	return &categoryRepo{
		fileName: fileName,
		file:     file,
	}
}

func (c *categoryRepo) Create(w http.ResponseWriter, r *http.Request) {
	// read and unmarshal category
	var categories []models.Category
	data, err := ioutil.ReadFile(c.fileName)
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(data, &categories)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	//  read request body unmarshal data
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}
	var category models.Category
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	// check existing of parent category  id in categories
	flag := false
	for _, ctg := range categories {
		if ctg.Id == category.ParentID {
			flag = true
			break
		}
	}
	if !flag {
		log.Println("parent id err: parent id not found")
		w.WriteHeader(400)
		w.Write([]byte("parent category not found"))
		return
	}

	uuid := uuid.New().String()
	categories = append(categories, models.Category{
		Id:       uuid,
		Name:     category.Name,
		ParentID: category.ParentID,
	})

	body, err = json.MarshalIndent(categories, "", " ")
	if err != nil {
		log.Println("Marshal err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}
	err = ioutil.WriteFile(c.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Write file err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created successfully!"))
}

func (u *categoryRepo) GetByID(w http.ResponseWriter, r *http.Request) {
	categories, err := u.Read()
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	//  read request body
	id := r.URL.Path[len("/category/"):]

	for _, v := range categories {
		if v.Id == id {
			for _, subCategory := range categories {
				if v.Id == subCategory.ParentID {
					v.SubCategories = append(v.SubCategories, subCategory)

				}
			}
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

// (models.GetListCategoryResponse, error)
func (u *categoryRepo) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := u.Read()
	if err != nil {
		log.Println("read file and unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	// get query params limit offset
	var (
		limit    int
		offset   int
		response *models.GetListCategoryResponse
		e        error
	)
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	if limitStr == "" {
		limit = len(categories)
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
	// -------------------------------------------------------------------------------------------------------------------------------------
	if limit+offset > len(categories) {
		if offset > len(categories) {
			response = &models.GetListCategoryResponse{
				Count:      len(categories),
				Categories: []models.Category{},
			}
		} else {
			response = &models.GetListCategoryResponse{
				Count:      len(categories),
				Categories: categories[offset:],
			}
		}

	} else {
		response = &models.GetListCategoryResponse{
			Count:      len(categories),
			Categories: categories[offset : limit+offset],
		}
	}

	for i, v := range response.Categories {
		for _, subCategory := range response.Categories {
			if v.Id == subCategory.ParentID {
				response.Categories[i].SubCategories = append(response.Categories[i].SubCategories, subCategory)
			}
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

func (u *categoryRepo) Update(w http.ResponseWriter, r *http.Request) {
	categories, err := u.Read()
	if err != nil {
		return
	}

	updatedCategory := models.UpdateCategory{}
	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	//  unmarshal data
	err = json.Unmarshal(body, &updatedCategory)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	flag := false
	for i, v := range categories {
		if v.Id == updatedCategory.Id {
			if updatedCategory.Name != "" {
				categories[i].Name = updatedCategory.Name
			}
			if updatedCategory.ParentID != "" {
				categories[i].ParentID = updatedCategory.ParentID
			}
			flag = true
		}
	}

	if !flag {
		res := "category with id" + " " + updatedCategory.Id + " " + "not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}
	// marshal updated  data and write it into  file
	body, err = json.MarshalIndent(categories, "", "   ")
	if err != nil {
		if err != nil {
			log.Println("Marshal err:", err)
			w.WriteHeader(500)
			w.Write([]byte("Incorrect data"))
			return
		}
	}
	err = ioutil.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		log.Println("Write file err:", err)
		w.WriteHeader(500)
		w.Write([]byte("Incorrect data"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("category updated successfully!"))
}

func (u *categoryRepo) Delete(w http.ResponseWriter, r *http.Request) {
	categories, err := u.Read()
	if err != nil {
		log.Println("read file err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	category := models.UpdateCategory{}
	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}

	//  unmarshal data
	err = json.Unmarshal(body, &category)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	flag := true
	for i, v := range categories {
		if v.Id == category.Id {
			categories = append(categories[:i], categories[i+1:]...)
			flag = false
			break
		}
	}

	if flag {
		res := "category with id" + " " + category.Id + " " + "not found"
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(res))
		return
	}

	body, err = json.MarshalIndent(categories, "", " ")
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

func (u *categoryRepo) Read() ([]models.Category, error) {
	data, err := ioutil.ReadFile(u.fileName)
	if err != nil {
		return []models.Category{}, err
	}

	var categories []models.Category
	err = json.Unmarshal(data, &categories)
	if err != nil {
		return []models.Category{}, err
	}
	return categories, nil
}
