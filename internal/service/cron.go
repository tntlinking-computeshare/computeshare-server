package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"time"
)

type CronJob struct {
	computeInstanceUC *biz.ComputeInstanceUsercase
	log               *log.Helper
}

func NewCronJob(computeInstanceUsercase *biz.ComputeInstanceUsercase, logger log.Logger) *CronJob {
	return &CronJob{
		computeInstanceUC: computeInstanceUsercase,
		log:               log.NewHelper(logger),
	}
}

func (c *CronJob) StartJob() {
	// 定时同步虚拟机的cpu和内存使用情况
	go c.syncComputeInstanceStatsTask()
}

func (c *CronJob) syncComputeInstanceStatsTask() {
	// 创建一个定时触发的通道，每隔一秒发送一个时间事件
	ticker := time.Tick(1 * time.Minute)

	// 使用 for 循环执行定时任务
	for {
		select {
		case <-ticker:
			// 在这里执行你的定时任务代码
			log.Info("开始同步虚拟机的cpu和内存使用情况")
			c.computeInstanceUC.SyncContainerStats()
			log.Info("结束同步虚拟机的cpu和内存使用情况")
		}
	}
}
