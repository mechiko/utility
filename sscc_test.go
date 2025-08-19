package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSsccTableDriven(t *testing.T) {
	// Defining the columns of the table
	var tests = []struct {
		name  string
		input string
		want  string
		err   bool
	}{
		// the table itself
		{"test 0", "", "", true},
		{"test 1", "46164463019900001", "00461644630199000016", false},
		{"test 2", "46164463019901", "00000461644630199016", false},
		{"test 3", "46164463A19901", "", true},
		{"test 4", "55546164463019900001", "00555461644630199004", false},
	}

	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans, err := Sscc(tt.input)
			if tt.err {
				assert.NotNil(t, err)
			} else {
				// ожидаем отсутствие ошибки
				assert.Nil(t, err)
			}
			// проверяем результат
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
