package estructuras

import "fmt"

type Tarjeta struct {
}

func NewTarjeta() *Tarjeta {
	return &Tarjeta{}
}

func (t *Tarjeta) Pagar() {
	fmt.Println("pagamos con targeta")
}
