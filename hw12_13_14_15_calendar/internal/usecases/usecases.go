package usecases

type (
	UseCases struct{}
)

func New() *UseCases {
	usecases := &UseCases{}
	return usecases
}
