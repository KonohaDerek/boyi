package types

type MerchantType int32

const (
	MerchantType__UNKNOWN MerchantType = iota
)

type MerchantAcctChangeType int32

const (
	MerchantAcctChangeType__UNKNOWN              MerchantAcctChangeType = iota // 未知
	MerchantAcctChangeType__UPDATE__BALANCE                                    // 更新余额
	MerchantAcctChangeType__UPDATE__BLOCK_AMOUNT                               // 更新圈存金额
	MerchantAcctChangeType__UPDATE__PROFIT                                     // 更新盈利
	MerchantAcctChangeType__UPDATE__BLOCK_PROFIT                               // 更新圈存盈利
)

var ConverterToMerchantAcctOperation = map[MerchantAcctChangeType]MerchantAcctOperation{
	MerchantAcctChangeType__UPDATE__BALANCE:      MerchantAcctOperation__UPDATE__BALANCE,
	MerchantAcctChangeType__UPDATE__BLOCK_AMOUNT: MerchantAcctOperation__UPDATE__BLOCK_AMOUNT,
	MerchantAcctChangeType__UPDATE__PROFIT:       MerchantAcctOperation__UPDATE__PROFIT,
	MerchantAcctChangeType__UPDATE__BLOCK_PROFIT: MerchantAcctOperation__UPDATE__BLOCK_PROFIT,
}

type MerchantAcctOperation int32

const (
	MerchantAcctOperation__UNKNOWN              MerchantAcctOperation = iota // 未知
	MerchantAcctOperation__CREATE                                            // 建立
	MerchantAcctOperation__DEPOSITING                                        // 存款中
	MerchantAcctOperation__DEPOSTED                                          // 存款完成
	MerchantAcctOperation__WITHDRAWING                                       // 提款中
	MerchantAcctOperation__WITHDRAWED                                        // 提款完成
	MerchantAcctOperation__TRANSFERING                                       // 转账中
	MerchantAcctOperation__TRANSFERED                                        // 转账完成
	MerchantAcctOperation__REFUNDING                                         // 退款中
	MerchantAcctOperation__REFUNDED                                          // 退款完成
	MerchantAcctOperation__UPDATE__BALANCE                                   // 更新余额
	MerchantAcctOperation__UPDATE__BLOCK_AMOUNT                              // 更新圈存金额
	MerchantAcctOperation__UPDATE__PROFIT                                    // 更新盈利
	MerchantAcctOperation__UPDATE__BLOCK_PROFIT                              // 更新圈存盈利
)
