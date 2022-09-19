package users

type User struct {
	Id        int64  `db:"id" json:"id"`
	Login     string `db:"login" json:"login"`
	Password  string `db:"password" json:"-"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
	Age       int64  `db:"age" json:"age"`
	Sex       int64  `db:"sex" json:"sex"`
	Interests string `db:"interests" json:"interests"`
	City      string `db:"city" json:"city"`
}
