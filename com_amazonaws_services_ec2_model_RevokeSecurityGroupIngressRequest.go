package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRevokeSecurityGroupIngressRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setGroupName(java.lang.String)
	SetGroupName(a string) 

	// public java.lang.String getGroupName()
	GetGroupName() string

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withGroupName(java.lang.String)
	WithGroupName(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public void setGroupId(java.lang.String)
	SetGroupId(a string) 

	// public java.lang.String getGroupId()
	GetGroupId() string

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withGroupId(java.lang.String)
	WithGroupId(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public void setSourceSecurityGroupName(java.lang.String)
	SetSourceSecurityGroupName(a string) 

	// public java.lang.String getSourceSecurityGroupName()
	GetSourceSecurityGroupName() string

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withSourceSecurityGroupName(java.lang.String)
	WithSourceSecurityGroupName(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public void setSourceSecurityGroupOwnerId(java.lang.String)
	SetSourceSecurityGroupOwnerId(a string) 

	// public java.lang.String getSourceSecurityGroupOwnerId()
	GetSourceSecurityGroupOwnerId() string

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withSourceSecurityGroupOwnerId(java.lang.String)
	WithSourceSecurityGroupOwnerId(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public void setIpProtocol(java.lang.String)
	SetIpProtocol(a string) 

	// public java.lang.String getIpProtocol()
	GetIpProtocol() string

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withIpProtocol(java.lang.String)
	WithIpProtocol(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public void setFromPort(java.lang.Integer)
	SetFromPort(a int) 

	// public java.lang.Integer getFromPort()
	GetFromPort() int

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withFromPort(java.lang.Integer)
	WithFromPort(a int) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public void setToPort(java.lang.Integer)
	SetToPort(a int) 

	// public java.lang.Integer getToPort()
	GetToPort() int

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withToPort(java.lang.Integer)
	WithToPort(a int) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public void setCidrIp(java.lang.String)
	SetCidrIp(a string) 

	// public java.lang.String getCidrIp()
	GetCidrIp() string

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withCidrIp(java.lang.String)
	WithCidrIp(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public java.util.List<com.amazonaws.services.ec2.model.IpPermission> getIpPermissions()
	GetIpPermissions() []*ServicesEc2ModelIpPermission

	// public void setIpPermissions(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
	SetIpPermissions(a []*ServicesEc2ModelIpPermission) 

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withIpPermissions(com.amazonaws.services.ec2.model.IpPermission...)
	WithIpPermissions(a ...*ServicesEc2ModelIpPermission) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withIpPermissions(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
	WithIpPermissions2(a []*ServicesEc2ModelIpPermission) *ServicesEc2ModelRevokeSecurityGroupIngressRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest clone()
	Clone3() *ServicesEc2ModelRevokeSecurityGroupIngressRequest
}

type ServicesEc2ModelRevokeSecurityGroupIngressRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest()
func NewServicesEc2ModelRevokeSecurityGroupIngressRequest() (*ServicesEc2ModelRevokeSecurityGroupIngressRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest(java.lang.String, java.util.List<com.amazonaws.services.ec2.model.IpPermission>)
func NewServicesEc2ModelRevokeSecurityGroupIngressRequest2(a string, b []*ServicesEc2ModelIpPermission) (*ServicesEc2ModelRevokeSecurityGroupIngressRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) SetGroupName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGroupName()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetGroupName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithGroupName(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupName", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGroupId(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) SetGroupId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGroupId()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetGroupId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withGroupId(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithGroupId(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupId", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceSecurityGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) SetSourceSecurityGroupName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSourceSecurityGroupName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSourceSecurityGroupName()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetSourceSecurityGroupName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSourceSecurityGroupName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withSourceSecurityGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithSourceSecurityGroupName(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceSecurityGroupName", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceSecurityGroupOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) SetSourceSecurityGroupOwnerId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSourceSecurityGroupOwnerId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSourceSecurityGroupOwnerId()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetSourceSecurityGroupOwnerId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSourceSecurityGroupOwnerId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withSourceSecurityGroupOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithSourceSecurityGroupOwnerId(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceSecurityGroupOwnerId", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIpProtocol(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) SetIpProtocol(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIpProtocol", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getIpProtocol()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetIpProtocol() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIpProtocol", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withIpProtocol(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithIpProtocol(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpProtocol", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFromPort(java.lang.Integer)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) SetFromPort(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFromPort", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getFromPort()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetFromPort() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFromPort", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withFromPort(java.lang.Integer)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithFromPort(a int) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFromPort", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setToPort(java.lang.Integer)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) SetToPort(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setToPort", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getToPort()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetToPort() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getToPort", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withToPort(java.lang.Integer)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithToPort(a int) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withToPort", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCidrIp(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) SetCidrIp(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCidrIp", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCidrIp()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetCidrIp() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCidrIp", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withCidrIp(java.lang.String)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithCidrIp(a string) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCidrIp", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.IpPermission> getIpPermissions()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetIpPermissions() []*ServicesEc2ModelIpPermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIpPermissions", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelIpPermission)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setIpPermissions(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) SetIpPermissions(a []*ServicesEc2ModelIpPermission)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIpPermissions", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withIpPermissions(com.amazonaws.services.ec2.model.IpPermission...)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithIpPermissions(a ...*ServicesEc2ModelIpPermission) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/IpPermission")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpPermissions", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/IpPermission"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest withIpPermissions(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) WithIpPermissions2(a []*ServicesEc2ModelIpPermission) *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpPermissions", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RevokeSecurityGroupIngressRequest clone()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) Clone3() *ServicesEc2ModelRevokeSecurityGroupIngressRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RevokeSecurityGroupIngressRequest")
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
	unique_x := &ServicesEc2ModelRevokeSecurityGroupIngressRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelRevokeSecurityGroupIngressRequest) Clone2() (*JavaLangObject, error) {
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


