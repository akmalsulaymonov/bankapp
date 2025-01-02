package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 40_000_00, Active: true}
	Withdraw(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: 4000000
}

func ExampleDeposit_incative() {
	card := types.Card{Balance: 40_000, Active: false}
	Deposit(&card, 10_000)
	fmt.Println(card.Balance)
	// Output: 40000
}

func ExampleDeposit_negative() {
	card := types.Card{Balance: -5_000, Active: true}
	Deposit(&card, 10_000)
	fmt.Println(card.Balance)
	// Output: 5000
}

func ExampleDeposit_max_limit() {
	card := types.Card{Balance: 1_000, Active: true}
	Deposit(&card, 70_000)
	fmt.Println(card.Balance)
	// Output: 1000
}

func ExampleBonus_add() {
	card := types.Card{Balance: 10_000, MinBalance: 10_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 12465
}

func ExampleBonus_inactive() {
	card := types.Card{Balance: 10_000, MinBalance: 10_000, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 10000
}

func ExampleBonus_negative() {
	card := types.Card{Balance: -10_000, MinBalance: 10_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: -10000
}

func ExampleBonus_limit() {
	card := types.Card{Balance: 10_000, MinBalance: 70_000, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	// Output: 15000
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 2_000_00,
			Active:  false,
		},
		{
			Balance: 5_000_00,
			Active:  true,
		},
		{
			Balance: -10_000_00,
			Active:  true,
		},
		{
			Balance: -5_000_00,
			Active:  true,
		},
	}

	fmt.Println(Total(cards))

	// Output: 1500000
}

func ExampleGetPANsFromPaymentSources() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
			PAN:     "4444 xxxx xxxx 0001",
		},
		{
			Balance: 2_000_00,
			Active:  false,
			PAN:     "4444 xxxx xxxx 0002",
		},
		{
			Balance: 5_000_00,
			Active:  true,
			PAN:     "4444 xxxx xxxx 0003",
		},
		{
			Balance: -10_000_00,
			Active:  true,
			PAN:     "4444 xxxx xxxx 0004",
		},
		{
			Balance: -5_000_00,
			Active:  true,
			PAN:     "4444 xxxx xxxx 0005",
		},
	}

	fmt.Println(PaymentSources(cards))

	// Output: [{card 4444 xxxx xxxx 0001 1000000} {card 4444 xxxx xxxx 0003 500000}]
}
