package service

func CheckDebit(transaction_type string, value int) int {
	if transaction_type == "d" {
		return value * -1
	}
	return value
}
