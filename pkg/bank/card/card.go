package card

import (
	"bank/pkg/bank/types"
)

const withdrawLimit = 20_000_00
const maxAmount = 50_000
const bonusLimit = 5_000

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       123,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  50_000_00,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return card
}

func Withdraw(card *types.Card, amount types.Money) {
	if card.Active && card.Balance >= amount && amount > 0 && amount <= withdrawLimit {
		card.Balance -= amount
	}
}

func Deposit(card *types.Card, amount types.Money) {
	if (card.Active && amount <= maxAmount) || card.Balance < 0 {
		card.Balance += amount
	}
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if card.Active && card.Balance > 0 {
		bonus := types.Money(int(card.MinBalance) * percent * daysInMonth / daysInYear)

		if bonus >= bonusLimit {
			bonus = bonusLimit
		}

		card.Balance += bonus
	}
}

func Total(cards []types.Card) types.Money {
	var sum types.Money
	for _, item := range cards {
		if item.Active && item.Balance > 0 {
			sum += item.Balance
		}
	}
	return sum
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	sources := []types.PaymentSource{}
	for _, item := range cards {
		if item.Active && item.Balance > 0 {
			sources = append(sources, types.PaymentSource{
				Type:    "card",
				Number:  item.PAN,
				Balance: item.Balance,
			})
		}
	}
	return sources
}
