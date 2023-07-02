package valueobject_test

import (
	"testing"
	"time"

	"github.com/oka311119/go-hexagonal-arch/internal/domain/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestNewDateRange(t *testing.T) {
	date1 := time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2023, 7, 3, 0, 0, 0, 0, time.UTC)

	// Test for correct creation of DateRange
	dateRange, err := valueobject.NewDateRange(date1, date2)

	assert.NoError(t, err)
	assert.Equal(t, date1, dateRange.GetDate1())
	assert.Equal(t, date2, dateRange.GetDate2())

	// Test for error when date1 is after date2
	_, err = valueobject.NewDateRange(date2, date1)
	assert.Error(t, err)
}
