package main 

import (
	"fmt"
	"os"
	"log"
	"github.com/zhangliangxiaohehanxin/finalexam/routes"
	"github.com/zhangliangxiaohehanxin/finalexam/service/customer"
)


func main() {
	hostName := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	logFile := createLogFile()

	cus := &customer.Customer{}
	route := routes.Route{ cus, hostName}
	r := route.Init()

	r.Run(fmt.Sprintf(":%s", port))
	defer route.DestroySession()
	defer logFile.Close()

}

func createLogFile() *os.File{
	// tail -f testlogfile
	logFile, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("can't Create Log file", err.Error())
	}

	log.SetOutput(logFile)
	return logFile
}
