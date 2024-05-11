package math

type API struct {
	equationService *EquationService
}

func New(repository EquationRepository) *API {
	return &API{
		equationService: NewEquationService(repository),
	}
}

