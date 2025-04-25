package utils

type key string

const (
	RabbitmqAmqpContextKey  key = "rabbitmq_context_key"
	RedisDBContextKey       key = "redis_db_context_key"
	ElasticsearchContextKey key = "elasticsearch_context_key"
	TracerContextKey        key = "tracer_context_key"
	CassandraDBContextKey   key = "cassandra_db_context_key"
	PostgresSQLContextKey   key = "postgres_sql_context_key"

	ProductSpanContextKey key = "product_service_span_key"

	UserUseCaseContextKey key = "user_use_case_context_key"
)
