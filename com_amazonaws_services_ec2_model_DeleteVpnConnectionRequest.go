package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeleteVpnConnectionRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpnConnectionId(java.lang.String)
	SetVpnConnectionId(a string) 

	// public java.lang.String getVpnConnectionId()
	GetVpnConnectionId() string

	// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest withVpnConnectionId(java.lang.String)
	WithVpnConnectionId(a string) *ServicesEc2ModelDeleteVpnConnectionRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest clone()
	Clone3() *ServicesEc2ModelDeleteVpnConnectionRequest
}

type ServicesEc2ModelDeleteVpnConnectionRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest()
func NewServicesEc2ModelDeleteVpnConnectionRequest() (*ServicesEc2ModelDeleteVpnConnectionRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteVpnConnectionRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDeleteVpnConnectionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest(java.lang.String)
func NewServicesEc2ModelDeleteVpnConnectionRequest2(a string) (*ServicesEc2ModelDeleteVpnConnectionRequest) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteVpnConnectionRequest", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelDeleteVpnConnectionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpnConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) SetVpnConnectionId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpnConnectionId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpnConnectionId()
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) GetVpnConnectionId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpnConnectionId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest withVpnConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) WithVpnConnectionId(a string) *ServicesEc2ModelDeleteVpnConnectionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnConnectionId", "com/amazonaws/services/ec2/model/DeleteVpnConnectionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeleteVpnConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRequest clone()
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) Clone3() *ServicesEc2ModelDeleteVpnConnectionRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DeleteVpnConnectionRequest")
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
	unique_x := &ServicesEc2ModelDeleteVpnConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRequest) Clone2() (*JavaLangObject, error) {
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


