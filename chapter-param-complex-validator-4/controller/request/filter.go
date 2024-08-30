package request

type QueryField struct {
	Sort   *SortField   `binding:"omitempty,alphanum"`
	Order  *OrderField  `binding:"omitempty,oneof=DESC ASC"`
	Offset *OffsetField `binding:"omitempty,int,min=0"`
	Limit  *LimitField  `binding:"omitempty,int,min=1,max=1000"`
}
