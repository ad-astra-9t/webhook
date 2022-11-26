package dbx

import "fmt"

const (
	webhookTableName = "webhooks"
)

var (
	PGQueryGetWebhook = fmt.Sprintf(`
SELECT
    callback
FROM
    %s
WHERE
    callback = $1`,
		webhookTableName,
	)
	PGQueryCreateWebhook = fmt.Sprintf(`
INSERT INTO
    %s (callback)
    VALUES
        ($1)`,
		webhookTableName,
	)
)
