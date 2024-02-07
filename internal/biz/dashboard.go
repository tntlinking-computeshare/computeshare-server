package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	pb "github.com/mohaijiang/computeshare-server/api/dashboard/v1"
	"github.com/mohaijiang/computeshare-server/internal/conf"
	"github.com/mohaijiang/computeshare-server/internal/global/consts"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

var (
	StorageSpaceTotal                = "sum(SeaweedFS_volumeServer_resource{type=\"all\"})"
	UsedStorageSpaceTotal            = "sum(SeaweedFS_volumeServer_total_disk_size) by (exported_instance)"
	StorageProvidersNum              = "count(SeaweedFS_volumeServer_resource{type=\"all\"})"
	UsedVolumeTotal                  = "sum(SeaweedFS_volumeServer_volumes{type=\"volume\"})"
	UnusedVolumeCount                = "sum(SeaweedFS_volumeServer_max_volumes)-sum(SeaweedFS_volumeServer_volumes{type=\"volume\"})"
	BucketsTotal                     = "count (count by(collection) (SeaweedFS_volumeServer_volumes{collection!=\"\"}))"
	BucketsCorrespondingVolumesNum   = "sum by (collection)(SeaweedFS_volumeServer_volumes{type=\"volume\",collection!=\"\"})"
	ProvidersCorrespondingVolumesNum = "sum by (instance)(SeaweedFS_volumeServer_volumes{type=\"volume\"})"
	S3CallTotal                      = "sum(SeaweedFS_s3_request_total)"
	S3CallWriteTotal                 = "sum(SeaweedFS_s3_request_total{type=~\"PUT|DELETE\"})"
	S3CallReadTotal                  = "sum(SeaweedFS_s3_request_total{type=~\"GET|LIST\"})"
)

type DashboardUseCase struct {
	agentRepo           AgentRepo
	gatewayRepo         GatewayRepo
	gatewayPortRepo     GatewayPortRepo
	cycleRedeemCodeRepo CycleRedeemCodeRepo
	cycleRechargeRepo   CycleRechargeRepo
	computeInstanceRepo ComputeInstanceRepo
	userRepo            UserRepo
	s3UserRepo          S3UserRepo
	dispose             *conf.Dispose
	logger              log.Logger
}

func NewDashboardUseCase(
	agentRepo AgentRepo,
	gatewayRepo GatewayRepo,
	gatewayPortRepo GatewayPortRepo,
	cycleRedeemCodeRepo CycleRedeemCodeRepo,
	cycleRechargeRepo CycleRechargeRepo,
	computeInstanceRepo ComputeInstanceRepo,
	userRepo UserRepo,
	dispose *conf.Dispose,
	s3UserRepo S3UserRepo,
	logger log.Logger,
) *DashboardUseCase {
	return &DashboardUseCase{
		agentRepo:           agentRepo,
		gatewayRepo:         gatewayRepo,
		gatewayPortRepo:     gatewayPortRepo,
		cycleRedeemCodeRepo: cycleRedeemCodeRepo,
		cycleRechargeRepo:   cycleRechargeRepo,
		computeInstanceRepo: computeInstanceRepo,
		userRepo:            userRepo,
		dispose:             dispose,
		s3UserRepo:          s3UserRepo,
		logger:              logger,
	}
}

func (d *DashboardUseCase) ProvidersCount(ctx context.Context) (count int, err error) {
	countAgent, err := d.agentRepo.CountAgent(ctx)
	if err != nil {
		return 0, err
	}
	return countAgent, nil
}

func (d *DashboardUseCase) ProvidersList(ctx context.Context) (list []*pb.ProvidersListReply_ProvidersList, err error) {
	agents, err := d.agentRepo.ListAgent(ctx)
	if err != nil {
		return nil, err
	}
	for _, agent := range agents {
		var providersList pb.ProvidersListReply_ProvidersList
		providersList.Type = consts.Agent
		providersList.Active = agent.Active
		providersList.Mac = agent.MAC
		providersList.Ip = agent.IP
		providersList.TotalCpu = agent.TotalCPU
		providersList.TotalMemory = agent.TotalMemory
		list = append(list, &providersList)
	}
	return list, nil
}

