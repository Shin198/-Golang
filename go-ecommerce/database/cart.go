package database

import (
	"errors"
)

var (
	ErrCantFindProduct   = errors.New("Can't find product")
	ErrCantDecodeProduct = errors.New("Can't find product")
	ErrUserNotValid      = errors.New("User is not valid")
	ErrCantUpdateUser    = errors.New("Can")
	ErrCantRemoveItem    = errors.New("Can't remove this item from cart")
	ErrCantGetItem       = errors.New("Unable to get item in cart")
	ErrCantBuyItem       = errors.New("Can't process the request")
)

func AddToCart() {

}

func RemoveFromCart() {

}

func BuyFromCart() {

}

func BuyNow() {

}
