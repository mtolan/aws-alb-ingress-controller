/*
Copyright 2017 The Kubernetes Authors.

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

package resolver

import (
	"github.com/kubernetes-sigs/aws-alb-ingress-controller/internal/ingress/controller/config"
)

// Resolver is an interface that knows how to extract information from a controller
type Resolver interface {
	// GetConfig returns the controller configuration
	GetConfig() *config.Configuration
	GetInstanceIDFromPodIP(string) (string, error)
}
