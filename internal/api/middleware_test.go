package api_test

// func TestJWTMiddleware(t *testing.T) {
// 	var bearer = "Bearer " + "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MzEsInVzZXJuYW1lIjoidGVzdHVzZXIxIn0.7AmboDarHgOiIJ9c7jkzIgAs3d2p6S6ZR7I1l4jOWsI"
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	req.Header.Add("Authorization", bearer)
// 	rr := httptest.NewRecorder()

// 	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		claims := r.Context().Value("claims").(map[string]interface{})
// 		val, ok := claims["username"].(string)
// 		if val != "testuser1" && !ok {
// 			t.Errorf("There were no claims in request context %q", val)
// 		}
// 	})

// 	handler := api.MiddlewareController.JWTmiddleware(testHandler)
// 	handler.ServeHTTP(rr, req)

// }
