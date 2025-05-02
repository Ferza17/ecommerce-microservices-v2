-- +goose Up
ALTER TABLE users ADD COLUMN is_verified BOOL DEFAULT FALSE;