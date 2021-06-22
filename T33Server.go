package main

import (
	"flag"
	"fmt"
	"github.com/aURORA-JC/T33Server/service"
	"github.com/aURORA-JC/T33Server/util"
	"github.com/siddontang/go-log/log"
)

func main() {
	// Server start
	fmt.Println("████████╗██████╗ ██████╗ ███████╗███████╗██████╗ ██╗   ██╗███████╗██████╗ ")
	fmt.Println("╚══██╔══╝╚════██╗╚════██╗██╔════╝██╔════╝██╔══██╗██║   ██║██╔════╝██╔══██╗")
	fmt.Println("   ██║    █████╔╝ █████╔╝███████╗█████╗  ██████╔╝██║   ██║█████╗  ██████╔╝")
	fmt.Println("   ██║    ╚═══██╗ ╚═══██╗╚════██║██╔══╝  ██╔══██╗╚██╗ ██╔╝██╔══╝  ██╔══██╗")
	fmt.Println("   ██║   ██████╔╝██████╔╝███████║███████╗██║  ██║ ╚████╔╝ ███████╗██║  ██║")
	fmt.Println("   ╚═╝   ╚═════╝ ╚═════╝ ╚══════╝╚══════╝╚═╝  ╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝")
	fmt.Println("==========================================================================")
	fmt.Println("T33Server is a teamwork for Computer Networking Design. (C) 2021 CUSTS 333")
	fmt.Println("==========================================================================")
	fmt.Println("================  DO NOT USE IN PRODUCTION ENVIRONMENTS  =================")
	fmt.Println("==========================================================================")

	// get main args
	serverPort := flag.String("port", "3330", "set T33Server Listen Port")
	optionFile := flag.String("conf", "", "set T33Server Config File Path")
	proxyMode := flag.Bool("proxy", false, "run T33Server under Proxy Mode")
	editMode := flag.Bool("edit", false, "open T33Server Config Guide")
	flag.Parse()

	// call edit mode & config guide
	if *editMode {
		log.Infoln("Turning to config edit mode")
		if util.EditGuide() {
			log.Infoln("Guide End, check T33File in this program's path")
		}
		return
	}

	// check args & choose server mode
	var port = *serverPort
	var path = *optionFile
	var mode = *proxyMode // true = proxy, false = web
	if mode {
		log.Infoln("T33Server will run under Proxy Mode")
	}

	if port == "3330" && path == "" {
		log.Warnln("T33Server will run with default config. Use 'Ctrl+C' to cancel")
	} else {
		log.Infoln("T33Server is searching for config file...")
	}

	// set config
	config := util.ParseT33(path)
	port = config.Port
	if port == "-1" {
		log.Errorln("Error Config, check T33File")
		return
	}
	util.Config = config
	log.Infoln("T33Server is going to listen port: " + port)

	// start server
	log.Infoln("T33Server Starting...")
	service.Start(port, mode)
}
