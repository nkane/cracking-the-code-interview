package build_order

import (
	"testing"
)

func TestFindBuildOrder(t *testing.T) {
	tests := []struct {
		projects     []string
		dependencies [][]string
		expected     []string
		expectError  bool
	}{
		{
			projects:     []string{"a", "b", "c", "d", "e", "f"},
			dependencies: [][]string{{"a", "d"}, {"f", "b"}, {"b", "d"}, {"f", "a"}, {"d", "c"}},
			expected:     []string{"f", "e", "a", "b", "d", "c"},
			expectError:  false,
		},
		{
			projects:     []string{"a", "b", "c"},
			dependencies: [][]string{{"a", "b"}, {"b", "c"}},
			expected:     []string{"a", "b", "c"},
			expectError:  false,
		},
		{
			projects:     []string{"a", "b", "c", "d"},
			dependencies: [][]string{{"a", "b"}, {"b", "c"}, {"c", "a"}},
			expected:     nil,
			expectError:  true,
		},
		{
			projects:     []string{"a"},
			dependencies: [][]string{},
			expected:     []string{"a"},
			expectError:  false,
		},
	}

	for _, test := range tests {
		order, err := FindBuildOrder(test.projects, test.dependencies)
		if test.expectError {
			if err == nil {
				t.Errorf("expected an error but got nil")
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error but got %v", err)
			}
			if !IsValidBuildOrder(order, test.dependencies) {
				t.Errorf("order %v does not respect dependencies %v", order, test.dependencies)
			}
		}
	}
}

func IsValidBuildOrder(order []string, dependencies [][]string) bool {
	position := make(map[string]int)
	for i, project := range order {
		position[project] = i
	}
	for _, dep := range dependencies {
		first, second := dep[0], dep[1]
		if position[first] > position[second] {
			return false
		}
	}
	return true
}
