package parseFlag

import "flag"

func ParseFlags() string {

	var appVersion string
	flag.StringVar(&appVersion, "appVersion", "sync", "sync or unsync version for app")

	flag.Parse()

	return appVersion
}
