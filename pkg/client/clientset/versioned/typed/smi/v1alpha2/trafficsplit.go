/*
Copyright The Flagger Authors.

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

package v1alpha2

import (
	"time"

	v1alpha2 "github.com/weaveworks/flagger/pkg/apis/smi/v1alpha2"
	scheme "github.com/weaveworks/flagger/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TrafficSplitsGetter has a method to return a TrafficSplitInterface.
// A group's client should implement this interface.
type TrafficSplitsGetter interface {
	TrafficSplits(namespace string) TrafficSplitInterface
}

// TrafficSplitInterface has methods to work with TrafficSplit resources.
type TrafficSplitInterface interface {
	Create(*v1alpha2.TrafficSplit) (*v1alpha2.TrafficSplit, error)
	Update(*v1alpha2.TrafficSplit) (*v1alpha2.TrafficSplit, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.TrafficSplit, error)
	List(opts v1.ListOptions) (*v1alpha2.TrafficSplitList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.TrafficSplit, err error)
	TrafficSplitExpansion
}

// trafficSplits implements TrafficSplitInterface
type trafficSplits struct {
	client rest.Interface
	ns     string
}

// newTrafficSplits returns a TrafficSplits
func newTrafficSplits(c *SplitV1alpha2Client, namespace string) *trafficSplits {
	return &trafficSplits{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the trafficSplit, and returns the corresponding trafficSplit object, and an error if there is any.
func (c *trafficSplits) Get(name string, options v1.GetOptions) (result *v1alpha2.TrafficSplit, err error) {
	result = &v1alpha2.TrafficSplit{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("trafficsplits").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TrafficSplits that match those selectors.
func (c *trafficSplits) List(opts v1.ListOptions) (result *v1alpha2.TrafficSplitList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.TrafficSplitList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("trafficsplits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested trafficSplits.
func (c *trafficSplits) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("trafficsplits").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a trafficSplit and creates it.  Returns the server's representation of the trafficSplit, and an error, if there is any.
func (c *trafficSplits) Create(trafficSplit *v1alpha2.TrafficSplit) (result *v1alpha2.TrafficSplit, err error) {
	result = &v1alpha2.TrafficSplit{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("trafficsplits").
		Body(trafficSplit).
		Do().
		Into(result)
	return
}

// Update takes the representation of a trafficSplit and updates it. Returns the server's representation of the trafficSplit, and an error, if there is any.
func (c *trafficSplits) Update(trafficSplit *v1alpha2.TrafficSplit) (result *v1alpha2.TrafficSplit, err error) {
	result = &v1alpha2.TrafficSplit{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("trafficsplits").
		Name(trafficSplit.Name).
		Body(trafficSplit).
		Do().
		Into(result)
	return
}

// Delete takes name of the trafficSplit and deletes it. Returns an error if one occurs.
func (c *trafficSplits) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("trafficsplits").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *trafficSplits) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("trafficsplits").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched trafficSplit.
func (c *trafficSplits) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.TrafficSplit, err error) {
	result = &v1alpha2.TrafficSplit{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("trafficsplits").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
