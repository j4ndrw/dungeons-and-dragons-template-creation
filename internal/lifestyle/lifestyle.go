package lifestyle

const (
	Wretched     = "Wretched"
	Squalid      = "Squalid"
	Poor         = "Poor"
	Modest       = "Modest"
	Comfortable  = "Comfortable"
	Wealthy      = "Wealthy"
	Aristocratic = "Aristocratic"
)

var Lifestyles []Lifestyle = []Lifestyle{
	Wretched,
	Squalid,
	Poor,
	Modest,
	Comfortable,
	Wealthy,
	Aristocratic,
}

type Lifestyle string
