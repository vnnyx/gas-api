package model

type (
	User struct {
		ID        string `gorm:"column:user_id;primaryKey;type:varchar(255)"`
		Username  string `gorm:"column:username;unique;type:varchar(50)"`
		Email     string `gorm:"column:email;unique;type:varchar(255)"`
		Handphone string `gorm:"column:handphone;type:varchar(50)"`
		Password  string `gorm:"column:password;type:varchar(255)"`
	}

	Kelontong struct {
		ID       string `gorm:"column:kelontong_id;primaryKey;varchar(255)"`
		Name     string `gorm:"column:name;varchar(100)"`
		Location string `gorm:"column:location;varchar(100)"`
	}

	Fashion struct {
		ID   string `gorm:"column:fashion_id;primaryKey;varchar(255)"`
		Name string `gorm:"column:name;varchar(100)"`
	}
)

func (User) TableName() string {
	return "user"
}

func (Kelontong) TableName() string {
	return "kelontong"
}

func (Fashion) TableName() string {
	return "fashion"
}
