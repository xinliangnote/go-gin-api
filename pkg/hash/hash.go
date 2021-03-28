package hash

var _ Hash = (*hash)(nil)

type Hash interface {
	i()

	// hashids
	HashidsEncode(params []int) (string, error)
	HashidsDecode(hash string) ([]int, error)
}

type hash struct {
	secret string
	length int
}

func New(secret string, length int) Hash {
	return &hash{
		secret: secret,
		length: length,
	}
}

func (h *hash) i() {}
