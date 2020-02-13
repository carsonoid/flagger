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

package fake

import (
	clientset "github.com/weaveworks/flagger/pkg/client/clientset/versioned"
	appmeshv1beta1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/appmesh/v1beta1"
	fakeappmeshv1beta1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/appmesh/v1beta1/fake"
	flaggerv1beta1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/flagger/v1beta1"
	fakeflaggerv1beta1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/flagger/v1beta1/fake"
	gloov1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/gloo/v1"
	fakegloov1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/gloo/v1/fake"
	networkingv1alpha3 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/istio/v1alpha3"
	fakenetworkingv1alpha3 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/istio/v1alpha3/fake"
	projectcontourv1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/projectcontour/v1"
	fakeprojectcontourv1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/projectcontour/v1/fake"
	splitv1alpha1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/smi/v1alpha1"
	fakesplitv1alpha1 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/smi/v1alpha1/fake"
	splitv1alpha2 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/smi/v1alpha2"
	fakesplitv1alpha2 "github.com/weaveworks/flagger/pkg/client/clientset/versioned/typed/smi/v1alpha2/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// AppmeshV1beta1 retrieves the AppmeshV1beta1Client
func (c *Clientset) AppmeshV1beta1() appmeshv1beta1.AppmeshV1beta1Interface {
	return &fakeappmeshv1beta1.FakeAppmeshV1beta1{Fake: &c.Fake}
}

// FlaggerV1beta1 retrieves the FlaggerV1beta1Client
func (c *Clientset) FlaggerV1beta1() flaggerv1beta1.FlaggerV1beta1Interface {
	return &fakeflaggerv1beta1.FakeFlaggerV1beta1{Fake: &c.Fake}
}

// GlooV1 retrieves the GlooV1Client
func (c *Clientset) GlooV1() gloov1.GlooV1Interface {
	return &fakegloov1.FakeGlooV1{Fake: &c.Fake}
}

// NetworkingV1alpha3 retrieves the NetworkingV1alpha3Client
func (c *Clientset) NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface {
	return &fakenetworkingv1alpha3.FakeNetworkingV1alpha3{Fake: &c.Fake}
}

// ProjectcontourV1 retrieves the ProjectcontourV1Client
func (c *Clientset) ProjectcontourV1() projectcontourv1.ProjectcontourV1Interface {
	return &fakeprojectcontourv1.FakeProjectcontourV1{Fake: &c.Fake}
}

// SplitV1alpha1 retrieves the SplitV1alpha1Client
func (c *Clientset) SplitV1alpha1() splitv1alpha1.SplitV1alpha1Interface {
	return &fakesplitv1alpha1.FakeSplitV1alpha1{Fake: &c.Fake}
}

// SplitV1alpha2 retrieves the SplitV1alpha2Client
func (c *Clientset) SplitV1alpha2() splitv1alpha2.SplitV1alpha2Interface {
	return &fakesplitv1alpha2.FakeSplitV1alpha2{Fake: &c.Fake}
}
