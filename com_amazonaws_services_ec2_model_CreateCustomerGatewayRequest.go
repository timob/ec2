package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateCustomerGatewayRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setType(java.lang.String)
	SetType2(a string) 

	// public java.lang.String getType()
	GetType() string

	// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest withType(java.lang.String)
	WithType2(a string) *ServicesEc2ModelCreateCustomerGatewayRequest

	// public void setType(com.amazonaws.services.ec2.model.GatewayType)
	SetType(a ServicesEc2ModelGatewayTypeInterface) 

	// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest withType(com.amazonaws.services.ec2.model.GatewayType)
	WithType(a ServicesEc2ModelGatewayTypeInterface) *ServicesEc2ModelCreateCustomerGatewayRequest

	// public void setPublicIp(java.lang.String)
	SetPublicIp(a string) 

	// public java.lang.String getPublicIp()
	GetPublicIp() string

	// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest withPublicIp(java.lang.String)
	WithPublicIp(a string) *ServicesEc2ModelCreateCustomerGatewayRequest

	// public void setBgpAsn(java.lang.Integer)
	SetBgpAsn(a int) 

	// public java.lang.Integer getBgpAsn()
	GetBgpAsn() int

	// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest withBgpAsn(java.lang.Integer)
	WithBgpAsn(a int) *ServicesEc2ModelCreateCustomerGatewayRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest clone()
	Clone3() *ServicesEc2ModelCreateCustomerGatewayRequest
}

type ServicesEc2ModelCreateCustomerGatewayRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest()
func NewServicesEc2ModelCreateCustomerGatewayRequest() (*ServicesEc2ModelCreateCustomerGatewayRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateCustomerGatewayRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateCustomerGatewayRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest(java.lang.String, java.lang.String, java.lang.Integer)
func NewServicesEc2ModelCreateCustomerGatewayRequest3(a string, b string, c int) (*ServicesEc2ModelCreateCustomerGatewayRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateCustomerGatewayRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &ServicesEc2ModelCreateCustomerGatewayRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest(com.amazonaws.services.ec2.model.GatewayType, java.lang.String, java.lang.Integer)
func NewServicesEc2ModelCreateCustomerGatewayRequest2(a ServicesEc2ModelGatewayTypeInterface, b string, c int) (*ServicesEc2ModelCreateCustomerGatewayRequest) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateCustomerGatewayRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GatewayType"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &ServicesEc2ModelCreateCustomerGatewayRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) SetType2(a string)  {
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
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) GetType() string {
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

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest withType(java.lang.String)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) WithType2(a string) *ServicesEc2ModelCreateCustomerGatewayRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/CreateCustomerGatewayRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateCustomerGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setType(com.amazonaws.services.ec2.model.GatewayType)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) SetType(a ServicesEc2ModelGatewayTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/GatewayType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest withType(com.amazonaws.services.ec2.model.GatewayType)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) WithType(a ServicesEc2ModelGatewayTypeInterface) *ServicesEc2ModelCreateCustomerGatewayRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/CreateCustomerGatewayRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GatewayType"))
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
	unique_x := &ServicesEc2ModelCreateCustomerGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) SetPublicIp(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPublicIp", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPublicIp()
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) GetPublicIp() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPublicIp", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest withPublicIp(java.lang.String)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) WithPublicIp(a string) *ServicesEc2ModelCreateCustomerGatewayRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublicIp", "com/amazonaws/services/ec2/model/CreateCustomerGatewayRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateCustomerGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBgpAsn(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) SetBgpAsn(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBgpAsn", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getBgpAsn()
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) GetBgpAsn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBgpAsn", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest withBgpAsn(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) WithBgpAsn(a int) *ServicesEc2ModelCreateCustomerGatewayRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBgpAsn", "com/amazonaws/services/ec2/model/CreateCustomerGatewayRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelCreateCustomerGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateCustomerGatewayRequest clone()
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) Clone3() *ServicesEc2ModelCreateCustomerGatewayRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateCustomerGatewayRequest")
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
	unique_x := &ServicesEc2ModelCreateCustomerGatewayRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateCustomerGatewayRequest) Clone2() (*JavaLangObject, error) {
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


