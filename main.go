package main

import (
	lep_audit "bupt/lep_audit/kitex_gen/lep_audit/lepaudit"
	"log"

	"github.com/zhongershashen/lep_lib/dal"
)

func main() {
	svr := lep_audit.NewServer(new(LepAuditImpl))
	dal.Init()
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
