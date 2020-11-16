package main

import (
	"net/http"
	"test/chart"
	"test/docker"
	"time"

	"github.com/gin-gonic/gin"
)

// func main() {
// 	http.HandleFunc("/images", images)

// 	http.ListenAndServe(":9000", nil)
// }

// func images(w http.ResponseWriter, r *http.Request) {
// 	for {
// 		docker.TransferListContainer()
// 		chart.MakePieChart()
// 		b, _ := ioutil.ReadFile("pie.html")

// 		fmt.Fprint(w, string(b))
// 		ioutil.WriteFile("pie.html", nil, os.ModeExclusive)
// 		time.Sleep(time.Second)
// 	}
// }

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("pie.html")
	r.GET("/images", images)
	r.Run(":9000")
}

func images(c *gin.Context) {

	docker.TransferListContainer()
	chart.MakePieChart()
	time.Sleep(time.Second)

	c.HTML(http.StatusOK, "pie.html", gin.H{})
}
