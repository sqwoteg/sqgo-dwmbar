package config

const SEPARATOR = "    "
const AFTER_STATUS_TEXT = " "
const DATE_FORMAT = "02.01.2006 Mon 15:04"

// If you use program other than xkb-switch,
// it should print layout to stdout on change.
// You should specify at least path and 1 arg.
var XKB_SWITCH_CMD = [...]string{
	"xkb-switch",
	"-W",
}

var MODULES_ORDER = []string{
	// module name from workers.go file
	"layout",
	"time",
}
