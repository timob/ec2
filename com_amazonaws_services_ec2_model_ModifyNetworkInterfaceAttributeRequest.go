package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyNetworkInterfaceAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setNetworkInterfaceId(java.lang.String)
	SetNetworkInterfaceId(a string) 

	// public java.lang.String getNetworkInterfaceId()
	GetNetworkInterfaceId() string

	// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withNetworkInterfaceId(java.lang.String)
	WithNetworkInterfaceId(a string) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest

	// public void setSourceDestCheck(java.lang.Boolean)
	SetSourceDestCheck(a bool) 

	// public java.lang.Boolean getSourceDestCheck()
	GetSourceDestCheck() bool

	// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withSourceDestCheck(java.lang.Boolean)
	WithSourceDestCheck(a bool) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest

	// public java.lang.Boolean isSourceDestCheck()
	IsSourceDestCheck() bool

	// public java.util.List<java.lang.String> getGroups()
	GetGroups() []string

	// public void setGroups(java.util.Collection<java.lang.String>)
	SetGroups(a []string) 

	// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withGroups(java.lang.String...)
	WithGroups(a ...string) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest

	// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withGroups(java.util.Collection<java.lang.String>)
	WithGroups2(a []string) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest

	// public void setAttachment(com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges)
	SetAttachment(a ServicesEc2ModelNetworkInterfaceAttachmentChangesInterface) 

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges getAttachment()
	GetAttachment() *ServicesEc2ModelNetworkInterfaceAttachmentChanges

	// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withAttachment(com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges)
	WithAttachment(a ServicesEc2ModelNetworkInterfaceAttachmentChangesInterface) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest clone()
	Clone3() *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest
}

type ServicesEc2ModelModifyNetworkInterfaceAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest()
func NewServicesEc2ModelModifyNetworkInterfaceAttributeRequest() (*ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyNetworkInterfaceAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyNetworkInterfaceAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) SetNetworkInterfaceId(a string)  {
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
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) GetNetworkInterfaceId() string {
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

// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) WithNetworkInterfaceId(a string) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaceId", "com/amazonaws/services/ec2/model/ModifyNetworkInterfaceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) WithDescription(a string) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ModifyNetworkInterfaceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) SetSourceDestCheck(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSourceDestCheck", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getSourceDestCheck()
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) GetSourceDestCheck() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSourceDestCheck", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) WithSourceDestCheck(a bool) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceDestCheck", "com/amazonaws/services/ec2/model/ModifyNetworkInterfaceAttributeRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifyNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isSourceDestCheck()
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) IsSourceDestCheck() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isSourceDestCheck", "java/lang/Boolean")
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

// public java.util.List<java.lang.String> getGroups()
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) GetGroups() []string {
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
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) SetGroups(a []string)  {
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

// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withGroups(java.lang.String...)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) WithGroups(a ...string) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/ModifyNetworkInterfaceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) WithGroups2(a []string) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/ModifyNetworkInterfaceAttributeRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttachment(com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) SetAttachment(a ServicesEc2ModelNetworkInterfaceAttachmentChangesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachment", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceAttachmentChanges"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges getAttachment()
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) GetAttachment() *ServicesEc2ModelNetworkInterfaceAttachmentChanges {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachment", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachmentChanges")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachmentChanges{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest withAttachment(com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges)
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) WithAttachment(a ServicesEc2ModelNetworkInterfaceAttachmentChangesInterface) *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachment", "com/amazonaws/services/ec2/model/ModifyNetworkInterfaceAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkInterfaceAttachmentChanges"))
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
	unique_x := &ServicesEc2ModelModifyNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyNetworkInterfaceAttributeRequest clone()
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) Clone3() *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyNetworkInterfaceAttributeRequest")
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
	unique_x := &ServicesEc2ModelModifyNetworkInterfaceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifyNetworkInterfaceAttributeRequest) Clone2() (*JavaLangObject, error) {
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


