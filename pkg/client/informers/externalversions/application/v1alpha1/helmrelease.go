/*
Copyright 2020 The KubeSphere Authors.

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

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	applicationv1alpha1 "kubesphere.io/openpitrix/pkg/api/application/v1alpha1"
	versioned "kubesphere.io/openpitrix/pkg/client/clientset/versioned"
	internalinterfaces "kubesphere.io/openpitrix/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubesphere.io/openpitrix/pkg/client/listers/application/v1alpha1"
)

// HelmReleaseInformer provides access to a shared informer and lister for
// HelmReleases.
type HelmReleaseInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HelmReleaseLister
}

type helmReleaseInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewHelmReleaseInformer constructs a new informer for HelmRelease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHelmReleaseInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHelmReleaseInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredHelmReleaseInformer constructs a new informer for HelmRelease type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHelmReleaseInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApplicationV1alpha1().HelmReleases().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApplicationV1alpha1().HelmReleases().Watch(context.TODO(), options)
			},
		},
		&applicationv1alpha1.HelmRelease{},
		resyncPeriod,
		indexers,
	)
}

func (f *helmReleaseInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHelmReleaseInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *helmReleaseInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&applicationv1alpha1.HelmRelease{}, f.defaultInformer)
}

func (f *helmReleaseInformer) Lister() v1alpha1.HelmReleaseLister {
	return v1alpha1.NewHelmReleaseLister(f.Informer().GetIndexer())
}
