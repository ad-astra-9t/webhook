package tx

import (
	"reflect"
	"testing"

	"github.com/ad-astra-9t/webhook/db"
	"github.com/ad-astra-9t/webhook/domain"
	mdl "github.com/ad-astra-9t/webhook/model"
)

func TestStorex_CreateWebhook(t *testing.T) {
	db := db.MustNewDB(
		"postgres",
		"host=localhost port=5431 user=test password=test dbname=testdb sslmode=disable",
	)
	dbx := NewDBX(db)
	model := NewDBXModel(dbx)
	modelx := NewModelx(model)
	adapt := mdl.ModelAdapt{}
	store := NewModelxStore(modelx, &adapt)
	storex := NewStorex(store)

	t.Run("Test create webhook", func(t *testing.T) {
		target := domain.Webhook{Callback: "https://callback.com"}

		err := storex.CreateWebhook(target)
		if err != nil {
			t.Fatalf("Failed to create webhook: %s\n", err.Error())
		}

		result, err := storex.GetWebhook(target)
		if err != nil || !reflect.DeepEqual(result, target) {
			t.Errorf("Webhook is created incorrectly, err: %#v, got: %#v, want: %#v\n", err, result, target)
		}
	})
}
