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
