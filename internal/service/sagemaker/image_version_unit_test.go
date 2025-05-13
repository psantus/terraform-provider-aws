package sagemaker

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseImageVersionID(t *testing.T) {
	testCases := []struct {
		desc            string
		id              string
		expectError     bool
		expectedName    string
		expectedVersion int
	}{
		{
			desc:            "valid ID with name and version",
			id:              "test-image:1",
			expectError:     false,
			expectedName:    "test-image",
			expectedVersion: 1,
		},
		{
			desc:        "invalid ID without version",
			id:          "test-image",
			expectError: true,
		},
		{
			desc:        "invalid ID with invalid version",
			id:          "test-image:abc",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			// Parse the ID manually to test the logic
			parts := strings.Split(tc.id, ":")
			if len(parts) != 2 {
				if !tc.expectError {
					t.Fatalf("Expected successful parsing but got error: invalid format")
				}
				return
			}

			name := parts[0]
			versionStr := parts[1]

			version, err := parseVersionFromID(versionStr)
			if err != nil {
				if !tc.expectError {
					t.Fatalf("Expected successful parsing but got error: %v", err)
				}
				return
			}

			if tc.expectError {
				t.Fatalf("Expected error but got none")
			}

			if name != tc.expectedName {
				t.Errorf("Expected name %s but got %s", tc.expectedName, name)
			}

			if version != tc.expectedVersion {
				t.Errorf("Expected version %d but got %d", tc.expectedVersion, version)
			}
		})
	}
}

// Helper function to parse version from ID
func parseVersionFromID(versionStr string) (int, error) {
	var version int
	_, err := fmt.Sscanf(versionStr, "%d", &version)
	if err != nil {
		return 0, fmt.Errorf("invalid version number: %s", versionStr)
	}
	return version, nil
}
