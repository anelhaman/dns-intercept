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
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	
	// Allow KUBECONFIG override
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		loadingRules.ExplicitPath = kubeconfig
	} else {
		loadingRules.ExplicitPath = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}

	configOverrides := &clientcmd.ConfigOverrides{}
	
	// Allow context override from environment
	if context := os.Getenv("KUBECONTEXT"); context != "" {
		configOverrides.CurrentContext = context
	}

	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		loadingRules,
		configOverrides,
	).ClientConfig()
	
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

// restartCoreDNS restarts the CoreDNS pods in the cluster
func restartCoreDNS(clientset *kubernetes.Clientset) error {
	// Get CoreDNS pods
	pods, err := clientset.CoreV1().Pods("kube-system").List(context.Background(), metav1.ListOptions{
		LabelSelector: "k8s-app=kube-dns",
	})
	if err != nil {
		return fmt.Errorf("failed to list CoreDNS pods: %v", err)
	}

	// Delete each pod to trigger a restart
	for _, pod := range pods.Items {
		err := clientset.CoreV1().Pods("kube-system").Delete(context.Background(), pod.Name, metav1.DeleteOptions{})
		if err != nil {
			return fmt.Errorf("failed to delete CoreDNS pod %s: %v", pod.Name, err)
		}
	}
	return nil
}
