package domain

import "github.com/pauldin91/gochain/src/internal"

var maxByTimestamp = func(k Transaction, t Transaction) Transaction {
	if k.Input.timestamp.UnixMilli() > t.Input.timestamp.Local().UnixMilli() {
		return k
	} else {
		return t
	}
}

var findTransactionByAddress = func(t Transaction, a string) bool {
	return t.Input.address == a
}
var findInputByAddress = func(t Input, a string) bool {
	return t.address == a
}

var findByAddressAndTimestamp = func(t Transaction, v TimestampAddressFilter) bool {
	return t.Input.timestamp.After(v.timestamp) &&
		internal.FindBy(t.Output, v.address, findInputByAddress) != nil
}
