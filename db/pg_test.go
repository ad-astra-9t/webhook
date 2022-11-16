package db

import "testing"

func TestMustNewDB(t *testing.T) {
	t.Run("Test connecting to PostgreSQL database", func(t *testing.T) {
		driver := "postgres"
		dsn := "host=localhost port=5431 user=test password=test dbname=testdb sslmode=disable"

		pgdb := MustNewDB(driver, dsn)
		err := pgdb.Ping()
		if err != nil {
			t.Errorf("Failed to connect to database. driver: %s, dsn: %s, err: %s\n", driver, dsn, err)
		}
	})
	t.Run("Test panicking when newing database", func(t *testing.T) {
		driver := "invalid"
		dsn := "host=localhost port=5431 user=test password=test dbname=testdb sslmode=disable"

		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("Failed to panick when newing database. driver: %s, dsn: %s\n", driver, dsn)
			}
		}()
		_ = MustNewDB(driver, dsn)
	})
}
