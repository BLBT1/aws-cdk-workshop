package awsec2


// Describes a set of filters for a path analysis.
//
// Use path filters to scope the analysis when there can be multiple resulting paths.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   pathFilterProperty := &PathFilterProperty{
//   	DestinationAddress: jsii.String("destinationAddress"),
//   	DestinationPortRange: &FilterPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   	SourceAddress: jsii.String("sourceAddress"),
//   	SourcePortRange: &FilterPortRangeProperty{
//   		FromPort: jsii.Number(123),
//   		ToPort: jsii.Number(123),
//   	},
//   }
//
type CfnNetworkInsightsPath_PathFilterProperty struct {
	// The destination IPv4 address.
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// The destination port range.
	DestinationPortRange interface{} `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// The source IPv4 address.
	SourceAddress *string `field:"optional" json:"sourceAddress" yaml:"sourceAddress"`
	// The source port range.
	SourcePortRange interface{} `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}

