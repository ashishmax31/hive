// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/hive/apis/hive/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterDeploymentCustomizationLister helps list ClusterDeploymentCustomizations.
// All objects returned here must be treated as read-only.
type ClusterDeploymentCustomizationLister interface {
	// List lists all ClusterDeploymentCustomizations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterDeploymentCustomization, err error)
	// ClusterDeploymentCustomizations returns an object that can list and get ClusterDeploymentCustomizations.
	ClusterDeploymentCustomizations(namespace string) ClusterDeploymentCustomizationNamespaceLister
	ClusterDeploymentCustomizationListerExpansion
}

// clusterDeploymentCustomizationLister implements the ClusterDeploymentCustomizationLister interface.
type clusterDeploymentCustomizationLister struct {
	indexer cache.Indexer
}

// NewClusterDeploymentCustomizationLister returns a new ClusterDeploymentCustomizationLister.
func NewClusterDeploymentCustomizationLister(indexer cache.Indexer) ClusterDeploymentCustomizationLister {
	return &clusterDeploymentCustomizationLister{indexer: indexer}
}

// List lists all ClusterDeploymentCustomizations in the indexer.
func (s *clusterDeploymentCustomizationLister) List(selector labels.Selector) (ret []*v1.ClusterDeploymentCustomization, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterDeploymentCustomization))
	})
	return ret, err
}

// ClusterDeploymentCustomizations returns an object that can list and get ClusterDeploymentCustomizations.
func (s *clusterDeploymentCustomizationLister) ClusterDeploymentCustomizations(namespace string) ClusterDeploymentCustomizationNamespaceLister {
	return clusterDeploymentCustomizationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterDeploymentCustomizationNamespaceLister helps list and get ClusterDeploymentCustomizations.
// All objects returned here must be treated as read-only.
type ClusterDeploymentCustomizationNamespaceLister interface {
	// List lists all ClusterDeploymentCustomizations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterDeploymentCustomization, err error)
	// Get retrieves the ClusterDeploymentCustomization from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterDeploymentCustomization, error)
	ClusterDeploymentCustomizationNamespaceListerExpansion
}

// clusterDeploymentCustomizationNamespaceLister implements the ClusterDeploymentCustomizationNamespaceLister
// interface.
type clusterDeploymentCustomizationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterDeploymentCustomizations in the indexer for a given namespace.
func (s clusterDeploymentCustomizationNamespaceLister) List(selector labels.Selector) (ret []*v1.ClusterDeploymentCustomization, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterDeploymentCustomization))
	})
	return ret, err
}

// Get retrieves the ClusterDeploymentCustomization from the indexer for a given namespace and name.
func (s clusterDeploymentCustomizationNamespaceLister) Get(name string) (*v1.ClusterDeploymentCustomization, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clusterdeploymentcustomization"), name)
	}
	return obj.(*v1.ClusterDeploymentCustomization), nil
}
