package system

import "time"

type LoginM struct {
	Username string `gorm:"type:varchar(64)" json:"username"`
	Password string `gorm:"type:varchar(128)" json:"password"`
}

type SysUserId struct {
	UserId int `gorm:"primary_key;AUTO_INCREMENT" json:"userId"`
}

type SysUserB struct {
	NickName string `gorm:"type:varchar(128)" json:"nickName"` // 昵称
	// Phone    string `gorm:"type:varchar(11)" json:"phone"`     // 手机号
	RoleId   int    `gorm:"type:int(11)" json:"roleId"`        // 角色编码
	// Salt     string `gorm:"type:varchar(255)" json:"salt"`     //盐
	Avatar   string `gorm:"type:varchar(255)" json:"avatar"`   //头像
	Sex      string `gorm:"type:varchar(255)" json:"sex"`      //性别
	Email    string `gorm:"type:varchar(128)" json:"email"`    //邮箱
	// DeptId   int    `gorm:"type:int(11)" json:"deptId"`        //部门编码
	// PostId   int    `gorm:"type:int(11)" json:"postId"`        //职位编码
	CreateBy string `gorm:"type:varchar(128)" json:"createBy"` //
	UpdateBy string `gorm:"type:varchar(128)" json:"updateBy"` //
	Remark   string `gorm:"type:varchar(255)" json:"remark"`   //备注
	Status   int `gorm:"type:int(1);" json:"status"`
	// Params   string `gorm:"-" json:"params"`
	BaseModel
}

type BaseModel struct {
	CreatedAt time.Time  `gorm:"column:create_time" json:"create_time" form:"create_time"`
	UpdatedAt time.Time  `gorm:"column:update_time" json:"update_time" form:"update_time"`
	DeletedAt *time.Time `gorm:"column:delete_time" sql:"index" json:"-"`
}

type SysUser struct {
	SysUserId
	SysUserB
	LoginM
}