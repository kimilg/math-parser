package math

import (
	"github.com/gavv/httpexpect/v2"
	"math-parser/api/router/requestlog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParse(t *testing.T) {
	eqnString := "G_{nm}(\\xi_2, \\tau; \\xi_1, 0)=G_{mn}(\\xi_1, \\tau; \\xi_2, 0)"
	
	mathAPI := New()
	handler := requestlog.NewHandler(mathAPI.Parse)
	
	server := httptest.NewServer(handler)
	defer server.Close()
	
	e := httpexpect.Default(t, server.URL)
	
	eqn := map[string]interface{}{
		"equation": eqnString,
	}
	e.POST("/parse").WithJSON(eqn).
		Expect().
		Status(http.StatusOK)
}