package card

import "bank/pkg/bank/types"

// Total func
func Total(card []types.Card) types.Money {
	sum := 0
	for _, operation := range card {
		if operation.Active && operation.Balance > 0 {
			sum += int(operation.Balance)
		}
	}
	return types.Money(sum)
}

func PaymentSources(cards []types.Card) []types.PaymentSource {

	var t []types.PaymentSource
	var p types.PaymentSource

	for _, operation := range cards {
		if operation.Active && operation.Balance > 0 {
			p.Balance = operation.Balance
			p.Number = string(operation.PAN)
			t = append(t, p)
		}
	}

	return t

}
