package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpcEndpointServicesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<java.lang.String> getServiceNames()
	GetServiceNames() []string

	// public void setServiceNames(java.util.Collection<java.lang.String>)
	SetServiceNames(a []string) 

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult withServiceNames(java.lang.String...)
	WithServiceNames(a ...string) *ServicesEc2ModelDescribeVpcEndpointServicesResult

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult withServiceNames(java.util.Collection<java.lang.String>)
	WithServiceNames2(a []string) *ServicesEc2ModelDescribeVpcEndpointServicesResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeVpcEndpointServicesResult

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult clone()
	Clone() *ServicesEc2ModelDescribeVpcEndpointServicesResult
}

type ServicesEc2ModelDescribeVpcEndpointServicesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult()
func NewServicesEc2ModelDescribeVpcEndpointServicesResult() (*ServicesEc2ModelDescribeVpcEndpointServicesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpcEndpointServicesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpcEndpointServicesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getServiceNames()
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) GetServiceNames() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getServiceNames", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setServiceNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) SetServiceNames(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setServiceNames", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult withServiceNames(java.lang.String...)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) WithServiceNames(a ...string) *ServicesEc2ModelDescribeVpcEndpointServicesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withServiceNames", "com/amazonaws/services/ec2/model/DescribeVpcEndpointServicesResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVpcEndpointServicesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult withServiceNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) WithServiceNames2(a []string) *ServicesEc2ModelDescribeVpcEndpointServicesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withServiceNames", "com/amazonaws/services/ec2/model/DescribeVpcEndpointServicesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVpcEndpointServicesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) WithNextToken(a string) *ServicesEc2ModelDescribeVpcEndpointServicesResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeVpcEndpointServicesResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVpcEndpointServicesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointServicesResult clone()
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) Clone() *ServicesEc2ModelDescribeVpcEndpointServicesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpcEndpointServicesResult")
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
	unique_x := &ServicesEc2ModelDescribeVpcEndpointServicesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVpcEndpointServicesResult) Clone2() (*JavaLangObject, error) {
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


