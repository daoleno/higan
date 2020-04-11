package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Note have many record
type Note struct {
	Records  []Record       `json:"records"`
	Recorder sdk.AccAddress `json:"recorder"`
}

// Record people who will be remembered
type Record struct {
	Name string    `json:"name"`
	Born time.Time `json:"born"`
	Died time.Time `json:"died"`
	Memo string    `json:"memo"`
}

func (r Record) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Name: %s
	Born: %s
	Died: %s
	Memo: %s`, r.Name, r.Born, r.Died, r.Memo))
}
