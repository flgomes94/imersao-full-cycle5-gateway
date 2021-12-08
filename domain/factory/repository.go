package factory

import "github.com/flgomes94/imersao-full-cycle5-gateway/domain/repository"

type RepositoryFactory interface {
	CreateTransactionRepository() repository.TransactionRepository
}
