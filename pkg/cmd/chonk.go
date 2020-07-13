/*
KUBECHONK
https://github.com/nickgerace/kubechonk

Apache 2.0 License
See "LICENSE" file for more information
*/

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	_ "k8s.io/client-go/plugin/pkg/client/auth" // Required for GKE and other vendors.
)

// ChonkerLists keeps track of all "chonkers" for each resource type.
type ChonkerLists struct {
	cpuList [][]string
	memList [][]string
}

// ChonkOptions provides information required to update the current context on a user's KUBECONFIG.
type ChonkOptions struct {
	configFlags *genericclioptions.ConfigFlags
	genericclioptions.IOStreams
	args []string
}

// NewChonkOptions provides an instance of ChonkOptions with default values.
func NewChonkOptions(streams genericclioptions.IOStreams) *ChonkOptions {
	return &ChonkOptions{
		configFlags: genericclioptions.NewConfigFlags(true),
		IOStreams:   streams,
	}
}

// NewCmdChonk provides a cobra command wrapping ChonkOptions.
func NewCmdChonk(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewChonkOptions(streams)

	cmd := &cobra.Command{
		Use:          "chonk [flags]",
		Short:        "Finds and displays the node(s) with the most CPU cores, and node(s) with the largest memory sizes.",
		Example:      "  kubectl chonk",
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Validate(); err != nil {
				return err
			}
			if err := o.Run(); err != nil {
				return err
			}
			return nil
		},
	}

	o.configFlags.AddFlags(cmd.Flags())

	return cmd
}

// Validate ensures that all required arguments and flag values are provided.
func (o *ChonkOptions) Validate() error {
	if len(o.args) > 0 {
		return fmt.Errorf("no arguments expected")
	}

	return nil
}

// Run finds and displays the CPU chonker node(s) and the RAM chonker node(s).
func (o *ChonkOptions) Run() error {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	// Switch on the compareWrapper function, which accounts for the "nil" condition.
	// We ignore any cases where the temp value is lesser than the max value.
	var cpuMaxValue *resource.Quantity
	var memMaxValue *resource.Quantity
	chonkerLists := ChonkerLists{}
	for _, node := range nodes.Items {
		cpuTempValue := node.Status.Capacity.Cpu()
		memTempValue := node.Status.Capacity.Memory()
		switch compareWrapper(cpuMaxValue, cpuTempValue) {
		case -1:
			cpuMaxValue = cpuTempValue
			chonkerLists.cpuList = append([][]string{}, buildChonker(node.Name, "cpu", cpuTempValue))
		case 0:
			chonkerLists.cpuList = append(chonkerLists.cpuList, buildChonker(node.Name, "cpu", cpuTempValue))
		}
		switch compareWrapper(memMaxValue, memTempValue) {
		case -1:
			memMaxValue = memTempValue
			chonkerLists.memList = append([][]string{}, buildChonker(node.Name, "memory", memTempValue))
		case 0:
			chonkerLists.memList = append(chonkerLists.memList, buildChonker(node.Name, "memory", memTempValue))
		}
	}

	// FIXME: Need to continue work here to print Memory in human readable units. Most likely
	// "GiB" and "MiB" will be necessary.
	// https://pkg.go.dev/k8s.io/apimachinery/pkg/api/resource?tab=doc#CanonicalValue
	// https://github.com/kubernetes/apimachinery/blob/v0.18.5/pkg/api/resource/quantity.go
	// https://github.com/kubernetes/apimachinery/blob/v0.18.5/pkg/api/resource/amount.go
	// The table padding is three spaces, which matches standard kubectl output.
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Node", "Resource", "Value"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("   ")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(append(chonkerLists.cpuList, chonkerLists.memList...))
	table.Render()

	return nil
}

// This function builds the chonker slice to be included in the final table.
func buildChonker(name string, resource string, value interface{}) []string {
	return []string{name, resource, fmt.Sprintf("%v", value)}
}

// This function is a wrapper around the included compare function for "Quantity". It accounts for
// the "nil" condition where the max value has not been set yet.
func compareWrapper(max *resource.Quantity, temp *resource.Quantity) int {
	if max == nil || max.Cmp(*temp) < 0 {
		return -1
	}
	if max.Cmp(*temp) == 0 {
		return 0
	}
	return 1
}
