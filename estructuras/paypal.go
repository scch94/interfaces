package estructuras

import "fmt"

type Paypal struct {
}

func NewPaypal() *Paypal {
	return &Paypal{}
}
func (p *Paypal) Pagar() {
	fmt.Println("pagamos por paypal")
}
