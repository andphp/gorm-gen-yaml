// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameUserOauth = "ims_user_oauth"

// UserOauth mapped from table <ims_user_oauth>
type UserOauth struct {
	ID     int32 `gorm:"column:id;type:INTEGER" json:"id"`
	UserID int32 `gorm:"column:user_id;type:INTEGER" json:"userId"`
}

// TableName UserOauth's table name
func (*UserOauth) TableName() string {
	return TableNameUserOauth
}