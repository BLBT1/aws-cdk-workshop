package awsec2


// Describes a port range.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   filterPortRangeProperty := &FilterPortRangeProperty{
//   	FromPort: jsii.Number(123),
//   	ToPort: jsii.Number(123),
//   }
//
type CfnNetworkInsightsPath_FilterPortRangeProperty struct {
	// The first port in the range.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The last port in the range.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

