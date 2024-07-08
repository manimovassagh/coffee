package types

type Role struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:enum('buyer', 'seller');not null"`
}

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null" json:"-"` // Exclude from JSON responses
	RoleID   uint
	Role     Role `gorm:"foreignKey:RoleID"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
	UserID      uint
	User        User `gorm:"foreignKey:UserID"`
}

type Order struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      User    `gorm:"foreignKey:UserID"`
	Total     float64 `gorm:"not null"`
	Status    string  `gorm:"not null"`
	CreatedAt uint64  `gorm:"autoCreateTime"`
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	Order     Order `gorm:"foreignKey:OrderID"`
	ProductID uint
	Product   Product `gorm:"foreignKey:ProductID"`
	Quantity  int     `gorm:"not null"`
	Price     float64 `gorm:"not null"`
}
