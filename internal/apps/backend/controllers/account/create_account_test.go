package account

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAccountController_Create(t *testing.T) {
	tsc := map[string]struct {
		*http.Request
		AccountController
		expectedStatusCode int
	}{
		"Given a valid non-existing account, a status code 201 is expected": {
			httptest.NewRequest(http.MethodPut, "/v1/account/create?id=94343721-6baa-4cd5-a0b4-6c5d0419c02d",
				strings.NewReader(`
						{
                        				"name": "Jared Nicolas V", 
                        				"last_name": "Mitchell",
							"email": "jared.gibson@gmail.com",
							"password": "7or2m27yw6zrkao",
                        				"details": {
                            					"permissions": [1, 2, 3]
                        				} 
						}`,
				)),
			NewAccountController(),
			http.StatusCreated,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			req := ts.Request
			resp := httptest.NewRecorder()

			ts.AccountController.Create(resp, req)

			if ts.expectedStatusCode != resp.Code {
				t.Errorf("status code was expected %d, but it was obtained %d", ts.expectedStatusCode, resp.Code)
			}
		})
	}
}
