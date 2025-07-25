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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta2

import (
	context "context"
	time "time"

	apisetcdv1beta2 "github.com/coreos/etcd-operator/pkg/apis/etcd/v1beta2"
	versioned "github.com/coreos/etcd-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/coreos/etcd-operator/pkg/generated/informers/externalversions/internalinterfaces"
	etcdv1beta2 "github.com/coreos/etcd-operator/pkg/generated/listers/etcd/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EtcdRestoreInformer provides access to a shared informer and lister for
// EtcdRestores.
type EtcdRestoreInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() etcdv1beta2.EtcdRestoreLister
}

type etcdRestoreInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEtcdRestoreInformer constructs a new informer for EtcdRestore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEtcdRestoreInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEtcdRestoreInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEtcdRestoreInformer constructs a new informer for EtcdRestore type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEtcdRestoreInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EtcdV1beta2().EtcdRestores(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EtcdV1beta2().EtcdRestores(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EtcdV1beta2().EtcdRestores(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EtcdV1beta2().EtcdRestores(namespace).Watch(ctx, options)
			},
		},
		&apisetcdv1beta2.EtcdRestore{},
		resyncPeriod,
		indexers,
	)
}

func (f *etcdRestoreInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEtcdRestoreInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *etcdRestoreInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisetcdv1beta2.EtcdRestore{}, f.defaultInformer)
}

func (f *etcdRestoreInformer) Lister() etcdv1beta2.EtcdRestoreLister {
	return etcdv1beta2.NewEtcdRestoreLister(f.Informer().GetIndexer())
}
