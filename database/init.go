package database

func Setup() {
	db := new(Mysql)
	db.Setup()
}

//数据库如果已经准备好了则不需要automigrate