package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRejectVpcPeeringConnectionRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpcPeeringConnectionId(java.lang.String)
	SetVpcPeeringConnectionId(a string) 

	// public java.lang.String getVpcPeeringConnectionId()
	GetVpcPeeringConnectionId() string

	// public com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionRequest withVpcPeeringConnectionId(java.lang.String)
	WithVpcPeeringConnectionId(a string) *ServicesEc2ModelRejectVpcPeeringConnectionRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionRequest clone()
	Clone3() *ServicesEc2ModelRejectVpcPeeringConnectionRequest
}

type ServicesEc2ModelRejectVpcPeeringConnectionRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionRequest()
func NewServicesEc2ModelRejectVpcPeeringConnectionRequest() (*ServicesEc2ModelRejectVpcPeeringConnectionRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RejectVpcPeeringConnectionRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRejectVpcPeeringConnectionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) SetVpcPeeringConnectionId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcPeeringConnectionId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcPeeringConnectionId()
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) GetVpcPeeringConnectionId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcPeeringConnectionId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionRequest withVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) WithVpcPeeringConnectionId(a string) *ServicesEc2ModelRejectVpcPeeringConnectionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnectionId", "com/amazonaws/services/ec2/model/RejectVpcPeeringConnectionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRejectVpcPeeringConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RejectVpcPeeringConnectionRequest clone()
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) Clone3() *ServicesEc2ModelRejectVpcPeeringConnectionRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RejectVpcPeeringConnectionRequest")
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
	unique_x := &ServicesEc2ModelRejectVpcPeeringConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelRejectVpcPeeringConnectionRequest) Clone2() (*JavaLangObject, error) {
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


