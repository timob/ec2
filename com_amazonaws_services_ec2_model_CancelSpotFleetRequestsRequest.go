package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelSpotFleetRequestsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getSpotFleetRequestIds()
	GetSpotFleetRequestIds() []string

	// public void setSpotFleetRequestIds(java.util.Collection<java.lang.String>)
	SetSpotFleetRequestIds(a []string) 

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest withSpotFleetRequestIds(java.lang.String...)
	WithSpotFleetRequestIds(a ...string) *ServicesEc2ModelCancelSpotFleetRequestsRequest

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest withSpotFleetRequestIds(java.util.Collection<java.lang.String>)
	WithSpotFleetRequestIds2(a []string) *ServicesEc2ModelCancelSpotFleetRequestsRequest

	// public void setTerminateInstances(java.lang.Boolean)
	SetTerminateInstances(a bool) 

	// public java.lang.Boolean getTerminateInstances()
	GetTerminateInstances() bool

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest withTerminateInstances(java.lang.Boolean)
	WithTerminateInstances(a bool) *ServicesEc2ModelCancelSpotFleetRequestsRequest

	// public java.lang.Boolean isTerminateInstances()
	IsTerminateInstances() bool

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest clone()
	Clone3() *ServicesEc2ModelCancelSpotFleetRequestsRequest
}

type ServicesEc2ModelCancelSpotFleetRequestsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest()
func NewServicesEc2ModelCancelSpotFleetRequestsRequest() (*ServicesEc2ModelCancelSpotFleetRequestsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelSpotFleetRequestsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelSpotFleetRequestsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getSpotFleetRequestIds()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) GetSpotFleetRequestIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotFleetRequestIds", "java/util/List")
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

// public void setSpotFleetRequestIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) SetSpotFleetRequestIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotFleetRequestIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest withSpotFleetRequestIds(java.lang.String...)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) WithSpotFleetRequestIds(a ...string) *ServicesEc2ModelCancelSpotFleetRequestsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestIds", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest withSpotFleetRequestIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) WithSpotFleetRequestIds2(a []string) *ServicesEc2ModelCancelSpotFleetRequestsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotFleetRequestIds", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTerminateInstances(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) SetTerminateInstances(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTerminateInstances", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getTerminateInstances()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) GetTerminateInstances() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTerminateInstances", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest withTerminateInstances(java.lang.Boolean)
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) WithTerminateInstances(a bool) *ServicesEc2ModelCancelSpotFleetRequestsRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTerminateInstances", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isTerminateInstances()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) IsTerminateInstances() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isTerminateInstances", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) GetDryRunRequest() *Request {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDryRunRequest", "com/amazonaws/Request")
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
	unique_x := &Request{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelSpotFleetRequestsRequest clone()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) Clone3() *ServicesEc2ModelCancelSpotFleetRequestsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelSpotFleetRequestsRequest")
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
	unique_x := &ServicesEc2ModelCancelSpotFleetRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) Clone() *AmazonWebServiceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/AmazonWebServiceRequest")
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
	unique_x := &AmazonWebServiceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCancelSpotFleetRequestsRequest) Clone2() (*JavaLangObject, error) {
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


