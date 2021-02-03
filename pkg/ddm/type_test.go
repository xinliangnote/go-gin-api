package ddm

import (
	"encoding/json"
	"testing"
)

type message struct {
	Name      IDName   `json:"name"`
	Mobile    Mobile   `json:"mobile"`
	IDCard    IDCard   `json:"id_card"`
	PassWord  PassWord `json:"password"`
	Email     Email    `json:"email"`
	BankCard1 BankCard `json:"bank_card_1"`
	BankCard2 BankCard `json:"bank_card_2"`
	BankCard3 BankCard `json:"bank_card_3"`
}

func TestMarshalJSON(t *testing.T) {
	msg := new(message)
	msg.Name = IDName("李鸿章")
	msg.Mobile = Mobile("13288887986")
	msg.IDCard = IDCard("125252525252525252")
	msg.PassWord = PassWord("123456")
	msg.Email = Email("xinliangnote@163.com")
	msg.BankCard1 = BankCard("6545654565456545")
	msg.BankCard2 = BankCard("65485269874569852")
	msg.BankCard3 = BankCard("6548526987456985298")

	marshal, _ := json.Marshal(msg)
	t.Log(string(marshal))
}
