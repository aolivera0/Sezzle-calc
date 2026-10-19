package httpapi

import (
	"encoding/json"
	"net/http"
	"sezzle-calc/backend/internal/calculator"
)

type CalculateRequest struct {
	Operation string  `json:"operation"`
	A         float64 `json:"a"`
	B         float64 `json:"b"`
}

type CalculateResponse struct {
	Result float64 `json:"result"`
}

func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSONError(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var request CalculateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)

	if err != nil {
		writeJSONError(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}

	var result float64

	switch request.Operation {
	case "add":
		result = calculator.Add(request.A, request.B)
	case "subtract":
		result = calculator.Subtract(request.A, request.B)
	case "multiply":
		result = calculator.Multiply(request.A, request.B)
	case "divide":
		result, err = calculator.Divide(request.A, request.B)

		if err != nil {
			writeJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "power":
		result = calculator.Power(request.A, request.B)
	case "sqrt":
		result, err = calculator.Sqrt(request.A)

		if err != nil {
			writeJSONError(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "percentage":
		result = calculator.Percentage(request.A, request.B)
	default:
		writeJSONError(w, "Unsupported operation", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response := CalculateResponse{Result: result}
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		writeJSONError(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
