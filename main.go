package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"go-template/etc"
	"go-template/svc"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

var configFile = flag.String("f", "etc/config.yaml", "this etc file")

func init() {
	flag.Parse()

	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	})

	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	//初始化配置
	etc.InitConfig(*configFile)

	//初始化服务
	g, err := svc.InitService()
	if err != nil {
		fmt.Println("init service fail", err.Error())
	}

	//优雅启动关闭
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	logrus.Info("server is starting...")

	httpServer := &http.Server{
		Addr:    etc.GetConfig().Addr,
		Handler: g,
	}
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Info("server is fail", err.Error())
		}
	}()

	<-ch
	logrus.Info("server is shutting down...")

	finish := make(chan struct{}, 1)

	// K8s 大概会给 30 秒关闭时间
	// @link  https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-phase
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()

			err := httpServer.Shutdown(ctx)
			if err != nil {
				logrus.Fatalf("Gin forced to shutdown: %s", err)
			}
		}()

		wg.Wait()
		finish <- struct{}{}
	}()

	select {
	case <-finish:
		logrus.Info("All service shutdown gracefully.")
	case <-ctx.Done():
		logrus.Error("Ungracefully shutdown.")
	}
}
