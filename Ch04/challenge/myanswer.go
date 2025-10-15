package main

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
	var sum float64
	for _, cartItem := range cart {
		itemTotal := cartItem.Price * float64(cartItem.Quantity)
		sum += itemTotal
	}

	return sum
}
