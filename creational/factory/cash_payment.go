package factory

type CashPM struct {
}

func (cp *CashPM) Pay(amount float32) string {
	return "paid with cash"
}
