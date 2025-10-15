# Calculate the value of a shopping cart

## Solution

func calculateTotal(cart []CartItem) float64 {
    var total float64 = 0
    for _, item := range cart {
        total += (item.Price * float64(item.Quantity))
    }
    return total
}

### Explanation

The solution’s calculateTotal() function receives one parameter typed as a slice of CartItem instances. The CartItem struct is already defined with three fields for the item’s name, price, and quantity.

type CartItem struct{
    Name string
    Price float64
    Quantity int
}

This solution starts by setting a variable named total, typed as a float64 value. 

var total float64 = 0

It then iterates through the slice of CartItem instances. For each item, it multiplies the item’s Quantity times Price, and then adds the result to the total. Note that the Quantity is converted to float64 before the math operation, so that all values are of the same type.

var total float64 = 0
for _, item := range cart {
    total += (item.Price * float64(item.Quantity))
}
return total

When you add 2 values of the same type, the result will always be of that type. Since the function declares its return as a float64 value, the result’s type satisfies that requirement.