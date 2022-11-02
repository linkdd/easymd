package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/linkdd/easymd/pkg/document"
)

func Serve(rootDocument string, host net.IP, port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testPaths := document.GetTestPaths(rootDocument, r.URL.Path)

		for _, testPath := range testPaths {
			if _, err := os.Stat(testPath); !os.IsNotExist(err) {
				page, success := document.RenderPath(testPath)

				if success {
					w.WriteHeader(http.StatusOK)
				} else {
					w.WriteHeader(http.StatusInternalServerError)
				}

				err := page.Render(w)
				if err != nil {
					log.Println("ERROR: ", err)
				}

				return
			}
		}

		if path, found := hasStaticFile(rootDocument, r.URL.Path); found {
			http.ServeFile(w, r, path)
		} else {
			page := document.Page{
				Params: document.PageParams{
					Title: "Not Found",
					Lang:  "en",
					Meta:  []document.MetaTag{},
					CSS:   []string{},
					JS:    []string{},
				},
				Content: fmt.Sprintf(
					`
						<div class="message is-danger m-5">
							<div class="message-header"><p>Not Found</p></div>
							<div class="message-body">%s</div>
						</div>
					`,
					r.URL.Path,
				),
			}

			w.WriteHeader(http.StatusNotFound)
			err := page.Render(w)
			if err != nil {
				log.Println("ERROR: ", err)
			}
		}
	})

	address := net.JoinHostPort(host.String(), strconv.Itoa(port))
	log.Println(fmt.Sprintf("Listening on http://%s", address))
	log.Fatal(http.ListenAndServe(address, nil))
}
