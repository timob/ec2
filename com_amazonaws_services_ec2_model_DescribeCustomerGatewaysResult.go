package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeCustomerGatewaysResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.CustomerGateway> getCustomerGateways()
	GetCustomerGateways() []*ServicesEc2ModelCustomerGateway

	// public void setCustomerGateways(java.util.Collection<com.amazonaws.services.ec2.model.CustomerGateway>)
	SetCustomerGateways(a []*ServicesEc2ModelCustomerGateway) 

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult withCustomerGateways(com.amazonaws.services.ec2.model.CustomerGateway...)
	WithCustomerGateways(a ...*ServicesEc2ModelCustomerGateway) *ServicesEc2ModelDescribeCustomerGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult withCustomerGateways(java.util.Collection<com.amazonaws.services.ec2.model.CustomerGateway>)
	WithCustomerGateways2(a []*ServicesEc2ModelCustomerGateway) *ServicesEc2ModelDescribeCustomerGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult clone()
	Clone() *ServicesEc2ModelDescribeCustomerGatewaysResult
}

type ServicesEc2ModelDescribeCustomerGatewaysResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult()
func NewServicesEc2ModelDescribeCustomerGatewaysResult() (*ServicesEc2ModelDescribeCustomerGatewaysResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeCustomerGatewaysResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeCustomerGatewaysResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.CustomerGateway> getCustomerGateways()
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysResult) GetCustomerGateways() []*ServicesEc2ModelCustomerGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCustomerGateways", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelCustomerGateway)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setCustomerGateways(java.util.Collection<com.amazonaws.services.ec2.model.CustomerGateway>)
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysResult) SetCustomerGateways(a []*ServicesEc2ModelCustomerGateway)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCustomerGateways", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult withCustomerGateways(com.amazonaws.services.ec2.model.CustomerGateway...)
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysResult) WithCustomerGateways(a ...*ServicesEc2ModelCustomerGateway) *ServicesEc2ModelDescribeCustomerGatewaysResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/CustomerGateway")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCustomerGateways", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CustomerGateway"))
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
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult withCustomerGateways(java.util.Collection<com.amazonaws.services.ec2.model.CustomerGateway>)
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysResult) WithCustomerGateways2(a []*ServicesEc2ModelCustomerGateway) *ServicesEc2ModelDescribeCustomerGatewaysResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCustomerGateways", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeCustomerGatewaysResult clone()
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysResult) Clone() *ServicesEc2ModelDescribeCustomerGatewaysResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeCustomerGatewaysResult")
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
	unique_x := &ServicesEc2ModelDescribeCustomerGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeCustomerGatewaysResult) Clone2() (*JavaLangObject, error) {
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


