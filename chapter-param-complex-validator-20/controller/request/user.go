package request

type UserMany struct {
	QueryField

	Amount     IntRangeField
	Level      IntsField
	Status     StringsField
	CreateTime DateTimeRangeField
	UpdateTime DateRangeField
	StartTime  TimeRangeField

	Filter struct {
		ID         *int    `binding:"omitempty,min=1"`
		Nickname   *string `binding:"omitempty,min=2"`
		Amount     IntRangeJsonField
		Level      IntsJsonField
		Status     StringsJsonField
		CreateTime DateTimeRangeJsonField
		UpdateTime DateRangeJsonField
		StartTime  TimeRangeJsonField
	}
}
