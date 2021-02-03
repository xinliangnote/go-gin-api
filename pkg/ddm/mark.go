package ddm

import (
	"fmt"
	"strings"
)

func (m Mobile) MarshalJSON() ([]byte, error) {
	if len(m) != 11 {
		return []byte(`"` + m + `"`), nil
	}

	v := fmt.Sprintf("%s****%s", m[:3], m[len(m)-4:])
	return []byte(`"` + v + `"`), nil
}

func (bc BankCard) MarshalJSON() ([]byte, error) {
	if len(bc) > 19 || len(bc) < 16 {
		return []byte(`"` + bc + `"`), nil
	}

	v := fmt.Sprintf("%s******%s", bc[:6], bc[len(bc)-4:])
	return []byte(`"` + v + `"`), nil
}

func (card IDCard) MarshalJSON() ([]byte, error) {
	if len(card) != 18 {
		return []byte(`"` + card + `"`), nil
	}

	v := fmt.Sprintf("%s******%s", card[:1], card[len(card)-1:])
	return []byte(`"` + v + `"`), nil
}

func (name IDName) MarshalJSON() ([]byte, error) {
	if len(name) < 1 {
		return []byte(`""`), nil
	}

	nameRune := []rune(name)
	v := fmt.Sprintf("*%s", string(nameRune[1:]))
	return []byte(`"` + v + `"`), nil
}

func (pw PassWord) MarshalJSON() ([]byte, error) {
	v := "******"
	return []byte(`"` + v + `"`), nil
}

func (e Email) MarshalJSON() ([]byte, error) {
	if !strings.Contains(string(e), "@") {
		return []byte(`"` + e + `"`), nil
	}

	split := strings.Split(string(e), "@")
	if len(split[0]) < 1 || len(split[1]) < 1 {
		return []byte(`"` + e + `"`), nil
	}

	v := fmt.Sprintf("%s***%s", split[0][:1], split[0][len(split[0])-1:])
	return []byte(`"` + v + "@" + split[1] + `"`), nil
}
