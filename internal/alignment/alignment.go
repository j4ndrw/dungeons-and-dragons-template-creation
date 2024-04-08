package alignment

const (
	LawfulGood  = "LawfulGood"
	NeutralGood = "NeutralGood"
	ChaoticGood = "ChaoticGood"

	LawfulEvil  = "LawfulEvil"
	NeutralEvil = "NeutralEvil"
	ChaoticEvil = "ChaoticEvil"

	LawfulNeutral  = "LawfulNeutral"
	TrueNeutral    = "TrueNeutral"
	ChaoticNeutral = "ChaoticNeutral"
)

var Alignments []Alignment = []Alignment{
	LawfulGood,
	NeutralGood,
	ChaoticGood,

	LawfulEvil,
	NeutralEvil,
	ChaoticEvil,

	LawfulNeutral,
	TrueNeutral,
	ChaoticNeutral,
}

type Alignment string
