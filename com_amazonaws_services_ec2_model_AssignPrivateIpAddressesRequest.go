package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAssignPrivateIpAddressesRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setNetworkInterfaceId(java.lang.String)
	SetNetworkInterfaceId(a string) 

	// public java.lang.String getNetworkInterfaceId()
	GetNetworkInterfaceId() string

	// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withNetworkInterfaceId(java.lang.String)
	WithNetworkInterfaceId(a string) *ServicesEc2ModelAssignPrivateIpAddressesRequest

	// public java.util.List<java.lang.String> getPrivateIpAddresses()
	GetPrivateIpAddresses() []string

	// public void setPrivateIpAddresses(java.util.Collection<java.lang.String>)
	SetPrivateIpAddresses(a []string) 

	// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withPrivateIpAddresses(java.lang.String...)
	WithPrivateIpAddresses(a ...string) *ServicesEc2ModelAssignPrivateIpAddressesRequest

	// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withPrivateIpAddresses(java.util.Collection<java.lang.String>)
	WithPrivateIpAddresses2(a []string) *ServicesEc2ModelAssignPrivateIpAddressesRequest

	// public void setSecondaryPrivateIpAddressCount(java.lang.Integer)
	SetSecondaryPrivateIpAddressCount(a int) 

	// public java.lang.Integer getSecondaryPrivateIpAddressCount()
	GetSecondaryPrivateIpAddressCount() int

	// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withSecondaryPrivateIpAddressCount(java.lang.Integer)
	WithSecondaryPrivateIpAddressCount(a int) *ServicesEc2ModelAssignPrivateIpAddressesRequest

	// public void setAllowReassignment(java.lang.Boolean)
	SetAllowReassignment(a bool) 

	// public java.lang.Boolean getAllowReassignment()
	GetAllowReassignment() bool

	// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withAllowReassignment(java.lang.Boolean)
	WithAllowReassignment(a bool) *ServicesEc2ModelAssignPrivateIpAddressesRequest

	// public java.lang.Boolean isAllowReassignment()
	IsAllowReassignment() bool

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest clone()
	Clone3() *ServicesEc2ModelAssignPrivateIpAddressesRequest
}

type ServicesEc2ModelAssignPrivateIpAddressesRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest()
func NewServicesEc2ModelAssignPrivateIpAddressesRequest() (*ServicesEc2ModelAssignPrivateIpAddressesRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AssignPrivateIpAddressesRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAssignPrivateIpAddressesRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) SetNetworkInterfaceId(a string)  {
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
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) GetNetworkInterfaceId() string {
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

// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) WithNetworkInterfaceId(a string) *ServicesEc2ModelAssignPrivateIpAddressesRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaceId", "com/amazonaws/services/ec2/model/AssignPrivateIpAddressesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAssignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getPrivateIpAddresses()
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) GetPrivateIpAddresses() []string {
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
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) SetPrivateIpAddresses(a []string)  {
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

// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withPrivateIpAddresses(java.lang.String...)
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) WithPrivateIpAddresses(a ...string) *ServicesEc2ModelAssignPrivateIpAddressesRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/AssignPrivateIpAddressesRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAssignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withPrivateIpAddresses(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) WithPrivateIpAddresses2(a []string) *ServicesEc2ModelAssignPrivateIpAddressesRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddresses", "com/amazonaws/services/ec2/model/AssignPrivateIpAddressesRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelAssignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSecondaryPrivateIpAddressCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) SetSecondaryPrivateIpAddressCount(a int)  {
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
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) GetSecondaryPrivateIpAddressCount() int {
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

// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withSecondaryPrivateIpAddressCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) WithSecondaryPrivateIpAddressCount(a int) *ServicesEc2ModelAssignPrivateIpAddressesRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSecondaryPrivateIpAddressCount", "com/amazonaws/services/ec2/model/AssignPrivateIpAddressesRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelAssignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAllowReassignment(java.lang.Boolean)
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) SetAllowReassignment(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAllowReassignment", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getAllowReassignment()
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) GetAllowReassignment() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAllowReassignment", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest withAllowReassignment(java.lang.Boolean)
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) WithAllowReassignment(a bool) *ServicesEc2ModelAssignPrivateIpAddressesRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAllowReassignment", "com/amazonaws/services/ec2/model/AssignPrivateIpAddressesRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelAssignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isAllowReassignment()
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) IsAllowReassignment() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isAllowReassignment", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AssignPrivateIpAddressesRequest clone()
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) Clone3() *ServicesEc2ModelAssignPrivateIpAddressesRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AssignPrivateIpAddressesRequest")
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
	unique_x := &ServicesEc2ModelAssignPrivateIpAddressesRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelAssignPrivateIpAddressesRequest) Clone2() (*JavaLangObject, error) {
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


