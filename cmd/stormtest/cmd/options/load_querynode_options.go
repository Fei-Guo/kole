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
package options

import "github.com/spf13/cobra"

type LoadQueryNodeFlags struct {
	*GlobalFlags
	Ns   string
	Name string
}

func NewLoadQueryNodeFlags(g *GlobalFlags) *LoadQueryNodeFlags {
	return &LoadQueryNodeFlags{
		GlobalFlags: g,
		Ns:          "infedge",
		Name:        "query-node",
	}
}

func (f *LoadQueryNodeFlags) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.Ns, "namespace", f.Ns, "Namespace where cr exist")
	cmd.Flags().StringVar(&f.Name, "name", f.Name, "the name of cr")
}
