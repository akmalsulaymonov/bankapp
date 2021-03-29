package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {

	cards := []types.Card{
		{
			Balance: 5_000_00,
			Active:  true,
			PAN:     "5058 xxxx xxxx 6666",
		},
		{
			Balance: 10_000_00,
			Active:  true,
			PAN:     "5058 xxxx xxxx 7777",
		},
		{
			Balance: 10_000_00,
			Active:  false,
			PAN:     "5058 xxxx xxxx 8888",
		},
		{
			Balance: 0,
			Active:  true,
			PAN:     "5058 xxxx xxxx 9999",
		},
	}

	res := PaymentSources(cards)
	for _, item := range res {
		fmt.Println(item.Number)
	}

	//Output:
	//5058 xxxx xxxx 6666
	//5058 xxxx xxxx 7777
}
