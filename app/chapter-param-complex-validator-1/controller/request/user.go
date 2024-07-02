package request

//	type GenderBySeparatorField struct {
//		NumbersBySeparatorField
//	}
//
//	func (req *GenderBySeparatorField) UnmarshalJSON(b []byte) error {
//		if err := req.NumbersBySeparatorField.UnmarshalJSON(b); err != nil {
//			return err
//		}
//
//		if err := validate.Var(req.Values, "dive,oneof=0 1 2"); err != nil {
//			return err
//		}
//
//		return nil
//	}
//
//	type StatusBySeparatorField struct {
//		StringsBySeparatorField
//	}
//
//	func (req *StatusBySeparatorField) UnmarshalJSON(b []byte) error {
//		if err := req.StringsBySeparatorField.UnmarshalJSON(b); err != nil {
//			return err
//		}
//
//		if err := validate.Var(req.Values, "dive,oneof=normal cancel invalid"); err != nil {
//			return err
//		}
//
//		return nil
//	}
//
//	type LevelRangeField struct {
//		NumberRangeField
//	}
//
//	func (req *LevelRangeField) UnmarshalJSON(b []byte) error {
//		if err := req.NumberRangeField.UnmarshalJSON(b); err != nil {
//			return err
//		}
//
//		if err := validate.Var([]int{req.Start, req.End}, "dive,oneof=0 10 20 30 40 50"); err != nil {
//			return err
//		}
//
//		return nil
//	}
type UserMany struct {
	QueryField

	Filter struct {
		ID         *int    `binding:"omitempty,min=1"`
		Nickname   *string `binding:"omitempty,min=2"`
		CreateTime TimeRangeFieldV91
		Level      NumberRangeFieldV9
	}
}
