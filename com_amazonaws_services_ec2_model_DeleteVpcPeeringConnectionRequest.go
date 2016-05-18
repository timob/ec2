package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeleteVpcPeeringConnectionRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpcPeeringConnectionId(java.lang.String)
	SetVpcPeeringConnectionId(a string) 

	// public java.lang.String getVpcPeeringConnectionId()
	GetVpcPeeringConnectionId() string

	// public com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionRequest withVpcPeeringConnectionId(java.lang.String)
	WithVpcPeeringConnectionId(a string) *ServicesEc2ModelDeleteVpcPeeringConnectionRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionRequest clone()
	Clone3() *ServicesEc2ModelDeleteVpcPeeringConnectionRequest
}

type ServicesEc2ModelDeleteVpcPeeringConnectionRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionRequest()
func NewServicesEc2ModelDeleteVpcPeeringConnectionRequest() (*ServicesEc2ModelDeleteVpcPeeringConnectionRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteVpcPeeringConnectionRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDeleteVpcPeeringConnectionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) SetVpcPeeringConnectionId(a string)  {
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
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) GetVpcPeeringConnectionId() string {
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

// public com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionRequest withVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) WithVpcPeeringConnectionId(a string) *ServicesEc2ModelDeleteVpcPeeringConnectionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnectionId", "com/amazonaws/services/ec2/model/DeleteVpcPeeringConnectionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeleteVpcPeeringConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DeleteVpcPeeringConnectionRequest clone()
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) Clone3() *ServicesEc2ModelDeleteVpcPeeringConnectionRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DeleteVpcPeeringConnectionRequest")
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
	unique_x := &ServicesEc2ModelDeleteVpcPeeringConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDeleteVpcPeeringConnectionRequest) Clone2() (*JavaLangObject, error) {
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


