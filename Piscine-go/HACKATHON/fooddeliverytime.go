package piscine

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		return 15
	} else if order == "chips" {
		return 10
	} else if order == "nuggets" {
		return 12
	} else {
		return 404
	}
}
