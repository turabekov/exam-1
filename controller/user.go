package controller

import (
	"app/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func (c *Controller) User(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c.store.User().CreateUser(w, r)
	}
	if r.Method == "GET" {
		path := strings.Split(r.URL.Path, "/")

		if len(path) > 2 {
			c.store.User().GetUserById(w, r)
		} else {
			c.store.User().GetList(w, r)
		}
	}
	if r.Method == "PUT" {
		c.store.User().UpdateUser(w, r)
	}
	if r.Method == "DELETE" {
		c.store.User().DeleteUser(w, r)
	}
}

// Task 10 Agar client 9 dan katta mahuslot sotib olgan bolsa,
// 1 tasi tekinga beriladi va 9 ta uchun pul hisoblanadi.
// 1 tasi eng arzon mahsulotni pulini hisoblamaysiz.
// Yangi korzinka qoshib tekshirib koring.
// func (c *Controller) WithdrawUserBalance(id string) error {

// 	user, err := c.store.User().GetUserById(&models.UserPrimaryKey{Id: id})
// 	if err != nil {
// 		return err
// 	}

// 	// Task10
// 	shopCarts, err := c.store.ShopCart().GetUserShopCarts(&models.UserPrimaryKey{
// 		Id: user.Id,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	userShopcarts := []models.ShopCart{}
// 	counter := 0
// 	for _, v := range shopCarts {
// 		if !v.Status {
// 			userShopcarts = append(userShopcarts, v)
// 			counter += v.Count
// 		}
// 	}

// 	totalBalance := 0
// 	// calc total
// 	for _, v := range userShopcarts {
// 		product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
// 			Id: v.ProductId,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		totalBalance += v.Count * product.Price
// 	}

// 	if counter > 9 {
// 		// declaring minimum price
// 		product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
// 			Id: userShopcarts[0].ProductId,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		dicountPrice := product.Price
// 		fmt.Println(dicountPrice)
// 		for _, v := range userShopcarts {
// 			product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
// 				Id: v.ProductId,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 			if product.Price < dicountPrice {
// 				dicountPrice = product.Price
// 			}
// 		}
// 		totalBalance = totalBalance - dicountPrice
// 	}

// 	newBalance := user.Balance - float64(totalBalance)
// 	if newBalance <= 0 {
// 		return errors.New("not enough money")
// 	}

// 	_, err = c.store.User().UpdateUser(&models.UpdateUser{
// 		Id:      user.Id,
// 		Name:    user.Name,
// 		Surname: user.Surname,
// 		Balance: newBalance,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	p, e := c.store.ShopCart().GetUserShopCarts(&models.UserPrimaryKey{
// 		Id: user.Id,
// 	})
// 	if e != nil {
// 		return e
// 	}

// 	for _, v := range p {
// 		_, err := c.store.ShopCart().UpdateShopCart(models.ShopCart{
// 			ProductId: v.ProductId,
// 			UserId:    v.UserId,
// 			Count:     v.Count,
// 			Status:    true,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // Homework
func (c *Controller) ExchangeMoney(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" || r.Method == "PUT" {
		//  read request body
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println("ioutil err:", err)
			w.WriteHeader(400)
			w.Write([]byte("Incorrect data"))
			return
		}
		//  unmarshal data
		var reqBody models.ReqExchangeMoney
		err = json.Unmarshal(body, &reqBody)
		if err != nil {
			log.Println("Unmarshal err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		//  get receiver user
		// -------------------------------------------------------------------------------------------------------------------------------
		res, err := http.Get("http://localhost:3000/user/" + reqBody.ReceiverId)
		if err != nil || res.StatusCode >= 400 {
			log.Println("Get user err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		var receiver models.User
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("ioutil err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(resBody, &receiver)
		if err != nil {
			log.Println("Unmarshal err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		// get sender user
		// -------------------------------------------------------------------------------------------------------------------------------
		res, err = http.Get("http://localhost:3000/user/" + reqBody.SenderId)
		if err != nil || res.StatusCode >= 400 {
			log.Println("Get user err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		var sender models.User
		resBody, err = ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("ioutil err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		err = json.Unmarshal(resBody, &sender)
		if err != nil {
			log.Println("Unmarshal err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Println(receiver, sender)

		komissiya, err := c.store.Komissiya().GetKomissiya()
		if err != nil {
			log.Println("comission err:", err)
			w.WriteHeader(400)
			w.Write([]byte(err.Error()))
			return
		}

		moneyWithKomissiya := float64(reqBody.Amount) + (float64(reqBody.Amount) * komissiya.Amount / 100)

		if sender.Balance > moneyWithKomissiya {
			// update sender  user
			user := &models.UpdateUser{
				Id:      sender.Id,
				Name:    sender.Name,
				Surname: sender.Surname,
				Balance: sender.Balance - moneyWithKomissiya,
			}
			data, err := json.Marshal(user)
			if err != nil {
				log.Println("Unmarshal err:", err)
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}
			bodyReader := bytes.NewReader(data)
			req, err := http.NewRequest(http.MethodPut, "http://localhost:3000/user", bodyReader)
			if err != nil {
				fmt.Printf("client: could not create request: %s\n", err)
				os.Exit(1)
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("client: error making http request: %s\n", err)
				os.Exit(1)
			}
			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Printf("client: could not read response body: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("client: response body: %s\n", resBody)
			// update receiver user
			user = &models.UpdateUser{
				Id:      receiver.Id,
				Name:    receiver.Name,
				Surname: receiver.Surname,
				Balance: receiver.Balance + float64(reqBody.Amount),
			}
			data, err = json.Marshal(user)
			if err != nil {
				log.Println("Unmarshal err:", err)
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}
			bodyReader = bytes.NewReader(data)
			req, err = http.NewRequest(http.MethodPut, "http://localhost:3000/user", bodyReader)
			if err != nil {
				fmt.Printf("client: could not create request: %s\n", err)
				os.Exit(1)
			}
			res, err = http.DefaultClient.Do(req)
			if err != nil {
				fmt.Printf("client: error making http request: %s\n", err)
				os.Exit(1)
			}
			resBody, err = ioutil.ReadAll(res.Body)
			if err != nil {
				fmt.Printf("client: could not read response body: %s\n", err)
				os.Exit(1)
			}
			fmt.Printf("client: response body: %s\n", resBody)
			// update comission
			err = c.store.Komissiya().UpdateBalanceKomissiya(models.Komissiya{
				Balance: komissiya.Balance + (float64(reqBody.Amount) * komissiya.Amount / 100),
				Amount:  komissiya.Amount,
			})
			if err != nil {
				log.Println("comission err:", err)
				w.WriteHeader(400)
				w.Write([]byte(err.Error()))
				return
			}

		} else {
			w.WriteHeader(400)
			w.Write([]byte("not enough money for sending"))
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("exchanged successfully!"))
	}
}
