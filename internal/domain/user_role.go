package domain

type UserRole struct {
	UserID int `gorm:"primaryKey"`
	RoleID int `gorm:"primaryKey"`
}