func (d *DashboardUseCase) GatewaysCount(ctx context.Context) (count int, err error) {
	countGateway, err := d.gatewayRepo.CountGateway(ctx)
	if err != nil {
		return 0, err
	}
	return countGateway, nil
}

func (d *DashboardUseCase) StoragesCount(ctx context.Context) (count *pb.StoragesCountReply_StoragesCount, err error) {
	var storagesCount pb.StoragesCountReply_StoragesCount
	//获取存储总容量
	storageSpaceTotalBytes, err := d.GetByteFromPrometheus(StorageSpaceTotal)
	if err != nil {
		return nil, err
	}
	storageSpaceTotal := ByteCountIEC(gjson.GetBytes(storageSpaceTotalBytes, "data.result.0.value.1").Int())

	//获取使用的存储总容量
	usedStorageSpaceTotalBytes, err := d.GetByteFromPrometheus(UsedStorageSpaceTotal)
	if err != nil {
		return nil, err
	}
	usedStorageSpaceTota := ByteCountIEC(gjson.GetBytes(usedStorageSpaceTotalBytes, "data.result.0.value.1").Int())
	storagesCount.StoragesTotal = usedStorageSpaceTota + "/" + storageSpaceTotal
	//获取存储提供者数量
	storageProvidersNumBytes, err := d.GetByteFromPrometheus(StorageProvidersNum)
	if err != nil {
		return nil, err
	}
	storagesCount.StorageProvidersNum = int32(gjson.GetBytes(storageProvidersNumBytes, "data.result.0.value.1").Int())
	//获取总使用的Volumes
	usedVolumeTotalBytes, err := d.GetByteFromPrometheus(UsedVolumeTotal)
	if err != nil {
		return nil, err
	}
	storagesCount.UsedVolumesTotal = int32(gjson.GetBytes(usedVolumeTotalBytes, "data.result.0.value.1").Int())
	//获取没有使用的Volumes数
	unusedVolumeCountBytes, err := d.GetByteFromPrometheus(UnusedVolumeCount)
	if err != nil {
		return nil, err
	}
	storagesCount.UnusedVolumeTotal = int32(gjson.GetBytes(unusedVolumeCountBytes, "data.result.0.value.1").Int())
	//获取总使用Buckets
	bucketsTotalBytes, err := d.GetByteFromPrometheus(BucketsTotal)
	if err != nil {
		return nil, err
	}
	storagesCount.BucketsTotal = int32(gjson.GetBytes(bucketsTotalBytes, "data.result.0.value.1").Int())
	return &storagesCount, nil
}

func (d *DashboardUseCase) StoragesProvidersList(ctx context.Context) ([]*pb.StoragesProvidersListReply_StoragesProviders, error) {
	var storagesProvidersList []*pb.StoragesProvidersListReply_StoragesProviders
	providersCorrespondingVolumesNumBytes, err := d.GetByteFromPrometheus(ProvidersCorrespondingVolumesNum)
	if err != nil {
		return nil, err
	}
	results := gjson.GetBytes(providersCorrespondingVolumesNumBytes, "data.result").Array()
	for i, manyByte := range results {
		var storagesProviders pb.StoragesProvidersListReply_StoragesProviders
		storagesProviders.Id = int32(i + 1)
		storagesProviders.Instance = manyByte.Get("metric.instance").String()
		storagesProviders.VolumeNum = int32(manyByte.Get("value.1").Int())
		storagesProvidersList = append(storagesProvidersList, &storagesProviders)
	}
	return storagesProvidersList, nil
}

