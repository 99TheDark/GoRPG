package connections

type Variation string

const (
	Default     Variation = "default"
	Top         Variation = "top"
	Bottom      Variation = "bottom"
	Left        Variation = "left"
	Right       Variation = "right"
	TopLeft     Variation = "top_left"
	BottomLeft  Variation = "bottom_left"
	TopRight    Variation = "top_right"
	BottomRight Variation = "bottom_right"
)

func (variation Variation) String() string {
	return string(variation)
}
