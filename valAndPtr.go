package gogb

type RealNumber interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func ToPtr[T any](in T) *T {
	return &in
}

// Return the underlying value or zero
func ToValOrZero[T any](in *T) T {
	if in == nil {
		var ret T
		return ret
	}
	return *in
}

func ToIntOrZero[T RealNumber](in *T) int {
	if in == nil {
		return 0
	}
	return int(*in)
}

func ToFloat32OrZero[T RealNumber](in *T) float32 {
	if in == nil {
		return 0
	}
	return float32(*in)
}

func ToFloat64OrZero[T RealNumber](in *T) float64 {
	if in == nil {
		return 0
	}
	return float64(*in)
}
