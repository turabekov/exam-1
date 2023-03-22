package controller

import (
	"app/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func (c *Controller) ShopCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c.AddShopCartController(w, r)
	}
	if r.Method == "GET" {
		path := strings.Split(r.URL.Path, "/")

		if len(path) > 2 {
			c.GetUserShopCarts(w, r)
		} else {
			// c.store.User().GetList(w, r)
		}
	}
	// if r.Method == "PUT" {
	// 	c.store.User().UpdateUser(w, r)
	// }
	if r.Method == "DELETE" {
		c.RemoveShopCartController(w, r)
	}
}

func (c *Controller) AddShopCartController(w http.ResponseWriter, r *http.Request) {
	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}
	//  unmarshal data
	var shopcartBody models.AddShopCart
	err = json.Unmarshal(body, &shopcartBody)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	// check existing of current user
	res, err := http.Get("http://localhost:3000/user/" + shopcartBody.UserId)
	if err != nil || res.StatusCode >= 400 {
		log.Println("Get user err:", err)
		w.WriteHeader(400)
		w.Write([]byte("User not found"))
		return
	}
	var user models.User
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	// check existing of current product
	res, err = http.Get("http://localhost:3000/product/" + shopcartBody.ProductId)
	if err != nil || res.StatusCode >= 400 {
		log.Println("Get product err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Product not found"))
		return
	}
	var product models.Product
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(resBody, &product)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	// if all ok  then add shopcart to db
	_, err = c.store.ShopCart().AddShopCart(&shopcartBody)
	if err != nil {
		log.Println("Add shop cart err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("added successfully!"))
}

func (c *Controller) RemoveShopCartController(w http.ResponseWriter, r *http.Request) {
	//  read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Incorrect data"))
		return
	}
	//  unmarshal data
	var shopcartBody models.RemoveShopCart
	err = json.Unmarshal(body, &shopcartBody)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	// check existing of current user
	res, err := http.Get("http://localhost:3000/user/" + shopcartBody.UserId)
	if err != nil || res.StatusCode >= 400 {
		log.Println("Get user err:", err)
		w.WriteHeader(400)
		w.Write([]byte("User not found"))
		return
	}
	var user models.User
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	// check existing of current product
	res, err = http.Get("http://localhost:3000/product/" + shopcartBody.ProductId)
	if err != nil || res.StatusCode >= 400 {
		log.Println("Get product err:", err)
		w.WriteHeader(400)
		w.Write([]byte("Product not found"))
		return
	}
	var product models.Product
	resBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(resBody, &product)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	_, err = c.store.ShopCart().RemoveShopCart(&shopcartBody)
	if err != nil {
		log.Println("Add shop cart err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("removed successfully!"))
}

func (c *Controller) GetUserShopCarts(w http.ResponseWriter, r *http.Request) {
	// check existing of current user
	id := r.URL.Path[len("/shopcart/"):]

	res, err := http.Get("http://localhost:3000/user/" + id)
	if err != nil || res.StatusCode >= 400 {
		log.Println("Get user err:", err)
		w.WriteHeader(400)
		w.Write([]byte("User not found"))
		return
	}
	var user models.User
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ioutil err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	err = json.Unmarshal(resBody, &user)
	if err != nil {
		log.Println("Unmarshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	// find from db
	userShopCarts, e := c.store.ShopCart().GetUserShopCarts(&models.UserPrimaryKey{
		Id: id,
	})
	if e != nil {
		log.Println("get user shop cart err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	data, err := json.Marshal(userShopCarts)
	if err != nil {
		log.Println("marshal err:", err)
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(data)
}

// models.CalculateShop
// not working !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
// func (c *Controller) CalcTotalPrice(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {

// 		id := r.URL.Path[len("/totalshopcart/"):]

// 		// check existing of current user
// 		res, err := http.Get("http://localhost:3000/user/" + id)
// 		if err != nil || res.StatusCode >= 400 {
// 			log.Println("Get user err:", err)
// 			w.WriteHeader(400)
// 			w.Write([]byte("User not found"))
// 			return
// 		}
// 		var user models.User
// 		resBody, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			log.Println("ioutil err:", err)
// 			w.WriteHeader(400)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		err = json.Unmarshal(resBody, &user)
// 		if err != nil {
// 			log.Println("Unmarshal err:", err)
// 			w.WriteHeader(400)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		// get users shop carts
// 		res, err = http.Get("http://localhost:3000/shopcarts/" + id)
// 		if err != nil || res.StatusCode >= 400 {
// 			log.Println("Get user err:", err)
// 			w.WriteHeader(400)
// 			w.Write([]byte("User not found"))
// 			return
// 		}
// 		var shopCarts []models.ShopCart
// 		resBody, err = ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			log.Println("ioutil err:", err)
// 			w.WriteHeader(400)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}
// 		err = json.Unmarshal(resBody, &shopCarts)
// 		if err != nil {
// 			log.Println("Unmarshal err:", err)
// 			w.WriteHeader(400)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}

// 		// ------------------------------------------------------------------------------------------------------
// 		var total float64

// 		// for _, v := range shopCarts {
// 		// 	if !v.Status {
// 		// 		res, err = http.Get("http://localhost:3000/product/" + v.ProductId)
// 		// 		if err != nil || res.StatusCode >= 400 {
// 		// 			log.Println("Get product err:", err)
// 		// 			w.WriteHeader(400)
// 		// 			w.Write([]byte("Product not found"))
// 		// 			return
// 		// 		}
// 		// 		var p models.Product
// 		// 		resBody, err = ioutil.ReadAll(res.Body)
// 		// 		if err != nil {
// 		// 			log.Println("ioutil err:", err)
// 		// 			w.WriteHeader(400)
// 		// 			w.Write([]byte(err.Error()))
// 		// 			return
// 		// 		}
// 		// 		err = json.Unmarshal(resBody, &p)
// 		// 		if err != nil {
// 		// 			log.Println("Unmarshal err:", err)
// 		// 			w.WriteHeader(400)
// 		// 			w.Write([]byte(err.Error()))
// 		// 			return
// 		// 		}

// 		// 		switch reqBody.DiscountStatus {
// 		// 		case "fixed":
// 		// 			total += float64(v.Count * (p.Price - reqBody.Discount))
// 		// 		case "precent":
// 		// 			if 0 <= reqBody.Discount && reqBody.Discount <= 100 {
// 		// 				total += float64(v.Count) * (float64(p.Price) - (float64(p.Price) * (float64(reqBody.Discount) / 100)))
// 		// 			} else {
// 		// 				log.Println("Unmarshal err:", err)
// 		// 				w.WriteHeader(400)
// 		// 				w.Write([]byte("out of range precent value"))
// 		// 				return
// 		// 			}
// 		// 		default:
// 		// 			log.Println("Unmarshal err:", err)
// 		// 			w.WriteHeader(400)
// 		// 			w.Write([]byte("not allowed status"))
// 		// 			return
// 		// 		}
// 		// 	}
// 		// }

// 		dataTotal, err := json.Marshal(&total)
// 		if err != nil {
// 			log.Println("marsha; err:", err)
// 			w.WriteHeader(400)
// 			w.Write([]byte(err.Error()))
// 			return
// 		}

// 		w.WriteHeader(http.StatusCreated)
// 		w.Write(dataTotal)
// 	}
// }
