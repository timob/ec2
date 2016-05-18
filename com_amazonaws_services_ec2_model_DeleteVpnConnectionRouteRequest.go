package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeleteVpnConnectionRouteRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpnConnectionId(java.lang.String)
	SetVpnConnectionId(a string) 

	// public java.lang.String getVpnConnectionId()
	GetVpnConnectionId() string

	// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest withVpnConnectionId(java.lang.String)
	WithVpnConnectionId(a string) *ServicesEc2ModelDeleteVpnConnectionRouteRequest

	// public void setDestinationCidrBlock(java.lang.String)
	SetDestinationCidrBlock(a string) 

	// public java.lang.String getDestinationCidrBlock()
	GetDestinationCidrBlock() string

	// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest withDestinationCidrBlock(java.lang.String)
	WithDestinationCidrBlock(a string) *ServicesEc2ModelDeleteVpnConnectionRouteRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest clone()
	Clone3() *ServicesEc2ModelDeleteVpnConnectionRouteRequest
}

type ServicesEc2ModelDeleteVpnConnectionRouteRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest()
func NewServicesEc2ModelDeleteVpnConnectionRouteRequest() (*ServicesEc2ModelDeleteVpnConnectionRouteRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DeleteVpnConnectionRouteRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDeleteVpnConnectionRouteRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpnConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) SetVpnConnectionId(a string)  {
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
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) GetVpnConnectionId() string {
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

// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest withVpnConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) WithVpnConnectionId(a string) *ServicesEc2ModelDeleteVpnConnectionRouteRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnConnectionId", "com/amazonaws/services/ec2/model/DeleteVpnConnectionRouteRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeleteVpnConnectionRouteRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDestinationCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) SetDestinationCidrBlock(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDestinationCidrBlock", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDestinationCidrBlock()
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) GetDestinationCidrBlock() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDestinationCidrBlock", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest withDestinationCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) WithDestinationCidrBlock(a string) *ServicesEc2ModelDeleteVpnConnectionRouteRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDestinationCidrBlock", "com/amazonaws/services/ec2/model/DeleteVpnConnectionRouteRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeleteVpnConnectionRouteRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DeleteVpnConnectionRouteRequest clone()
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) Clone3() *ServicesEc2ModelDeleteVpnConnectionRouteRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DeleteVpnConnectionRouteRequest")
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
	unique_x := &ServicesEc2ModelDeleteVpnConnectionRouteRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelDeleteVpnConnectionRouteRequest) Clone2() (*JavaLangObject, error) {
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


