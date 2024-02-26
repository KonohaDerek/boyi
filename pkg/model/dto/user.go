package dto

import (
	"boyi/internal/claims"
	"boyi/pkg/model/enums/types"
	"boyi/pkg/model/option/common"
	"encoding/json"
	"fmt"
	"time"

	"gopkg.in/vmihailenco/msgpack.v2"
	"gorm.io/plugin/soft_delete"
)

const (
	UserCacheKey = "user:%d"
	// user:{user_id}:online
	UserOnlineKey = "user:%d:online"

	LockCreateTouristKey = "lock_create_tourist_key"
)

// User admin 管理员
type User struct {
	ID              uint64                `gorm:"autoIncrement;primary_key"`                                   // id
	AccountType     types.AccountType     `gorm:"type:tinyint(4);DEFAULT:0;NOT NULL;idx_users_account_type"`   // 帳號類型
	Status          types.UserStatus      `gorm:"type:tinyint(4);DEFAULT:1;NOT NULL;"`                         // 狀態
	Username        string                `gorm:"type:nvarchar(100);DEFAULT:'';NOT NULL;uniqueIndex:udx_name"` // 用户名
	Password        string                `gorm:"type:nvarchar(255);NOT NULL"`                                 // 密碼
	AliasName       string                `gorm:"type:nvarchar(100);DEFAULT:''"`                               // 別名（聊天室顯示用）
	Email           string                `gorm:"type:nvarchar(255);DEFAULT:''"`                               // 電子郵件
	Area            string                `gorm:"type:nvarchar(100);DEFAULT:''"`                               // 居住地區
	Notes           string                `gorm:"type:nvarchar(100);DEFAULT:''"`                               // 備註
	IsNeedChangePwd common.YesNo          `gorm:"type:tinyint;NOT NULL;DEFAULT:2"`                             // 是否需要修改密碼
	AvatarKey       FileKey               `gorm:"type:nvarchar(255);NOT NULL"`                                 // 頭像 s3 的 key
	LastLoginAt     time.Time             `gorm:"type:TIMESTAMP;DEFAULT:current_timestamp"`                    // 最後登陸時間
	LastLoginIP     string                `gorm:"type:varchar(40);"`                                           // 最後登陸 IP
	CreatedAt       time.Time             `gorm:"type:TIMESTAMP;NOT NULL;"`                                    // 创建时间
	UpdatedAt       time.Time             `gorm:"type:TIMESTAMP;NOT NULL;"`                                    // 更新时间
	DeletedAt       soft_delete.DeletedAt `gorm:"uniqueIndex:udx_name"`
	UpdateUserID    uint64                `gorm:""`             // 更新人
	CreateUserID    uint64                `gorm:"type:int(11)"` // 建立者ID

	Roles      []UserRole      `gorm:"PRELOAD:false;foreignKey:UserID"`
	Whitelists []UserWhitelist `gorm:"PRELOAD:false;foreignKey:UserID"`
	Tags       []UserTag       `gorm:"PRELOAD:false;foreignKey:UserID" msgpack:"-"`

	Menu Authority `gorm:"-" msgpack:"-"` // 該 User 擁有的 menu tree
}

// TableName return database table name
func (u User) TableName() string {
	return "users"
}

func (c User) Marshal() []byte {
	b, _ := msgpack.Marshal(c)
	return b
}

func (c *User) Unmarshal(s string) error {
	if err := msgpack.Unmarshal([]byte(s), &c); err != nil {
		return err
	}
	return nil
}

func (u User) ToJSON() string {
	b, _ := json.Marshal(u)
	return string(b)
}

func (u *User) FromJSON(s string) error {
	if err := json.Unmarshal([]byte(s), &u); err != nil {
		return err
	}
	return nil
}

func (u *User) ToClaims(expires time.Time) claims.Claims {
	cc := claims.Claims{
		Id:          u.ID,
		Username:    u.Username,
		AccountType: uint64(u.AccountType),
		ExpiredAt:   expires.Unix(),
	}

	cc.Competences = make(map[string]bool)

	// 如果為 admin 時 擁有所有權限
	if u.AccountType == types.AccountType__System {
		for key := range GetMenuMap() {
			cc.Competences[key.String()] = true
		}
	} else {
		for i := range u.Roles {
			if u.Roles[i].Role.SupportAccountType != u.AccountType {
				continue
			}
			dfsToMap(cc.Competences, u.Roles[i].Role.Authority)
		}
	}

	return cc
}

func (u *User) GenerateOnlineKey() string {
	return fmt.Sprintf(UserOnlineKey, u.ID)
}

func (u *User) VerifyAllowCreate(c claims.Claims) error {
	var (
		key ManagerMenuKey
	)

	if c.AccountType == uint64(types.AccountType__Admin) {
		return nil
	}

	switch u.AccountType {
	case types.AccountType__Manager:
		key = API_Manager_Create
	case types.AccountType__CustomerService:
		key = API_CustomerService_Create
	}
	return c.VerifyRole(key.String())
}

func (u *User) VerifyAllowUpdate(c claims.Claims) error {
	var (
		key ManagerMenuKey
	)

	if c.AccountType == uint64(types.AccountType__Admin) {
		return nil
	}

	if c.Id == u.ID {
		return nil
	}

	switch u.AccountType {
	case types.AccountType__Manager:
		key = API_Manager_Update
	case types.AccountType__CustomerService:
		key = API_CustomerService_Update
	}
	return c.VerifyRole(key.String())
}

func (u *User) VerifyAllowDelete(c claims.Claims) error {
	var (
		key ManagerMenuKey
	)

	if c.AccountType == uint64(types.AccountType__Admin) {
		return nil
	}

	switch u.AccountType {
	case types.AccountType__Manager:
		key = API_Manager_Delete
	case types.AccountType__CustomerService:
		key = API_CustomerService_Delete
	}
	return c.VerifyRole(key.String())
}

func (u *User) VerifyAllowUpdatePassword(c claims.Claims) error {
	var (
		key ManagerMenuKey
	)

	if u.ID == c.Id {
		return nil
	}

	if c.AccountType == uint64(types.AccountType__Admin) {
		return nil
	}

	switch u.AccountType {
	case types.AccountType__Manager:
		key = API_Manager_Password_Update
	case types.AccountType__CustomerService:
		key = API_CustomerService_Password_Update
	}

	return c.VerifyRole(key.String())
}
