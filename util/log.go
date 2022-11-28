package util

import "log"

var Log *log.Logger

func init() {
	Log = log.Default()
	Log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}
