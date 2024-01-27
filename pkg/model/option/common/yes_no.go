package common

type YesNo int32

const (
	YesNo__UNKNOWN YesNo = iota
	YesNo__YES
	YesNo__NO
)

func YesNoAndOperator(a, b YesNo) YesNo {
	if a == b {
		return a
	}
	return YesNo__NO
}
