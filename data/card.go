package data

type Card struct {
	ID            string `gorm:"primaryKey,column:id"`
	Amount        int64  `gorm:"column:amount"`
	BlockedAmount int64  `gorm:"column:blocked_amount"`
}
