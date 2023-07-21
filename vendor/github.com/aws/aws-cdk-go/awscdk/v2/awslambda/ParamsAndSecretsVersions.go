package awslambda


// Parameters and Secrets Extension versions.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import sm "github.com/aws/aws-cdk-go/awscdk"
//   import ssm "github.com/aws/aws-cdk-go/awscdk"
//
//
//   secret := sm.NewSecret(stack, jsii.String("Secret"))
//   parameter := ssm.NewStringParameter(stack, jsii.String("Parameter"), &StringParameterProps{
//   	ParameterName: jsii.String("mySsmParameterName"),
//   	StringValue: jsii.String("mySsmParameterValue"),
//   })
//
//   paramsAndSecrets := lambda.ParamsAndSecretsLayerVersion_FromVersion(lambda.ParamsAndSecretsVersions_V1_0_103, &ParamsAndSecretsOptions{
//   	CacheSize: jsii.Number(500),
//   	LogLevel: lamabda.paramsAndSecretsLogLevel_DEBUG,
//   })
//
//   lambdaFunction := lambda.NewFunction(this, jsii.String("MyFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_NODEJS_18_X(),
//   	Handler: jsii.String("index.handler"),
//   	Architecture: lambda.Architecture_ARM_64(),
//   	Code: lambda.Code_FromAsset(path.join(__dirname, jsii.String("lambda-handler"))),
//   	ParamsAndSecrets: ParamsAndSecrets,
//   })
//
//   secret.grantRead(lambdaFunction)
//   parameter.grantRead(lambdaFunction)
//
type ParamsAndSecretsVersions string

const (
	// Version 1.0.103.
	//
	// Note: This is the latest version.
	ParamsAndSecretsVersions_V1_0_103 ParamsAndSecretsVersions = "V1_0_103"
)

