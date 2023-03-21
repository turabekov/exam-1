package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsondb"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	jsondb, err := jsondb.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}
	defer jsondb.CloseDB()
	c := controller.NewController(&cfg, jsondb)

	// user crud api
	http.HandleFunc("/user", c.User)
	http.HandleFunc("/user/", c.User)

	// exchange money api
	http.HandleFunc("/exchange", c.ExchangeMoney)
	// category crud api
	http.HandleFunc("/category", c.Category)
	http.HandleFunc("/category/", c.Category)

	// product api crud
	http.HandleFunc("/product", c.Product)
	http.HandleFunc("/product/", c.Product)

	// Running server
	fmt.Println("Server running on port 3000")
	err = http.ListenAndServe("localhost:3000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func Exam() {
	// ======Exam===================================================================================================================================
	// Task1
	// res, err := c.FilterByDateShopCarts("2022-02-22", "2022-02-22")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, v := range res {
	// 	fmt.Println(v)
	// }

	// Task2
	// err = c.ClientHistory("abea6d88-820e-4863-8f69-e91f891b92b0")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Task3
	// err = c.ClientStats("abea6d88-820e-4863-8f69-e91f891b92b0")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Task4
	// err = c.ProductsSoldStats()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// Task5
	// err = c.TopMaxSoldProducts()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Task6
	// err = c.TopMinSoldProducts()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Task7
	// err = c.MaxSoldProductsDate()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Task 8
	// err = c.CategoryTable()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Task 9
	// err = c.ActiveClient()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Task 10
	// Add data to shop cart
	// sh, e := c.AddShopCart(&models.AddShopCart{
	// 	ProductId: "0a38e50d-0749-4bef-b6bc-ec9f3206f09a",
	// 	UserId:    "abea6d88-820e-4863-8f69-e91f891b92b0",
	// 	Count:     10,
	// })
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// fmt.Println("Shop cart added", sh)
	// draw from user balance
	// err = c.WithdrawUserBalance("abea6d88-820e-4863-8f69-e91f891b92b0")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Task 11
	// Branch(c)
}
