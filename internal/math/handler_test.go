package math

import (
	"context"
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"math-parser/api/router/requestlog"
	"math-parser/internal/math/domain/formula"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	const nihongo = "안녕하세요"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}


type mockEquationRepository struct {}

func (r *mockEquationRepository) Create(ctx context.Context, value string) (*formula.Equation, error) {
	return &formula.Equation{
		Id: 0,
		Value: value,
	}, nil
}
func (r *mockEquationRepository) Get(ctx context.Context, id formula.ID) (*formula.Equation, error) {
	return nil, nil
}
func (r *mockEquationRepository) GetFromValue(ctx context.Context, value string) (*formula.Equation, error) {
	return nil, nil
}
func (r *mockEquationRepository) List(ctx context.Context) ([]*formula.Equation, error) {
	return nil, nil
}
func (r *mockEquationRepository) Update(ctx context.Context, id formula.ID, value string) (*formula.Equation, error) {
	return nil, nil
}
func (r *mockEquationRepository) Delete(ctx context.Context, id formula.ID) error {
	return nil
}

func TestParse(t *testing.T) {
	equation := "G_{nm}(\\xi_2, \\tau; \\xi_1, 0)=G_{mn}(\\xi_1, \\tau; \\xi_2, 0)"
	
	repo := &mockEquationRepository{}
	mathAPI := New(repo)
	handler := requestlog.NewHandler(mathAPI.Parse)
	
	server := httptest.NewServer(handler)
	defer server.Close()
	
	e := httpexpect.Default(t, server.URL)
	
	eqn := map[string]interface{} {
		"equation": equation,
	}
	e.POST("/parse").WithJSON(eqn).
		Expect().
		Status(http.StatusOK)
}