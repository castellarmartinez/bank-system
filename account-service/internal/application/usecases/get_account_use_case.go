package usecases

type GetAccountBalanceUseCase interface {
	Execute(id int64) (float64, error)
}
