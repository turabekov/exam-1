package main

import (
	"app/config"
	"app/controller"
	"app/models"
	"app/storage/jsondb"
	"fmt"
	"log"
)

func main() {

	cfg := config.Load()
	jsondb, err := jsondb.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}
	defer jsondb.CloseDB()
	c := controller.NewController(&cfg, jsondb)

	// User(c)
	// Product(c)
	// ShopCart(c)
	// Category(c)

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
	err = c.WithdrawUserBalance("abea6d88-820e-4863-8f69-e91f891b92b0")
	if err != nil {
		log.Fatal(err)
	}

	// Task 11
	// Branch(c)
}

// Task 11
func Branch(c *controller.Controller) {
	// Create branch
	id, err := c.CreateBranch(models.BranchReq{
		Name: "Mirobod",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

	// Get List
	branches, err := c.GetBranchList(&models.GetBranchListRequest{
		Offset: 1,
		Limit:  100,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(branches.Branches)

	// Get by id
	branch, err := c.GetBranchByIdController(&models.BranchPrimaryKey{
		Id: "4101931f-c1f3-45a2-9e54-e2699b14bd97",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(branch)

	// Update branch
	b, e := c.UpdateBranchController(&models.Branch{
		Id:   "9ec528e7-4756-4005-9c4a-fa480ff63f9f",
		Name: "Beruniy",
	})
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(b)

	// Remove branch
	b, e = c.DeleteBranchController(&models.BranchPrimaryKey{
		Id: "b4bc93a4-e501-4ba1-9d20-232967649276",
	})
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(b)

}

func Category(c *controller.Controller) {
	// c.CreateCategory(&models.CreateCategory{
	// 	Name: "Xumoyun",
	// })
	// cat, e := c.GetByIdCategory(&models.CategoryPrimaryKey{
	// 	Id: "117bc391-ce09-4976-b5e9-7fdde869895b",
	// })
	// if e != nil {
	// 	log.Fatal(e)
	// }
	// fmt.Println(cat)
}

func User(c *controller.Controller) {}

func Product(c *controller.Controller) {}

func ShopCart(c *controller.Controller) {
	// Add data to shop cart
	sh, e := c.AddShopCart(&models.AddShopCart{
		ProductId: "ffa888f7-e0cb-44e9-9cae-8c4ca2a115b9",
		UserId:    "27457ac2-74dd-4656-b9b0-0d46b1af10dc",
		Count:     61,
	})
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Shop cart added", sh)

}
