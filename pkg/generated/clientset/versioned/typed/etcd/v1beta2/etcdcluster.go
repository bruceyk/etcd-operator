/*
Copyright 2025 The etcd-operator Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	context "context"

	etcdv1beta2 "github.com/coreos/etcd-operator/pkg/apis/etcd/v1beta2"
	applyconfigurationetcdv1beta2 "github.com/coreos/etcd-operator/pkg/generated/applyconfiguration/etcd/v1beta2"
	scheme "github.com/coreos/etcd-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// EtcdClustersGetter has a method to return a EtcdClusterInterface.
// A group's client should implement this interface.
type EtcdClustersGetter interface {
	EtcdClusters(namespace string) EtcdClusterInterface
}

// EtcdClusterInterface has methods to work with EtcdCluster resources.
type EtcdClusterInterface interface {
	Create(ctx context.Context, etcdCluster *etcdv1beta2.EtcdCluster, opts v1.CreateOptions) (*etcdv1beta2.EtcdCluster, error)
	Update(ctx context.Context, etcdCluster *etcdv1beta2.EtcdCluster, opts v1.UpdateOptions) (*etcdv1beta2.EtcdCluster, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, etcdCluster *etcdv1beta2.EtcdCluster, opts v1.UpdateOptions) (*etcdv1beta2.EtcdCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*etcdv1beta2.EtcdCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*etcdv1beta2.EtcdClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *etcdv1beta2.EtcdCluster, err error)
	Apply(ctx context.Context, etcdCluster *applyconfigurationetcdv1beta2.EtcdClusterApplyConfiguration, opts v1.ApplyOptions) (result *etcdv1beta2.EtcdCluster, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, etcdCluster *applyconfigurationetcdv1beta2.EtcdClusterApplyConfiguration, opts v1.ApplyOptions) (result *etcdv1beta2.EtcdCluster, err error)
	EtcdClusterExpansion
}

// etcdClusters implements EtcdClusterInterface
type etcdClusters struct {
	*gentype.ClientWithListAndApply[*etcdv1beta2.EtcdCluster, *etcdv1beta2.EtcdClusterList, *applyconfigurationetcdv1beta2.EtcdClusterApplyConfiguration]
}

// newEtcdClusters returns a EtcdClusters
func newEtcdClusters(c *EtcdV1beta2Client, namespace string) *etcdClusters {
	return &etcdClusters{
		gentype.NewClientWithListAndApply[*etcdv1beta2.EtcdCluster, *etcdv1beta2.EtcdClusterList, *applyconfigurationetcdv1beta2.EtcdClusterApplyConfiguration](
			"etcdclusters",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *etcdv1beta2.EtcdCluster { return &etcdv1beta2.EtcdCluster{} },
			func() *etcdv1beta2.EtcdClusterList { return &etcdv1beta2.EtcdClusterList{} },
		),
	}
}
