package main 

import (
	"fmt"
	"os"
	"github.com/zhangliangxiaohehanxin/finalexam/routes"
	"github.com/zhangliangxiaohehanxin/finalexam/service/customer"
)

func main() {
	hostName := os.Getenv("DATABASE_URL")
	port := os.Getenv("PORT")
	
	cus := &customer.Customer{}
	r := routes.Route{ cus, hostName}.Init()
	r.Run(fmt.Sprintf(":%s", port))
}
