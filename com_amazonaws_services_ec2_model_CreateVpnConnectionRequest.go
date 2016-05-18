package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpnConnectionRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setType(java.lang.String)
	SetType(a string) 

	// public java.lang.String getType()
	GetType() string

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest withType(java.lang.String)
	WithType(a string) *ServicesEc2ModelCreateVpnConnectionRequest

	// public void setCustomerGatewayId(java.lang.String)
	SetCustomerGatewayId(a string) 

	// public java.lang.String getCustomerGatewayId()
	GetCustomerGatewayId() string

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest withCustomerGatewayId(java.lang.String)
	WithCustomerGatewayId(a string) *ServicesEc2ModelCreateVpnConnectionRequest

	// public void setVpnGatewayId(java.lang.String)
	SetVpnGatewayId(a string) 

	// public java.lang.String getVpnGatewayId()
	GetVpnGatewayId() string

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest withVpnGatewayId(java.lang.String)
	WithVpnGatewayId(a string) *ServicesEc2ModelCreateVpnConnectionRequest

	// public void setOptions(com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification)
	SetOptions(a ServicesEc2ModelVpnConnectionOptionsSpecificationInterface) 

	// public com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification getOptions()
	GetOptions() *ServicesEc2ModelVpnConnectionOptionsSpecification

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest withOptions(com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification)
	WithOptions(a ServicesEc2ModelVpnConnectionOptionsSpecificationInterface) *ServicesEc2ModelCreateVpnConnectionRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpnConnectionRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest clone()
	Clone3() *ServicesEc2ModelCreateVpnConnectionRequest
}

type ServicesEc2ModelCreateVpnConnectionRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest()
func NewServicesEc2ModelCreateVpnConnectionRequest() (*ServicesEc2ModelCreateVpnConnectionRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpnConnectionRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpnConnectionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest(java.lang.String, java.lang.String, java.lang.String)
func NewServicesEc2ModelCreateVpnConnectionRequest2(a string, b string, c string) (*ServicesEc2ModelCreateVpnConnectionRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpnConnectionRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &ServicesEc2ModelCreateVpnConnectionRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) SetType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getType()
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) GetType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest withType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) WithType(a string) *ServicesEc2ModelCreateVpnConnectionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/CreateVpnConnectionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCustomerGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) SetCustomerGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCustomerGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCustomerGatewayId()
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) GetCustomerGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCustomerGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest withCustomerGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) WithCustomerGatewayId(a string) *ServicesEc2ModelCreateVpnConnectionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCustomerGatewayId", "com/amazonaws/services/ec2/model/CreateVpnConnectionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpnGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) SetVpnGatewayId(a string)  {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) GetVpnGatewayId() string {
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

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest withVpnGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) WithVpnGatewayId(a string) *ServicesEc2ModelCreateVpnConnectionRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnGatewayId", "com/amazonaws/services/ec2/model/CreateVpnConnectionRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOptions(com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) SetOptions(a ServicesEc2ModelVpnConnectionOptionsSpecificationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOptions", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnConnectionOptionsSpecification"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification getOptions()
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) GetOptions() *ServicesEc2ModelVpnConnectionOptionsSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOptions", "com/amazonaws/services/ec2/model/VpnConnectionOptionsSpecification")
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
	unique_x := &ServicesEc2ModelVpnConnectionOptionsSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest withOptions(com.amazonaws.services.ec2.model.VpnConnectionOptionsSpecification)
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) WithOptions(a ServicesEc2ModelVpnConnectionOptionsSpecificationInterface) *ServicesEc2ModelCreateVpnConnectionRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOptions", "com/amazonaws/services/ec2/model/CreateVpnConnectionRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnConnectionOptionsSpecification"))
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpnConnectionRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpnConnectionRequest clone()
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) Clone3() *ServicesEc2ModelCreateVpnConnectionRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpnConnectionRequest")
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
	unique_x := &ServicesEc2ModelCreateVpnConnectionRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateVpnConnectionRequest) Clone2() (*JavaLangObject, error) {
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


