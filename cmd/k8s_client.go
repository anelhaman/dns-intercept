package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// getK8sClient returns a Kubernetes clientset using the default kubeconfig
func getK8sClient() (*kubernetes.Clientset, error) {
	home := homedir.HomeDir()

	kubeconfig := os.Getenv("KUBECONFIG")

	if kubeconfig == "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("error building kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error creating kubernetes client: %v", err)
	}

	return clientset, nil
}

// getCoreDNSConfigMap retrieves the CoreDNS ConfigMap from the cluster
func getCoreDNSConfigMap(clientset *kubernetes.Clientset) (*corev1.ConfigMap, error) {
	return clientset.CoreV1().ConfigMaps("kube-system").Get(
		context.Background(),
		"coredns",
		metav1.GetOptions{},
	)
}

// updateCoreDNSConfigMap updates the CoreDNS ConfigMap in the cluster
func updateCoreDNSConfigMap(clientset *kubernetes.Clientset, configMap *corev1.ConfigMap) error {
	_, err := clientset.CoreV1().ConfigMaps("kube-system").Update(
		context.Background(),
		configMap,
		metav1.UpdateOptions{},
	)
	return err
}
