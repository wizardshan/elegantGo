package request

type UserMany struct {
	QueryField

	Amount     IntRange
	Level      Ints    `binding:"omitempty,dive,required,oneof=10 20 30 40 50"`
	Status     Strings `binding:"omitempty,dive,required,oneof=normal cancel invalid"`
	CreateTime TimeRange

	Filter struct {
		ID         *int    `binding:"omitempty,min=1"`
		Nickname   *string `binding:"omitempty,min=2"`
		Amount     IntRange
		Level      Ints    `binding:"omitempty,dive,required,oneof=10 20 30 40 50"`
		Status     Strings `binding:"omitempty,dive,required,oneof=normal cancel invalid"`
		CreateTime TimeRange
	}
}
