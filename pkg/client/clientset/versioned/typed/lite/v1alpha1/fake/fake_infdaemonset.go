/*
Copyright 2022 The OpenYurt Authors.

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
	"context"

	v1alpha1 "github.com/openyurtio/kole/pkg/apis/lite/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInfDaemonSets implements InfDaemonSetInterface
type FakeInfDaemonSets struct {
	Fake *FakeLiteV1alpha1
	ns   string
}

var infdaemonsetsResource = schema.GroupVersionResource{Group: "lite.openyurt.io", Version: "v1alpha1", Resource: "infdaemonsets"}

var infdaemonsetsKind = schema.GroupVersionKind{Group: "lite.openyurt.io", Version: "v1alpha1", Kind: "InfDaemonSet"}

// Get takes name of the infDaemonSet, and returns the corresponding infDaemonSet object, and an error if there is any.
func (c *FakeInfDaemonSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InfDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(infdaemonsetsResource, c.ns, name), &v1alpha1.InfDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InfDaemonSet), err
}

// List takes label and field selectors, and returns the list of InfDaemonSets that match those selectors.
func (c *FakeInfDaemonSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InfDaemonSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(infdaemonsetsResource, infdaemonsetsKind, c.ns, opts), &v1alpha1.InfDaemonSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InfDaemonSetList{ListMeta: obj.(*v1alpha1.InfDaemonSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.InfDaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested infDaemonSets.
func (c *FakeInfDaemonSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(infdaemonsetsResource, c.ns, opts))

}

// Create takes the representation of a infDaemonSet and creates it.  Returns the server's representation of the infDaemonSet, and an error, if there is any.
func (c *FakeInfDaemonSets) Create(ctx context.Context, infDaemonSet *v1alpha1.InfDaemonSet, opts v1.CreateOptions) (result *v1alpha1.InfDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(infdaemonsetsResource, c.ns, infDaemonSet), &v1alpha1.InfDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InfDaemonSet), err
}

// Update takes the representation of a infDaemonSet and updates it. Returns the server's representation of the infDaemonSet, and an error, if there is any.
func (c *FakeInfDaemonSets) Update(ctx context.Context, infDaemonSet *v1alpha1.InfDaemonSet, opts v1.UpdateOptions) (result *v1alpha1.InfDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(infdaemonsetsResource, c.ns, infDaemonSet), &v1alpha1.InfDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InfDaemonSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInfDaemonSets) UpdateStatus(ctx context.Context, infDaemonSet *v1alpha1.InfDaemonSet, opts v1.UpdateOptions) (*v1alpha1.InfDaemonSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(infdaemonsetsResource, "status", c.ns, infDaemonSet), &v1alpha1.InfDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InfDaemonSet), err
}

// Delete takes name of the infDaemonSet and deletes it. Returns an error if one occurs.
func (c *FakeInfDaemonSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(infdaemonsetsResource, c.ns, name), &v1alpha1.InfDaemonSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInfDaemonSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(infdaemonsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InfDaemonSetList{})
	return err
}

// Patch applies the patch and returns the patched infDaemonSet.
func (c *FakeInfDaemonSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InfDaemonSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(infdaemonsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.InfDaemonSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InfDaemonSet), err
}
