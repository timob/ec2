package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyVpcEndpointRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setVpcEndpointId(java.lang.String)
	SetVpcEndpointId(a string) 

	// public java.lang.String getVpcEndpointId()
	GetVpcEndpointId() string

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withVpcEndpointId(java.lang.String)
	WithVpcEndpointId(a string) *ServicesEc2ModelModifyVpcEndpointRequest

	// public void setResetPolicy(java.lang.Boolean)
	SetResetPolicy(a bool) 

	// public java.lang.Boolean getResetPolicy()
	GetResetPolicy() bool

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withResetPolicy(java.lang.Boolean)
	WithResetPolicy(a bool) *ServicesEc2ModelModifyVpcEndpointRequest

	// public java.lang.Boolean isResetPolicy()
	IsResetPolicy() bool

	// public void setPolicyDocument(java.lang.String)
	SetPolicyDocument(a string) 

	// public java.lang.String getPolicyDocument()
	GetPolicyDocument() string

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withPolicyDocument(java.lang.String)
	WithPolicyDocument(a string) *ServicesEc2ModelModifyVpcEndpointRequest

	// public java.util.List<java.lang.String> getAddRouteTableIds()
	GetAddRouteTableIds() []string

	// public void setAddRouteTableIds(java.util.Collection<java.lang.String>)
	SetAddRouteTableIds(a []string) 

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withAddRouteTableIds(java.lang.String...)
	WithAddRouteTableIds(a ...string) *ServicesEc2ModelModifyVpcEndpointRequest

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withAddRouteTableIds(java.util.Collection<java.lang.String>)
	WithAddRouteTableIds2(a []string) *ServicesEc2ModelModifyVpcEndpointRequest

	// public java.util.List<java.lang.String> getRemoveRouteTableIds()
	GetRemoveRouteTableIds() []string

	// public void setRemoveRouteTableIds(java.util.Collection<java.lang.String>)
	SetRemoveRouteTableIds(a []string) 

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withRemoveRouteTableIds(java.lang.String...)
	WithRemoveRouteTableIds(a ...string) *ServicesEc2ModelModifyVpcEndpointRequest

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withRemoveRouteTableIds(java.util.Collection<java.lang.String>)
	WithRemoveRouteTableIds2(a []string) *ServicesEc2ModelModifyVpcEndpointRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest clone()
	Clone3() *ServicesEc2ModelModifyVpcEndpointRequest
}

type ServicesEc2ModelModifyVpcEndpointRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest()
func NewServicesEc2ModelModifyVpcEndpointRequest() (*ServicesEc2ModelModifyVpcEndpointRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyVpcEndpointRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcEndpointId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) SetVpcEndpointId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcEndpointId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcEndpointId()
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) GetVpcEndpointId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcEndpointId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withVpcEndpointId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) WithVpcEndpointId(a string) *ServicesEc2ModelModifyVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcEndpointId", "com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setResetPolicy(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) SetResetPolicy(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResetPolicy", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getResetPolicy()
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) GetResetPolicy() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResetPolicy", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withResetPolicy(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) WithResetPolicy(a bool) *ServicesEc2ModelModifyVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResetPolicy", "com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isResetPolicy()
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) IsResetPolicy() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isResetPolicy", "java/lang/Boolean")
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

// public void setPolicyDocument(java.lang.String)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) SetPolicyDocument(a string)  {
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
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) GetPolicyDocument() string {
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

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withPolicyDocument(java.lang.String)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) WithPolicyDocument(a string) *ServicesEc2ModelModifyVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPolicyDocument", "com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getAddRouteTableIds()
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) GetAddRouteTableIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAddRouteTableIds", "java/util/List")
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

// public void setAddRouteTableIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) SetAddRouteTableIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAddRouteTableIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withAddRouteTableIds(java.lang.String...)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) WithAddRouteTableIds(a ...string) *ServicesEc2ModelModifyVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAddRouteTableIds", "com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withAddRouteTableIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) WithAddRouteTableIds2(a []string) *ServicesEc2ModelModifyVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAddRouteTableIds", "com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getRemoveRouteTableIds()
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) GetRemoveRouteTableIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRemoveRouteTableIds", "java/util/List")
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

// public void setRemoveRouteTableIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) SetRemoveRouteTableIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRemoveRouteTableIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withRemoveRouteTableIds(java.lang.String...)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) WithRemoveRouteTableIds(a ...string) *ServicesEc2ModelModifyVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRemoveRouteTableIds", "com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest withRemoveRouteTableIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) WithRemoveRouteTableIds2(a []string) *ServicesEc2ModelModifyVpcEndpointRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRemoveRouteTableIds", "com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyVpcEndpointRequest clone()
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) Clone3() *ServicesEc2ModelModifyVpcEndpointRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyVpcEndpointRequest")
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
	unique_x := &ServicesEc2ModelModifyVpcEndpointRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifyVpcEndpointRequest) Clone2() (*JavaLangObject, error) {
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


