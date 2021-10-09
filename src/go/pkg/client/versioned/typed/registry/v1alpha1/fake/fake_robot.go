// Copyright 2021 The Cloud Robotics Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/googlecloudrobotics/core/src/go/pkg/apis/registry/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRobots implements RobotInterface
type FakeRobots struct {
	Fake *FakeRegistryV1alpha1
	ns   string
}

var robotsResource = schema.GroupVersionResource{Group: "registry.cloudrobotics.com", Version: "v1alpha1", Resource: "robots"}

var robotsKind = schema.GroupVersionKind{Group: "registry.cloudrobotics.com", Version: "v1alpha1", Kind: "Robot"}

// Get takes name of the robot, and returns the corresponding robot object, and an error if there is any.
func (c *FakeRobots) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Robot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(robotsResource, c.ns, name), &v1alpha1.Robot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Robot), err
}

// List takes label and field selectors, and returns the list of Robots that match those selectors.
func (c *FakeRobots) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RobotList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(robotsResource, robotsKind, c.ns, opts), &v1alpha1.RobotList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RobotList{ListMeta: obj.(*v1alpha1.RobotList).ListMeta}
	for _, item := range obj.(*v1alpha1.RobotList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested robots.
func (c *FakeRobots) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(robotsResource, c.ns, opts))

}

// Create takes the representation of a robot and creates it.  Returns the server's representation of the robot, and an error, if there is any.
func (c *FakeRobots) Create(ctx context.Context, robot *v1alpha1.Robot, opts v1.CreateOptions) (result *v1alpha1.Robot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(robotsResource, c.ns, robot), &v1alpha1.Robot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Robot), err
}

// Update takes the representation of a robot and updates it. Returns the server's representation of the robot, and an error, if there is any.
func (c *FakeRobots) Update(ctx context.Context, robot *v1alpha1.Robot, opts v1.UpdateOptions) (result *v1alpha1.Robot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(robotsResource, c.ns, robot), &v1alpha1.Robot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Robot), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRobots) UpdateStatus(ctx context.Context, robot *v1alpha1.Robot, opts v1.UpdateOptions) (*v1alpha1.Robot, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(robotsResource, "status", c.ns, robot), &v1alpha1.Robot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Robot), err
}

// Delete takes name of the robot and deletes it. Returns an error if one occurs.
func (c *FakeRobots) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(robotsResource, c.ns, name), &v1alpha1.Robot{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRobots) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(robotsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RobotList{})
	return err
}

// Patch applies the patch and returns the patched robot.
func (c *FakeRobots) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Robot, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(robotsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Robot{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Robot), err
}
