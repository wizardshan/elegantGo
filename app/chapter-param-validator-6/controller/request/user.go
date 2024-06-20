package request

type GenderBySeparatorField struct {
	NumberBySeparatorFieldV1
}

func (req *GenderBySeparatorField) UnmarshalJSON(b []byte) error {
	if err := req.NumberBySeparatorFieldV1.UnmarshalJSON(b); err != nil {
		return err
	}

	if err := validate.Var(req.Values, "dive,oneof=0 1 2"); err != nil {
		return err
	}

	return nil
}

type StatusBySeparatorField struct {
	StringBySeparatorFieldV1
}

func (req *StatusBySeparatorField) UnmarshalJSON(b []byte) error {
	if err := req.StringBySeparatorFieldV1.UnmarshalJSON(b); err != nil {
		return err
	}

	if err := validate.Var(req.Values, "dive,oneof=normal cancel invalid"); err != nil {
		return err
	}

	return nil
}

type LevelRangeField struct {
	NumberRangeFieldV1
}

func (req *LevelRangeField) UnmarshalJSON(b []byte) error {
	if err := req.NumberRangeFieldV1.UnmarshalJSON(b); err != nil {
		return err
	}

	if err := validate.Var([]int{req.Start, req.End}, "dive,oneof=0 10 20 30 40 50"); err != nil {
		return err
	}

	return nil
}

type UserMany struct {
	QueryField

	Filter struct {
		ID         *int    `binding:"omitempty,min=1"`
		Nickname   *string `binding:"omitempty,min=2"`
		Level      NumberBySeparatorFieldV7
		Status     StringBySeparatorFieldV7
		Amount     NumberRangeFieldV1
		CreateTime TimeRangeFieldV7
	}
}
