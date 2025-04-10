package permission

type Permission struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"size:100;not null;unique"`
}
