package controller

import (
	"app/models"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"
)

// Task 1 Shop cartlar Date boyicha filter qoyish kerak
func (c *Controller) FilterByDateShopCarts(fromDate string, toDate string) ([]models.ShopCart, error) {
	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return []models.ShopCart{}, err
	}

	firstDate, err := time.Parse("2006-01-02", fromDate)
	if err != nil {
		return []models.ShopCart{}, err
	}
	secondDate, err := time.Parse("2006-01-02", toDate)
	if err != nil {
		return []models.ShopCart{}, err
	}

	if !secondDate.After(firstDate) && !secondDate.Equal(firstDate) {
		return []models.ShopCart{}, errors.New("from date is invalid")
	}

	filteredShopCarts := []models.ShopCart{}
	for _, v := range shopCarts {
		str := strings.Split(v.Time, " ")
		parsedDate, err := time.Parse("2006-01-02", str[0])
		if err != nil {
			return []models.ShopCart{}, err
		}

		if (firstDate.Before(parsedDate) && secondDate.After(parsedDate)) || secondDate.Equal(parsedDate) {
			filteredShopCarts = append(filteredShopCarts, v)
		}
	}

	return filteredShopCarts, nil
}

// Task2 Client history chiqish kerak. Ya'ni sotib olgan mahsulotlari korsatish kerak
func (c *Controller) ClientHistory(userId string) error {
	// check user exists
	user, err := c.store.User().GetUserById(&models.UserPrimaryKey{
		Id: userId,
	})
	if err != nil {
		return err
	}

	shopCarts, err := c.store.ShopCart().GetUserShopCarts(&models.UserPrimaryKey{
		Id: userId,
	})
	if err != nil {
		return err
	}

	fmt.Println("Client Name:", user.Name)
	i := 1
	for _, v := range shopCarts {
		product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
			Id: v.ProductId,
		})
		if err != nil {
			return err
		}
		if v.Status {
			fmt.Printf("%d. Name: %s Price: %d Count: %d Total: %d Time: %s\n", i, product.Name, product.Price, v.Count, product.Price*v.Count, v.Time)
			i++
		}
	}
	return nil
}

// Task3 Client qancha pul mahsulot sotib olganligi haqida hisobot.
func (c *Controller) ClientStats(userId string) error {
	// check user exists
	user, err := c.store.User().GetUserById(&models.UserPrimaryKey{
		Id: userId,
	})
	if err != nil {
		return err
	}

	shopCarts, err := c.store.ShopCart().GetUserShopCarts(&models.UserPrimaryKey{
		Id: userId,
	})
	if err != nil {
		return err
	}

	total := 0
	for _, v := range shopCarts {
		product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
			Id: v.ProductId,
		})
		if err != nil {
			return err
		}
		if v.Status {
			total += product.Price * v.Count
		}
	}

	fmt.Printf("Name: %s Total Buy Price: %d\n", user.Name, total)

	return nil
}

// Task 4. Productlarni Qancha sotilgan boyicha hisobot

func (c *Controller) ProductsSoldStats() error {
	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return err
	}

	hash := map[string]int{}
	for _, v := range shopCarts {
		if v.Status {
			hash[v.ProductId] = hash[v.ProductId] + v.Count
		}
	}

	for key, val := range hash {
		product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
			Id: key,
		})
		if err != nil {
			return err
		}
		fmt.Printf(" Name: %s count: %d\n", product.Name, val)
	}

	return nil
}

// Task5. Top 10 ta sotilayotgan mahsulotlarni royxati.
type ProductCount struct {
	Name  string
	Count int
}

func (c *Controller) TopMaxSoldProducts() error {
	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return err
	}

	hash := map[string]int{}
	for _, v := range shopCarts {
		if v.Status {
			hash[v.ProductId] = hash[v.ProductId] + v.Count
		}
	}

	productCounts := []ProductCount{}
	for key, val := range hash {
		product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
			Id: key,
		})
		if err != nil {
			return err
		}
		productCounts = append(productCounts, ProductCount{
			Name:  product.Name,
			Count: val,
		})
	}

	sort.Slice(productCounts, func(i, j int) bool {
		return productCounts[i].Count > productCounts[j].Count
	})

	fmt.Println("Top max sold products:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d Name: %s count: %d\n", i+1, productCounts[i].Name, productCounts[i].Count)
	}

	return nil
}

