package database

import(

)

var(
	ErrCantFindProduct = errors.New("Can't find product")
	ErrCantDecodeProducts = errors.New("Can't decode products")
	ErrUserIdIsNotValid = 		errors.New("User id is not valid")
	ErrCantRemoveItemCart = 		errors.New("Can't remove item from cart")
	ErrCantGetItem = 		errors.New("Can't get item from cart")
	ErrCantBuyCartItem = 			errors.New("Can't buy cart item")

)

func AddProductToCart(){

}

func RemoveCartItem(){

}

func BuyItemFromCart(){

}

func InstantBuy(){

}

