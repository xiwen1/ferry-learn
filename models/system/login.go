package system

import (
	orm "ferry-learn/global"
	"ferry-learn/tools"
)

type Login struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	LoginType int    `json:"loginType"`
}

func (u *Login) GetUser() (user SysUser, role SysRole, err error) {
	err = orm.Eloquent.Table("sys_user").Where("username = ?", u.Username).Find(&user).Error
	if err != nil {return}
	//check the password
	if u.LoginType == 0 {
		_, err = tools.CompareHashAndPassword(user.Password, u.Password)
		if err != nil {return}
	}
	//get user's role
	err = orm.Eloquent.Table("sys_role").Where("role_id = ?", user.RoleId).First(&role).Error
	if err != nil {return}
	return

}