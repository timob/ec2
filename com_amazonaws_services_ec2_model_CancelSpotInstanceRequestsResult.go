package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelSpotInstanceRequestsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest> getCancelledSpotInstanceRequests()
	GetCancelledSpotInstanceRequests() []*ServicesEc2ModelCancelledSpotInstanceRequest

	// public void setCancelledSpotInstanceRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest>)
	SetCancelledSpotInstanceRequests(a []*ServicesEc2ModelCancelledSpotInstanceRequest) 

	// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsResult withCancelledSpotInstanceRequests(com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest...)
	WithCancelledSpotInstanceRequests(a ...*ServicesEc2ModelCancelledSpotInstanceRequest) *ServicesEc2ModelCancelSpotInstanceRequestsResult

	// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsResult withCancelledSpotInstanceRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest>)
	WithCancelledSpotInstanceRequests2(a []*ServicesEc2ModelCancelledSpotInstanceRequest) *ServicesEc2ModelCancelSpotInstanceRequestsResult

	// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsResult clone()
	Clone() *ServicesEc2ModelCancelSpotInstanceRequestsResult
}

type ServicesEc2ModelCancelSpotInstanceRequestsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsResult()
func NewServicesEc2ModelCancelSpotInstanceRequestsResult() (*ServicesEc2ModelCancelSpotInstanceRequestsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelSpotInstanceRequestsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest> getCancelledSpotInstanceRequests()
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsResult) GetCancelledSpotInstanceRequests() []*ServicesEc2ModelCancelledSpotInstanceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCancelledSpotInstanceRequests", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelCancelledSpotInstanceRequest)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setCancelledSpotInstanceRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest>)
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsResult) SetCancelledSpotInstanceRequests(a []*ServicesEc2ModelCancelledSpotInstanceRequest)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCancelledSpotInstanceRequests", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsResult withCancelledSpotInstanceRequests(com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest...)
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsResult) WithCancelledSpotInstanceRequests(a ...*ServicesEc2ModelCancelledSpotInstanceRequest) *ServicesEc2ModelCancelSpotInstanceRequestsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/CancelledSpotInstanceRequest")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCancelledSpotInstanceRequests", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CancelledSpotInstanceRequest"))
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsResult withCancelledSpotInstanceRequests(java.util.Collection<com.amazonaws.services.ec2.model.CancelledSpotInstanceRequest>)
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsResult) WithCancelledSpotInstanceRequests2(a []*ServicesEc2ModelCancelledSpotInstanceRequest) *ServicesEc2ModelCancelSpotInstanceRequestsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCancelledSpotInstanceRequests", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsResult clone()
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsResult) Clone() *ServicesEc2ModelCancelSpotInstanceRequestsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsResult")
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsResult) Clone2() (*JavaLangObject, error) {
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


