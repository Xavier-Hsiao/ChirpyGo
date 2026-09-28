package main

import "sync/atomic"

type apiConfig struct {
	fileServerHits atomic.Int32
}
