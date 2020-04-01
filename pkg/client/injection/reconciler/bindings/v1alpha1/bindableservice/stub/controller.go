/*
Copyright 2019 the original author or authors.

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

// Code generated by injection-gen. DO NOT EDIT.

package bindableservice

import (
	context "context"

	bindableservice "github.com/projectriff/bindings/pkg/client/injection/informers/bindings/v1alpha1/bindableservice"
	v1alpha1bindableservice "github.com/projectriff/bindings/pkg/client/injection/reconciler/bindings/v1alpha1/bindableservice"
	configmap "knative.dev/pkg/configmap"
	controller "knative.dev/pkg/controller"
	logging "knative.dev/pkg/logging"
)

// TODO: PLEASE COPY AND MODIFY THIS FILE AS A STARTING POINT

// NewController creates a Reconciler for BindableService and returns the result of NewImpl.
func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	logger := logging.FromContext(ctx)

	bindableserviceInformer := bindableservice.Get(ctx)

	// TODO: setup additional informers here.

	r := &Reconciler{}
	impl := v1alpha1bindableservice.NewImpl(ctx, r)

	logger.Info("Setting up event handlers.")

	bindableserviceInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	// TODO: add additional informer event handlers here.

	return impl
}
