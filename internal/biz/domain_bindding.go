package biz

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"github.com/mohaijiang/computeshare-server/api/global"
	pb "github.com/mohaijiang/computeshare-server/api/network_mapping/v1"
	"github.com/mohaijiang/computeshare-server/internal/utils"
	"github.com/samber/lo"
	apicorev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	"strings"
	"time"
)

const DEFAULT_NAMESPACE = "example"

type DomainBinding struct {

	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// 用户ID
	UserID uuid.UUID `json:"user_id,omitempty"`
	// 实例ID
	FkComputeInstanceID uuid.UUID `json:"fk_compute_instance_id,omitempty"`
	// 网络映射id
	FkNetworkMappingID uuid.UUID `json:"fk_network_mapping_id,omitempty"`
	// 映射名
	Name string `json:"name,omitempty"`
	// 域名
	Domain string `json:"domain,omitempty"`
	// 映射到gateway的端口
	GatewayPort int32 `json:"gateway_port,omitempty"`
	// 创建时间
	CreateTime time.Time `json:"create_time,omitempty"`
}

type DomainBindingRepository interface {
	PageQuery(ctx context.Context, userId, networkMappingId uuid.UUID, page, size int32) (*global.Page[*DomainBinding], error)
	Save(ctx context.Context, domainBinding *DomainBinding) error
	Get(ctx context.Context, id uuid.UUID) (*DomainBinding, error)
	Delete(ctx context.Context, id uuid.UUID) error
	ListByNetworkMappingId(ctx context.Context, id uuid.UUID) ([]*DomainBinding, error)
}

func NewDomainBindingUseCase(domainBindingRepository DomainBindingRepository,
	networkMappingRepo NetworkMappingRepo,
	gatewayRepo GatewayRepo,
	logger log.Logger) (*DomainBindingUseCase, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	//use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	return &DomainBindingUseCase{
		domainBindingRepository: domainBindingRepository,
		networkMappingRepo:      networkMappingRepo,
		gatewayRepo:             gatewayRepo,
		clientset:               clientset,
		log:                     log.NewHelper(logger),
	}, err
}

type DomainBindingUseCase struct {
	domainBindingRepository DomainBindingRepository
	clientset               *kubernetes.Clientset
	networkMappingRepo      NetworkMappingRepo
	gatewayRepo             GatewayRepo
	log                     *log.Helper
}

func (uc *DomainBindingUseCase) List(ctx context.Context, fkUserId, networkMappingId uuid.UUID, page, size int32) (*global.Page[*pb.DomainBindingVO], error) {
	resp, err := uc.domainBindingRepository.PageQuery(ctx, fkUserId, networkMappingId, page, size)
	if err != nil {
		return nil, err
	}
	result := global.Map[*DomainBinding, *pb.DomainBindingVO](resp, uc.toBiz)
	return result, err
}

func (uc *DomainBindingUseCase) ListByNetworkMappingId(ctx context.Context, networkMappingId uuid.UUID) ([]*pb.DomainBindingVO, error) {
	list, err := uc.domainBindingRepository.ListByNetworkMappingId(ctx, networkMappingId)
	if err != nil {
		return nil, err
	}

	return lo.Map(list, uc.toBiz), err
}

