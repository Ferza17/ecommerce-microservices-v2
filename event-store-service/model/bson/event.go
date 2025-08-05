package bson

import "time"

// Base Event structure
type BaseEvent struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"`
	Timestamp time.Time              `json:"timestamp"`
	SagaID    string                 `json:"saga_id"`
	Data      map[string]interface{} `json:"data"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// Saga State
type SagaState struct {
	ID        string                 `json:"id"`
	Status    string                 `json:"status"` // STARTED, IN_PROGRESS, COMPLETED, FAILED, ROLLING_BACK, ROLLED_BACK
	Steps     []SagaStep             `json:"steps"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	Data      map[string]interface{} `json:"data"`
}

type SagaStep struct {
	Service    string    `json:"service"`
	Action     string    `json:"action"`
	Status     string    `json:"status"` // PENDING, COMPLETED, FAILED, COMPENSATED
	Timestamp  time.Time `json:"timestamp"`
	RetryCount int       `json:"retry_count"`
	MaxRetries int       `json:"max_retries"`
}

// Service State for tracking last known state
type ServiceState struct {
	ServiceName string                 `json:"service_name"`
	EntityID    string                 `json:"entity_id"`
	LastState   map[string]interface{} `json:"last_state"`
	Version     int64                  `json:"version"`
	UpdatedAt   time.Time              `json:"updated_at"`
}
