package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a LogGroup.
//
// Example:
//   var vpc vpc
//
//   kmsKey := kms.NewKey(this, jsii.String("KmsKey"))
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the log group
//   logGroup := logs.NewLogGroup(this, jsii.String("LogGroup"), &LogGroupProps{
//   	EncryptionKey: kmsKey,
//   })
//
//   // Pass the KMS key in the `encryptionKey` field to associate the key to the S3 bucket
//   execBucket := s3.NewBucket(this, jsii.String("EcsExecBucket"), &BucketProps{
//   	EncryptionKey: kmsKey,
//   })
//
//   cluster := ecs.NewCluster(this, jsii.String("Cluster"), &ClusterProps{
//   	Vpc: Vpc,
//   	ExecuteCommandConfiguration: &ExecuteCommandConfiguration{
//   		KmsKey: *KmsKey,
//   		LogConfiguration: &ExecuteCommandLogConfiguration{
//   			CloudWatchLogGroup: logGroup,
//   			CloudWatchEncryptionEnabled: jsii.Boolean(true),
//   			S3Bucket: execBucket,
//   			S3EncryptionEnabled: jsii.Boolean(true),
//   			S3KeyPrefix: jsii.String("exec-command-output"),
//   		},
//   		Logging: ecs.ExecuteCommandLogging_OVERRIDE,
//   	},
//   })
//
type LogGroupProps struct {
	// Data Protection Policy for this log group.
	DataProtectionPolicy DataProtectionPolicy `field:"optional" json:"dataProtectionPolicy" yaml:"dataProtectionPolicy"`
	// The KMS customer managed key to encrypt the log group with.
	EncryptionKey awskms.IKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Name of the log group.
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// Determine the removal policy of this log group.
	//
	// Normally you want to retain the log group so you can diagnose issues
	// from logs even after a deployment that no longer includes the log group.
	// In that case, use the normal date-based retention policy to age out your
	// logs.
	RemovalPolicy awscdk.RemovalPolicy `field:"optional" json:"removalPolicy" yaml:"removalPolicy"`
	// How long, in days, the log contents will be retained.
	//
	// To retain all logs, set this value to RetentionDays.INFINITE.
	Retention RetentionDays `field:"optional" json:"retention" yaml:"retention"`
}

