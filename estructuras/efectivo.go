package estructuras

import "fmt"

type Efectivo struct {
}

func NewEfectivo() *Efectivo {
	return &Efectivo{}
}
func (efectivo *Efectivo) Pagar() {
	fmt.Println("pagamos por efectivo")
}
