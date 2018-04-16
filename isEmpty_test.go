package isEmpty

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestIsEmpty(t *testing.T) {
	var (
		logger    logrus.FieldLogger = logrus.New()
		str       interface{}        = "asdsd"
		integer   interface{}        = 1234
		sliceStr  interface{}        = []string{"a", "b"}
		sliceByte interface{}        = []byte{1, 2, 3, 4}
		sliceInt  interface{}        = []int{1, 2, 3, 4}

		emptyLogger logrus.FieldLogger
	)
	tests := []struct {
		name   string
		values []interface{}
		want   bool
	}{
		{"Not empty logger", []interface{}{logger}, false},
		{"Not empty string", []interface{}{str}, false},
		{"Not empty int", []interface{}{integer}, false},
		{"Not empty []string", []interface{}{sliceStr}, false},
		{"Not empty []byte", []interface{}{sliceByte}, false},
		{"Not empty []int", []interface{}{sliceInt}, false},
		{"Not empty all", []interface{}{logger, str, sliceByte, sliceInt, sliceStr, integer}, false},

		{"Empty logger", []interface{}{emptyLogger}, true},
		{"Empty string", []interface{}{logger, str, sliceByte, sliceInt, sliceStr, integer, ""}, true},
		{"Empty int", []interface{}{logger, str, sliceByte, sliceInt, sliceStr, integer, 0}, true},
		{"Empty interface{}", []interface{}{logger, str, sliceByte, sliceInt, sliceStr, integer, nil}, true},
		{"Empty []string", []interface{}{logger, str, sliceByte, sliceInt, sliceStr, integer, []string{}}, true},
		{"Empty []int", []interface{}{logger, str, sliceByte, sliceInt, sliceStr, integer, []int{}}, true},
		{"Empty []byte", []interface{}{logger, str, sliceByte, sliceInt, sliceStr, integer, []byte{}}, true},
		{"All empty", []interface{}{emptyLogger, "", []byte{}, []int{}, []string{}, 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := Values(tt.values...); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
