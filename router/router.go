package router

import (
	"GoPolyglot/controllers"
	"GoPolyglot/controllers/auth"
	"GoPolyglot/controllers/tasks"
	"GoPolyglot/libs/common/error_wrapper"
	"GoPolyglot/libs/configs"
	"GoPolyglot/libs/logger"
	"GoPolyglot/router/middlewares"
	"GoPolyglot/router/middlewares/trace"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func InitRouter() {
	gin.SetMode(configs.GConfig.Server.Mode)

	if configs.GConfig.Server.Mode != "debug" {
		gin.DefaultWriter = io.Discard
	}

	router := gin.Default()
	router.Use(trace.SetUp())

	router.GET("/", func(c *gin.Context) { c.JSON(http.StatusOK, error_wrapper.WithSuccess(c)) })
	router.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "PONG") })

	r := router.Group("/api")
	{
		r.POST("/", error_wrapper.WrapperError(controllers.Index))
	}

	router.POST("/auth/users", error_wrapper.WrapperError(auth.RegisterUser)) // 用户注册
	router.POST("/auth/login", error_wrapper.WrapperError(auth.LoginUser))    // 用户登录

	// 任务管理相关路由
	//withAuth := router.Group("", middleware.StrongAuthMiddleware())
	limitAuth := router.Group("", middleware.StrongAuthMiddleware(), middleware.LogMiddleware(), middleware.RateLimitMiddleware())
	taskGroup := limitAuth.Group("/tasks")
	{
		taskGroup.POST("", error_wrapper.WrapperError(tasks.CreateTask)) // 创建翻译任务
		//taskGroup.POST("/:task_id/translate", error_wrapper.WrapperError(tasks.TranslateTask)) // 执行翻译任务
		//taskGroup.GET("/:task_id", error_wrapper.WrapperError(tasks.GetTaskStatus))        // 获取任务状态
		//taskGroup.GET("/:task_id/download", error_wrapper.WrapperError(tasks.DownloadTask)) // 下载翻译文档
	}

	srv := &http.Server{
		Addr:    configs.GConfig.Server.Address + ":" + configs.GConfig.Server.Port,
		Handler: router,
	}
	go func() {
		// 服务连接
		logger.GLogger.Info("服务器启动成功!")
		f, err := os.OpenFile("pid.txt", os.O_CREATE|os.O_RDWR, 0666)
		if err == nil {
			_, err = f.WriteString(strconv.Itoa(syscall.Getpid()))
			if err != nil {
				defer f.Close()
			}
		} else {
			logger.GLogger.Fatal(err.Error())
		}

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.GLogger.Fatal(err.Error())
		}

	}()
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.GLogger.Info("服务器关闭中 ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.GLogger.Fatal(err.Error())
	}
	logger.GLogger.Info("服务器关闭！")

}
