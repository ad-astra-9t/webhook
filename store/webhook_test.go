package store

import (
	"reflect"
	"testing"

	"github.com/ad-astra-9t/webhook/db"
	"github.com/ad-astra-9t/webhook/domain"
	mdl "github.com/ad-astra-9t/webhook/model"
)

func TestCreateWebhook(t *testing.T) {
	db := db.MustNewDB(
		"postgres",
		"host=localhost port=5431 user=test password=test dbname=testdb sslmode=disable",
	)
	model := mdl.NewDefaultModel(db)
	adapt := &mdl.ModelAdapt{}
	store := NewWebhookStore(model, adapt)

	t.Run("Test create webhook", func(t *testing.T) {
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
