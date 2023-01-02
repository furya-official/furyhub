package tibc_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	migratetibc "github.com/furya-official/furyhub/migrate/tibc"
	"github.com/furya-official/furyhub/simapp"
)

func TestLoadClient(t *testing.T) {
	app := simapp.Setup(t, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Height: 1})
	migratetibc.CreateClient(ctx, app.AppCodec(), "v1.3", app.TIBCKeeper.ClientKeeper)
}
