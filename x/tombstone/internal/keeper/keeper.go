package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/daoleno/higan/x/tombstone/internal/types"
)

// Keeper of the tombstone store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
}

// NewKeeper creates a tombstone keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramspace types.ParamSubspace) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetNote get note from store
func (k Keeper) GetNote(ctx sdk.Context, recorder sdk.AccAddress) (types.Note, error) {
	records, err := k.GetRecords(ctx, recorder)
	if err != nil {
		return types.Note{}, nil
	}

	return types.Note{
		Records:  records,
		Recorder: recorder}, nil

}

// SetNote set note to store
func (k Keeper) SetNote(ctx sdk.Context, msg types.MsgSetRecord) error {
	store := ctx.KVStore(k.storeKey)
	record := types.Record{
		Name: msg.Name,
		Born: msg.Born,
		Died: msg.Died,
		Memo: msg.Memo,
	}
	oldNote, err := k.GetNote(ctx, msg.Recorder)
	if err != nil {
		return err
	}
	oldRecords := oldNote.GetRecords()
	note := types.Note{
		Records:  append(oldRecords, record),
		Recorder: msg.Recorder,
	}

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(note)
	store.Set([]byte(msg.Recorder), bz)

	return nil
}

// GetRecords get many records recording by recorder
func (k Keeper) GetRecords(ctx sdk.Context, recorder sdk.AccAddress) ([]types.Record, error) {
	store := ctx.KVStore(k.storeKey)
	var records []types.Record
	byteKey := []byte(recorder)
	byteValue := store.Get(byteKey)
	if byteValue == nil {
		return nil, fmt.Errorf("Can not find any records")
	}
	err := k.cdc.UnmarshalBinaryLengthPrefixed(byteValue, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}
