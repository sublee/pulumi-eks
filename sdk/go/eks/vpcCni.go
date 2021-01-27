// *** WARNING: this file was generated by pulumi-gen-eks. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VpcCni manages the configuration of the Amazon VPC CNI plugin for Kubernetes by applying its YAML chart.
type VpcCni struct {
	pulumi.CustomResourceState
}

// NewVpcCni registers a new resource with the given unique name, arguments, and options.
func NewVpcCni(ctx *pulumi.Context,
	name string, args *VpcCniArgs, opts ...pulumi.ResourceOption) (*VpcCni, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kubeconfig == nil {
		return nil, errors.New("invalid value for required argument 'Kubeconfig'")
	}
	var resource VpcCni
	err := ctx.RegisterResource("eks:index:VpcCni", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcCni gets an existing VpcCni resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcCni(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcCniState, opts ...pulumi.ResourceOption) (*VpcCni, error) {
	var resource VpcCni
	err := ctx.ReadResource("eks:index:VpcCni", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcCni resources.
type vpcCniState struct {
}

type VpcCniState struct {
}

func (VpcCniState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcCniState)(nil)).Elem()
}

type vpcCniArgs struct {
	// Specifies that your pods may use subnets and security groups (within the same VPC as your control plane resources) that are independent of your cluster's `resourcesVpcConfig`.
	//
	// Defaults to false.
	CustomNetworkConfig *bool `pulumi:"customNetworkConfig"`
	// Specifies the ENI_CONFIG_LABEL_DEF environment variable value for worker nodes. This is used to tell Kubernetes to automatically apply the ENIConfig for each Availability Zone
	// Ref: https://docs.aws.amazon.com/eks/latest/userguide/cni-custom-network.html (step 5(c))
	//
	// Defaults to the official AWS CNI image in ECR.
	EniConfigLabelDef *string `pulumi:"eniConfigLabelDef"`
	// Used to configure the MTU size for attached ENIs. The valid range is from 576 to 9001.
	//
	// Defaults to 9001.
	EniMtu *int `pulumi:"eniMtu"`
	// Specifies whether an external NAT gateway should be used to provide SNAT of secondary ENI IP addresses. If set to true, the SNAT iptables rule and off-VPC IP rule are not applied, and these rules are removed if they have already been applied.
	//
	// Defaults to false.
	ExternalSnat *bool `pulumi:"externalSnat"`
	// Specifies the container image to use in the AWS CNI cluster DaemonSet.
	//
	// Defaults to the official AWS CNI image in ECR.
	Image *string `pulumi:"image"`
	// The kubeconfig to use when setting the VPC CNI options.
	Kubeconfig interface{} `pulumi:"kubeconfig"`
	// Specifies the file path used for logs.
	//
	// Defaults to "stdout" to emit Pod logs for `kubectl logs`.
	LogFile *string `pulumi:"logFile"`
	// Specifies the log level used for logs.
	//
	// Defaults to "DEBUG"
	// Valid values: "DEBUG", "INFO", "WARN", "ERROR", or "FATAL".
	LogLevel *string `pulumi:"logLevel"`
	// Specifies whether NodePort services are enabled on a worker node's primary network interface. This requires additional iptables rules and that the kernel's reverse path filter on the primary interface is set to loose.
	//
	// Defaults to true.
	NodePortSupport *bool `pulumi:"nodePortSupport"`
	// Specifies the veth prefix used to generate the host-side veth device name for the CNI.
	//
	// The prefix can be at most 4 characters long.
	//
	// Defaults to "eni".
	VethPrefix *string `pulumi:"vethPrefix"`
	// Specifies the number of free elastic network interfaces (and all of their available IP addresses) that the ipamD daemon should attempt to keep available for pod assignment on the node.
	//
	// Defaults to 1.
	WarmEniTarget *int `pulumi:"warmEniTarget"`
	// Specifies the number of free IP addresses that the ipamD daemon should attempt to keep available for pod assignment on the node.
	WarmIpTarget *int `pulumi:"warmIpTarget"`
}

// The set of arguments for constructing a VpcCni resource.
type VpcCniArgs struct {
	// Specifies that your pods may use subnets and security groups (within the same VPC as your control plane resources) that are independent of your cluster's `resourcesVpcConfig`.
	//
	// Defaults to false.
	CustomNetworkConfig pulumi.BoolPtrInput
	// Specifies the ENI_CONFIG_LABEL_DEF environment variable value for worker nodes. This is used to tell Kubernetes to automatically apply the ENIConfig for each Availability Zone
	// Ref: https://docs.aws.amazon.com/eks/latest/userguide/cni-custom-network.html (step 5(c))
	//
	// Defaults to the official AWS CNI image in ECR.
	EniConfigLabelDef pulumi.StringPtrInput
	// Used to configure the MTU size for attached ENIs. The valid range is from 576 to 9001.
	//
	// Defaults to 9001.
	EniMtu pulumi.IntPtrInput
	// Specifies whether an external NAT gateway should be used to provide SNAT of secondary ENI IP addresses. If set to true, the SNAT iptables rule and off-VPC IP rule are not applied, and these rules are removed if they have already been applied.
	//
	// Defaults to false.
	ExternalSnat pulumi.BoolPtrInput
	// Specifies the container image to use in the AWS CNI cluster DaemonSet.
	//
	// Defaults to the official AWS CNI image in ECR.
	Image pulumi.StringPtrInput
	// The kubeconfig to use when setting the VPC CNI options.
	Kubeconfig pulumi.Input
	// Specifies the file path used for logs.
	//
	// Defaults to "stdout" to emit Pod logs for `kubectl logs`.
	LogFile pulumi.StringPtrInput
	// Specifies the log level used for logs.
	//
	// Defaults to "DEBUG"
	// Valid values: "DEBUG", "INFO", "WARN", "ERROR", or "FATAL".
	LogLevel pulumi.StringPtrInput
	// Specifies whether NodePort services are enabled on a worker node's primary network interface. This requires additional iptables rules and that the kernel's reverse path filter on the primary interface is set to loose.
	//
	// Defaults to true.
	NodePortSupport pulumi.BoolPtrInput
	// Specifies the veth prefix used to generate the host-side veth device name for the CNI.
	//
	// The prefix can be at most 4 characters long.
	//
	// Defaults to "eni".
	VethPrefix pulumi.StringPtrInput
	// Specifies the number of free elastic network interfaces (and all of their available IP addresses) that the ipamD daemon should attempt to keep available for pod assignment on the node.
	//
	// Defaults to 1.
	WarmEniTarget pulumi.IntPtrInput
	// Specifies the number of free IP addresses that the ipamD daemon should attempt to keep available for pod assignment on the node.
	WarmIpTarget pulumi.IntPtrInput
}

func (VpcCniArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcCniArgs)(nil)).Elem()
}

