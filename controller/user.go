package controller

import (
	"app/models"
	"errors"
	"fmt"
)

func (c *Controller) CreateUser(req *models.CreateUser) (id string, err error) {

	id, err = c.store.User().Create(req)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Controller) GetList(req *models.GetListRequest) (*models.GetListResponse, error) {

	users, err := c.store.User().GetList(req)
	if err != nil {
		return &models.GetListResponse{}, err
	}

	return users, nil
}

func (c *Controller) GetUserByIdController(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User().GetUserById(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (c *Controller) UpdateUserController(req *models.UpdateUser) (models.User, error) {
	user, err := c.store.User().UpdateUser(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
func (c *Controller) DeleteUserController(req *models.UserPrimaryKey) (models.User, error) {
	user, err := c.store.User().DeleteUser(req)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Task 10 Agar client 9 dan katta mahuslot sotib olgan bolsa,
// 1 tasi tekinga beriladi va 9 ta uchun pul hisoblanadi.
// 1 tasi eng arzon mahsulotni pulini hisoblamaysiz.
// Yangi korzinka qoshib tekshirib koring.
func (c *Controller) WithdrawUserBalance(id string) error {

	user, err := c.store.User().GetUserById(&models.UserPrimaryKey{Id: id})
	if err != nil {
		return err
	}

	// Task10
	shopCarts, err := c.store.ShopCart().GetUserShopCarts(&models.UserPrimaryKey{
		Id: user.Id,
	})
	if err != nil {
		return err
	}

	userShopcarts := []models.ShopCart{}
	counter := 0
	for _, v := range shopCarts {
		if !v.Status {
			userShopcarts = append(userShopcarts, v)
			counter += v.Count
		}
	}

	totalBalance := 0
	// calc total
	for _, v := range userShopcarts {
		product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
			Id: v.ProductId,
		})
		if err != nil {
			return err
		}
		totalBalance += v.Count * product.Price
	}

	if counter > 9 {
		// declaring minimum price
		product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
			Id: userShopcarts[0].ProductId,
		})
		if err != nil {
			return err
		}
		dicountPrice := product.Price
		fmt.Println(dicountPrice)
		for _, v := range userShopcarts {
			product, err := c.store.Product().GetProductById(&models.ProductPrimaryKey{
				Id: v.ProductId,
			})
			if err != nil {
				return err
			}
			if product.Price < dicountPrice {
				dicountPrice = product.Price
			}
		}
		totalBalance = totalBalance - dicountPrice
	}

	newBalance := user.Balance - float64(totalBalance)
	if newBalance <= 0 {
		return errors.New("not enough money")
	}

	_, err = c.store.User().UpdateUser(&models.UpdateUser{
		Id:      user.Id,
		Name:    user.Name,
		Surname: user.Surname,
		Balance: newBalance,
	})
	if err != nil {
		return err
	}

	p, e := c.store.ShopCart().GetUserShopCarts(&models.UserPrimaryKey{
		Id: user.Id,
	})
	if e != nil {
		return e
	}

	for _, v := range p {
		_, err := c.store.ShopCart().UpdateShopCart(models.ShopCart{
			ProductId: v.ProductId,
			UserId:    v.UserId,
			Count:     v.Count,
			Status:    true,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// Homework
func (c *Controller) ExchangeMoney(req models.ReqExchangeMoney) error {
	sender, err := c.store.User().GetUserById(&models.UserPrimaryKey{
		Id: req.SenderId,
	})
	if err != nil {
		return err
	}
	receiver, err := c.store.User().GetUserById(&models.UserPrimaryKey{
		Id: req.ReceiverId,
	})
	if err != nil {
		return err
	}

	komissiya, err := c.store.Komissiya().GetKomissiya()
	if err != nil {
		return err
	}

	moneyWithKomissiya := float64(req.Amount) + (float64(req.Amount) * komissiya.Amount / 100)

	if sender.Balance > moneyWithKomissiya {
		_, err = c.store.User().UpdateUser(&models.UpdateUser{
			Id:      sender.Id,
			Name:    sender.Name,
			Surname: sender.Surname,
			Balance: sender.Balance - moneyWithKomissiya,
		})
		if err != nil {
			return err
		}
		_, err = c.store.User().UpdateUser(&models.UpdateUser{
			Id:      receiver.Id,
			Name:    receiver.Name,
			Surname: receiver.Surname,
			Balance: receiver.Balance + float64(req.Amount),
		})
		if err != nil {
			return err
		}
		err := c.store.Komissiya().UpdateBalanceKomissiya(models.Komissiya{
			Balance: komissiya.Balance + (float64(req.Amount) * komissiya.Amount / 100),
			Amount:  komissiya.Amount,
		})
		if err != nil {
			return nil
		}

	} else {
		return errors.New("not enough money for sending")
	}

	return nil
}
