package types

type MerchantType int32

const (
	MerchantType__UNKNOWN MerchantType = iota
)

type MerchantAcctChangeType int32

const (
	MerchantAcctChangeType__UNKNOWN  MerchantAcctChangeType = iota // 未知
	MerchantAcctChangeType__DEPOSIT                                // 存款
	MerchantAcctChangeType__WITHDRAW                               // 提款
	MerchantAcctChangeType__TRANSFER                               // 转账
	MerchantAcctChangeType__REFUND                                 // 退款
	MerchantAcctChangeType__RECHARGE                               // 充值
)
