package store

import (
	"reflect"
	"testing"

	"github.com/ad-astra-9t/webhook/db"
	"github.com/ad-astra-9t/webhook/domain"
)

func TestCreateWebhook(t *testing.T) {
	store := WebhookStore{
		db: db.MustNewDB(
			"postgres",
			"host=localhost port=5431 user=test password=test dbname=testdb sslmode=disable",
		),
		tableName: "webhooks",
		dbAdapter: db.AdaptWebhook,
	}
	w1 := domain.Webhook{Callback: "https://callback.com"}

	err := store.CreateWebhook(w1)
	if err != nil {
		t.Fatalf("Failed to create webhook: %s\n", err.Error())
	}

	w2, err := store.GetWebhook(w1)
	if err != nil || !reflect.DeepEqual(w2, w1) {
		t.Errorf("Webhook is created incorrectly, err: %#v, got: %#v, want: %#v\n", err, w2, w1)
	}
}
