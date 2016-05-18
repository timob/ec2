package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpnGatewaysResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.VpnGateway> getVpnGateways()
	GetVpnGateways() []*ServicesEc2ModelVpnGateway

	// public void setVpnGateways(java.util.Collection<com.amazonaws.services.ec2.model.VpnGateway>)
	SetVpnGateways(a []*ServicesEc2ModelVpnGateway) 

	// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult withVpnGateways(com.amazonaws.services.ec2.model.VpnGateway...)
	WithVpnGateways(a ...*ServicesEc2ModelVpnGateway) *ServicesEc2ModelDescribeVpnGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult withVpnGateways(java.util.Collection<com.amazonaws.services.ec2.model.VpnGateway>)
	WithVpnGateways2(a []*ServicesEc2ModelVpnGateway) *ServicesEc2ModelDescribeVpnGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult clone()
	Clone() *ServicesEc2ModelDescribeVpnGatewaysResult
}

type ServicesEc2ModelDescribeVpnGatewaysResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult()
func NewServicesEc2ModelDescribeVpnGatewaysResult() (*ServicesEc2ModelDescribeVpnGatewaysResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpnGatewaysResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpnGatewaysResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.VpnGateway> getVpnGateways()
func (jbobject *ServicesEc2ModelDescribeVpnGatewaysResult) GetVpnGateways() []*ServicesEc2ModelVpnGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpnGateways", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVpnGateway)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVpnGateways(java.util.Collection<com.amazonaws.services.ec2.model.VpnGateway>)
func (jbobject *ServicesEc2ModelDescribeVpnGatewaysResult) SetVpnGateways(a []*ServicesEc2ModelVpnGateway)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpnGateways", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult withVpnGateways(com.amazonaws.services.ec2.model.VpnGateway...)
func (jbobject *ServicesEc2ModelDescribeVpnGatewaysResult) WithVpnGateways(a ...*ServicesEc2ModelVpnGateway) *ServicesEc2ModelDescribeVpnGatewaysResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VpnGateway")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnGateways", "com/amazonaws/services/ec2/model/DescribeVpnGatewaysResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnGateway"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpnGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult withVpnGateways(java.util.Collection<com.amazonaws.services.ec2.model.VpnGateway>)
func (jbobject *ServicesEc2ModelDescribeVpnGatewaysResult) WithVpnGateways2(a []*ServicesEc2ModelVpnGateway) *ServicesEc2ModelDescribeVpnGatewaysResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnGateways", "com/amazonaws/services/ec2/model/DescribeVpnGatewaysResult", conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpnGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVpnGatewaysResult) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean equals(java.lang.Object)
func (jbobject *ServicesEc2ModelDescribeVpnGatewaysResult) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *ServicesEc2ModelDescribeVpnGatewaysResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpnGatewaysResult clone()
func (jbobject *ServicesEc2ModelDescribeVpnGatewaysResult) Clone() *ServicesEc2ModelDescribeVpnGatewaysResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpnGatewaysResult")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &ServicesEc2ModelDescribeVpnGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVpnGatewaysResult) Clone2() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}


