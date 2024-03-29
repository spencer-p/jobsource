/*
Copyright 2019 The Knative Authors.

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

package v1alpha1

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"

	"knative.dev/pkg/apis"
)

// Validate implements apis.Validatable
func (js *JobSource) Validate(ctx context.Context) *apis.FieldError {
	return js.Spec.Validate(ctx).ViaField("spec")
}

// Validate implements apis.Validatable
func (jss *JobSourceSpec) Validate(ctx context.Context) *apis.FieldError {
	if jss.Template != nil && equality.Semantic.DeepEqual(jss.Template, &corev1.PodTemplateSpec{}) {
		return apis.ErrMissingField("template")
	}
	return nil
}
