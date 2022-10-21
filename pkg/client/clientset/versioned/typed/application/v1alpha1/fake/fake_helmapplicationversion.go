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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubesphere.io/openpitrix/pkg/api/application/v1alpha1"
)

// FakeHelmApplicationVersions implements HelmApplicationVersionInterface
type FakeHelmApplicationVersions struct {
	Fake *FakeApplicationV1alpha1
}

var helmapplicationversionsResource = schema.GroupVersionResource{Group: "application.kubesphere.io", Version: "v1alpha1", Resource: "helmapplicationversions"}

var helmapplicationversionsKind = schema.GroupVersionKind{Group: "application.kubesphere.io", Version: "v1alpha1", Kind: "HelmApplicationVersion"}

// Get takes name of the helmApplicationVersion, and returns the corresponding helmApplicationVersion object, and an error if there is any.
func (c *FakeHelmApplicationVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HelmApplicationVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(helmapplicationversionsResource, name), &v1alpha1.HelmApplicationVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplicationVersion), err
}

// List takes label and field selectors, and returns the list of HelmApplicationVersions that match those selectors.
func (c *FakeHelmApplicationVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HelmApplicationVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(helmapplicationversionsResource, helmapplicationversionsKind, opts), &v1alpha1.HelmApplicationVersionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HelmApplicationVersionList{ListMeta: obj.(*v1alpha1.HelmApplicationVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.HelmApplicationVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested helmApplicationVersions.
func (c *FakeHelmApplicationVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(helmapplicationversionsResource, opts))
}

// Create takes the representation of a helmApplicationVersion and creates it.  Returns the server's representation of the helmApplicationVersion, and an error, if there is any.
func (c *FakeHelmApplicationVersions) Create(ctx context.Context, helmApplicationVersion *v1alpha1.HelmApplicationVersion, opts v1.CreateOptions) (result *v1alpha1.HelmApplicationVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(helmapplicationversionsResource, helmApplicationVersion), &v1alpha1.HelmApplicationVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplicationVersion), err
}

// Update takes the representation of a helmApplicationVersion and updates it. Returns the server's representation of the helmApplicationVersion, and an error, if there is any.
func (c *FakeHelmApplicationVersions) Update(ctx context.Context, helmApplicationVersion *v1alpha1.HelmApplicationVersion, opts v1.UpdateOptions) (result *v1alpha1.HelmApplicationVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(helmapplicationversionsResource, helmApplicationVersion), &v1alpha1.HelmApplicationVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplicationVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHelmApplicationVersions) UpdateStatus(ctx context.Context, helmApplicationVersion *v1alpha1.HelmApplicationVersion, opts v1.UpdateOptions) (*v1alpha1.HelmApplicationVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(helmapplicationversionsResource, "status", helmApplicationVersion), &v1alpha1.HelmApplicationVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplicationVersion), err
}

// Delete takes name of the helmApplicationVersion and deletes it. Returns an error if one occurs.
func (c *FakeHelmApplicationVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(helmapplicationversionsResource, name), &v1alpha1.HelmApplicationVersion{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHelmApplicationVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(helmapplicationversionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HelmApplicationVersionList{})
	return err
}

// Patch applies the patch and returns the patched helmApplicationVersion.
func (c *FakeHelmApplicationVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HelmApplicationVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(helmapplicationversionsResource, name, pt, data, subresources...), &v1alpha1.HelmApplicationVersion{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HelmApplicationVersion), err
}
