#!/bin/sh

generate-orm:
	@dsync -i src/model/diesel/schema.rs \
	-o src/model/diesel \
	--schema-path crate::model::diesel::schema:: \
	--no-crud \
	--model-path crate::model::diesel:: \
	--connection-type diesel::pg::PgConnection

migration-run:
	diesel migration run \
	--migration-dir ./migrations

insert-mock:
	cargo run insert-mock \
	 --direction local
