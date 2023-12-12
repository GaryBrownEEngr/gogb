package gogb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetAge(t *testing.T) {
	got, err := GetAge(DateUtc(2000, 1, 2), DateUtc(2001, 1, 1))
	require.NoError(t, err)
	require.Equal(t, 0, got)
	got, err = GetAge(DateUtc(2000, 1, 2), DateUtc(2001, 1, 2))
	require.NoError(t, err)
	require.Equal(t, 1, got)

	got, err = GetAge(DateUtc(2000, 1, 2), DateUtc(2000, 1, 1))
	require.Error(t, err)
	require.Empty(t, got)
}

func TestTimePtr(t *testing.T) {
	t1 := TimePtr(2000, 12, 28, 23, 59, 59, 9000, time.UTC)
	require.NotEmpty(t, t1)
	require.Equal(t, 2000, t1.Year())
	require.Equal(t, 23, t1.Hour())
	require.Equal(t, 9000, t1.Nanosecond())

	t1 = DateUtcPtr(2001, 1, 2)
	require.NotEmpty(t, t1)
	require.Equal(t, 2001, t1.Year())
	require.Equal(t, time.Month(1), t1.Month())
	require.Equal(t, 2, t1.Day())
}