// CreateDomainBinding 创建端口绑定
func (uc *DomainBindingUseCase) CreateDomainBinding(ctx context.Context, domainBinding *DomainBinding) error {

	networkMapping, err := uc.networkMappingRepo.GetNetworkMapping(ctx, domainBinding.FkNetworkMappingID)
	if err != nil {
		return err
	}

	gateway, err := uc.gatewayRepo.GetGateway(ctx, networkMapping.FkGatewayID)
	if err != nil {
		return err
	}

	hostIp := gateway.IP
	hostPort := networkMapping.GatewayPort
	hostDomain := domainBinding.Domain

	domainBinding.FkComputeInstanceID = networkMapping.FkComputerID
	domainBinding.GatewayPort = networkMapping.GatewayPort
	domainBinding.CreateTime = time.Now()

	err = uc.domainBindingRepository.Save(ctx, domainBinding)

	ingressName := strings.ReplaceAll(domainBinding.Name, ".", "-")

	namespace, err := uc.clientset.CoreV1().Namespaces().Get(ctx, DEFAULT_NAMESPACE, v1.GetOptions{})
	if err != nil {
		namespace = &apicorev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: DEFAULT_NAMESPACE,
			},
		}
		// 创建namespace
		namespace, err = uc.clientset.CoreV1().Namespaces().Create(ctx, namespace, v1.CreateOptions{})
	}
	fmt.Println(namespace)
	if err != nil {
		return err
	}

	// 检查endpoint是否重复并创建
	endpointClient := uc.clientset.CoreV1().Endpoints(DEFAULT_NAMESPACE)
	endpoint, err := endpointClient.Get(ctx, ingressName, v1.GetOptions{})

	endpoint = &apicorev1.Endpoints{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingressName,
			Namespace: DEFAULT_NAMESPACE,
		},
		Subsets: []apicorev1.EndpointSubset{
			{
				Addresses: []apicorev1.EndpointAddress{
					{
						IP: hostIp,
					},
				},
				Ports: []apicorev1.EndpointPort{
					{
						Port: hostPort,
					},
				},
			},
		},
	}
	if err != nil {
		endpoint, err = endpointClient.Create(ctx, endpoint, v1.CreateOptions{})
		if err != nil {
			uc.log.Error(err)
			return err
		}
	} else {
		endpoint, err = endpointClient.Update(ctx, endpoint, v1.UpdateOptions{})
		if err != nil {
			uc.log.Error(err)
			return err
		}
	}

	if err != nil {
		return err
	}

	// 检查service是否重复并创建
	serviceClient := uc.clientset.CoreV1().Services(DEFAULT_NAMESPACE)
	service, err := serviceClient.Get(ctx, ingressName, v1.GetOptions{})

	service = &apicorev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingressName,
			Namespace: DEFAULT_NAMESPACE,
		},
		Spec: apicorev1.ServiceSpec{
			Ports: []apicorev1.ServicePort{
				{
					Protocol: apicorev1.ProtocolTCP,
					Port:     hostPort,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: hostPort,
					},
				},
			},
		},
	}

	if err != nil {
		service, err = serviceClient.Create(ctx, service, v1.CreateOptions{})
		if err != nil {
			uc.log.Error(err)
			return err
		}
	} else {
		service, err = serviceClient.Update(ctx, service, v1.UpdateOptions{})
		if err != nil {
			uc.log.Error(err)
			return err
		}
	}

	// 检查ingress 是否存在，并创建
	ingressClient := uc.clientset.NetworkingV1().Ingresses(DEFAULT_NAMESPACE)

	ingress, err := ingressClient.Get(ctx, ingressName, v1.GetOptions{})
	pathType := networkingv1.PathTypePrefix
	ingress = &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingressName,
			Namespace: DEFAULT_NAMESPACE,
			Annotations: map[string]string{
				"cert-manager.io/cluster-issuer":              "letsencrypt-prod",
				"nginx.ingress.kubernetes.io/proxy-body-size": "0",
			},
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: utils.ToAddress("nginx"),
			Rules: []networkingv1.IngressRule{
				{
					Host: hostDomain,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathType,
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: ingressName,
											Port: networkingv1.ServiceBackendPort{
												Number: int32(hostPort),
											},
										},
									},
								},
							},
						},
					},
				},
			},
			TLS: []networkingv1.IngressTLS{
				{
					Hosts:      []string{hostDomain},
					SecretName: ingressName,
				},
			},
		},
	}
	if err != nil {
		ingress, err = ingressClient.Create(ctx, ingress, v1.CreateOptions{})
		if err != nil {
			uc.log.Error(err)
			return err
		}
	} else {
		ingress, err = ingressClient.Update(ctx, ingress, v1.UpdateOptions{})
		if err != nil {
			uc.log.Error(err)
			return err
		}
	}

	return err
}

func (uc *DomainBindingUseCase) DeleteDomainBinding(ctx context.Context, id uuid.UUID, userId uuid.UUID) error {
	domainBinding, err := uc.domainBindingRepository.Get(ctx, id)
	if err != nil {
		return err
	}
	if userId != domainBinding.UserID {
		return errors.New("no permission")
	}

	ingressName := strings.ReplaceAll(domainBinding.Name, ".", "-")
	namespace, err := uc.clientset.CoreV1().Namespaces().Get(ctx, DEFAULT_NAMESPACE, v1.GetOptions{})
	if err != nil {
		namespace = &apicorev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: DEFAULT_NAMESPACE,
			},
		}
		// 创建namespace
		namespace, err = uc.clientset.CoreV1().Namespaces().Create(ctx, namespace, v1.CreateOptions{})

		return err
	}

	// 检查ingress 是否存在，并删除
	ingressClient := uc.clientset.NetworkingV1().Ingresses(DEFAULT_NAMESPACE)

	_, err = ingressClient.Get(ctx, ingressName, v1.GetOptions{})
	if err == nil {
		_ = ingressClient.Delete(ctx, ingressName, v1.DeleteOptions{})
	}

	// 检查service是否重复并删除
	serviceClient := uc.clientset.CoreV1().Services(DEFAULT_NAMESPACE)
	_, err = serviceClient.Get(ctx, ingressName, v1.GetOptions{})
	if err == nil {
		_ = serviceClient.Delete(ctx, ingressName, v1.DeleteOptions{})
	}

	// 检查endpoint是否重复并删除
	endpointClient := uc.clientset.CoreV1().Endpoints(DEFAULT_NAMESPACE)
	_, err = endpointClient.Get(ctx, ingressName, v1.GetOptions{})

	if err == nil {
		err = endpointClient.Delete(ctx, ingressName, v1.DeleteOptions{})
		fmt.Println(err)
	}

	err = uc.domainBindingRepository.Delete(ctx, id)

	return err
}

func (uc *DomainBindingUseCase) toBiz(item *DomainBinding, _ int) *pb.DomainBindingVO {

	if item == nil {
		return nil
	}

	return &pb.DomainBindingVO{
		Id:                item.ID.String(),
		ComputeInstanceId: item.FkComputeInstanceID.String(),
		NetworkMappingId:  item.FkNetworkMappingID.String(),
		Name:              item.Name,
		Domain:            item.Domain,
		GatewayPort:       int32(item.GatewayPort),
		CreateTime:        item.CreateTime.UnixMilli(),
	}
}

func (uc *DomainBindingUseCase) Get(ctx context.Context, id uuid.UUID) (*DomainBinding, error) {
	return uc.domainBindingRepository.Get(ctx, id)
}
