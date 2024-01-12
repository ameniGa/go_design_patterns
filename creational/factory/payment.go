package factory

import "errors"

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(method int) (PaymentMethod, error) {
	switch method {
	case 1:
		return &CashPM{}, nil
	case 2:
		return &DebitCardPM{}, nil
	default:
		return nil, errors.New("unsupported payment method")
	}
}
