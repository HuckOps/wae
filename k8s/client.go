package client

import (
	"context"
	"sync"
	"wae/config"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	ClientSet *kubernetes.Clientset
	mutex     sync.Mutex
	ctx       context.Context
}

func NewClient(kc config.KubeConfig) (*Client, error) {
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", kc.Path)
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	return &Client{ClientSet: clientset, ctx: context.Background()}, nil
}

func (c *Client) GetClusterNamespaces() ([]v1.Namespace, error) {
	namespaces, err := c.ClientSet.CoreV1().Namespaces().List(c.ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces.Items, nil
}

func (c *Client) CreateNamespace(namespaceName string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
			Labels: map[string]string{
				"managed-by": "wae",
				"created-at": metav1.Now().Format("20060102150405"),
				"updated-at": metav1.Now().Format("20060102150405"),
			},
		},
	}
	_, err := c.ClientSet.CoreV1().Namespaces().Create(c.ctx, namespace, metav1.CreateOptions{})
	return err
}

func (c *Client) DeleteNamespace(namespaceName string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.ClientSet.CoreV1().Namespaces().Delete(c.ctx, namespaceName, metav1.DeleteOptions{})
}