type VpcCniInput interface {
	pulumi.Input

	ToVpcCniOutput() VpcCniOutput
	ToVpcCniOutputWithContext(ctx context.Context) VpcCniOutput
}

func (*VpcCni) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcCni)(nil))
}

func (i *VpcCni) ToVpcCniOutput() VpcCniOutput {
	return i.ToVpcCniOutputWithContext(context.Background())
}

func (i *VpcCni) ToVpcCniOutputWithContext(ctx context.Context) VpcCniOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcCniOutput)
}

func (i *VpcCni) ToVpcCniPtrOutput() VpcCniPtrOutput {
	return i.ToVpcCniPtrOutputWithContext(context.Background())
}

func (i *VpcCni) ToVpcCniPtrOutputWithContext(ctx context.Context) VpcCniPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcCniPtrOutput)
}

type VpcCniPtrInput interface {
	pulumi.Input

	ToVpcCniPtrOutput() VpcCniPtrOutput
	ToVpcCniPtrOutputWithContext(ctx context.Context) VpcCniPtrOutput
}

type vpcCniPtrType VpcCniArgs

func (*vpcCniPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcCni)(nil))
}

func (i *vpcCniPtrType) ToVpcCniPtrOutput() VpcCniPtrOutput {
	return i.ToVpcCniPtrOutputWithContext(context.Background())
}

func (i *vpcCniPtrType) ToVpcCniPtrOutputWithContext(ctx context.Context) VpcCniPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcCniPtrOutput)
}

