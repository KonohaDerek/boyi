package http

import (
	"time"
)

// REQUEST_TIME_OUT
const (
	RequestTimeout = 10 * time.Second
)

// Http methods
const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)
