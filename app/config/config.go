package config

import "os"

var firstrun bool

func init() {

	if _, err := os.Stat("block.db"); os.IsNotExist(err) {
		firstrun = true
	}
}

func IsFirstRun() bool {
	return firstrun
}
