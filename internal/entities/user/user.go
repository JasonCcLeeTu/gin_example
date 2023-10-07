package entityuser

type User struct {
	Account  string `json:"account" gorm:"column:account"`
	Password string `json:"password" gorm:"column:password"`
}
