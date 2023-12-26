package azure

import (
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/core"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func CreateResourceGroup(ctx *pulumi.Context) error  {

	// _, rsg_err := resources.NewResourceGroup(ctx, "resourceGroup", &resources.ResourceGroupArgs{
	// 	Location:          pulumi.String("eastus"),
	// 	ResourceGroupName: pulumi.String("pulumi-jenkins-resource-group"),
	// })
	// if rsg_err != nil {
	// 	return rsg_err
	// }

	_, rsg_err := core.NewResourceGroup(ctx, "pulumi_jenkins_resource_group", &core.ResourceGroupArgs{
		Location: pulumi.String("East US"),
	})
	if rsg_err != nil {
		return rsg_err
	}
	
	return nil
}