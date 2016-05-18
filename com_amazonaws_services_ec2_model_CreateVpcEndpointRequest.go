package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVpcEndpointRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelCreateVpcEndpointRequest

	// public void setServiceName(java.lang.String)
	SetServiceName(a string) 

	// public java.lang.String getServiceName()
	GetServiceName() string

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withServiceName(java.lang.String)
	WithServiceName(a string) *ServicesEc2ModelCreateVpcEndpointRequest

	// public void setPolicyDocument(java.lang.String)
	SetPolicyDocument(a string) 

	// public java.lang.String getPolicyDocument()
	GetPolicyDocument() string

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withPolicyDocument(java.lang.String)
	WithPolicyDocument(a string) *ServicesEc2ModelCreateVpcEndpointRequest

	// public java.util.List<java.lang.String> getRouteTableIds()
	GetRouteTableIds() []string

	// public void setRouteTableIds(java.util.Collection<java.lang.String>)
	SetRouteTableIds(a []string) 

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withRouteTableIds(java.lang.String...)
	WithRouteTableIds(a ...string) *ServicesEc2ModelCreateVpcEndpointRequest

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withRouteTableIds(java.util.Collection<java.lang.String>)
	WithRouteTableIds2(a []string) *ServicesEc2ModelCreateVpcEndpointRequest

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelCreateVpcEndpointRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpcEndpointRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest clone()
	Clone3() *ServicesEc2ModelCreateVpcEndpointRequest
}

type ServicesEc2ModelCreateVpcEndpointRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest()
func NewServicesEc2ModelCreateVpcEndpointRequest() (*ServicesEc2ModelCreateVpcEndpointRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVpcEndpointRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVpcEndpointRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) WithVpcId(a string) *ServicesEc2ModelCreateVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/CreateVpcEndpointRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setServiceName(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) SetServiceName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setServiceName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getServiceName()
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) GetServiceName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getServiceName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withServiceName(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) WithServiceName(a string) *ServicesEc2ModelCreateVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withServiceName", "com/amazonaws/services/ec2/model/CreateVpcEndpointRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPolicyDocument(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) SetPolicyDocument(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPolicyDocument", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPolicyDocument()
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) GetPolicyDocument() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPolicyDocument", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withPolicyDocument(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) WithPolicyDocument(a string) *ServicesEc2ModelCreateVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPolicyDocument", "com/amazonaws/services/ec2/model/CreateVpcEndpointRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getRouteTableIds()
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) GetRouteTableIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRouteTableIds", "java/util/List")
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

// public void setRouteTableIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) SetRouteTableIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRouteTableIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withRouteTableIds(java.lang.String...)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) WithRouteTableIds(a ...string) *ServicesEc2ModelCreateVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTableIds", "com/amazonaws/services/ec2/model/CreateVpcEndpointRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withRouteTableIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) WithRouteTableIds2(a []string) *ServicesEc2ModelCreateVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRouteTableIds", "com/amazonaws/services/ec2/model/CreateVpcEndpointRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) SetClientToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getClientToken()
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) GetClientToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) WithClientToken(a string) *ServicesEc2ModelCreateVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/CreateVpcEndpointRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateVpcEndpointRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVpcEndpointRequest clone()
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) Clone3() *ServicesEc2ModelCreateVpcEndpointRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVpcEndpointRequest")
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
	unique_x := &ServicesEc2ModelCreateVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateVpcEndpointRequest) Clone2() (*JavaLangObject, error) {
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


