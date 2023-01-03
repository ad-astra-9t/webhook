package etx

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/ad-astra-9t/webhook/db"
	"github.com/ad-astra-9t/webhook/domain"
	"github.com/ad-astra-9t/webhook/etx/cptx"
	"github.com/ad-astra-9t/webhook/model"
)

func TestStoreEtx_CreateWebhook(t *testing.T) {
	db := db.MustNewDB(
		"postgres",
		"host=localhost port=5431 user=test password=test dbname=testdb sslmode=disable",
	)
	txOptions := &sql.TxOptions{}
	dbCptx := cptx.NewDBCptx(db, txOptions)
	modelCptx := cptx.NewModelCptx(dbCptx)
	storeCptx := cptx.NewStoreCptx(modelCptx, &model.ModelAdapt{})
	storeEtx := NewStoreEtx(storeCptx)

	t.Run("Test create webhook", func(t *testing.T) {
		target := domain.Webhook{Callback: "https://callback.com"}

		err := storeEtx.CreateWebhook(target)
		if err != nil {
			t.Fatalf("Failed to create webhook: %s\n", err.Error())
		}

		result, err := storeEtx.GetWebhook(target)
		if err != nil || !reflect.DeepEqual(result, target) {
			t.Errorf("Webhook is created incorrectly, err: %#v, got: %#v, want: %#v\n", err, result, target)
		}
	})

	t.Run("Test cancel transaction when creating webhooks", func(t *testing.T) {
		ctx := context.Background()

		err := storeEtx.Etx(ctx)
		if err != nil {
			t.Fatalf("Store failed to start transaction. err: %s\n", err.Error())
		}

		targets := []domain.Webhook{
			{Callback: "https://callback1.com"},
			{Callback: "https://callback2.com"},
		}
		for _, target := range targets {
			err := storeEtx.CreateWebhook(target)
			if err != nil {
				t.Fatalf("Store failed to create webhook. err: %s, target: %#v\n", err.Error(), target)
			}
			result, err := storeEtx.GetWebhook(target)
			if err != nil {
				t.Fatalf("Store failed to get webhook. err: %s, target: %#v\n", err.Error(), target)
			}
			if target != result {
				t.Fatalf("Webhook is incorrectly created. got: %#v, want: %#v\n", result, target)
			}
		}

		err = storeEtx.Cancel()
		if err != nil {
			t.Fatalf("Store failed to cancel transaction. err: %s\n", err.Error())
		}
		result, err := storeEtx.GetWebhook(targets[0])
		if err == nil {
			t.Fatalf("Transaction is incorrectly cancelled. result: %#v\n", result)
		}
		if err != sql.ErrNoRows {
			t.Fatalf("Store failed to get webhook. err: %s, target: %#v\n", err.Error(), targets[0])
		}
	})
}
