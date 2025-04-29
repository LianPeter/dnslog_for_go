package example

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//go:embed static/* templates/*
var embedFS embed.FS

func StartExampleServer() {
	// 初始化 Gin 引擎
	r := gin.Default()

	// 设置 zap 日志
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	// 嵌入并注册静态资源，例如 CSS、JS 等
	staticFiles, err := fs.Sub(embedFS, "static")
	if err != nil {
		logger.Error("无法加载静态文件", zap.Error(err))
	}
	r.StaticFS("/static", http.FS(staticFiles))

	// 嵌入并注册 HTML 模板
	tmplFiles, err := fs.Sub(embedFS, "templates")
	if err != nil {
		logger.Error("无法加载模板文件", zap.Error(err))
	}
	tmpl, err := template.ParseFS(tmplFiles, "*.html")
	if err != nil {
		logger.Error("模板解析失败", zap.Error(err))
	}
	r.SetHTMLTemplate(tmpl)

	// 设置路由：展示表单页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "欢迎访问 DNSLog 示例",
		})
	})

	// 启动 HTTP 服务
	if err := r.Run(":8080"); err != nil {
		logger.Fatal("服务器启动失败", zap.Error(err))
	}
}