func (d *DashboardUseCase) StorageBucketsVolumeNumList(ctx context.Context) ([]*pb.StorageBucketsVolumeNumListReply_BucketsVolume, error) {
	var bucketsVolumeList []*pb.StorageBucketsVolumeNumListReply_BucketsVolume
	bucketsCorrespondingVolumesNumBytes, err := d.GetByteFromPrometheus(BucketsCorrespondingVolumesNum)
	if err != nil {
		return nil, err
	}
	results := gjson.GetBytes(bucketsCorrespondingVolumesNumBytes, "data.result").Array()
	for i, manyByte := range results {
		var bucketsVolume pb.StorageBucketsVolumeNumListReply_BucketsVolume
		bucketsVolume.Id = int32(i + 1)
		bucketsVolume.Name = manyByte.Get("metric.collection").String()
		bucketsVolume.VolumeNum = int32(manyByte.Get("value.1").Int())
		bucketsVolumeList = append(bucketsVolumeList, &bucketsVolume)
	}
	return bucketsVolumeList, nil
}

func (d *DashboardUseCase) StorageS3KeyCallCount(ctx context.Context) (*pb.StorageS3KeyCallCountReply_S3KeyCallCount, error) {
	var count pb.StorageS3KeyCallCountReply_S3KeyCallCount
	s3CallTotalBytes, err := d.GetByteFromPrometheus(S3CallTotal)
	if err != nil {
		return nil, err
	}
	count.S3CallTotal = int32(gjson.GetBytes(s3CallTotalBytes, "data.result.0.value.1").Int())

	s3CallWriteTotalBytes, err := d.GetByteFromPrometheus(S3CallWriteTotal)
	if err != nil {
		return nil, err
	}
	count.S3WriteCallTotal = int32(gjson.GetBytes(s3CallWriteTotalBytes, "data.result.0.value.1").Int())

	s3CallReadTotalBytes, err := d.GetByteFromPrometheus(S3CallReadTotal)
	if err != nil {
		return nil, err
	}
	count.S3ReadCallTotal = int32(gjson.GetBytes(s3CallReadTotalBytes, "data.result.0.value.1").Int())
	//获取总s3key数
	countS3Key, err := d.s3UserRepo.CountS3Key(ctx)
	if err != nil {
		return nil, err
	}
	count.S3KeyTotal = int32(countS3Key)
	return &count, nil
}

func (d *DashboardUseCase) GetByteFromPrometheus(query string) ([]byte, error) {
	params := url.Values{}
	parseURL, err := url.Parse(d.dispose.Prometheus.Host + d.dispose.Prometheus.QueryApi)
	if err != nil {
		return nil, err
	}
	params.Set("query", query)
	parseURL.RawQuery = params.Encode()
	urlPathWithParams := parseURL.String()
	resp, err := http.Get(urlPathWithParams)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil

}

func (d *DashboardUseCase) GatewaysList(ctx context.Context) (list []*pb.GatewaysListReply_GatewaysList, err error) {
	gatewayList, err := d.gatewayRepo.ListGateway(ctx)
	if err != nil {
		return nil, err
	}
	gatewayPorts, err := d.gatewayPortRepo.CountGatewayPort(ctx)
	if err != nil {
		return nil, err
	}
	gatewayPortByIsUsed, err := d.gatewayPortRepo.CountIntranetGatewayPortByIsUsed(ctx, true)
	if err != nil {
		return nil, err
	}
	publicGatewayPortByIsUsed, err := d.gatewayPortRepo.CountPublicGatewayPortByIsUsed(ctx, true)
	if err != nil {
		return nil, err
	}
	publicGatewayPorts, err := d.gatewayPortRepo.CountPublicGatewayPort(ctx)
	if err != nil {
		return nil, err
	}
	intranetGatewayPorts, err := d.gatewayPortRepo.CountIntranetGatewayPort(ctx)
	if err != nil {
		return nil, err
	}
	gatewayPortMap := make(map[uuid.UUID]int32)
	gatewayPortByIsUsedMap := make(map[uuid.UUID]int32)
	publicGatewayPortByIsUsedMap := make(map[uuid.UUID]int32)
	publicGatewayPortMap := make(map[uuid.UUID]int32)
	intranetGatewayPortMap := make(map[uuid.UUID]int32)
	for _, gatewayPort := range gatewayPorts {
		gatewayPortMap[gatewayPort.FkGatewayID] = gatewayPort.Count
	}
	for _, gatewayPortCount := range gatewayPortByIsUsed {
		gatewayPortByIsUsedMap[gatewayPortCount.FkGatewayID] = gatewayPortCount.Count
	}
	for _, publicGatewayPortCount := range publicGatewayPortByIsUsed {
		publicGatewayPortByIsUsedMap[publicGatewayPortCount.FkGatewayID] = publicGatewayPortCount.Count
	}
	for _, publicGatewayPort := range publicGatewayPorts {
		publicGatewayPortMap[publicGatewayPort.FkGatewayID] = publicGatewayPort.Count
	}
	for _, intranetGatewayPort := range intranetGatewayPorts {
		intranetGatewayPortMap[intranetGatewayPort.FkGatewayID] = intranetGatewayPort.Count
	}
	for _, gateway := range gatewayList {
		var gatewaysList pb.GatewaysListReply_GatewaysList
		gatewaysList.Ip = gateway.IP
		gatewaysList.Name = gateway.Name
		gatewaysList.TotalPort = gatewayPortMap[gateway.ID]
		gatewaysList.UseIntranetPort = strconv.Itoa(int(gatewayPortByIsUsedMap[gateway.ID])) + " / " + strconv.Itoa(int(intranetGatewayPortMap[gateway.ID]))
		gatewaysList.UsePublicPort = strconv.Itoa(int(publicGatewayPortByIsUsedMap[gateway.ID])) + "/" + strconv.Itoa(int(publicGatewayPortMap[gateway.ID]))
		list = append(list, &gatewaysList)
	}
	return list, nil
}

