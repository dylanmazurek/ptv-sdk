package ptv

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dylanmazurek/ptv-sdk/pkg/ptv/models"
	"github.com/stretchr/testify/assert"
)

func TestClient_RouteTypes(t *testing.T) {
	tests := []struct {
		name          string
		handlerFunc   func(w http.ResponseWriter, r *http.Request)
		expectedTypes []models.RouteType
		expectedError string
	}{
		{
			name: "successful response",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, `{
					"route_types": [
						{"route_type_name": "Train", "route_type": 0},
						{"route_type_name": "Tram", "route_type": 1}
					],
					"status": {"version": "3.0", "health": 1}
				}`)
			},
			expectedTypes: []models.RouteType{
				{RouteTypeName: "Train", RouteType: 0},
				{RouteTypeName: "Tram", RouteType: 1},
			},
		},
		{
			name: "API error response",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, `{"message": "Internal Server Error", "status": {"version": "3.0", "health": 0}}`)
			},
			expectedError: "failed to execute request: unexpected status code: 500 Internal Server Error",
		},
		{
			name: "invalid JSON response",
			handlerFunc: func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintln(w, `{"route_types": [`)
			},
			expectedError: "failed to execute request: unexpected end of JSON input",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.handlerFunc))
			defer server.Close()

			clientOpts := []Option{
				WithBaseURL(server.URL),
			}

			client := New(t.Context(), clientOpts...)

			routeTypes, err := client.RouteTypes()
			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTypes, routeTypes)
			}
		})
	}
}
