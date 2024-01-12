package service

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mohaijiang/computeshare-server/internal/biz"
	"github.com/mohaijiang/computeshare-server/internal/data/ent"
	"time"
)

type CronJob struct {
	computeInstanceUC   *biz.ComputeInstanceUsercase
	agentUsecase        *biz.AgentUsecase
	cycleRenewalUseCase *biz.CycleRenewalUseCase
	db                  *ent.Client
	log                 *log.Helper
}

func NewCronJob(
	computeInstanceUsercase *biz.ComputeInstanceUsercase,
	agentUsecase *biz.AgentUsecase,
	cycleRenewalUseCase *biz.CycleRenewalUseCase,
	db *ent.Client,
	logger log.Logger,
) *CronJob {
	return &CronJob{
		computeInstanceUC:   computeInstanceUsercase,
		agentUsecase:        agentUsecase,
		cycleRenewalUseCase: cycleRenewalUseCase,
		db:                  db,
		log:                 log.NewHelper(logger),
	}
}

func (c *CronJob) StartJob() {
	// 定时同步虚拟机的cpu和内存使用情况
	//go c.syncComputeInstanceStatsTask()
	// 同步agent状态
	//go c.syncAgentStatusTask()
	// 同步虚拟机过期
	go c.syncContainerOverdue()
	// 每日同步续费管理
	go c.syncRenewalOrder(c.db)

}

//// 同步计算资源CPU和内存
//func (c *CronJob) syncComputeInstanceStatsTask() {
//	// 创建一个定时触发的通道，每隔一秒发送一个时间事件
//	ticker := time.Tick(1 * time.Minute)
//
//	// 使用 for 循环执行定时任务
//	for {
//		select {
//		case <-ticker:
//			// 在这里执行你的定时任务代码
//			log.Info("开始同步虚拟机的cpu和内存使用情况")
//			c.computeInstanceUC.SyncContainerStats()
//			log.Info("结束同步虚拟机的cpu和内存使用情况")
//		}
//	}
//}

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

func (c *CronJob) syncRenewalOrder(db *ent.Client) {
	// 获取当前时间
	currentTime := time.Now()

	// 计算距离下一个 23:00 的时间差
	// 如果当前时间已经过了 23:00，则计算到明天 23:00 的时间差
	nextTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 0, 0, 0, currentTime.Location())
	if currentTime.After(nextTime) {
		nextTime = nextTime.Add(24 * time.Hour)
	}
	timeUntilNext := nextTime.Sub(currentTime)
	fmt.Println(timeUntilNext)

	// 创建定时器
	timer := time.NewTimer(timeUntilNext)

	// 执行定时任务
	for {
		select {
		case <-timer.C:
			// 在这里执行你的定时任务逻辑
			fmt.Println("执行定时任务：每日23点")

			c.cycleRenewalUseCase.DailyCheck(db)

			// 重新计算下一个 23:00 的时间差
			nextTime = nextTime.Add(24 * time.Hour)
			timeUntilNext = nextTime.Sub(time.Now())
			timer.Reset(timeUntilNext)
		}
	}
}
