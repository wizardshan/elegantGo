package request

import "elegantGo/param-complex-validate-1/controller/request/user"

type UserMany struct {
	SortField   `binding:"omitempty"`
	OrderField  `binding:"omitempty"`
	OffsetField `binding:"omitempty"`
	LimitField  `binding:"omitempty"`

	IDField            `binding:"omitempty"`
	user.NicknameField `binding:"omitempty"`
	Amount             IntRangeField
	Level              IntsField
	Status             StringsField
	CreateTime         DateTimeRangeField
	UpdateTime         DateRangeField
	StartTime          TimeRangeField
}
