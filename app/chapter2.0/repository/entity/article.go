package entity

type Article struct {
	HashID  string
	Title   string
	Content string

	HashIDQuery string
	SQL         string
	Err         string
}
