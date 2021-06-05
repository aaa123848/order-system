package domain

import "time"

type Transaction struct {
	CreateTime time.Time
	Price      uint
}

// type TransactionRepository interface {
// 	GetAllTransaction() [][]byte
// 	CreateTransaction(price uint) error
// }

// type TransactionUseCase interface {
// 	GetAllTransaction() [][]byte
// 	CreateTransaction(price uint)
// }

// type TransactionRepositoryOnlyCreate interface {
// 	CreateTransaction(price uint)
// }

// type TransactionUseCaseOnlyCreate interface {
// 	CreateTransaction(price uint)
// }
