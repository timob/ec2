package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeNatGatewaysResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.NatGateway> getNatGateways()
	GetNatGateways() []*ServicesEc2ModelNatGateway

	// public void setNatGateways(java.util.Collection<com.amazonaws.services.ec2.model.NatGateway>)
	SetNatGateways(a []*ServicesEc2ModelNatGateway) 

	// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult withNatGateways(com.amazonaws.services.ec2.model.NatGateway...)
	WithNatGateways(a ...*ServicesEc2ModelNatGateway) *ServicesEc2ModelDescribeNatGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult withNatGateways(java.util.Collection<com.amazonaws.services.ec2.model.NatGateway>)
	WithNatGateways2(a []*ServicesEc2ModelNatGateway) *ServicesEc2ModelDescribeNatGatewaysResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeNatGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult clone()
	Clone() *ServicesEc2ModelDescribeNatGatewaysResult
}

type ServicesEc2ModelDescribeNatGatewaysResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult()
func NewServicesEc2ModelDescribeNatGatewaysResult() (*ServicesEc2ModelDescribeNatGatewaysResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeNatGatewaysResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeNatGatewaysResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.NatGateway> getNatGateways()
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) GetNatGateways() []*ServicesEc2ModelNatGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNatGateways", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelNatGateway)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setNatGateways(java.util.Collection<com.amazonaws.services.ec2.model.NatGateway>)
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) SetNatGateways(a []*ServicesEc2ModelNatGateway)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNatGateways", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult withNatGateways(com.amazonaws.services.ec2.model.NatGateway...)
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) WithNatGateways(a ...*ServicesEc2ModelNatGateway) *ServicesEc2ModelDescribeNatGatewaysResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/NatGateway")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNatGateways", "com/amazonaws/services/ec2/model/DescribeNatGatewaysResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NatGateway"))
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
	unique_x := &ServicesEc2ModelDescribeNatGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult withNatGateways(java.util.Collection<com.amazonaws.services.ec2.model.NatGateway>)
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) WithNatGateways2(a []*ServicesEc2ModelNatGateway) *ServicesEc2ModelDescribeNatGatewaysResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNatGateways", "com/amazonaws/services/ec2/model/DescribeNatGatewaysResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeNatGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) SetNextToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNextToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNextToken()
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) GetNextToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNextToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) WithNextToken(a string) *ServicesEc2ModelDescribeNatGatewaysResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeNatGatewaysResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeNatGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeNatGatewaysResult clone()
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) Clone() *ServicesEc2ModelDescribeNatGatewaysResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeNatGatewaysResult")
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
	unique_x := &ServicesEc2ModelDescribeNatGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeNatGatewaysResult) Clone2() (*JavaLangObject, error) {
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


