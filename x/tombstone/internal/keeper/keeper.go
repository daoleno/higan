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
	store := ctx.KVStore(k.storeKey)
	var note types.Note
	byteValue := store.Get(recorder)
	if byteValue == nil {
		return types.Note{
			Records:  nil,
			Recorder: recorder}, nil
	}
	err := k.cdc.UnmarshalBinaryLengthPrefixed(byteValue, &note)
	if err != nil {
		return types.Note{}, err
	}

	return note, nil

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
	oldRecords := oldNote.Records
	note := types.Note{
		Records:  append(oldRecords, record),
		Recorder: msg.Recorder,
	}

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(note)
	store.Set(msg.Recorder, bz)

	return nil
}

// ListRecoder - Get all recorder
func (k Keeper) ListRecoder(ctx sdk.Context) ([]sdk.AccAddress, error) {
	store := ctx.KVStore(k.storeKey)
	var itr = store.Iterator(nil, nil)
	defer itr.Close()

	var allRecorder []sdk.AccAddress
	for ; itr.Valid(); itr.Next() {
		allRecorder = append(allRecorder, itr.Key())
	}

	return allRecorder, nil
}

// ListNote - Get all records
func (k Keeper) ListNote(ctx sdk.Context) ([]types.Note, error) {
	store := ctx.KVStore(k.storeKey)
	var itr = store.Iterator(nil, nil)
	defer itr.Close()

	var (
		note  types.Note
		notes []types.Note
	)
	for ; itr.Valid(); itr.Next() {
		err := k.cdc.UnmarshalBinaryLengthPrefixed(itr.Value(), &note)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	return notes, nil
}
