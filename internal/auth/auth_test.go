package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input http.Header
		out1  string
		out2  error
	}{
		"simple":         {input: http.Header{"Authorization": []string{"ApiKey abc123"}}, out1: "abc123", out2: nil},
		"no header":      {input: http.Header{}, out1: "", out2: ErrNoAuthHeaderIncluded},
		"invalid format": {input: http.Header{"Authorization": []string{"Bearer abc123"}}, out1: "", out2: ErrMalformedAuthHeader},
		"test failure":   {input: http.Header{"Authorization": []string{"ApiKey xyz357"}}, out1: "abc123", out2: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got1, got2 := GetAPIKey(tc.input)
			if tc.out2 != nil && !reflect.DeepEqual(tc.out2.Error(), got2.Error()) {
				t.Fatalf("expected: %v, got: %v", tc.out2, got2)
			} else if tc.out1 != "" && !reflect.DeepEqual(tc.out1, got1) {
				t.Fatalf("expected: %v, got: %v", tc.out1, got1)
			}
		})
	}
}