// Task6. TOP 10 ta Eng past sotilayotgan mahsulotlar royxati
func (c *Controller) TopMinSoldProducts() error {
	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return err
	}

	hash := map[string]int{}
	for _, v := range shopCarts {
		if v.Status {
			hash[v.ProductId] = hash[v.ProductId] + v.Count
		}
	}

	productCounts := []ProductCount{}
	for key, val := range hash {
		product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
			Id: key,
		})
		if err != nil {
			return err
		}
		productCounts = append(productCounts, ProductCount{
			Name:  product.Name,
			Count: val,
		})
	}

	sort.Slice(productCounts, func(i, j int) bool {
		return productCounts[i].Count < productCounts[j].Count
	})
	fmt.Println("Top min sold products:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d Name: %s count: %d\n", i+1, productCounts[i].Name, productCounts[i].Count)
	}

	return nil
}

// Task7 Qaysi Sanada eng kop mahsulot sotilganligi boyicha jadval
type ProductTimeCount struct {
	Name  string
	Count int
	Time  string
}

func (c *Controller) MaxSoldProductsDate() error {
	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return err
	}

	hash := map[string]int{}
	for _, v := range shopCarts {
		if v.Status {
			date := strings.Split(v.Time, " ")
			hash[date[0]] = hash[date[0]] + v.Count
		}
	}

	// max date sold
	var maxSoldDate string
	var count int = 0
	for key, val := range hash {
		if count < val {
			count = val
			maxSoldDate = key
		}
	}

	productTimeCounts := []ProductTimeCount{}
	for _, v := range shopCarts {
		if v.Status {
			date := strings.Split(v.Time, " ")
			if maxSoldDate == date[0] {
				product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
					Id: v.ProductId,
				})
				if err != nil {
					return err
				}
				productTimeCounts = append(productTimeCounts, ProductTimeCount{
					Name:  product.Name,
					Count: v.Count,
					Time:  v.Time,
				})
			}
		}
	}

	sort.Slice(productTimeCounts, func(i, j int) bool {
		return productTimeCounts[i].Count > productTimeCounts[j].Count
	})

	fmt.Printf("In this day %s products sold most count %d:\n", maxSoldDate, count)
	for i := 0; i < len(productTimeCounts); i++ {
		fmt.Printf("%d Name: %s Sana: %s count: %d\n", i+1, productTimeCounts[i].Name, productTimeCounts[i].Time, productTimeCounts[i].Count)
	}

	return nil
}

// Task8  Qaysi category larda qancha mahsulot sotilgan boyicha jadval Name: Electronika Count: 12

func (c *Controller) CategoryTable() error {
	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return err
	}

	hashCategory := map[string]int{}
	hashSubCategory := map[string]int{}
	for _, v := range shopCarts {
		if v.Status {
			product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
				Id: v.ProductId,
			})
			if err != nil {
				return err
			}
			category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{
				Id: product.CategoryId,
			})
			if err != nil {
				return err
			}

			hashSubCategory[product.CategoryId] = hashSubCategory[product.CategoryId] + v.Count

			if len(category.ParentID) > 0 {
				categoryParent, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{
					Id: category.ParentID,
				})
				if err != nil {
					return err
				}
				hashCategory[categoryParent.Id] = hashCategory[categoryParent.Id] + v.Count
			} else {
				hashCategory[category.Id] = hashCategory[category.Id] + v.Count
			}
		}
	}

	for key, val := range hashCategory {
		category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{
			Id: key,
		})
		if err != nil {
			return err
		}

		fmt.Printf("Catogry Name: %s Count: %d\n", category.Name, val)
		for key2, val2 := range hashSubCategory {
			subcategory, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{
				Id: key2,
			})
			if err != nil {
				return err
			}
			if subcategory.ParentID == category.Id {
				fmt.Printf("\tName: %s Count: %d\n", subcategory.Name, val2)
			}
		}

	}

	return nil
}

// Task 9 Qaysi Client eng Active xaridor. Bitta ma'lumot chiqsa yetarli.

func (c *Controller) ActiveClient() error {
	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
	if err != nil {
		return err
	}

	hash := map[string]int{}
	for _, v := range shopCarts {
		hash[v.UserId] = hash[v.UserId] + v.Count
	}

	maxCount := 0
	var userId string
	for key, val := range hash {
		if maxCount < val {
			maxCount = val
			userId = key
		}
	}

	user, err := c.store.User().GetUserById(&models.UserPrimaryKey{
		Id: userId,
	})
	if err != nil {
		return err
	}

	fmt.Println("Active xaridor:\n\tName", user.Name, "\n\tMaxsulotlar soni umumiy:", maxCount)

	return nil
}

// 10. Agar client 9 dan katta mahuslot sotib olgan bolsa,
//     1 tasi tekinga beriladi va 9 ta uchun pul hisoblanadi.
//     1 tasi eng arzon mahsulotni pulini hisoblamaysiz.
//     Yangi korzinka qoshib tekshirib koring.
// Task10 in User Controller in WithdrawUserBalance function
