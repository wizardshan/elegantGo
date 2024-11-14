package hashid

import (
	"errors"
	"github.com/speps/go-hashids/v2"
)

const (
	CategoryPost    = iota
	CategoryArticle = iota
	CategoryUser
	CategoryAuthor
)

var h *hashids.HashID

func init() {
	hd := hashids.NewData()
	hd.Salt = "elegantGo!@#"
	hd.MinLength = 10

	h, _ = hashids.NewWithData(hd)
}

func EncodeUserID(id int) (string, error) {
	return h.Encode([]int{id, CategoryUser})
}

func DecodeUserID(hashID string) (int, error) {
	id, category, err := decode(hashID)
	if err != nil {
		return id, err
	}
	if category != CategoryUser {
		return id, errors.New("the id category error")
	}

	return id, err
}

func EncodePostID(id int) (string, error) {
	return h.Encode([]int{id, CategoryPost})
}

func DecodePostID(hashID string) (int, error) {
	id, category, err := decode(hashID)
	if err != nil {
		return id, err
	}
	if category != CategoryPost {
		return id, errors.New("the id category error")
	}

	return id, err
}

func decode(hashID string) (int, int, error) {
	numbers, err := h.DecodeWithError(hashID)
	id, category := 0, 0
	if err != nil {
		return id, category, err
	}

	if len(numbers) != 2 {
		return id, category, errors.New("the numbers capacity expected value is 2")
	}

	id, category = numbers[0], numbers[1]

	return id, category, nil
}
