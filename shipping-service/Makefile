#!/bin/sh

generate-orm:
	@dsync -i src/model/diesel/schema.rs \
	-o src/model/diesel \
	--single-model-file \
	--schema-path crate::model::diesel::schema:: \
	--no-crud \
	--model-path crate::model::diesel:: \
	--connection-type diesel::pg::PgConnection

