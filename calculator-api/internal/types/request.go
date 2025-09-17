package types

import (
	"fmt"
	"strings"
)

// CalcRequest models the incoming JSON for /calculate.
type CalcRequest struct {
	Op string  `json:"op"`
	A  float64 `json:"a"`
	B  float64 `json:"b"`
}

// ErrorBody is the JSON shape used for error responses.
type ErrorBody struct {
	Code    string `json:"code"`    // machine-readable error code
	Message string `json:"message"` // human-friendly message
	Field   string `json:"field,omitempty"`
}

// AllowedOps returns the canonical list of supported operations.
func AllowedOps() []string { return []string{"add", "sub", "mul", "div"} }

// NormalizeOp lowercases and trims an op string.
func NormalizeOp(op string) string { return strings.ToLower(strings.TrimSpace(op)) }

// Validate checks the semantic rules for CalcRequest and returns an *ErrorBody
// when validation fails, or nil when the request is valid.
// - ensures op is one of allowed ops
// - ensures division by zero is rejected
func (c *CalcRequest) Validate() *ErrorBody {
	op := NormalizeOp(c.Op)
	ok := false
	for _, a := range AllowedOps() {
		if op == a {
			ok = true
			break
		}
	}
	if !ok {
		return &ErrorBody{
			Code:    "invalid_op",
			Message: fmt.Sprintf("invalid op; must be one of %v", AllowedOps()),
			Field:   "op",
		}
	}

	if op == "div" && c.B == 0 {
		return &ErrorBody{
			Code:    "division_by_zero",
			Message: "division by zero is not allowed",
			Field:   "b",
		}
	}

	return nil
}
