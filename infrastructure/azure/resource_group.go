package azure

import (
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateResourceGroup(ctx *pulumi.Context) (*core.ResourceGroup, error)   {

	rsg, rsg_err := core.NewResourceGroup(ctx, "pulumi_jenkins_resource_group", &core.ResourceGroupArgs{
		Location: pulumi.String("East US"),
	})
	
	return rsg, rsg_err
}