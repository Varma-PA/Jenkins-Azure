package main

import (
	"infrastructure/azure"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return azure.CreateResourceGroup(ctx);
	})
}
