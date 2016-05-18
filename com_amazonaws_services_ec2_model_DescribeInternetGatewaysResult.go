package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeInternetGatewaysResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.InternetGateway> getInternetGateways()
	GetInternetGateways() []*ServicesEc2ModelInternetGateway

	// public void setInternetGateways(java.util.Collection<com.amazonaws.services.ec2.model.InternetGateway>)
	SetInternetGateways(a []*ServicesEc2ModelInternetGateway) 

	// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult withInternetGateways(com.amazonaws.services.ec2.model.InternetGateway...)
	WithInternetGateways(a ...*ServicesEc2ModelInternetGateway) *ServicesEc2ModelDescribeInternetGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult withInternetGateways(java.util.Collection<com.amazonaws.services.ec2.model.InternetGateway>)
	WithInternetGateways2(a []*ServicesEc2ModelInternetGateway) *ServicesEc2ModelDescribeInternetGatewaysResult

	// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult clone()
	Clone() *ServicesEc2ModelDescribeInternetGatewaysResult
}

type ServicesEc2ModelDescribeInternetGatewaysResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult()
func NewServicesEc2ModelDescribeInternetGatewaysResult() (*ServicesEc2ModelDescribeInternetGatewaysResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeInternetGatewaysResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeInternetGatewaysResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.InternetGateway> getInternetGateways()
func (jbobject *ServicesEc2ModelDescribeInternetGatewaysResult) GetInternetGateways() []*ServicesEc2ModelInternetGateway {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInternetGateways", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInternetGateway)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setInternetGateways(java.util.Collection<com.amazonaws.services.ec2.model.InternetGateway>)
func (jbobject *ServicesEc2ModelDescribeInternetGatewaysResult) SetInternetGateways(a []*ServicesEc2ModelInternetGateway)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInternetGateways", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult withInternetGateways(com.amazonaws.services.ec2.model.InternetGateway...)
func (jbobject *ServicesEc2ModelDescribeInternetGatewaysResult) WithInternetGateways(a ...*ServicesEc2ModelInternetGateway) *ServicesEc2ModelDescribeInternetGatewaysResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InternetGateway")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInternetGateways", "com/amazonaws/services/ec2/model/DescribeInternetGatewaysResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InternetGateway"))
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
	unique_x := &ServicesEc2ModelDescribeInternetGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult withInternetGateways(java.util.Collection<com.amazonaws.services.ec2.model.InternetGateway>)
func (jbobject *ServicesEc2ModelDescribeInternetGatewaysResult) WithInternetGateways2(a []*ServicesEc2ModelInternetGateway) *ServicesEc2ModelDescribeInternetGatewaysResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInternetGateways", "com/amazonaws/services/ec2/model/DescribeInternetGatewaysResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeInternetGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeInternetGatewaysResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeInternetGatewaysResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeInternetGatewaysResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeInternetGatewaysResult clone()
func (jbobject *ServicesEc2ModelDescribeInternetGatewaysResult) Clone() *ServicesEc2ModelDescribeInternetGatewaysResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeInternetGatewaysResult")
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
	unique_x := &ServicesEc2ModelDescribeInternetGatewaysResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeInternetGatewaysResult) Clone2() (*JavaLangObject, error) {
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


