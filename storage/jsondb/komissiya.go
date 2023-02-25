package jsondb

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"os"
)

type komissiyaRepo struct {
	fileName string
	file     *os.File
}

func NewKomissiyaRepo(fileName string, file *os.File) *komissiyaRepo {
	return &komissiyaRepo{
		fileName: fileName,
		file:     file,
	}
}

func (k *komissiyaRepo) GetKomissiya() (models.Komissiya, error) {
	komissiya := []models.Komissiya{}

	data, err := ioutil.ReadFile(k.fileName)
	if err != nil {
		return models.Komissiya{}, err
	}

	err = json.Unmarshal(data, &komissiya)
	if err != nil {
		return models.Komissiya{}, err
	}

	return komissiya[0], nil
}

func (k *komissiyaRepo) UpdateBalanceKomissiya(komissiya models.Komissiya) error {
	komissiyaArr := []models.Komissiya{}

	data, err := ioutil.ReadFile(k.fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &komissiyaArr)
	if err != nil {
		return err
	}

	komissiyaArr[0].Balance = komissiya.Balance

	body, err := json.MarshalIndent(komissiyaArr, "", "   ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(k.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