// VpcCniArrayInput is an input type that accepts VpcCniArray and VpcCniArrayOutput values.
// You can construct a concrete instance of `VpcCniArrayInput` via:
//
//          VpcCniArray{ VpcCniArgs{...} }
type VpcCniArrayInput interface {
	pulumi.Input

	ToVpcCniArrayOutput() VpcCniArrayOutput
	ToVpcCniArrayOutputWithContext(context.Context) VpcCniArrayOutput
}

type VpcCniArray []VpcCniInput

func (VpcCniArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*VpcCni)(nil))
}

func (i VpcCniArray) ToVpcCniArrayOutput() VpcCniArrayOutput {
	return i.ToVpcCniArrayOutputWithContext(context.Background())
}

func (i VpcCniArray) ToVpcCniArrayOutputWithContext(ctx context.Context) VpcCniArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcCniArrayOutput)
}

// VpcCniMapInput is an input type that accepts VpcCniMap and VpcCniMapOutput values.
// You can construct a concrete instance of `VpcCniMapInput` via:
//
//          VpcCniMap{ "key": VpcCniArgs{...} }
type VpcCniMapInput interface {
	pulumi.Input

	ToVpcCniMapOutput() VpcCniMapOutput
	ToVpcCniMapOutputWithContext(context.Context) VpcCniMapOutput
}

type VpcCniMap map[string]VpcCniInput

func (VpcCniMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*VpcCni)(nil))
}

func (i VpcCniMap) ToVpcCniMapOutput() VpcCniMapOutput {
	return i.ToVpcCniMapOutputWithContext(context.Background())
}

func (i VpcCniMap) ToVpcCniMapOutputWithContext(ctx context.Context) VpcCniMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcCniMapOutput)
}

type VpcCniOutput struct {
	*pulumi.OutputState
}

func (VpcCniOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcCni)(nil))
}

func (o VpcCniOutput) ToVpcCniOutput() VpcCniOutput {
	return o
}

func (o VpcCniOutput) ToVpcCniOutputWithContext(ctx context.Context) VpcCniOutput {
	return o
}

func (o VpcCniOutput) ToVpcCniPtrOutput() VpcCniPtrOutput {
	return o.ToVpcCniPtrOutputWithContext(context.Background())
}

func (o VpcCniOutput) ToVpcCniPtrOutputWithContext(ctx context.Context) VpcCniPtrOutput {
	return o.ApplyT(func(v VpcCni) *VpcCni {
		return &v
	}).(VpcCniPtrOutput)
}

type VpcCniPtrOutput struct {
	*pulumi.OutputState
}

func (VpcCniPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcCni)(nil))
}

func (o VpcCniPtrOutput) ToVpcCniPtrOutput() VpcCniPtrOutput {
	return o
}

func (o VpcCniPtrOutput) ToVpcCniPtrOutputWithContext(ctx context.Context) VpcCniPtrOutput {
	return o
}

type VpcCniArrayOutput struct{ *pulumi.OutputState }

func (VpcCniArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VpcCni)(nil))
}

func (o VpcCniArrayOutput) ToVpcCniArrayOutput() VpcCniArrayOutput {
	return o
}

func (o VpcCniArrayOutput) ToVpcCniArrayOutputWithContext(ctx context.Context) VpcCniArrayOutput {
	return o
}

func (o VpcCniArrayOutput) Index(i pulumi.IntInput) VpcCniOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VpcCni {
		return vs[0].([]VpcCni)[vs[1].(int)]
	}).(VpcCniOutput)
}

type VpcCniMapOutput struct{ *pulumi.OutputState }

func (VpcCniMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]VpcCni)(nil))
}

func (o VpcCniMapOutput) ToVpcCniMapOutput() VpcCniMapOutput {
	return o
}

func (o VpcCniMapOutput) ToVpcCniMapOutputWithContext(ctx context.Context) VpcCniMapOutput {
	return o
}

func (o VpcCniMapOutput) MapIndex(k pulumi.StringInput) VpcCniOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) VpcCni {
		return vs[0].(map[string]VpcCni)[vs[1].(string)]
	}).(VpcCniOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcCniOutput{})
	pulumi.RegisterOutputType(VpcCniPtrOutput{})
	pulumi.RegisterOutputType(VpcCniArrayOutput{})
	pulumi.RegisterOutputType(VpcCniMapOutput{})
}
