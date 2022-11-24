package model

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ad-astra-9t/webhook/dbx"
	"github.com/ad-astra-9t/webhook/domain"
)

const webhookTableName = "webhooks"

type Webhook struct {
	ID       uint   `db:"id"`
	Callback string `db:"callback"`
}

type WebhookModel struct {
	dbx *dbx.DBX
}

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
	query := fmt.Sprintf(`
	SELECT
		callback
	FROM
		%s
	WHERE
		callback = $1`,
		webhookTableName,
	)

	args, err := m.getWebhookArgs(target)
	if err != nil {
		return result, err
	}

	ext, err := m.dbx.ToExt()
	if err != nil {
		return result, err
	}

	row := ext.QueryRowx(query, args...)
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
	query := fmt.Sprintf(`
INSERT INTO
	%s (callback)
	VALUES
		($1)`,
		webhookTableName,
	)

	args, err := m.createWebhookArgs(target)
	if err != nil {
		return err
	}

	ext, err := m.dbx.ToExt()
	if err != nil {
		return err
	}

	_, err = ext.Exec(query, args...)

	return err
}

func (m WebhookModel) AdaptModel(domainwebhook domain.Webhook) (modelwebhook Webhook) {
	modelwebhook = Webhook{
		ID:       domainwebhook.ID,
		Callback: domainwebhook.Callback,
	}
	return
}

func (m WebhookModel) AdaptDomain(modelwebhook Webhook) (domainwebhook domain.Webhook) {
	domainwebhook = domain.Webhook{
		ID:       modelwebhook.ID,
		Callback: modelwebhook.Callback,
	}
	return
}

func NewWebhookModel(dbx *dbx.DBX) WebhookModel {
	return WebhookModel{
		dbx: dbx,
	}
}
