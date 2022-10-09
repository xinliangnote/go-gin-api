package order

import "time"

// Order
//
//go:generate gormgen -structs Order -input .
type Order struct {
	Id          int32     //
	OrderNo     string    //
	OrderFee    int32     // ()
	Status      int32     //  1:  2:
	IsDeleted   int32     //  1:  -1:
	CreatedAt   time.Time `gorm:"time"` //
	CreatedUser string    //
	UpdatedAt   time.Time `gorm:"time"` //
	UpdatedUser string    //
}
