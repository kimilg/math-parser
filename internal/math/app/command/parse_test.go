package command

import (
	"context"
	"math-parser/internal/math/domain/formula"
	"math-parser/internal/math/domain/formula/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	//given
	mockRepository := mocks.NewMockRepository(t)

	mockRepository.On("GetFromValue", mock.Anything, mock.Anything).
		Return(&formula.Equation{
			Id:       1,
			Value:    "2=1+1",
			Category: "mock_category",
		}, nil)

	mockRepository.On("Update", mock.Anything, mock.Anything).
		Return(&formula.Equation{
			Id:       1,
			Value:    "2=1+1",
			Category: "mock_category",
		}, nil)

	//when
	equationMemory := formula.NewEquationMemory(mockRepository)

	parseHandler := NewParseHandler(mockRepository, equationMemory)
	ctx := context.Background()

	parseHandler.Handle(ctx, Parse{Equation: &formula.Equation{}})

	//then
	resultExpression, err := equationMemory.Get(1)
	require.NoError(t, err)
	assert.Equal(t, formula.ID(1), resultExpression.EquationId)
	assert.Equal(t, "mock_category", resultExpression.Category)
	assert.Equal(t, 2, int(getLeftFirstConstant(resultExpression, 0, 0)))
}

func getLeftFirstConstant(expression *formula.Expression, i int, j int) float64 {
	return ((expression.Elements[i].(*formula.Expression)).Elements[j].(*formula.Constant)).Value
}
