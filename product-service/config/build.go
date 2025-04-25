package config

import "time"

var (
	AppName,
	Version,
	Commit,
	Build string
	Now = time.Now()
)