func (d *DashboardUseCase) CyclesCount(ctx context.Context) (count *pb.CyclesCountReply_CyclesCount, err error) {
	var cyclesCount pb.CyclesCountReply_CyclesCount
	countCycleRecoveryTotal, err := d.cycleRedeemCodeRepo.CountCycleGrantTotal(ctx)
	if err != nil {
		return nil, err
	}
	cyclesCount.GrantTotal = countCycleRecoveryTotal.StringFixed(2)
	countCycleUseTotal, err := d.cycleRedeemCodeRepo.CountCycleUseTotal(ctx)
	if err != nil {
		return nil, err
	}
	countCycleRedeemCodeTotal, err := d.cycleRedeemCodeRepo.CountCycleRedeemCodeTotal(ctx)
	if err != nil {
		return nil, err
	}
	cyclesCount.GrantVouchersTotal = strconv.Itoa(countCycleRedeemCodeTotal)
	cyclesCount.RecoveryTotal = countCycleUseTotal.StringFixed(2)
	countRechargeCycle, err := d.cycleRechargeRepo.CountRechargeCycle(ctx)
	if err != nil {
		return nil, err
	}
	cyclesCount.RechargedTotal = countRechargeCycle.StringFixed(2)
	return &cyclesCount, nil
}

func (d *DashboardUseCase) LastComputeInstancesCount(ctx context.Context) (count []*pb.LastComputeInstancesCountReply_ComputeInstances, err error) {
	computeInstances := d.computeInstanceRepo.ListLastTop10(ctx)
	var ids []uuid.UUID
	idMap := make(map[uuid.UUID]int32)
	for _, computeInstance := range computeInstances {
		id, err := uuid.Parse(computeInstance.Owner)
		if err != nil {
			return nil, err
		}
		idMap[id] = 1
	}
	for key, _ := range idMap {
		ids = append(ids, key)
	}
	users, err := d.userRepo.ListUserByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	idNameMap := make(map[string]string)
	for _, user := range users {
		idNameMap[user.ID.String()] = user.Name
	}
	for _, instance := range computeInstances {
		var reply pb.LastComputeInstancesCountReply_ComputeInstances
		reply.Id = instance.ID.String()
		reply.Name = instance.Name
		reply.Specs = strconv.Itoa(instance.Core) + "C" + strconv.Itoa(instance.Memory) + "G"
		reply.Owner = idNameMap[instance.Owner]
		reply.CreateTime = instance.CreateTime.UnixMilli()
		count = append(count, &reply)
	}
	return count, nil
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "KMGTPE"[exp])
}
