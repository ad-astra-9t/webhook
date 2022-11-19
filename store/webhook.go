package store

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	"github.com/ad-astra-9t/webhook/db"
	"github.com/ad-astra-9t/webhook/domain"

	"github.com/jmoiron/sqlx"
)

type WebhookStore struct {
	db        *sqlx.DB
	dbAdapter webhookDBAdapter
	tableName string
}

type webhookDBAdapter func(webhook domain.Webhook) db.Webhook

func (s *WebhookStore) createWebhookArgs(webhook db.Webhook) ([]interface{}, error) {
	args := make([]interface{}, 0)

	v1 := reflect.ValueOf(webhook)
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

func (s *WebhookStore) getWebhookArgs(webhook db.Webhook) ([]interface{}, error) {
	args := make([]interface{}, 0)

	v1 := reflect.ValueOf(webhook)
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

func (s *WebhookStore) createWebhook(args ...interface{}) (sql.Result, error) {
	query := fmt.Sprintf(`
INSERT INTO
	%s (callback)
	VALUES
		($1)`,
		s.tableName,
	)
	return s.db.Exec(query, args...)
}

func (s *WebhookStore) CreateWebhook(webhook domain.Webhook) error {
	w := s.dbAdapter(webhook)

	args, err := s.createWebhookArgs(w)
	if err != nil {
		return err
	}

	_, err = s.createWebhook(args...)

	return err
}

func (s *WebhookStore) getWebhook(args ...interface{}) *sqlx.Row {
	query := fmt.Sprintf(`
SELECT
	callback
FROM
	%s
WHERE
	callback = $1`,
		s.tableName,
	)
	return s.db.QueryRowx(query, args...)
}

func (s *WebhookStore) GetWebhook(webhook domain.Webhook) (domain.Webhook, error) {
	var result domain.Webhook

	w := s.dbAdapter(webhook)

	args, err := s.getWebhookArgs(w)
	if err != nil {
		return result, err
	}

	row := s.getWebhook(args...)

	err = row.StructScan(&result)
	if err != nil {
		return result, err
	}

	return result, err
}
