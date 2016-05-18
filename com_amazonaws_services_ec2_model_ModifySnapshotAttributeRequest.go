package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifySnapshotAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public void setAttribute(java.lang.String)
	SetAttribute2(a string) 

	// public java.lang.String getAttribute()
	GetAttribute() string

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withAttribute(java.lang.String)
	WithAttribute2(a string) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public void setAttribute(com.amazonaws.services.ec2.model.SnapshotAttributeName)
	SetAttribute(a ServicesEc2ModelSnapshotAttributeNameInterface) 

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withAttribute(com.amazonaws.services.ec2.model.SnapshotAttributeName)
	WithAttribute(a ServicesEc2ModelSnapshotAttributeNameInterface) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public void setOperationType(java.lang.String)
	SetOperationType2(a string) 

	// public java.lang.String getOperationType()
	GetOperationType() string

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withOperationType(java.lang.String)
	WithOperationType2(a string) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public void setOperationType(com.amazonaws.services.ec2.model.OperationType)
	SetOperationType(a ServicesEc2ModelOperationTypeInterface) 

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withOperationType(com.amazonaws.services.ec2.model.OperationType)
	WithOperationType(a ServicesEc2ModelOperationTypeInterface) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public java.util.List<java.lang.String> getUserIds()
	GetUserIds() []string

	// public void setUserIds(java.util.Collection<java.lang.String>)
	SetUserIds(a []string) 

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withUserIds(java.lang.String...)
	WithUserIds(a ...string) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withUserIds(java.util.Collection<java.lang.String>)
	WithUserIds2(a []string) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public java.util.List<java.lang.String> getGroupNames()
	GetGroupNames() []string

	// public void setGroupNames(java.util.Collection<java.lang.String>)
	SetGroupNames(a []string) 

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withGroupNames(java.lang.String...)
	WithGroupNames(a ...string) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withGroupNames(java.util.Collection<java.lang.String>)
	WithGroupNames2(a []string) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public void setCreateVolumePermission(com.amazonaws.services.ec2.model.CreateVolumePermissionModifications)
	SetCreateVolumePermission(a ServicesEc2ModelCreateVolumePermissionModificationsInterface) 

	// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications getCreateVolumePermission()
	GetCreateVolumePermission() *ServicesEc2ModelCreateVolumePermissionModifications

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withCreateVolumePermission(com.amazonaws.services.ec2.model.CreateVolumePermissionModifications)
	WithCreateVolumePermission(a ServicesEc2ModelCreateVolumePermissionModificationsInterface) *ServicesEc2ModelModifySnapshotAttributeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest clone()
	Clone3() *ServicesEc2ModelModifySnapshotAttributeRequest
}

type ServicesEc2ModelModifySnapshotAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest()
func NewServicesEc2ModelModifySnapshotAttributeRequest() (*ServicesEc2ModelModifySnapshotAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest(java.lang.String, java.lang.String, java.lang.String)
func NewServicesEc2ModelModifySnapshotAttributeRequest3(a string, b string, c string) (*ServicesEc2ModelModifySnapshotAttributeRequest) {
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

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest(java.lang.String, com.amazonaws.services.ec2.model.SnapshotAttributeName, com.amazonaws.services.ec2.model.OperationType)
func NewServicesEc2ModelModifySnapshotAttributeRequest2(a string, b ServicesEc2ModelSnapshotAttributeNameInterface, c ServicesEc2ModelOperationTypeInterface) (*ServicesEc2ModelModifySnapshotAttributeRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("com/amazonaws/services/ec2/model/SnapshotAttributeName"), conv_c.Value().Cast("com/amazonaws/services/ec2/model/OperationType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) SetSnapshotId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshotId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSnapshotId()
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) GetSnapshotId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshotId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithSnapshotId(a string) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) SetAttribute2(a string)  {
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
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) GetAttribute() string {
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

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithAttribute2(a string) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(com.amazonaws.services.ec2.model.SnapshotAttributeName)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) SetAttribute(a ServicesEc2ModelSnapshotAttributeNameInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/SnapshotAttributeName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withAttribute(com.amazonaws.services.ec2.model.SnapshotAttributeName)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithAttribute(a ServicesEc2ModelSnapshotAttributeNameInterface) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SnapshotAttributeName"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOperationType(java.lang.String)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) SetOperationType2(a string)  {
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
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) GetOperationType() string {
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

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withOperationType(java.lang.String)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithOperationType2(a string) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOperationType", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOperationType(com.amazonaws.services.ec2.model.OperationType)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) SetOperationType(a ServicesEc2ModelOperationTypeInterface)  {
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

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withOperationType(com.amazonaws.services.ec2.model.OperationType)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithOperationType(a ServicesEc2ModelOperationTypeInterface) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOperationType", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/OperationType"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getUserIds()
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) GetUserIds() []string {
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
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) SetUserIds(a []string)  {
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

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withUserIds(java.lang.String...)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithUserIds(a ...string) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserIds", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withUserIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithUserIds2(a []string) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserIds", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getGroupNames()
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) GetGroupNames() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupNames", "java/util/List")
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

// public void setGroupNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) SetGroupNames(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupNames", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withGroupNames(java.lang.String...)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithGroupNames(a ...string) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupNames", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withGroupNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithGroupNames2(a []string) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupNames", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreateVolumePermission(com.amazonaws.services.ec2.model.CreateVolumePermissionModifications)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) SetCreateVolumePermission(a ServicesEc2ModelCreateVolumePermissionModificationsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreateVolumePermission", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVolumePermissionModifications"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications getCreateVolumePermission()
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) GetCreateVolumePermission() *ServicesEc2ModelCreateVolumePermissionModifications {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreateVolumePermission", "com/amazonaws/services/ec2/model/CreateVolumePermissionModifications")
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
	unique_x := &ServicesEc2ModelCreateVolumePermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest withCreateVolumePermission(com.amazonaws.services.ec2.model.CreateVolumePermissionModifications)
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) WithCreateVolumePermission(a ServicesEc2ModelCreateVolumePermissionModificationsInterface) *ServicesEc2ModelModifySnapshotAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateVolumePermission", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVolumePermissionModifications"))
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifySnapshotAttributeRequest clone()
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) Clone3() *ServicesEc2ModelModifySnapshotAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifySnapshotAttributeRequest")
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
	unique_x := &ServicesEc2ModelModifySnapshotAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifySnapshotAttributeRequest) Clone2() (*JavaLangObject, error) {
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


