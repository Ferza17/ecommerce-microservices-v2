package kafka_schema

type Fields struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Optional bool   `json:"optional"`
}
type Schema struct {
	Type   string   `json:"type"`
	Fields []Fields `json:"fields"`
}

type KafkaSinkSchema struct {
	Schema  Schema         `json:"schema"`
	Payload map[string]any `json:"payload"`
}
