package hash

import (
	"github.com/speps/go-hashids"
)

func (h *hash) HashidsEncode(params []int) (string, error) {
	hd := hashids.NewData()
	hd.Salt = h.secret
	hd.MinLength = h.length

	hashStr, err := hashids.NewWithData(hd).Encode(params)
	if err != nil {
		return "", err
	}

	return hashStr, nil
}

func (h *hash) HashidsDecode(hash string) ([]int, error) {
	hd := hashids.NewData()
	hd.Salt = h.secret
	hd.MinLength = h.length

	ids, err := hashids.NewWithData(hd).DecodeWithError(hash)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
