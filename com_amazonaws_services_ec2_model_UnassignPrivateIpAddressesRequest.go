package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelUnassignPrivateIpAddressesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setNetworkInterfaceId(java.lang.String)
	SetNetworkInterfaceId(a string) 

	// public java.lang.String getNetworkInterfaceId()
	GetNetworkInterfaceId() string

	// public com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest withNetworkInterfaceId(java.lang.String)
	WithNetworkInterfaceId(a string) *ServicesEc2ModelUnassignPrivateIpAddressesRequest

	// public java.util.List<java.lang.String> getPrivateIpAddresses()
	GetPrivateIpAddresses() []string

	// public void setPrivateIpAddresses(java.util.Collection<java.lang.String>)
	SetPrivateIpAddresses(a []string) 

	// public com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest withPrivateIpAddresses(java.lang.String...)
	WithPrivateIpAddresses(a ...string) *ServicesEc2ModelUnassignPrivateIpAddressesRequest

	// public com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest withPrivateIpAddresses(java.util.Collection<java.lang.String>)
	WithPrivateIpAddresses2(a []string) *ServicesEc2ModelUnassignPrivateIpAddressesRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest clone()
	Clone3() *ServicesEc2ModelUnassignPrivateIpAddressesRequest
}

type ServicesEc2ModelUnassignPrivateIpAddressesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest()
func NewServicesEc2ModelUnassignPrivateIpAddressesRequest() (*ServicesEc2ModelUnassignPrivateIpAddressesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/UnassignPrivateIpAddressesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelUnassignPrivateIpAddressesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) SetNetworkInterfaceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkInterfaceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNetworkInterfaceId()
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) GetNetworkInterfaceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkInterfaceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest withNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) WithNetworkInterfaceId(a string) *ServicesEc2ModelUnassignPrivateIpAddressesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaceId", "com/amazonaws/services/ec2/model/UnassignPrivateIpAddressesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUnassignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getPrivateIpAddresses()
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) GetPrivateIpAddresses() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateIpAddresses", "java/util/List")
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

// public void setPrivateIpAddresses(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) SetPrivateIpAddresses(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrivateIpAddresses", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest withPrivateIpAddresses(java.lang.String...)
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) WithPrivateIpAddresses(a ...string) *ServicesEc2ModelUnassignPrivateIpAddressesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/UnassignPrivateIpAddressesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUnassignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest withPrivateIpAddresses(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) WithPrivateIpAddresses2(a []string) *ServicesEc2ModelUnassignPrivateIpAddressesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/UnassignPrivateIpAddressesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelUnassignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.UnassignPrivateIpAddressesRequest clone()
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) Clone3() *ServicesEc2ModelUnassignPrivateIpAddressesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/UnassignPrivateIpAddressesRequest")
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
	unique_x := &ServicesEc2ModelUnassignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelUnassignPrivateIpAddressesRequest) Clone2() (*JavaLangObject, error) {
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


