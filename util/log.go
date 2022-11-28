package util

import "log"

var Log *log.Logger

func init() {
	Log = log.Default()
	log.SetFlags(log.Ldate | log.Llongfile | log.Ltime)
}
