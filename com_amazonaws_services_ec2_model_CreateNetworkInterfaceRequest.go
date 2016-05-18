package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateNetworkInterfaceRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelCreateNetworkInterfaceRequest

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelCreateNetworkInterfaceRequest

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelCreateNetworkInterfaceRequest

	// public java.util.List<java.lang.String> getGroups()
	GetGroups() []string

	// public void setGroups(java.util.Collection<java.lang.String>)
	SetGroups(a []string) 

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withGroups(java.lang.String...)
	WithGroups(a ...string) *ServicesEc2ModelCreateNetworkInterfaceRequest

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withGroups(java.util.Collection<java.lang.String>)
	WithGroups2(a []string) *ServicesEc2ModelCreateNetworkInterfaceRequest

	// public java.util.List<com.amazonaws.services.ec2.model.PrivateIpAddressSpecification> getPrivateIpAddresses()
	GetPrivateIpAddresses() []*ServicesEc2ModelPrivateIpAddressSpecification

	// public void setPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.PrivateIpAddressSpecification>)
	SetPrivateIpAddresses(a []*ServicesEc2ModelPrivateIpAddressSpecification) 

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withPrivateIpAddresses(com.amazonaws.services.ec2.model.PrivateIpAddressSpecification...)
	WithPrivateIpAddresses(a ...*ServicesEc2ModelPrivateIpAddressSpecification) *ServicesEc2ModelCreateNetworkInterfaceRequest

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.PrivateIpAddressSpecification>)
	WithPrivateIpAddresses2(a []*ServicesEc2ModelPrivateIpAddressSpecification) *ServicesEc2ModelCreateNetworkInterfaceRequest

	// public void setSecondaryPrivateIpAddressCount(java.lang.Integer)
	SetSecondaryPrivateIpAddressCount(a int) 

	// public java.lang.Integer getSecondaryPrivateIpAddressCount()
	GetSecondaryPrivateIpAddressCount() int

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withSecondaryPrivateIpAddressCount(java.lang.Integer)
	WithSecondaryPrivateIpAddressCount(a int) *ServicesEc2ModelCreateNetworkInterfaceRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest clone()
	Clone3() *ServicesEc2ModelCreateNetworkInterfaceRequest
}

type ServicesEc2ModelCreateNetworkInterfaceRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest()
func NewServicesEc2ModelCreateNetworkInterfaceRequest() (*ServicesEc2ModelCreateNetworkInterfaceRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) SetSubnetId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSubnetId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSubnetId()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) GetSubnetId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSubnetId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) WithSubnetId(a string) *ServicesEc2ModelCreateNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) SetDescription(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDescription()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) GetDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) WithDescription(a string) *ServicesEc2ModelCreateNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) SetPrivateIpAddress(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrivateIpAddress", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrivateIpAddress()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) GetPrivateIpAddress() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateIpAddress", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) WithPrivateIpAddress(a string) *ServicesEc2ModelCreateNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getGroups()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) GetGroups() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroups", "java/util/List")
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

// public void setGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) SetGroups(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withGroups(java.lang.String...)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) WithGroups(a ...string) *ServicesEc2ModelCreateNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) WithGroups2(a []string) *ServicesEc2ModelCreateNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.PrivateIpAddressSpecification> getPrivateIpAddresses()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) GetPrivateIpAddresses() []*ServicesEc2ModelPrivateIpAddressSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateIpAddresses", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelPrivateIpAddressSpecification)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.PrivateIpAddressSpecification>)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) SetPrivateIpAddresses(a []*ServicesEc2ModelPrivateIpAddressSpecification)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrivateIpAddresses", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withPrivateIpAddresses(com.amazonaws.services.ec2.model.PrivateIpAddressSpecification...)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) WithPrivateIpAddresses(a ...*ServicesEc2ModelPrivateIpAddressSpecification) *ServicesEc2ModelCreateNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/PrivateIpAddressSpecification")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PrivateIpAddressSpecification"))
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withPrivateIpAddresses(java.util.Collection<com.amazonaws.services.ec2.model.PrivateIpAddressSpecification>)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) WithPrivateIpAddresses2(a []*ServicesEc2ModelPrivateIpAddressSpecification) *ServicesEc2ModelCreateNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSecondaryPrivateIpAddressCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) SetSecondaryPrivateIpAddressCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSecondaryPrivateIpAddressCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getSecondaryPrivateIpAddressCount()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) GetSecondaryPrivateIpAddressCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSecondaryPrivateIpAddressCount", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest withSecondaryPrivateIpAddressCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) WithSecondaryPrivateIpAddressCount(a int) *ServicesEc2ModelCreateNetworkInterfaceRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecondaryPrivateIpAddressCount", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateNetworkInterfaceRequest clone()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) Clone3() *ServicesEc2ModelCreateNetworkInterfaceRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateNetworkInterfaceRequest")
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
	unique_x := &ServicesEc2ModelCreateNetworkInterfaceRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelCreateNetworkInterfaceRequest) Clone2() (*JavaLangObject, error) {
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


