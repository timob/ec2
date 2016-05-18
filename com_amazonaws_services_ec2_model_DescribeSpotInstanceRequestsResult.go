package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSpotInstanceRequestsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.SpotInstanceRequest> getSpotInstanceRequests()
	GetSpotInstanceRequests() []*ServicesEc2ModelSpotInstanceRequest

	// public void setSpotInstanceRequests(java.util.Collection<com.amazonaws.services.ec2.model.SpotInstanceRequest>)
	SetSpotInstanceRequests(a []*ServicesEc2ModelSpotInstanceRequest) 

	// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult withSpotInstanceRequests(com.amazonaws.services.ec2.model.SpotInstanceRequest...)
	WithSpotInstanceRequests(a ...*ServicesEc2ModelSpotInstanceRequest) *ServicesEc2ModelDescribeSpotInstanceRequestsResult

	// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult withSpotInstanceRequests(java.util.Collection<com.amazonaws.services.ec2.model.SpotInstanceRequest>)
	WithSpotInstanceRequests2(a []*ServicesEc2ModelSpotInstanceRequest) *ServicesEc2ModelDescribeSpotInstanceRequestsResult

	// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult clone()
	Clone() *ServicesEc2ModelDescribeSpotInstanceRequestsResult
}

type ServicesEc2ModelDescribeSpotInstanceRequestsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult()
func NewServicesEc2ModelDescribeSpotInstanceRequestsResult() (*ServicesEc2ModelDescribeSpotInstanceRequestsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSpotInstanceRequestsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSpotInstanceRequestsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.SpotInstanceRequest> getSpotInstanceRequests()
func (jbobject *ServicesEc2ModelDescribeSpotInstanceRequestsResult) GetSpotInstanceRequests() []*ServicesEc2ModelSpotInstanceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotInstanceRequests", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelSpotInstanceRequest)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSpotInstanceRequests(java.util.Collection<com.amazonaws.services.ec2.model.SpotInstanceRequest>)
func (jbobject *ServicesEc2ModelDescribeSpotInstanceRequestsResult) SetSpotInstanceRequests(a []*ServicesEc2ModelSpotInstanceRequest)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotInstanceRequests", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult withSpotInstanceRequests(com.amazonaws.services.ec2.model.SpotInstanceRequest...)
func (jbobject *ServicesEc2ModelDescribeSpotInstanceRequestsResult) WithSpotInstanceRequests(a ...*ServicesEc2ModelSpotInstanceRequest) *ServicesEc2ModelDescribeSpotInstanceRequestsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/SpotInstanceRequest")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotInstanceRequests", "com/amazonaws/services/ec2/model/DescribeSpotInstanceRequestsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SpotInstanceRequest"))
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
	unique_x := &ServicesEc2ModelDescribeSpotInstanceRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult withSpotInstanceRequests(java.util.Collection<com.amazonaws.services.ec2.model.SpotInstanceRequest>)
func (jbobject *ServicesEc2ModelDescribeSpotInstanceRequestsResult) WithSpotInstanceRequests2(a []*ServicesEc2ModelSpotInstanceRequest) *ServicesEc2ModelDescribeSpotInstanceRequestsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotInstanceRequests", "com/amazonaws/services/ec2/model/DescribeSpotInstanceRequestsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSpotInstanceRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSpotInstanceRequestsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSpotInstanceRequestsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSpotInstanceRequestsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSpotInstanceRequestsResult clone()
func (jbobject *ServicesEc2ModelDescribeSpotInstanceRequestsResult) Clone() *ServicesEc2ModelDescribeSpotInstanceRequestsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSpotInstanceRequestsResult")
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
	unique_x := &ServicesEc2ModelDescribeSpotInstanceRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSpotInstanceRequestsResult) Clone2() (*JavaLangObject, error) {
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


