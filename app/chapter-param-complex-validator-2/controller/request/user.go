package request

type UserMany struct {
	QueryField

	Filter struct {
		ID         *int    `binding:"omitempty,min=1"`
		Nickname   *string `binding:"omitempty,min=2"`
		Amount     NumberRange
		Status     Strings `binding:"omitempty,dive,required,oneof=normal cancel invalid"`
		Level      Numbers `binding:"omitempty,dive,required,oneof=10 20 30 40 50"`
		CreateTime TimeRange
	}
}
