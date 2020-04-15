package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// Query endpoints supported by the tombstone querier
const (
	QueryRecord      = "Record"
	QueryAllRecord   = "Records"
	QueryAllRecorder = "Recorder"
)

type QueryRecordRes Note
type QueryRecorderRes []sdk.AccAddress
type QueryAllNoteRes []Note
