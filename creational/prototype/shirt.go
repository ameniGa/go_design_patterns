package prototype

import "fmt"

type shirtColor byte

type shirt struct {
	Price float32
	SKU   string
	Color shirtColor
}

func (sh *shirt) GetInfo() string {
	return fmt.Sprintf("shirt of Color: %v, SKU: %s, Price: %v", sh.Color, sh.SKU, sh.Price)
}

func (sh *shirt) GetPrice() float32 {
	return sh.Price
}
