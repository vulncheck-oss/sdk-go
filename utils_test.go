package sdk

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFormatCVE(t *testing.T) {
	t.Run("cve is formatted correctly", func(t *testing.T) {
		tests := []struct {
			name     string
			cve      string
			expected string
		}{
			{
				name:     "cve is already formatted correctly",
				cve:      "CVE-2024-38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly with lowercase",
				cve:      "cve-2024-38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly with mixed case",
				cve:      "cVe-2024-38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly without prefix",
				cve:      "2024-38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly if underscores are used instead of dashes",
				cve:      "CVE_2024_38077",
				expected: "CVE-2024-38077",
			},
			{
				name:     "cve is formatted correctly if special characters are used",
				cve:      "CVE-2024-38077!",
				expected: "CVE-2024-38077",
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				formattedCve := FormatCVE(test.cve)
				if formattedCve != test.expected {
					t.Errorf("expected %s, got %s", test.expected, formattedCve)
				}
			})
		}

	})
}

func TestHandleErrorResponse(t *testing.T) {
	metaError := MetaError{
		Error:  true,
		Errors: []string{"error1", "error2"},
	}
	metaErrorJSON, _ := json.Marshal(metaError)

	resp := httptest.NewRecorder()
	resp.WriteHeader(http.StatusBadRequest)
	resp.Body.Write(metaErrorJSON)

	err := handleErrorResponse(resp.Result())

	reqErr, ok := err.(ReqError)
	if !ok {
		t.Fatalf("expected ReqError, got %T", err)
	}

	if reqErr.StatusCode != http.StatusBadRequest {
		t.Errorf("expected status code %d, got %d", http.StatusBadRequest, reqErr.StatusCode)
	}

	expectedErrors := []string{"error1", "error2"}
	if len(reqErr.Reason.Errors) != len(expectedErrors) {
		t.Fatalf("expected %d errors, got %d", len(expectedErrors), len(reqErr.Reason.Errors))
	}
	for i, expectedError := range expectedErrors {
		if reqErr.Reason.Errors[i] != expectedError {
			t.Errorf("expected error %q, got %q", expectedError, reqErr.Reason.Errors[i])
		}
	}
}
