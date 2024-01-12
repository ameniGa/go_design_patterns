package factory

type DebitCardPM struct {
}

func (dcp *DebitCardPM) Pay(amount float32) string {
	return "paid with debit card"
}
