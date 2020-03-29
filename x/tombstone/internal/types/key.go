package types

const (
	// ModuleName is the name of the module
	ModuleName = "tombstone"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querierer msgs
	QuerierRoute = ModuleName

	// LayoutDate Time format
	LayoutDate = "02/02/2020"

	// MemoMaxLength max length of memo
	MemoMaxLength = 140
)
