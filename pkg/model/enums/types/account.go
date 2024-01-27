package types

type AccountType int32

const (
	AccountType__UNKNOWN AccountType = iota
	AccountType__Admin
	AccountType__System
	AccountType__Manager
	AccountType__CustomerService
	AccountType__Member
	AccountType__Tourist
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
