package utility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testsParse = []struct {
	name   string
	err    bool
	prefix string
	i      int
	result string
}{
	// the table itself
	{"test 0", true, "123456", 0, ""},
	{"test 1", false, "1234567", 0, "00123456700000000000"},
	{"test 3", false, "1234567890", 0, "00123456789000000005"},
	{"test 4", false, "1234567890", 1234567, "00123456789012345675"},
	{"test 5", true, "1234567890", 1234567890, "00123456789001234560"},
	{"test 6", false, "123456789", 12345, "00123456789000123452"},
}

func TestGenerateSSCC(t *testing.T) {
	// The execution loop
	for _, tt := range testsParse {
		t.Run(tt.name, func(t *testing.T) {
			sscc, err := GenerateSSCC(tt.i, tt.prefix)
			if tt.err {
				assert.NotNil(t, err, "ожидаем ошибку")
			} else {
				// ожидаем отсутствие ошибки
				assert.Nil(t, err)
				assert.Equal(t, sscc, tt.result, "ожидаемое значение")
			}
		})
	}
}
