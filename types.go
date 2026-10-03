package main

import (
	"sync/atomic"

	"github.com/Xavier-Hsiao/chirpy/internal/database"
)

type apiConfig struct {
	fileServerHits atomic.Int32
	dbQueries      *database.Queries
}
