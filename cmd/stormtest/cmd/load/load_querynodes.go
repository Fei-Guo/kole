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
package load

import (
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog/v2"

	"github.com/openyurtio/kole/cmd/stormtest/cmd/options"
	"github.com/openyurtio/kole/pkg/stormtest"
)

func NewLoadQueryNodeCommand() *cobra.Command {

	ops := options.NewLoadQueryNodeFlags(&globalOptions)
	c := &cobra.Command{
		Use:   "querynode",
		Short: "The time required to test load QueryNode",
		Long:  "The time required to test load QueryNode",
		Run: func(cmd *cobra.Command, args []string) {
			klog.V(4).Infof("Stormtest load querynode %s config: %#v", ops.Name, *ops)
			if err := RunLoadQueryNode(ops); err != nil {
				klog.Fatal(err)
			}
		},
	}

	ops.AddFlags(c)
	return c
}
func init() {
	subrootCmd.AddCommand(NewLoadQueryNodeCommand())
}

func RunLoadQueryNode(opt *options.LoadQueryNodeFlags) error {
	defer runtime.HandleCrash()

	lite, err := stormtest.NewLoadQueryNode(opt)
	if err != nil {
		klog.Errorf("NewLoadQueryNode [%s] error %v", opt.Name, err)
		return err
	}
	return lite.Run()
}
