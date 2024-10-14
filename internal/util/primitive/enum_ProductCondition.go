package primitive

type ProductCondition string

const (
	ProductConditionNew    ProductCondition = "new"
	ProductConditionSecond ProductCondition = "second-hand"
)

func (p ProductCondition) IsValid() bool {
	switch p {
	case ProductConditionNew, ProductConditionSecond:
		return true
	default:
		return false
	}
}
