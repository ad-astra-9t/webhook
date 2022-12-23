TEST_PGDB_NAME=testdb
TEST_PGDB_INIT_SCRIPT_DIR=$(PWD)/initdb

start-test-db-container:
	docker run \
		--env POSTGRES_USER=test \
		--env POSTGRES_PASSWORD=test \
		--env POSTGRES_DB=$(TEST_PGDB_NAME) \
		--publish=5431:5432 \
		--detach=true \
		--name=$(TEST_PGDB_NAME) \
		--rm \
		postgres:15.0

init-test-db-schema:
	for INIT_SCRIPT in $$(ls $(TEST_PGDB_INIT_SCRIPT_DIR)); do \
		docker cp \
			$(TEST_PGDB_INIT_SCRIPT_DIR)/$$INIT_SCRIPT \
			$(TEST_PGDB_NAME):/docker-entrypoint-initdb.d/$$INIT_SCRIPT; \
	done

start-test-db: start-test-db-container init-test-db-schema

stop-test-db:
	docker stop testdb
