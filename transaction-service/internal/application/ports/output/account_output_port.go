package output

type AccountOutputPort interface {
	GetBalance(accountID int64) (float64, error)
}
