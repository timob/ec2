package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyImageAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setImageId(java.lang.String)
	SetImageId(a string) 

	// public java.lang.String getImageId()
	GetImageId() string

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withImageId(java.lang.String)
	WithImageId(a string) *ServicesEc2ModelModifyImageAttributeRequest

	// public void setAttribute(java.lang.String)
	SetAttribute(a string) 

	// public java.lang.String getAttribute()
	GetAttribute() string

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withAttribute(java.lang.String)
	WithAttribute(a string) *ServicesEc2ModelModifyImageAttributeRequest

	// public void setOperationType(java.lang.String)
	SetOperationType2(a string) 

	// public java.lang.String getOperationType()
	GetOperationType() string

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withOperationType(java.lang.String)
	WithOperationType2(a string) *ServicesEc2ModelModifyImageAttributeRequest

	// public void setOperationType(com.amazonaws.services.ec2.model.OperationType)
	SetOperationType(a ServicesEc2ModelOperationTypeInterface) 

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withOperationType(com.amazonaws.services.ec2.model.OperationType)
	WithOperationType(a ServicesEc2ModelOperationTypeInterface) *ServicesEc2ModelModifyImageAttributeRequest

	// public java.util.List<java.lang.String> getUserIds()
	GetUserIds() []string

	// public void setUserIds(java.util.Collection<java.lang.String>)
	SetUserIds(a []string) 

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withUserIds(java.lang.String...)
	WithUserIds(a ...string) *ServicesEc2ModelModifyImageAttributeRequest

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withUserIds(java.util.Collection<java.lang.String>)
	WithUserIds2(a []string) *ServicesEc2ModelModifyImageAttributeRequest

	// public java.util.List<java.lang.String> getUserGroups()
	GetUserGroups() []string

	// public void setUserGroups(java.util.Collection<java.lang.String>)
	SetUserGroups(a []string) 

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withUserGroups(java.lang.String...)
	WithUserGroups(a ...string) *ServicesEc2ModelModifyImageAttributeRequest

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withUserGroups(java.util.Collection<java.lang.String>)
	WithUserGroups2(a []string) *ServicesEc2ModelModifyImageAttributeRequest

	// public java.util.List<java.lang.String> getProductCodes()
	GetProductCodes() []string

	// public void setProductCodes(java.util.Collection<java.lang.String>)
	SetProductCodes(a []string) 

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withProductCodes(java.lang.String...)
	WithProductCodes(a ...string) *ServicesEc2ModelModifyImageAttributeRequest

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withProductCodes(java.util.Collection<java.lang.String>)
	WithProductCodes2(a []string) *ServicesEc2ModelModifyImageAttributeRequest

	// public void setValue(java.lang.String)
	SetValue(a string) 

	// public java.lang.String getValue()
	GetValue() string

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withValue(java.lang.String)
	WithValue(a string) *ServicesEc2ModelModifyImageAttributeRequest

	// public void setLaunchPermission(com.amazonaws.services.ec2.model.LaunchPermissionModifications)
	SetLaunchPermission(a ServicesEc2ModelLaunchPermissionModificationsInterface) 

	// public com.amazonaws.services.ec2.model.LaunchPermissionModifications getLaunchPermission()
	GetLaunchPermission() *ServicesEc2ModelLaunchPermissionModifications

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withLaunchPermission(com.amazonaws.services.ec2.model.LaunchPermissionModifications)
	WithLaunchPermission(a ServicesEc2ModelLaunchPermissionModificationsInterface) *ServicesEc2ModelModifyImageAttributeRequest

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelModifyImageAttributeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyImageAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest clone()
	Clone3() *ServicesEc2ModelModifyImageAttributeRequest
}

type ServicesEc2ModelModifyImageAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest()
func NewServicesEc2ModelModifyImageAttributeRequest() (*ServicesEc2ModelModifyImageAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyImageAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyImageAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest(java.lang.String, java.lang.String)
func NewServicesEc2ModelModifyImageAttributeRequest2(a string, b string) (*ServicesEc2ModelModifyImageAttributeRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelModifyImageAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImageId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetImageId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImageId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImageId()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetImageId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withImageId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithImageId(a string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageId", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetAttribute(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttribute", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAttribute()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetAttribute() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttribute", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithAttribute(a string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOperationType(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetOperationType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOperationType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOperationType()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetOperationType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOperationType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withOperationType(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithOperationType2(a string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOperationType", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOperationType(com.amazonaws.services.ec2.model.OperationType)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetOperationType(a ServicesEc2ModelOperationTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOperationType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/OperationType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withOperationType(com.amazonaws.services.ec2.model.OperationType)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithOperationType(a ServicesEc2ModelOperationTypeInterface) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOperationType", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/OperationType"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getUserIds()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetUserIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserIds", "java/util/List")
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

// public void setUserIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetUserIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withUserIds(java.lang.String...)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithUserIds(a ...string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserIds", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withUserIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithUserIds2(a []string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserIds", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getUserGroups()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetUserGroups() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserGroups", "java/util/List")
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

// public void setUserGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetUserGroups(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withUserGroups(java.lang.String...)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithUserGroups(a ...string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserGroups", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withUserGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithUserGroups2(a []string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserGroups", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getProductCodes()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetProductCodes() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductCodes", "java/util/List")
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

// public void setProductCodes(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetProductCodes(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductCodes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withProductCodes(java.lang.String...)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithProductCodes(a ...string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withProductCodes(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithProductCodes2(a []string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setValue(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetValue(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValue", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getValue()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetValue() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValue", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withValue(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithValue(a string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValue", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLaunchPermission(com.amazonaws.services.ec2.model.LaunchPermissionModifications)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetLaunchPermission(a ServicesEc2ModelLaunchPermissionModificationsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLaunchPermission", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/LaunchPermissionModifications"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.LaunchPermissionModifications getLaunchPermission()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetLaunchPermission() *ServicesEc2ModelLaunchPermissionModifications {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLaunchPermission", "com/amazonaws/services/ec2/model/LaunchPermissionModifications")
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
	unique_x := &ServicesEc2ModelLaunchPermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withLaunchPermission(com.amazonaws.services.ec2.model.LaunchPermissionModifications)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithLaunchPermission(a ServicesEc2ModelLaunchPermissionModificationsInterface) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLaunchPermission", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/LaunchPermissionModifications"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) WithDescription(a string) *ServicesEc2ModelModifyImageAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyImageAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyImageAttributeRequest clone()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) Clone3() *ServicesEc2ModelModifyImageAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyImageAttributeRequest")
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
	unique_x := &ServicesEc2ModelModifyImageAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifyImageAttributeRequest) Clone2() (*JavaLangObject, error) {
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


