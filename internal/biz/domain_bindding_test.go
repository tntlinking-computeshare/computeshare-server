package biz

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/tj/assert"
	apicorev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func GetDomainBindingUseCase() *DomainBindingUseCase {

	logger := log.With(log.NewStdLogger(os.Stdout),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", 1,
		"service.name", "test",
		"service.version", "Testv1",
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
	uc, err := NewDomainBindingUseCase(nil, nil, nil, logger)
	if err != nil {
		panic(err)
	}
	return uc
}

func TestDomainBindingUseCase_CreateDomainBinding(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)
	uc := GetDomainBindingUseCase()
	err := uc.CreateDomainBinding(ctx, &DomainBinding{})

	assert.NoError(t, err)
}

func TestDomainBindingUseCase_DeleteDomainBinding(t *testing.T) {
	ctx := context.Background()
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config_computeshare"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	//use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	ingressName := strings.ReplaceAll("nginx.mohaijiang.computeshare.newtouch.com", ".", "-")
	namespace, err := clientset.CoreV1().Namespaces().Get(ctx, DEFAULT_NAMESPACE, v1.GetOptions{})
	if err != nil {
		namespace = &apicorev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: DEFAULT_NAMESPACE,
			},
		}
		// 创建namespace
		namespace, err = clientset.CoreV1().Namespaces().Create(ctx, namespace, v1.CreateOptions{})

	}

	// 检查ingress 是否存在，并删除
	ingressClient := clientset.NetworkingV1().Ingresses(DEFAULT_NAMESPACE)

	ingress, err := ingressClient.Get(ctx, ingressName, v1.GetOptions{})
	fmt.Println(ingress)
	if err != nil {
		err = ingressClient.Delete(ctx, ingressName, v1.DeleteOptions{})
	}

	// 检查service是否重复并删除
	serviceClient := clientset.CoreV1().Services(DEFAULT_NAMESPACE)
	_, err = serviceClient.Get(ctx, ingressName, v1.GetOptions{})
	if err != nil {
		_ = serviceClient.Delete(ctx, ingressName, v1.DeleteOptions{})
	}

	// 检查endpoint是否重复并删除
	endpointClient := clientset.CoreV1().Endpoints(DEFAULT_NAMESPACE)
	_, err = endpointClient.Get(ctx, ingressName, v1.GetOptions{})

	if err == nil {
		err = endpointClient.Delete(ctx, ingressName, v1.DeleteOptions{})
		fmt.Println(err)
	}

}
