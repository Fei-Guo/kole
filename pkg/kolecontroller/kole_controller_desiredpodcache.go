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

package kolecontroller

import (
	"sync"

	"github.com/openyurtio/kole/pkg/data"
)

type DesiredPodsCache struct {
	*sync.RWMutex
	// nodeName / pod key / Pod
	Cache map[string]map[string]*data.Pod
}

func (c *DesiredPodsCache) ReadRange(f func(nodename string, desiredPodsMap map[string]*data.Pod)) {
	c.RLock()
	for nodeName, _ := range c.Cache {
		f(nodeName, c.Cache[nodeName])
	}
	c.RUnlock()
}

func (c *DesiredPodsCache) WriteRange(f func(nodename string, desiredPodsMap map[string]*data.Pod)) {
	c.Lock()
	for nodeName, _ := range c.Cache {
		f(nodeName, c.Cache[nodeName])
	}
	c.Unlock()
}

func (c *DesiredPodsCache) SafeOperate(f func()) {
	c.RLock()
	f()
	c.RUnlock()
}

func (c *DesiredPodsCache) Len() int {
	var l int
	c.RLock()
	l = len(c.Cache)
	c.RUnlock()
	return l
}
