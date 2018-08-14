package data

var Cards map[string]Card

func init() {
	Cards = make(map[string]Card)
}

type Card struct {
	Amount Amount
	ID     string
}

type Amount struct {
	Value float32
}

func TopUp(cardID string, amount Amount) Amount {
	card := Cards[cardID]
	card.Amount.Value = card.Amount.Value + amount.Value

	return card.Amount
}
