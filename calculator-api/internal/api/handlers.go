package api

import (
	"calculator-api/internal/calc"
	"calculator-api/internal/types"
	"encoding/json"
	"net/http"
)

// CalcResponse is the JSON shape for successful results.
type CalcResponse struct {
	Result float64 `json:"result"`
}

// CalculateHandler handles POST /calculate requests.
func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request body into CalcRequest
	var req types.CalcRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.ErrorBody{
			Code:    "invalid_json",
			Message: "invalid JSON body",
		})
		return
	}

	// Validate semantic rules
	if vErr := req.Validate(); vErr != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(vErr)
		return
	}

	// Normalize op and run calculation
	op := types.NormalizeOp(req.Op)
	res, err := calc.Do(op, req.A, req.B)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.ErrorBody{
			Code:    "calc_error",
			Message: err.Error(),
		})
		return
	}

	// Success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CalcResponse{Result: res})
}
