package request

type UserMany struct {
	QueryField

	AmountRange     IntRangeField
	Levels          IntsField
	Genders         StringsField
	CreateTimeRange DateTimeRangeField
	UpdateTimeRange DateRangeField
	StartTimeRange  TimeRangeField

	Filter struct {
		ID              *int    `binding:"omitempty,min=1"`
		Nickname        *string `binding:"omitempty,min=2"`
		AmountRange     IntRangeJsonField
		Levels          IntsJsonField
		Genders         StringsJsonField
		CreateTimeRange DateTimeRangeJsonField
		UpdateTimeRange DateRangeJsonField
		StartTimeRange  TimeRangeJsonField
	}
}
