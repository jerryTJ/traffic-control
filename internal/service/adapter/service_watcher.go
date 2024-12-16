package adapter

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func RunServiceInformer(clientset *kubernetes.Clientset, stopCh <-chan struct{}) {
	// 1. 创建 SharedInformerFactory
	factory := informers.NewSharedInformerFactory(clientset, time.Minute)
	// 2. 获取 Service informer
	serviceInformer := factory.Core().V1().Services().Informer()
	// 3. 添加事件处理函数
	serviceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			service := obj.(*corev1.Service)
			fmt.Printf("New Service Created: Name=%s, Namespace=%s, ClusterIP=%s\n",
				service.Name, service.Namespace, service.Spec.ClusterIP)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldService := oldObj.(*corev1.Service)
			newService := newObj.(*corev1.Service)
			fmt.Printf("Service Updated: Name=%s, Old ClusterIP=%s, New ClusterIP=%s\n",
				oldService.Name, oldService.Spec.ClusterIP, newService.Spec.ClusterIP)
		},
		DeleteFunc: func(obj interface{}) {
			service := obj.(*corev1.Service)
			fmt.Printf("Service Deleted: Name=%s, Namespace=%s\n",
				service.Name, service.Namespace)
		},
	})

	// 4. 捕获运行时错误
	defer runtime.HandleCrash()

	// 5. 启动 informer
	fmt.Println("Starting Service Informer...")
	factory.Start(stopCh)

	// 6. 等待缓存同步
	if !cache.WaitForCacheSync(stopCh, serviceInformer.HasSynced) {
		log.Fatalf("Failed to sync caches")
	}

	// 7. 保持 Goroutine 运行，直到 stopCh 关闭
	<-stopCh
	fmt.Println("Stopping Service Informer...")
}
func Start(kubeconfig string) {
	var clientset *kubernetes.Clientset
	if kubeconfig != "" {
		// 加载 KubeConfig
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			log.Fatalf("Error building kubeconfig: %v", err)
		}

		// 创建 Kubernetes 客户端
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalf("Error creating Kubernetes client: %v", err)
		}

	} else {
		// 使用集群内的配置
		config, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalf("Error building in-cluster config: %v", err)
		}

		// 创建 Kubernetes 客户端
		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalf("Error creating Kubernetes client: %v", err)
		}
	}

	// 创建 stopCh 通道，用于控制 informer 停止
	stopCh := make(chan struct{})

	// 捕获系统信号，优雅关闭程序
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	go func() {
		<-signalCh
		fmt.Println("Received termination signal. Shutting down...")
		close(stopCh)
	}()

	// 5. 在 Goroutine 中运行 informer
	go RunServiceInformer(clientset, stopCh)

	// 6. 主线程等待 stopCh 关闭
	//<-stopCh
	//fmt.Println("Program stopped.")
}
