package main

import (
	"fmt"
	"github.com/bradrydzewski/go.stripe"
)

func main() {
	stripe.SetKey("SECRET")

	customers, err := stripe.Customers.List()

	for _, c := range customers {
		fmt.Println(c, err)
	}

	paramsCharge := stripe.ChargeParams{
		Desc:     "Calzone",
		Amount:   400,
		Currency: "usd",
		Card: &stripe.CardParams{
			Name:     "George Costanza",
			Number:   "4242424242424242",
			ExpYear:  2015,
			ExpMonth: 5,
			CVC:      "226",
		},
	}

	charge, err := stripe.Charges.Create(&paramsCharge)

	fmt.Println(charge, err)
}
