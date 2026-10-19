package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculateHandler_Add(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodPost,
		"/api/calculate",
		strings.NewReader(`{"operation":"add","a":10,"b":5}`),
	)

	response := httptest.NewRecorder()

	CalculateHandler(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	if !strings.Contains(response.Body.String(), `"result":15`) {
		t.Fatalf("expected response %s, got %s", `"result":15`, response.Body.String())
	}
}
func TestCalculateHandler_UnsupportedOperation(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodPost,
		"/api/calculate",
		strings.NewReader(`
		{"operation":"log","a":10,"b":2}`),
	)

	response := httptest.NewRecorder()

	CalculateHandler(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}
}

func TestCalculateHandler_DivisionByZero(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/api/calculate", strings.NewReader(`{"operation":"divide","a":5,"b":0}`))

	response := httptest.NewRecorder()

	CalculateHandler(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}
}

func TestSqrtOfNegativeNumber(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/api/calculate", strings.NewReader(`{"operation":"sqrt","a":-3,"b":0}`))

	response := httptest.NewRecorder()

	CalculateHandler(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, response.Code)
	}
}

func TestCalculateHandler_InvalidJSON(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "/api/calculate", strings.NewReader(`{"operation":"add","a":1,"b":4`))

	response := httptest.NewRecorder()

	CalculateHandler(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf(`expected status %d, got %d`, http.StatusBadRequest, response.Code)
	}
}

func TestCalculateHandler_MethodNotAllowed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/api/calculate", nil)

	response := httptest.NewRecorder()

	CalculateHandler(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf(`expected status %d, got %d`, http.StatusMethodNotAllowed, response.Code)
	}
}

func TestCalculateHandler_Operations(t *testing.T) {
	tests := []struct {
		name     string
		body     string
		expected float64
	}{
		{
			name:     "subtraction",
			body:     `{"operation":"subtract","a":10,"b":5}`,
			expected: 5,
		}, {
			name:     "multiplication",
			body:     `{"operation":"multiply", "a":10,"b":5}`,
			expected: 50,
		}, {
			name:     "division",
			body:     `{"operation":"divide","a":10,"b":5}`,
			expected: 2,
		}, {
			name:     "power",
			body:     `{"operation":"power","a":4,"b":2}`,
			expected: 16,
		}, {
			name:     "sqrt",
			body:     `{"operation":"sqrt","a":4,"b":0}`,
			expected: 2,
		}, {
			name:     "percentage",
			body:     `{"operation":"percentage","a":10,"b":20}`,
			expected: 2,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/api/calculate", strings.NewReader(test.body))

			response := httptest.NewRecorder()

			CalculateHandler(response, request)

			var responseBody CalculateResponse
			err := json.NewDecoder(response.Body).Decode(&responseBody)
			if err != nil {
				t.Fatalf("failed to decode response %v", err)
			}

			if responseBody.Result != test.expected {
				t.Fatalf(`expected result %v, got %v`, test.expected, responseBody.Result)
			}
		})
	}
}

func TestCalculateHandler_EmptyBody(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodPost,
		"/api/calculate",
		strings.NewReader(""),
	)
	response := httptest.NewRecorder()

	CalculateHandler(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d",
			http.StatusBadRequest, response.Code)
	}
}
