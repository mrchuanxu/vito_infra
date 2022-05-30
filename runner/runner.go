package runner

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mrchuanxu/vito_infra/logger"
)

// Runable 可运行、停止的服务
type Runable interface {
	Start()
	Stop()
}

// Run 运行
// 运行服务，并检测系统中断
// 本方法适合在main gorouting中需要运行后台服务时调用，调用此方法后，此方法会自动异步调用传入的Runable.Start方法，并检测系统中断，当检测到系统中断则调用Runable.Stop方法来终止服务
// @Author Trans
// @param startFn 启动服务的方法，方法应当是同步方法
// @param stopFn 停止服务的方法，方法应当是同步方法
func Run(r Runable) {
	csignal := make(chan os.Signal, 1)
	signal.Notify(csignal, syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	for {
		// 运行传入的startFn
		go r.Start()
		// 等待结束信号
		s := <-csignal
		// close(config.StopSignal)
		switch s {
		case syscall.SIGINT:
			logger.TransLogger.Warn("Process exit with SIGINT!")
		case syscall.SIGQUIT:
			logger.TransLogger.Warn("process exit with SIGQUIT!")
		case syscall.SIGHUP:
			time.Sleep(time.Second)
			logger.TransLogger.Warn("process restart with SIGHUP...")
			continue // continue可以重启服务
		case syscall.SIGKILL:
			logger.TransLogger.Warn("process killed with SIGKILL!")
		case syscall.SIGTERM:
			logger.TransLogger.Warn("process exit with SIGTERM!")
		default:
			logger.TransLogger.Warn("process unknown exit")
		}
		// 收到结束信号
		r.Stop()
		// 运行至此处代表关闭服务
		break
	}
}
