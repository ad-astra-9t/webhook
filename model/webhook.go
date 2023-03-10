package model

import (
	"errors"
	"reflect"

	"github.com/ad-astra-9t/webhook/db"
	query "github.com/ad-astra-9t/webhook/db/query"
)

type Webhook struct {
	ID       uint   `db:"id"`
	Callback string `db:"callback"`
}

type WebhookModel struct {
	db db.DB
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
	query := query.PGQueryGetWebhook

	args, err := m.getWebhookArgs(target)
	if err != nil {
		return result, err
	}

	row := m.db.QueryRowx(query, args...)
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

	_, err = m.db.Exec(query, args...)

	return err
}

func NewWebhookModel(db db.DB) WebhookModel {
	return WebhookModel{db}
}
