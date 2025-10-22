package auth

import (
	"net/http"
	"strings"
	"testing"
)

// TestGetAPIKey -
func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		key       string
		value     string
		expect    string
		expectErr string
	}{
		"correct":                    {key: "Authorization", value: "ApiKey CorrectKey", expect: "CorrectKey", expectErr: "not expecting an error"},
		"empty header":               {expectErr: "no authorization header"},
		"empty authorization header": {key: "Authorization", expectErr: "no authorization header"},
		"malformed":                  {key: "Authorization", value: "ApiKey-CorrectKey", expectErr: "malformed authorization header"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if strings.Contains(err.Error(), test.expectErr) {
					return
				}
				t.Errorf("Unexpected error:%v\n", err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected:%s", output)
				return
			}
		})
	}
}
