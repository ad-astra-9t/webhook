package modelx

import (
	"errors"
	"reflect"

	query "github.com/ad-astra-9t/webhook/dbx/query"
	"github.com/ad-astra-9t/webhook/domain"
)

type Webhook struct {
	ID       uint   `db:"id"`
	Callback string `db:"callback"`
}

type WebhookModel struct {
	autoTxDB AutoTxDB
}

type WebhookAdapt struct{}

func (m WebhookModel) getWebhookArgs(modelwebhook Webhook) ([]interface{}, error) {
	args := make([]interface{}, 0)

	v1 := reflect.ValueOf(modelwebhook)
	if v1.Kind() != reflect.Struct {
		return args, errors.New("failed to generate args")
	}

	for i := 0; i < v1.NumField(); i++ {
		v2 := v1.Field(i)
		if !v2.CanInterface() {
			return args, errors.New("failed to generate args")
		}
		if !v2.IsZero() {
			args = append(args, v2.Interface())
		}
	}

	return args, nil
}

func (m WebhookModel) GetWebhook(target Webhook) (result Webhook, err error) {
	query := query.PGQueryGetWebhook

	args, err := m.getWebhookArgs(target)
	if err != nil {
		return result, err
	}

	txdb, err := m.autoTxDB.AutoTx()
	if err != nil {
		return result, err
	}

	row := txdb.QueryRowx(query, args...)
	err = row.StructScan(&result)

	return result, err
}

func (m WebhookModel) createWebhookArgs(modelwebhook Webhook) ([]interface{}, error) {
	args := make([]interface{}, 0)

	v1 := reflect.ValueOf(modelwebhook)
	if v1.Kind() != reflect.Struct {
		return args, errors.New("failed to generate args")
	}

	requiredFields := []string{"Callback"}
	for _, name := range requiredFields {
		v2 := v1.FieldByName(name)
		if v2.IsZero() || !v2.CanInterface() {
			return args, errors.New("failed to generate args")
		}
		args = append(args, v2.Interface())
	}

	return args, nil
}

func (m WebhookModel) CreateWebhook(target Webhook) error {
	query := query.PGQueryCreateWebhook

	args, err := m.createWebhookArgs(target)
	if err != nil {
		return err
	}

	txdb, err := m.autoTxDB.AutoTx()
	if err != nil {
		return err
	}

	_, err = txdb.Exec(query, args...)

	return err
}

func (a WebhookAdapt) AdaptModel(domainwebhook domain.Webhook) (modelwebhook Webhook) {
	modelwebhook = Webhook{
		ID:       domainwebhook.ID,
		Callback: domainwebhook.Callback,
	}
	return
}

func (a WebhookAdapt) AdaptDomain(modelwebhook Webhook) (domainwebhook domain.Webhook) {
	domainwebhook = domain.Webhook{
		ID:       modelwebhook.ID,
		Callback: modelwebhook.Callback,
	}
	return
}

func NewWebhookModel(autoTxDB AutoTxDB) WebhookModel {
	return WebhookModel{autoTxDB}
}
