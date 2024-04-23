package main

import (
	"errors"
	"fmt"

	"github.com/scch94/interfaces.git/estructuras"
)

type MetodoPago interface {
	Pagar()
}

func getMetodo(metodo string) (MetodoPago, error) {
	switch metodo {
	case "targeta":
		return estructuras.NewTarjeta(), nil
	case "efectivo":
		return estructuras.NewEfectivo(), nil
	case "paypal":
		return estructuras.NewPaypal(), nil
	default:
		return nil, errors.New("metodo de pago no soportado")
	}

}

func main() {
	var metodo string = "1231"
	metododepago, err := getMetodo(metodo)
	if err != nil {
		fmt.Println(err)
		return
	}
	metododepago.Pagar()
}
