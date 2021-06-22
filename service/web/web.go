package web

import (
	"github.com/aURORA-JC/T33Server/util"
	"github.com/siddontang/go-log/log"
	"io/ioutil"
	"mime"
	"net/http"
	"path"
	"strings"
)

func T33WebService(w http.ResponseWriter, r *http.Request) {
	// get request file path & turn to local path
	filePath := r.RequestURI
	if filePath == "/" {
		filePath = "." + util.Config.Service.Path + "/" + util.Config.Service.File
	} else {
		filePath = "." + util.Config.Service.Path + "/" + filePath
	}
	filePath = strings.Split(filePath, "?")[0]

	// set content-type and server in header
	fileType := path.Ext(filePath)
	setHeader(w, fileType)

	// read file & write to response body
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		// file not found
		if r.RequestURI == "/" && util.Config.Service.File == "index.html" {
			// but maybe example site
			_, err := w.Write([]byte(t33ServerHtml))
			if err != nil {
				log.Errorln("Write data failed")
				return
			}
			return
		}

		// just not this file
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(t33NotFoundHtml))
		if err != nil {
			log.Errorln("Write data failed")
			return
		}
	} else {
		// file exist
		_, err := w.Write(file)
		if err != nil {
			log.Errorln("Write data failed")
			return
		}
	}

	// print log to console
	log.Infoln("["+util.Config.Tag+"]", r.Method, r.RequestURI, strings.Split(r.RemoteAddr, ":")[0], r.UserAgent())
}

func setHeader(w http.ResponseWriter, t string) {
	// set content-type
	w.Header().Set("Content-Type", mime.TypeByExtension(t))
	// set server
	w.Header().Set("Server", "T33Server 0.1.1-dev [web]")
}

// default html
const t33ServerHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>T33Server | Demo</title>
    <link href="https://cdn.bootcdn.net/ajax/libs/twitter-bootstrap/5.0.1/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
<main>
    <div class="container py-4">
        <header class="pb-3 mb-4 border-bottom">
            <a href="#" class="d-flex align-items-center text-dark text-decoration-none">
                <span class="fs-4">T33 Server</span>
            </a>
        </header>

        <div class="p-5 mb-4 bg-light rounded-3">
            <div class="container-fluid py-5">
                <h1 class="display-5 fw-bold">Hello World</h1>
                <p class="col-md-8 fs-4">
                    Just a TeamWork for Computer Network.
                </p>
            </div>
        </div>

        <footer class="pt-3 mt-4 text-muted border-top">
            &copy; 2021 T33Server 09/18/36 - T33Server based on Golang!!!
        </footer>
    </div>
</main>
</body>
</html>
`
const t33NotFoundHtml = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>404 Not Found</title>
    <link href="https://cdn.bootcdn.net/ajax/libs/twitter-bootstrap/5.0.1/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
<main>
    <div class="container py-4">
        <header class="pb-3 mb-4 border-bottom">
            <a href="#" class="d-flex align-items-center text-dark text-decoration-none">
                <span class="fs-4">T33 Server</span>
            </a>
        </header>

        <div class="p-5 mb-4 bg-light rounded-3">
            <div class="container-fluid py-5">
                <h1 class="display-5 fw-bold">404 Not Found</h1>
                <p class="col-md-8 fs-4">
                    The page you requested is not available.
                </p>
            </div>
        </div>

        <footer class="pt-3 mt-4 text-muted border-top">
            &copy; 2021 T33Server 09/18/36 - T33Server based on Golang!!!
        </footer>
    </div>
</main>
</body>
</html>
`
