package calc

import (
	"testing"
)

func TestDo(t *testing.T) {
	test := []struct {
		name    string
		op      string
		a, b    float64
		want    float64
		wantErr bool
	}{
		{"add ints", "add", 2, 4, 6, false},
		{"sub ints", "sub", 4, 2, 2, false},
		{"mul ints", "mul", 2, 4, 8, false},
		{"div ints", "div", 8, 4, 2, false},
		{"div zero", "div", 2, 0, 0, true},
		{"Unknown operation", "pow", 2, 4, 0, true},
	}
	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Do(tc.op, tc.a, tc.b)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Do(%q,%v,%v) error = %v, wantErr=%v", tc.op, tc.a, tc.b, err, tc.wantErr)
			}
			if !tc.wantErr && got != tc.want {
				t.Fatalf("Do(%q,%v,%v) = %v, want %v", tc.op, tc.a, tc.b, got, tc.want)
			}
		})
	}
}
