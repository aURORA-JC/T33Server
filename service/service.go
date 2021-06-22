package service

import (
	"github.com/aURORA-JC/T33Server/service/proxy"
	"github.com/aURORA-JC/T33Server/service/web"
	"github.com/aURORA-JC/T33Server/util"
	"github.com/siddontang/go-log/log"
	"net/http"
)

func Start(port string, mode bool) {
	// select service mode
	if !mode {
		http.HandleFunc("/", web.T33WebService)
	} else {
		http.HandleFunc(util.Config.Service.Path, proxy.T33ProxyService)
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Errorln("T33Server Start Failed: ", err)
	}
}
