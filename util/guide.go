package util

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func EditGuide() bool {
	// call struct
	var f T33File

	// guide here
	fmt.Println("#1/2 service Part")
	fmt.Println(" - 1/2 input service tag")
	_, _ = fmt.Scan(&f.Tag)
	fmt.Println(" - 2/2 input service port (Port Range: 0 ~ 65535)")
	_, _ = fmt.Scan(&f.Port)
	fmt.Println("#2/2 Setting Part")
	fmt.Println(" - 1/2 input file path (eg. /xxx)")
	_, _ = fmt.Scan(&f.Service.Path)
	fmt.Println(" - 2/2 input default file (eg. index.html)")
	_, _ = fmt.Scan(&f.Service.File)
	fmt.Println("Editing...")

	// struct to json
	b, err := json.Marshal(f)
	if err != nil {
		fmt.Println("turn to json failed")
		return false
	}

	// write to file
	file, _ := os.OpenFile("T33File_"+time.Now().Format("0405"), os.O_CREATE, 0666)
	_, err = file.Write(b)
	_ = file.Close()

	if err != nil {
		return false
	}

	return true
}
