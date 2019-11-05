package cmd

import (
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
	"io"
)

// NewCommand returns cobra.Command to run command
func NewCommand(in io.Reader, out, err io.Writer) *cobra.Command {

	cmds := &cobra.Command{
		Use:   "kubeadm",
		Short: "kubeadm: easily bootstrap a secure Kubernetes cluster",
		Long: dedent.Dedent(`

			    ┌──────────────────────────────────────────────────────────┐
			    │ KUBEADM                                                  │
			    │ Easily bootstrap a secure Kubernetes cluster             │
			    │                                                          │
			    │ Please give us feedback at:                              │
			    │ https://github.com/kubernetes/kubeadm/issues             │
			    └──────────────────────────────────────────────────────────┘

			Example usage:

			    Create a two-machine cluster with one control-plane node
			    (which controls the cluster), and one worker node
			    (where your workloads, like Pods and Deployments run).

			    ┌──────────────────────────────────────────────────────────┐
			    │ On the first machine:                                    │
			    ├──────────────────────────────────────────────────────────┤
			    │ control-plane# kubeadm init                              │
			    └──────────────────────────────────────────────────────────┘

			    ┌──────────────────────────────────────────────────────────┐
			    │ On the second machine:                                   │
			    ├──────────────────────────────────────────────────────────┤
			    │ worker# kubeadm join <arguments-returned-from-init>      │
			    └──────────────────────────────────────────────────────────┘

			    You can then repeat the second step on as many other machines as you like.

		`),

		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	cmds.ResetFlags()

	//cmds.AddCommand(alpha.NewCmdAlpha(in, out))
	return cmds
}
