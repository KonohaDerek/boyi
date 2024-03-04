package types

type AccountType int32

const (
	AccountType__UNKNOWN         AccountType = iota // 未知
	AccountType__Admin                              // 系統管理員
	AccountType__System                             // 系統
	AccountType__Manager                            // 管理員
	AccountType__CustomerService                    // 客服
	AccountType__Merchant                           // 商戶
	AccountType__Agent                              // 代理
	AccountType__Member                             // 會員
	AccountType__Tourist                            // 訪客
)

type UserStatus int32

const (
	UserStatus__UNKNOWN UserStatus = iota
	UserStatus__UnVerified
	UserStatus__Actived
	UserStatus__Locked
	UserStatus__Disabled
	UserStatus__Deleted
)
