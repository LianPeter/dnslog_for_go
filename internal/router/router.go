package router

import (
	"dnslog_for_go/internal/domain"
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
		log.Fatalf("Failed to embed static files: %v", err)
	}
	r.StaticFS("/static", http.FS(staticFiles))

	// 嵌入html模板
	tmplFiles, err := fs.Sub(embedFS, "templates")
	if err != nil {
		log.Fatalf("Failed to embed template files: %v", err)
	}
	tmpl, err := template.ParseFS(tmplFiles, "*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
	r.SetHTMLTemplate(tmpl)

	// 路由处理
	r.GET("/dnslog", domain.ShowForm)
	r.POST("/submit", domain.SubmitDomain)

	// 监听并启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
