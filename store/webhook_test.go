package store

import (
	"reflect"
	"testing"

	"github.com/ad-astra-9t/webhook/autotx"
	"github.com/ad-astra-9t/webhook/db"
	"github.com/ad-astra-9t/webhook/domain"
	mdx "github.com/ad-astra-9t/webhook/modelx"
)

func TestCreateWebhook(t *testing.T) {
	t.Run("Test create webhook", func(t *testing.T) {
		db := db.MustNewDB(
			"postgres",
			"host=localhost port=5431 user=test password=test dbname=testdb sslmode=disable",
		)
		dbx := autotx.NewDBX(db)
		model := mdx.NewModel(dbx)
		modelx := mdx.NewModelx(model)
		adapt := &mdx.ModelAdapt{}
		store := NewWebhookStore(modelx, adapt)

		target := domain.Webhook{Callback: "https://callback.com"}

		err := store.CreateWebhook(target)
		if err != nil {
			t.Fatalf("Failed to create webhook: %s\n", err.Error())
		}

		result, err := store.GetWebhook(target)
		if err != nil || !reflect.DeepEqual(result, target) {
			t.Errorf("Webhook is created incorrectly, err: %#v, got: %#v, want: %#v\n", err, result, target)
		}
	})
}
