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
	"os"
	"os/signal"
	"syscall"
	"time"
)

func StartServer(embedFS embed.FS) {
	r := gin.Default()

	// 嵌入静态文件
	staticFiles, err := fs.Sub(embedFS, "static")
	if err != nil {
		log_write.Error("Failed to embed static files", zap.Error(err))
		return
	}
	r.StaticFS("/static", http.FS(staticFiles))

	// 嵌入 HTML 模板
	tmplFiles, err := fs.Sub(embedFS, "templates")
	if err != nil {
		log_write.Error("Failed to embed template files", zap.Error(err))
		return
	}
	tmpl, err := template.ParseFS(tmplFiles, "*.html")
	if err != nil {
		log_write.Error("Failed to parse template files", zap.Error(err))
		return
	}
	r.SetHTMLTemplate(tmpl)

	// 路由处理
	r.GET("/dnslog", domain.ShowForm)
	r.POST("/submit", domain.SubmitDomain)
	r.POST("/random-domain", domain.RandomDomain)
	r.POST("/change", domain.ChangeServer)
	r.POST("/change-pact", domain.ChangePact)

	// 启动服务器
	go func() {
		if err := r.Run(":8080"); err != nil {
			log_write.Error("Failed to run server", zap.Error(err))
		}
	}()

	log_write.Info("Server started on :8080")

	// 退出
	exitChannel := make(chan os.Signal, 1)
	signal.Notify(exitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-exitChannel

	// 退出前清理
	log_write.Info("Shutting down server gracefully...")
	signal.Stop(exitChannel)
	time.Sleep(2 * time.Second) // 延迟处理
}
