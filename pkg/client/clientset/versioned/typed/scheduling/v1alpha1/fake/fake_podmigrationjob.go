/*
Copyright 2022 The Koordinator Authors.

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

	v1alpha1 "github.com/koordinator-sh/koordinator/apis/scheduling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodMigrationJobs implements PodMigrationJobInterface
type FakePodMigrationJobs struct {
	Fake *FakeSchedulingV1alpha1
}

var podmigrationjobsResource = schema.GroupVersionResource{Group: "scheduling.koordinator.sh", Version: "v1alpha1", Resource: "podmigrationjobs"}

var podmigrationjobsKind = schema.GroupVersionKind{Group: "scheduling.koordinator.sh", Version: "v1alpha1", Kind: "PodMigrationJob"}

// Get takes name of the podMigrationJob, and returns the corresponding podMigrationJob object, and an error if there is any.
func (c *FakePodMigrationJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PodMigrationJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(podmigrationjobsResource, name), &v1alpha1.PodMigrationJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodMigrationJob), err
}

// List takes label and field selectors, and returns the list of PodMigrationJobs that match those selectors.
func (c *FakePodMigrationJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PodMigrationJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(podmigrationjobsResource, podmigrationjobsKind, opts), &v1alpha1.PodMigrationJobList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodMigrationJobList{ListMeta: obj.(*v1alpha1.PodMigrationJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.PodMigrationJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podMigrationJobs.
func (c *FakePodMigrationJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(podmigrationjobsResource, opts))
}

// Create takes the representation of a podMigrationJob and creates it.  Returns the server's representation of the podMigrationJob, and an error, if there is any.
func (c *FakePodMigrationJobs) Create(ctx context.Context, podMigrationJob *v1alpha1.PodMigrationJob, opts v1.CreateOptions) (result *v1alpha1.PodMigrationJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(podmigrationjobsResource, podMigrationJob), &v1alpha1.PodMigrationJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodMigrationJob), err
}

// Update takes the representation of a podMigrationJob and updates it. Returns the server's representation of the podMigrationJob, and an error, if there is any.
func (c *FakePodMigrationJobs) Update(ctx context.Context, podMigrationJob *v1alpha1.PodMigrationJob, opts v1.UpdateOptions) (result *v1alpha1.PodMigrationJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(podmigrationjobsResource, podMigrationJob), &v1alpha1.PodMigrationJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodMigrationJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodMigrationJobs) UpdateStatus(ctx context.Context, podMigrationJob *v1alpha1.PodMigrationJob, opts v1.UpdateOptions) (*v1alpha1.PodMigrationJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(podmigrationjobsResource, "status", podMigrationJob), &v1alpha1.PodMigrationJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodMigrationJob), err
}

// Delete takes name of the podMigrationJob and deletes it. Returns an error if one occurs.
func (c *FakePodMigrationJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(podmigrationjobsResource, name), &v1alpha1.PodMigrationJob{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodMigrationJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(podmigrationjobsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodMigrationJobList{})
	return err
}

// Patch applies the patch and returns the patched podMigrationJob.
func (c *FakePodMigrationJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PodMigrationJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(podmigrationjobsResource, name, pt, data, subresources...), &v1alpha1.PodMigrationJob{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodMigrationJob), err
}