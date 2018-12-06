package main

/*
 * param package handles the cli parameters
 */

import (
	"flag"
	"sync"
)

// Config is the cmd parameter output
type Config struct {
	File *string
}

var (
	config     *Config
	configLock sync.RWMutex
)

func init() {
	c := Config{
		File: flag.String("file", "", "file to parse"),
	}
	flag.Parse()
	config = &c
}

// Get Allows you to get a parameter
func Get() *Config {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}
