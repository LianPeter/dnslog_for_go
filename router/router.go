package router

import (
	"dnslog_for_go/domain"
	"embed"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func StartServer(embedFS embed.FS) {
	r := gin.Default()

	// 嵌入静态文件
	staticFiles, err := fs.Sub(embedFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	r.StaticFS("/static", http.FS(staticFiles))

	// 嵌入html模板
	tmplFiles, err := fs.Sub(embedFS, "templates")
	if err != nil {
		log.Fatal(err)
	}
	tmpl, err := template.ParseFS(tmplFiles, "*.html")
	if err != nil {
		log.Fatal(err)
	}
	r.SetHTMLTemplate(tmpl)

	r.GET("/dnslog", domain.ShowForm)
	r.POST("/submit", domain.SubmitDomain)

	r.Run(":8080")
}
