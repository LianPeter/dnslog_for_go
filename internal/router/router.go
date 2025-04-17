package router

import (
	"dnslog_for_go/internal/domain"
	"dnslog_for_go/internal/log_write"
	"embed"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"html/template"
	"io/fs"
	"net/http"
)

func StartServer(embedFS embed.FS) {
	r := gin.Default()

	// 嵌入静态文件
	staticFiles, err := fs.Sub(embedFS, "static")
	if err != nil {
		log_write.Error("Failed to embed static files: %v", zap.Error(err))
	}
	r.StaticFS("/static", http.FS(staticFiles))

	// 嵌入html模板
	tmplFiles, err := fs.Sub(embedFS, "templates")
	if err != nil {
		log_write.Error("Failed to embed template files: %v", zap.Error(err))
	}
	tmpl, err := template.ParseFS(tmplFiles, "*.html")
	if err != nil {
		log_write.Error("Failed to embed template files: %v", zap.Error(err))
	}
	r.SetHTMLTemplate(tmpl)

	// 路由处理
	r.GET("/dnslog", domain.ShowForm)
	r.POST("/submit", domain.SubmitDomain)

	// 监听并启动服务器
	if err := r.Run(":8080"); err != nil {
		log_write.Error("Failed to run server: %v", zap.Error(err))
	}
	
	log_write.Info("Server started on :8080")
}
