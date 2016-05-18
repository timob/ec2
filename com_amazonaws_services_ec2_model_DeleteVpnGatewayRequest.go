package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeleteVpnGatewayRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpnGatewayId(java.lang.String)
	SetVpnGatewayId(a string) 

	// public java.lang.String getVpnGatewayId()
	GetVpnGatewayId() string

	// public com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest withVpnGatewayId(java.lang.String)
	WithVpnGatewayId(a string) *ServicesEc2ModelDeleteVpnGatewayRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest clone()
	Clone3() *ServicesEc2ModelDeleteVpnGatewayRequest
}

type ServicesEc2ModelDeleteVpnGatewayRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest()
func NewServicesEc2ModelDeleteVpnGatewayRequest() (*ServicesEc2ModelDeleteVpnGatewayRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteVpnGatewayRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDeleteVpnGatewayRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest(java.lang.String)
func NewServicesEc2ModelDeleteVpnGatewayRequest2(a string) (*ServicesEc2ModelDeleteVpnGatewayRequest) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteVpnGatewayRequest", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ServicesEc2ModelDeleteVpnGatewayRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpnGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) SetVpnGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpnGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpnGatewayId()
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) GetVpnGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpnGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest withVpnGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) WithVpnGatewayId(a string) *ServicesEc2ModelDeleteVpnGatewayRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnGatewayId", "com/amazonaws/services/ec2/model/DeleteVpnGatewayRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeleteVpnGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DeleteVpnGatewayRequest clone()
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) Clone3() *ServicesEc2ModelDeleteVpnGatewayRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DeleteVpnGatewayRequest")
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
	unique_x := &ServicesEc2ModelDeleteVpnGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDeleteVpnGatewayRequest) Clone2() (*JavaLangObject, error) {
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


