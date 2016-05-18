package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelSpotInstanceRequestsRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public java.util.List<java.lang.String> getSpotInstanceRequestIds()
	GetSpotInstanceRequestIds() []string

	// public void setSpotInstanceRequestIds(java.util.Collection<java.lang.String>)
	SetSpotInstanceRequestIds(a []string) 

	// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest withSpotInstanceRequestIds(java.lang.String...)
	WithSpotInstanceRequestIds(a ...string) *ServicesEc2ModelCancelSpotInstanceRequestsRequest

	// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest withSpotInstanceRequestIds(java.util.Collection<java.lang.String>)
	WithSpotInstanceRequestIds2(a []string) *ServicesEc2ModelCancelSpotInstanceRequestsRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest clone()
	Clone3() *ServicesEc2ModelCancelSpotInstanceRequestsRequest
}

type ServicesEc2ModelCancelSpotInstanceRequestsRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest()
func NewServicesEc2ModelCancelSpotInstanceRequestsRequest() (*ServicesEc2ModelCancelSpotInstanceRequestsRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelSpotInstanceRequestsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest(java.util.List<java.lang.String>)
func NewServicesEc2ModelCancelSpotInstanceRequestsRequest2(a []string) (*ServicesEc2ModelCancelSpotInstanceRequestsRequest) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsRequest", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelCancelSpotInstanceRequestsRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<java.lang.String> getSpotInstanceRequestIds()
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) GetSpotInstanceRequestIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSpotInstanceRequestIds", "java/util/List")
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

// public void setSpotInstanceRequestIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) SetSpotInstanceRequestIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSpotInstanceRequestIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest withSpotInstanceRequestIds(java.lang.String...)
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) WithSpotInstanceRequestIds(a ...string) *ServicesEc2ModelCancelSpotInstanceRequestsRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotInstanceRequestIds", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest withSpotInstanceRequestIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) WithSpotInstanceRequestIds2(a []string) *ServicesEc2ModelCancelSpotInstanceRequestsRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSpotInstanceRequestIds", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelSpotInstanceRequestsRequest clone()
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) Clone3() *ServicesEc2ModelCancelSpotInstanceRequestsRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestsRequest")
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestsRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestsRequest) Clone2() (*JavaLangObject, error) {
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


