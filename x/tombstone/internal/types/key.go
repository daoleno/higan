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
	LayoutDate = "01/02/2006"

	// MemoMaxLength max length of memo
	MemoMaxLength = 140
)
