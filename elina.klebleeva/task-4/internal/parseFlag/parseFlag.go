package parseFlag

import "flag"

func ParseFlags() (string, string) {

	var (
		appVersion string
		logFile    string
	)

	flag.StringVar(&appVersion, "appVersion", "sync", "sync or unsync version for app")
	flag.StringVar(&logFile, "logFile", "logs/operation_logs.txt", "path to log file")

	flag.Parse()

	return appVersion, logFile
}
