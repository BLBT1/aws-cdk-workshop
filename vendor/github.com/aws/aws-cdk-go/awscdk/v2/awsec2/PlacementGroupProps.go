package awsec2


// Props for a PlacementGroup.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   placementGroupProps := &PlacementGroupProps{
//   	Partitions: jsii.Number(123),
//   	PlacementGroupName: jsii.String("placementGroupName"),
//   	SpreadLevel: awscdk.Aws_ec2.PlacementGroupSpreadLevel_HOST,
//   	Strategy: awscdk.*Aws_ec2.PlacementGroupStrategy_CLUSTER,
//   }
//
type PlacementGroupProps struct {
	// The number of partitions.
	//
	// Valid only when Strategy is set to partition.
	Partitions *float64 `field:"optional" json:"partitions" yaml:"partitions"`
	// the name of this placement group.
	PlacementGroupName *string `field:"optional" json:"placementGroupName" yaml:"placementGroupName"`
	// Places instances on distinct hardware.
	//
	// Spread placement groups are recommended for applications
	// that have a small number of critical instances that should be kept separate from each other.
	// Launching instances in a spread level placement group reduces the risk of simultaneous failures
	// that might occur when instances share the same equipment.
	// Spread level placement groups provide access to distinct hardware,
	// and are therefore suitable for mixing instance types or launching instances over time.
	// If you start or launch an instance in a spread placement group and there is insufficient
	// unique hardware to fulfill the request, the request fails. Amazon EC2 makes more distinct hardware
	// available over time, so you can try your request again later.
	// Placement groups can spread instances across racks or hosts.
	// You can use host level spread placement groups only with AWS Outposts.
	SpreadLevel PlacementGroupSpreadLevel `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
	// Which strategy to use when launching instances.
	Strategy PlacementGroupStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

