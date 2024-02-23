package model

type Users []*User

func (users Users) SumAmount() int {
	total := 0
	for _, u := range users {
		total += u.Amount
	}

	return total
}
