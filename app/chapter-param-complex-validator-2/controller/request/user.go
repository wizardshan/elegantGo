package request

type UserMany struct {
	QueryField

	Filter struct {
		ID         *int    `binding:"omitempty,min=1"`
		Nickname   *string `binding:"omitempty,min=2"`
		Amount     NumberRange
		Status     Strings
		Level      Numbers
		CreateTime TimeRange
	}
}
