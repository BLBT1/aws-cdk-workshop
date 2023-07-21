package awsec2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an Amazon EC2 security group within a VPC.
//
// Security Groups act like a firewall with a set of rules, and are associated
// with any AWS resource that has or creates Elastic Network Interfaces (ENIs).
// A typical example of a resource that has a security group is an Instance (or
// Auto Scaling Group of instances)
//
// If you are defining new infrastructure in CDK, there is a good chance you
// won't have to interact with this class at all. Like IAM Roles, Security
// Groups need to exist to control access between AWS resources, but CDK will
// automatically generate and populate them with least-privilege permissions
// for you so you can concentrate on your business logic.
//
// All Constructs that require Security Groups will create one for you if you
// don't specify one at construction. After construction, you can selectively
// allow connections to and between constructs via--for example-- the `instance.connections`
// object. Think of it as "allowing connections to your instance", rather than
// "adding ingress rules a security group". See the [Allowing
// Connections](https://docs.aws.amazon.com/cdk/api/latest/docs/aws-cdk-lib.aws_ec2-readme.html#allowing-connections)
// section in the library documentation for examples.
//
// Direct manipulation of the Security Group through `addIngressRule` and
// `addEgressRule` is possible, but mutation through the `.connections` object
// is recommended. If you peer two constructs with security groups this way,
// appropriate rules will be created in both.
//
// If you have an existing security group you want to use in your CDK application,
// you would import it like this:
//
// ```ts
// const securityGroup = ec2.SecurityGroup.fromSecurityGroupId(this, 'SG', 'sg-12345', {
//   mutable: false
// });
// ```.
//
// Example:
//   var vpc vpc
//
//
//   template := ec2.NewLaunchTemplate(this, jsii.String("LaunchTemplate"), &LaunchTemplateProps{
//   	MachineImage: ec2.MachineImage_LatestAmazonLinux2022(),
//   	SecurityGroup: ec2.NewSecurityGroup(this, jsii.String("LaunchTemplateSG"), &SecurityGroupProps{
//   		Vpc: vpc,
//   	}),
//   })
//
type SecurityGroup interface {
	awscdk.Resource
	ISecurityGroup
	// Whether the SecurityGroup has been configured to allow all outbound ipv6 traffic.
	AllowAllIpv6Outbound() *bool
	// Whether the SecurityGroup has been configured to allow all outbound traffic.
	AllowAllOutbound() *bool
	// Whether the rule can be inlined into a SecurityGroup or not.
	CanInlineRule() *bool
	// The network connections associated with this resource.
	Connections() Connections
	DefaultPort() Port
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// The tree node.
	Node() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The ID of the security group.
	SecurityGroupId() *string
	// The VPC ID this security group is part of.
	SecurityGroupVpcId() *string
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// A unique identifier for this connection peer.
	UniqueId() *string
	// Add an egress rule for the current security group.
	//
	// `remoteRule` controls where the Rule object is created if the peer is also a
	// securityGroup and they are in different stack. If false (default) the
	// rule object is created under the current SecurityGroup object. If true and the
	// peer is also a SecurityGroup, the rule object is created under the remote
	// SecurityGroup object.
	AddEgressRule(peer IPeer, connection Port, description *string, remoteRule *bool)
	// Add an ingress rule for the current security group.
	//
	// `remoteRule` controls where the Rule object is created if the peer is also a
	// securityGroup and they are in different stack. If false (default) the
	// rule object is created under the current SecurityGroup object. If true and the
	// peer is also a SecurityGroup, the rule object is created under the remote
	// SecurityGroup object.
	AddIngressRule(peer IPeer, connection Port, description *string, remoteRule *bool)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Determine where to parent a new ingress/egress rule.
	//
	// A SecurityGroup rule is parented under the group it's related to, UNLESS
	// we're in a cross-stack scenario with another Security Group. In that case,
	// we respect the 'remoteRule' flag and will parent under the other security
	// group.
	//
	// This is necessary to avoid cyclic dependencies between stacks, since both
	// ingress and egress rules will reference both security groups, and a naive
	// parenting will lead to the following situation:
	//
	//   ╔════════════════════╗         ╔════════════════════╗
	//   ║  ┌───────────┐     ║         ║    ┌───────────┐   ║
	//   ║  │  GroupA   │◀────╬─┐   ┌───╬───▶│  GroupB   │   ║
	//   ║  └───────────┘     ║ │   │   ║    └───────────┘   ║
	//   ║        ▲           ║ │   │   ║          ▲         ║
	//   ║        │           ║ │   │   ║          │         ║
	//   ║        │           ║ │   │   ║          │         ║
	//   ║  ┌───────────┐     ║ └───┼───╬────┌───────────┐   ║
	//   ║  │  EgressA  │─────╬─────┘   ║    │ IngressB  │   ║
	//   ║  └───────────┘     ║         ║    └───────────┘   ║
	//   ║                    ║         ║                    ║
	//   ╚════════════════════╝         ╚════════════════════╝
	//
	// By having the ability to switch the parent, we avoid the cyclic reference by
	// keeping all rules in a single stack.
	//
	// If this happens, we also have to change the construct ID, because
	// otherwise we might have two objects with the same ID if we have
	// multiple reversed security group relationships.
	//
	//   ╔═══════════════════════════════════╗
	//   ║┌───────────┐                      ║
	//   ║│  GroupB   │                      ║
	//   ║└───────────┘                      ║
	//   ║      ▲                            ║
	//   ║      │              ┌───────────┐ ║
	//   ║      ├────"from A"──│ IngressB  │ ║
	//   ║      │              └───────────┘ ║
	//   ║      │              ┌───────────┐ ║
	//   ║      ├─────"to B"───│  EgressA  │ ║
	//   ║      │              └───────────┘ ║
	//   ║      │              ┌───────────┐ ║
	//   ║      └─────"to B"───│  EgressC  │ ║  <-- oops
	//   ║                     └───────────┘ ║
	// ╚═══════════════════════════════════╝.
	DetermineRuleScope(peer IPeer, connection Port, fromTo *string, remoteRule *bool) *RuleScope
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Produce the egress rule JSON for the given connection.
	ToEgressRuleConfig() interface{}
	// Produce the ingress rule JSON for the given connection.
	ToIngressRuleConfig() interface{}
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for SecurityGroup
type jsiiProxy_SecurityGroup struct {
	internal.Type__awscdkResource
	jsiiProxy_ISecurityGroup
}

func (j *jsiiProxy_SecurityGroup) AllowAllIpv6Outbound() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowAllIpv6Outbound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) AllowAllOutbound() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"allowAllOutbound",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) CanInlineRule() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canInlineRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) Connections() Connections {
	var returns Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) DefaultPort() Port {
	var returns Port
	_jsii_.Get(
		j,
		"defaultPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) SecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) SecurityGroupVpcId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityGroupVpcId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityGroup) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}


func NewSecurityGroup(scope constructs.Construct, id *string, props *SecurityGroupProps) SecurityGroup {
	_init_.Initialize()

	if err := validateNewSecurityGroupParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityGroup{}

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewSecurityGroup_Override(s SecurityGroup, scope constructs.Construct, id *string, props *SecurityGroupProps) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		[]interface{}{scope, id, props},
		s,
	)
}

// Look up a security group by id.
func SecurityGroup_FromLookupById(scope constructs.Construct, id *string, securityGroupId *string) ISecurityGroup {
	_init_.Initialize()

	if err := validateSecurityGroup_FromLookupByIdParameters(scope, id, securityGroupId); err != nil {
		panic(err)
	}
	var returns ISecurityGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		"fromLookupById",
		[]interface{}{scope, id, securityGroupId},
		&returns,
	)

	return returns
}

