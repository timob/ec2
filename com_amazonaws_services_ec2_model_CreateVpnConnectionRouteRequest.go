package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpnConnectionRouteRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpnConnectionId(java.lang.String)
	SetVpnConnectionId(a string) 

	// public java.lang.String getVpnConnectionId()
	GetVpnConnectionId() string

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest withVpnConnectionId(java.lang.String)
	WithVpnConnectionId(a string) *ServicesEc2ModelCreateVpnConnectionRouteRequest

	// public void setDestinationCidrBlock(java.lang.String)
	SetDestinationCidrBlock(a string) 

	// public java.lang.String getDestinationCidrBlock()
	GetDestinationCidrBlock() string

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest withDestinationCidrBlock(java.lang.String)
	WithDestinationCidrBlock(a string) *ServicesEc2ModelCreateVpnConnectionRouteRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest clone()
	Clone3() *ServicesEc2ModelCreateVpnConnectionRouteRequest
}

type ServicesEc2ModelCreateVpnConnectionRouteRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest()
func NewServicesEc2ModelCreateVpnConnectionRouteRequest() (*ServicesEc2ModelCreateVpnConnectionRouteRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpnConnectionRouteRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpnConnectionRouteRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpnConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) SetVpnConnectionId(a string)  {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) GetVpnConnectionId() string {
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

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest withVpnConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) WithVpnConnectionId(a string) *ServicesEc2ModelCreateVpnConnectionRouteRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnConnectionId", "com/amazonaws/services/ec2/model/CreateVpnConnectionRouteRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionRouteRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDestinationCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) SetDestinationCidrBlock(a string)  {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) GetDestinationCidrBlock() string {
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

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest withDestinationCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) WithDestinationCidrBlock(a string) *ServicesEc2ModelCreateVpnConnectionRouteRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDestinationCidrBlock", "com/amazonaws/services/ec2/model/CreateVpnConnectionRouteRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionRouteRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRouteRequest clone()
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) Clone3() *ServicesEc2ModelCreateVpnConnectionRouteRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpnConnectionRouteRequest")
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionRouteRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRouteRequest) Clone2() (*JavaLangObject, error) {
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


