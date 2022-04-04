package shipping

import (
	"ex04/coubang"
)

type Sender interface {
	Send()
}

func shipping(sender Sender) {
	switch sender := sender.(type) {
	default:
		sender.Send()
	case *coubang.Coubang:
		sender.SendRocket()
	}

}