// Look up a security group by name.
func SecurityGroup_FromLookupByName(scope constructs.Construct, id *string, securityGroupName *string, vpc IVpc) ISecurityGroup {
	_init_.Initialize()

	if err := validateSecurityGroup_FromLookupByNameParameters(scope, id, securityGroupName, vpc); err != nil {
		panic(err)
	}
	var returns ISecurityGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		"fromLookupByName",
		[]interface{}{scope, id, securityGroupName, vpc},
		&returns,
	)

	return returns
}

// Import an existing security group into this app.
//
// This method will assume that the Security Group has a rule in it which allows
// all outbound traffic, and so will not add egress rules to the imported Security
// Group (only ingress rules).
//
// If your existing Security Group needs to have egress rules added, pass the
// `allowAllOutbound: false` option on import.
func SecurityGroup_FromSecurityGroupId(scope constructs.Construct, id *string, securityGroupId *string, options *SecurityGroupImportOptions) ISecurityGroup {
	_init_.Initialize()

	if err := validateSecurityGroup_FromSecurityGroupIdParameters(scope, id, securityGroupId, options); err != nil {
		panic(err)
	}
	var returns ISecurityGroup

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		"fromSecurityGroupId",
		[]interface{}{scope, id, securityGroupId, options},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SecurityGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func SecurityGroup_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSecurityGroup_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func SecurityGroup_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateSecurityGroup_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return whether the indicated object is a security group.
func SecurityGroup_IsSecurityGroup(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityGroup_IsSecurityGroupParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_ec2.SecurityGroup",
		"isSecurityGroup",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) AddEgressRule(peer IPeer, connection Port, description *string, remoteRule *bool) {
	if err := s.validateAddEgressRuleParameters(peer, connection); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addEgressRule",
		[]interface{}{peer, connection, description, remoteRule},
	)
}

func (s *jsiiProxy_SecurityGroup) AddIngressRule(peer IPeer, connection Port, description *string, remoteRule *bool) {
	if err := s.validateAddIngressRuleParameters(peer, connection); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addIngressRule",
		[]interface{}{peer, connection, description, remoteRule},
	)
}

func (s *jsiiProxy_SecurityGroup) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := s.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (s *jsiiProxy_SecurityGroup) DetermineRuleScope(peer IPeer, connection Port, fromTo *string, remoteRule *bool) *RuleScope {
	if err := s.validateDetermineRuleScopeParameters(peer, connection, fromTo); err != nil {
		panic(err)
	}
	var returns *RuleScope

	_jsii_.Invoke(
		s,
		"determineRuleScope",
		[]interface{}{peer, connection, fromTo, remoteRule},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := s.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) GetResourceNameAttribute(nameAttr *string) *string {
	if err := s.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) ToEgressRuleConfig() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toEgressRuleConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) ToIngressRuleConfig() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toIngressRuleConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

