package ballot
import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"

	sdk "github.com/cosmos/cosmos-sdk/types"
)
type Keeper struct {
	coinKeeper bank.Keeper

	ballotStoreKey  sdk.StoreKey // Unexposed key to access ballot store from sdk.Context
	ownersStoreKey sdk.StoreKey // Unexposed key to access owners store from sdk.Context
	historyStoreKey sdk.StoreKey // Unexposed key to access history store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the nservice Keeper
func NewKeeper(coinKeeper bank.Keeper, ballotStoreKey sdk.StoreKey, ownersStoreKey sdk.StoreKey, historyStoreKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		coinKeeper:     coinKeeper,
		ballotStoreKey:  ballotStoreKey,
		ownersStoreKey: ownersStoreKey,
		historyStoreKey: historyStoreKey,
		cdc:            cdc,
	}
}