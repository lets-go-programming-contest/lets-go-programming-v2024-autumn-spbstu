package config

import "flag"

var dbFileName string
var serverFileName string

func init() {
	flag.StringVar(&dbFileName, "dbconfig", "C:/Users/User/lets-go-programming-v2024-autumn-spbstu/mesropyan.artyom/task-9/configs/dbConfig.yaml", "Read file with configuration data")
	flag.StringVar(&serverFileName, "serverconfig", "C:/Users/User/lets-go-programming-v2024-autumn-spbstu/mesropyan.artyom/task-9/configs/serverConfig.yaml", "Read file with configuration data")
	flag.Parse()
}
