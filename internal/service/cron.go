package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"time"
)

type CronJob struct {
	computeInstanceUC *biz.ComputeInstanceUsercase
	agentUsecase      *biz.AgentUsecase
	log               *log.Helper
}

func NewCronJob(computeInstanceUsercase *biz.ComputeInstanceUsercase, agentUsecase *biz.AgentUsecase, logger log.Logger) *CronJob {
	return &CronJob{
		computeInstanceUC: computeInstanceUsercase,
		agentUsecase:      agentUsecase,
		log:               log.NewHelper(logger),
	}
}

func (c *CronJob) StartJob() {
	// 定时同步虚拟机的cpu和内存使用情况
	go c.syncComputeInstanceStatsTask()
	// 同步agent状态
	//go c.syncAgentStatusTask()
	// 同步虚拟机过期
	go c.syncContainerOverdue()

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

func (c *CronJob) syncAgentStatusTask() {
	// 创建一个定时触发的通道，每隔一秒发送一个时间事件
	ticker := time.Tick(10 * time.Minute)

	// 使用 for 循环执行定时任务
	for {
		select {
		case <-ticker:
			// 在这里执行你的定时任务代码
			log.Info("开始同步agent状态节点情况")
			c.agentUsecase.SyncAgentStatus()
			log.Info("结束同步agent状态节点情况")
		}
	}
}

func (c *CronJob) syncContainerOverdue() {
	// 创建一个定时触发的通道，每隔一秒发送一个时间事件
	ticker := time.Tick(10 * time.Minute)

	// 使用 for 循环执行定时任务
	for {
		select {
		case <-ticker:
			// 在这里执行你的定时任务代码
			log.Info("开始同步instance节点过期")
			c.computeInstanceUC.SyncContainerOverdue()
			log.Info("结束同步instance节点过期")
		}
	}
}
