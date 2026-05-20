package client

import (
	"testing"
	"wae/config"
)

func init() {
	err := config.LoadConfig("../config.yaml")
	if err != nil {
		panic(err)
	}
}

func TestGetNamespaces(t *testing.T) {
	kc := config.Config.KubeConfigs[0]
	client, err := NewClient(kc)
	if err != nil {
		t.Errorf("NewClient failed: %v", err)
	}
	namespaces, err := client.GetClusterNamespaces()
	if err != nil {
		t.Errorf("GetClusterNamespaces failed: %v", err)
	}
	if len(namespaces) == 0 {
		t.Errorf("GetClusterNamespaces failed: no namespaces")
	}
	for _, ns := range namespaces {
		t.Logf("namespace: %s", ns.Name)
	}
}

func TestCreateAndDeleteNamespace(t *testing.T) {
	kc := config.Config.KubeConfigs[0]
	client, err := NewClient(kc)
	if err != nil {
		t.Errorf("NewClient failed: %v", err)
	}
	err = client.CreateNamespace("test-namespace")
	if err != nil {
		t.Errorf("CreateNamespace failed: %v", err)
	}
	err = client.DeleteNamespace("test-namespace")
	if err != nil {
		t.Errorf("DeleteNamespace failed: %v", err)
	}
}
