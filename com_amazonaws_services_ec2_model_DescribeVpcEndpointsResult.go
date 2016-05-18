package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpcEndpointsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.VpcEndpoint> getVpcEndpoints()
	GetVpcEndpoints() []*ServicesEc2ModelVpcEndpoint

	// public void setVpcEndpoints(java.util.Collection<com.amazonaws.services.ec2.model.VpcEndpoint>)
	SetVpcEndpoints(a []*ServicesEc2ModelVpcEndpoint) 

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult withVpcEndpoints(com.amazonaws.services.ec2.model.VpcEndpoint...)
	WithVpcEndpoints(a ...*ServicesEc2ModelVpcEndpoint) *ServicesEc2ModelDescribeVpcEndpointsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult withVpcEndpoints(java.util.Collection<com.amazonaws.services.ec2.model.VpcEndpoint>)
	WithVpcEndpoints2(a []*ServicesEc2ModelVpcEndpoint) *ServicesEc2ModelDescribeVpcEndpointsResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeVpcEndpointsResult

	// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult clone()
	Clone() *ServicesEc2ModelDescribeVpcEndpointsResult
}

type ServicesEc2ModelDescribeVpcEndpointsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult()
func NewServicesEc2ModelDescribeVpcEndpointsResult() (*ServicesEc2ModelDescribeVpcEndpointsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpcEndpointsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpcEndpointsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.VpcEndpoint> getVpcEndpoints()
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) GetVpcEndpoints() []*ServicesEc2ModelVpcEndpoint {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcEndpoints", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVpcEndpoint)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVpcEndpoints(java.util.Collection<com.amazonaws.services.ec2.model.VpcEndpoint>)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) SetVpcEndpoints(a []*ServicesEc2ModelVpcEndpoint)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcEndpoints", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult withVpcEndpoints(com.amazonaws.services.ec2.model.VpcEndpoint...)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) WithVpcEndpoints(a ...*ServicesEc2ModelVpcEndpoint) *ServicesEc2ModelDescribeVpcEndpointsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VpcEndpoint")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcEndpoints", "com/amazonaws/services/ec2/model/DescribeVpcEndpointsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcEndpoint"))
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
	unique_x := &ServicesEc2ModelDescribeVpcEndpointsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult withVpcEndpoints(java.util.Collection<com.amazonaws.services.ec2.model.VpcEndpoint>)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) WithVpcEndpoints2(a []*ServicesEc2ModelVpcEndpoint) *ServicesEc2ModelDescribeVpcEndpointsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcEndpoints", "com/amazonaws/services/ec2/model/DescribeVpcEndpointsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVpcEndpointsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) SetNextToken(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) GetNextToken() string {
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

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) WithNextToken(a string) *ServicesEc2ModelDescribeVpcEndpointsResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeVpcEndpointsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVpcEndpointsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpcEndpointsResult clone()
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) Clone() *ServicesEc2ModelDescribeVpcEndpointsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpcEndpointsResult")
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
	unique_x := &ServicesEc2ModelDescribeVpcEndpointsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVpcEndpointsResult) Clone2() (*JavaLangObject, error) {
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


