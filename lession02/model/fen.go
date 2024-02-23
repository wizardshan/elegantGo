package model

type Fen struct {
	Value int
}

func (fen *Fen) ToYuan() int {
	return fen.Value / 100
}

func (fen *Fen) YuanRate() int {
	return 100
}
