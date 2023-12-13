package gogb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToPtr(t *testing.T) {
	intPtr := ToPtr(1)
	require.NotNil(t, intPtr)
	require.Equal(t, 1, *intPtr)

	intVal := ToValOrZero(intPtr)
	require.Equal(t, 1, intVal)
	intPtr = nil
	intVal = ToValOrZero(intPtr)
	require.Empty(t, intVal)

	intVal = ToIntOrZero(ToPtr(10.1))
	require.Equal(t, 10, intVal)
	var f32Ptr *float32
	intVal = ToIntOrZero(f32Ptr)
	require.Empty(t, intVal)

	f32Val := ToFloat32OrZero(ToPtr(10.1))
	require.Equal(t, float32(10.1), f32Val)
	f32Val = ToFloat32OrZero(f32Ptr)
	require.Empty(t, f32Val)

	f64Val := ToFloat64OrZero(ToPtr(10))
	require.Equal(t, 10.0, f64Val)
	f64Val = ToFloat64OrZero(f32Ptr)
	require.Empty(t, f64Val)

	var stringPtr *string
	require.Empty(t, ToValOrZero(stringPtr))
	stringPtr = ToPtr("abc")
	require.Equal(t, "abc", ToValOrZero(stringPtr))
}

func TestFirstNotZero(t *testing.T) {
	require.Equal(t, 0, FirstNotZero(0))
	require.Equal(t, 0, FirstNotZero(0, 0, 0))
	require.Equal(t, 1, FirstNotZero(0, 1, 0))
	require.Equal(t, "", FirstNotZero(""))
	require.Equal(t, "", FirstNotZero("", "", ""))
	require.Equal(t, "a", FirstNotZero("", "a", ""))

	var f64Ptr *float64
	require.Nil(t, FirstNotZero(nil, f64Ptr, nil, nil))
	f64Ptr = ToPtr(1.1)
	require.Equal(t, ToPtr(1.1), FirstNotZero(nil, f64Ptr, nil, nil))
}
