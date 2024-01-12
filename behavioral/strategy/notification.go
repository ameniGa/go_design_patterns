package strategy

// this is the contract to implement
// each entity has different algorithm to implement this contract
type MessagingHandler interface {
	Send(msg string, to string)
}

/*****************************************/
// sms is an implementation of MessagingHandler
type sms struct{}

func newSms() *sms {
	return &sms{}
}

func (s *sms) Send(msg string, to string) {
	// buisness logic for messaging via sms
}

/********************************************/
// pushNotif is an implementation of MessagingHandler
type pushNotif struct{}

func newPushNotif() *pushNotif {
	return &pushNotif{}
}

func (s *pushNotif) Send(msg string, to string) {
	// buisness logic for messaging via push notification
}

/********************************************/

// CreateMessagingHandler returns an implementation of the contract for the given messaging type
func CreateMessagingHandler(mesType string) MessagingHandler {
	switch mesType {
	case "sms":
		return newSms()
	case "push":
		return newPushNotif()
	default:
		return nil
	}
}
