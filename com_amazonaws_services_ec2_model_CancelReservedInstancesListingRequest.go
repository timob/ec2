package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelReservedInstancesListingRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setReservedInstancesListingId(java.lang.String)
	SetReservedInstancesListingId(a string) 

	// public java.lang.String getReservedInstancesListingId()
	GetReservedInstancesListingId() string

	// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingRequest withReservedInstancesListingId(java.lang.String)
	WithReservedInstancesListingId(a string) *ServicesEc2ModelCancelReservedInstancesListingRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelReservedInstancesListingRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingRequest clone()
	Clone3() *ServicesEc2ModelCancelReservedInstancesListingRequest
}

type ServicesEc2ModelCancelReservedInstancesListingRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingRequest()
func NewServicesEc2ModelCancelReservedInstancesListingRequest() (*ServicesEc2ModelCancelReservedInstancesListingRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CancelReservedInstancesListingRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCancelReservedInstancesListingRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesListingId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) SetReservedInstancesListingId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesListingId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservedInstancesListingId()
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) GetReservedInstancesListingId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesListingId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingRequest withReservedInstancesListingId(java.lang.String)
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) WithReservedInstancesListingId(a string) *ServicesEc2ModelCancelReservedInstancesListingRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesListingId", "com/amazonaws/services/ec2/model/CancelReservedInstancesListingRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelReservedInstancesListingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CancelReservedInstancesListingRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CancelReservedInstancesListingRequest clone()
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) Clone3() *ServicesEc2ModelCancelReservedInstancesListingRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CancelReservedInstancesListingRequest")
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
	unique_x := &ServicesEc2ModelCancelReservedInstancesListingRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCancelReservedInstancesListingRequest) Clone2() (*JavaLangObject, error) {
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


