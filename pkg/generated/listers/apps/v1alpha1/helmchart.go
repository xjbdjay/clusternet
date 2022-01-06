/*
Copyright The Clusternet Authors.

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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HelmChartLister helps list HelmCharts.
// All objects returned here must be treated as read-only.
type HelmChartLister interface {
	// List lists all HelmCharts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HelmChart, err error)
	// HelmCharts returns an object that can list and get HelmCharts.
	HelmCharts(namespace string) HelmChartNamespaceLister
	HelmChartListerExpansion
}

// helmChartLister implements the HelmChartLister interface.
type helmChartLister struct {
	indexer cache.Indexer
}

// NewHelmChartLister returns a new HelmChartLister.
func NewHelmChartLister(indexer cache.Indexer) HelmChartLister {
	return &helmChartLister{indexer: indexer}
}

// List lists all HelmCharts in the indexer.
func (s *helmChartLister) List(selector labels.Selector) (ret []*v1alpha1.HelmChart, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HelmChart))
	})
	return ret, err
}

// HelmCharts returns an object that can list and get HelmCharts.
func (s *helmChartLister) HelmCharts(namespace string) HelmChartNamespaceLister {
	return helmChartNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HelmChartNamespaceLister helps list and get HelmCharts.
// All objects returned here must be treated as read-only.
type HelmChartNamespaceLister interface {
	// List lists all HelmCharts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HelmChart, err error)
	// Get retrieves the HelmChart from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HelmChart, error)
	HelmChartNamespaceListerExpansion
}

// helmChartNamespaceLister implements the HelmChartNamespaceLister
// interface.
type helmChartNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HelmCharts in the indexer for a given namespace.
func (s helmChartNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HelmChart, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HelmChart))
	})
	return ret, err
}

// Get retrieves the HelmChart from the indexer for a given namespace and name.
func (s helmChartNamespaceLister) Get(name string) (*v1alpha1.HelmChart, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("helmchart"), name)
	}
	return obj.(*v1alpha1.HelmChart), nil
}
