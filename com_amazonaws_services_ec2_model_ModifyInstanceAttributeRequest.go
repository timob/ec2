package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelModifyInstanceAttributeRequestInterface interface {
	AmazonWebServiceRequestInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public void setAttribute(java.lang.String)
	SetAttribute2(a string) 

	// public java.lang.String getAttribute()
	GetAttribute() string

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withAttribute(java.lang.String)
	WithAttribute2(a string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public void setAttribute(com.amazonaws.services.ec2.model.InstanceAttributeName)
	SetAttribute(a ServicesEc2ModelInstanceAttributeNameInterface) 

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withAttribute(com.amazonaws.services.ec2.model.InstanceAttributeName)
	WithAttribute(a ServicesEc2ModelInstanceAttributeNameInterface) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public void setValue(java.lang.String)
	SetValue(a string) 

	// public java.lang.String getValue()
	GetValue() string

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withValue(java.lang.String)
	WithValue(a string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification> getBlockDeviceMappings()
	GetBlockDeviceMappings() []*ServicesEc2ModelInstanceBlockDeviceMappingSpecification

	// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification>)
	SetBlockDeviceMappings(a []*ServicesEc2ModelInstanceBlockDeviceMappingSpecification) 

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withBlockDeviceMappings(com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification...)
	WithBlockDeviceMappings(a ...*ServicesEc2ModelInstanceBlockDeviceMappingSpecification) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification>)
	WithBlockDeviceMappings2(a []*ServicesEc2ModelInstanceBlockDeviceMappingSpecification) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public void setSourceDestCheck(java.lang.Boolean)
	SetSourceDestCheck(a bool) 

	// public java.lang.Boolean getSourceDestCheck()
	GetSourceDestCheck() bool

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withSourceDestCheck(java.lang.Boolean)
	WithSourceDestCheck(a bool) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public java.lang.Boolean isSourceDestCheck()
	IsSourceDestCheck() bool

	// public void setDisableApiTermination(java.lang.Boolean)
	SetDisableApiTermination(a bool) 

	// public java.lang.Boolean getDisableApiTermination()
	GetDisableApiTermination() bool

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withDisableApiTermination(java.lang.Boolean)
	WithDisableApiTermination(a bool) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public java.lang.Boolean isDisableApiTermination()
	IsDisableApiTermination() bool

	// public void setInstanceType(java.lang.String)
	SetInstanceType(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withInstanceType(java.lang.String)
	WithInstanceType(a string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public void setKernel(java.lang.String)
	SetKernel(a string) 

	// public java.lang.String getKernel()
	GetKernel() string

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withKernel(java.lang.String)
	WithKernel(a string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public void setRamdisk(java.lang.String)
	SetRamdisk(a string) 

	// public java.lang.String getRamdisk()
	GetRamdisk() string

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withRamdisk(java.lang.String)
	WithRamdisk(a string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public void setUserData(java.lang.String)
	SetUserData(a string) 

	// public java.lang.String getUserData()
	GetUserData() string

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withUserData(java.lang.String)
	WithUserData(a string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public void setInstanceInitiatedShutdownBehavior(java.lang.String)
	SetInstanceInitiatedShutdownBehavior(a string) 

	// public java.lang.String getInstanceInitiatedShutdownBehavior()
	GetInstanceInitiatedShutdownBehavior() string

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withInstanceInitiatedShutdownBehavior(java.lang.String)
	WithInstanceInitiatedShutdownBehavior(a string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public java.util.List<java.lang.String> getGroups()
	GetGroups() []string

	// public void setGroups(java.util.Collection<java.lang.String>)
	SetGroups(a []string) 

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withGroups(java.lang.String...)
	WithGroups(a ...string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withGroups(java.util.Collection<java.lang.String>)
	WithGroups2(a []string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public void setEbsOptimized(java.lang.Boolean)
	SetEbsOptimized(a bool) 

	// public java.lang.Boolean getEbsOptimized()
	GetEbsOptimized() bool

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withEbsOptimized(java.lang.Boolean)
	WithEbsOptimized(a bool) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public java.lang.Boolean isEbsOptimized()
	IsEbsOptimized() bool

	// public void setSriovNetSupport(java.lang.String)
	SetSriovNetSupport(a string) 

	// public java.lang.String getSriovNetSupport()
	GetSriovNetSupport() string

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withSriovNetSupport(java.lang.String)
	WithSriovNetSupport(a string) *ServicesEc2ModelModifyInstanceAttributeRequest

	// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest> getDryRunRequest()
	GetDryRunRequest() *Request

	// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest clone()
	Clone3() *ServicesEc2ModelModifyInstanceAttributeRequest
}

type ServicesEc2ModelModifyInstanceAttributeRequest struct {
	AmazonWebServiceRequest
}

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest()
func NewServicesEc2ModelModifyInstanceAttributeRequest() (*ServicesEc2ModelModifyInstanceAttributeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest(java.lang.String, java.lang.String)
func NewServicesEc2ModelModifyInstanceAttributeRequest3(a string, b string) (*ServicesEc2ModelModifyInstanceAttributeRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest(java.lang.String, com.amazonaws.services.ec2.model.InstanceAttributeName)
func NewServicesEc2ModelModifyInstanceAttributeRequest2(a string, b ServicesEc2ModelInstanceAttributeNameInterface) (*ServicesEc2ModelModifyInstanceAttributeRequest) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("com/amazonaws/services/ec2/model/InstanceAttributeName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithInstanceId(a string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetAttribute2(a string)  {
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
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetAttribute() string {
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withAttribute(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithAttribute2(a string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttribute(com.amazonaws.services.ec2.model.InstanceAttributeName)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetAttribute(a ServicesEc2ModelInstanceAttributeNameInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceAttributeName"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withAttribute(com.amazonaws.services.ec2.model.InstanceAttributeName)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithAttribute(a ServicesEc2ModelInstanceAttributeNameInterface) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttribute", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceAttributeName"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setValue(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetValue(a string)  {
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
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetValue() string {
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withValue(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithValue(a string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValue", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification> getBlockDeviceMappings()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetBlockDeviceMappings() []*ServicesEc2ModelInstanceBlockDeviceMappingSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockDeviceMappings", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceBlockDeviceMappingSpecification)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification>)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetBlockDeviceMappings(a []*ServicesEc2ModelInstanceBlockDeviceMappingSpecification)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBlockDeviceMappings", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withBlockDeviceMappings(com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification...)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithBlockDeviceMappings(a ...*ServicesEc2ModelInstanceBlockDeviceMappingSpecification) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceBlockDeviceMappingSpecification")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceBlockDeviceMappingSpecification"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceBlockDeviceMappingSpecification>)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithBlockDeviceMappings2(a []*ServicesEc2ModelInstanceBlockDeviceMappingSpecification) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetSourceDestCheck(a bool)  {
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
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetSourceDestCheck() bool {
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withSourceDestCheck(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithSourceDestCheck(a bool) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSourceDestCheck", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isSourceDestCheck()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) IsSourceDestCheck() bool {
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

// public void setDisableApiTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetDisableApiTermination(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDisableApiTermination", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getDisableApiTermination()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetDisableApiTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisableApiTermination", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withDisableApiTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithDisableApiTermination(a bool) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDisableApiTermination", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDisableApiTermination()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) IsDisableApiTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDisableApiTermination", "java/lang/Boolean")
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

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetInstanceType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceType()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetInstanceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithInstanceType(a string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKernel(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetKernel(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKernel", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKernel()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetKernel() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKernel", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withKernel(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithKernel(a string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKernel", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRamdisk(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetRamdisk(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRamdisk", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRamdisk()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetRamdisk() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRamdisk", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withRamdisk(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithRamdisk(a string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRamdisk", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserData(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetUserData(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserData", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getUserData()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetUserData() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserData", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withUserData(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithUserData(a string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserData", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceInitiatedShutdownBehavior(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetInstanceInitiatedShutdownBehavior(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceInitiatedShutdownBehavior", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceInitiatedShutdownBehavior()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetInstanceInitiatedShutdownBehavior() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceInitiatedShutdownBehavior", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withInstanceInitiatedShutdownBehavior(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithInstanceInitiatedShutdownBehavior(a string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceInitiatedShutdownBehavior", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getGroups()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetGroups() []string {
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
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetGroups(a []string)  {
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withGroups(java.lang.String...)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithGroups(a ...string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withGroups(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithGroups2(a []string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetEbsOptimized(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEbsOptimized", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEbsOptimized()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetEbsOptimized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEbsOptimized", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withEbsOptimized(java.lang.Boolean)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithEbsOptimized(a bool) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEbsOptimized", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEbsOptimized()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) IsEbsOptimized() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEbsOptimized", "java/lang/Boolean")
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

// public void setSriovNetSupport(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) SetSriovNetSupport(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSriovNetSupport", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSriovNetSupport()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetSriovNetSupport() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSriovNetSupport", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest withSriovNetSupport(java.lang.String)
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) WithSriovNetSupport(a string) *ServicesEc2ModelModifyInstanceAttributeRequest {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSriovNetSupport", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.Request<com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest> getDryRunRequest()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) GetDryRunRequest() *Request {
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
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ModifyInstanceAttributeRequest clone()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) Clone3() *ServicesEc2ModelModifyInstanceAttributeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ModifyInstanceAttributeRequest")
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
	unique_x := &ServicesEc2ModelModifyInstanceAttributeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.AmazonWebServiceRequest clone()
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) Clone() *AmazonWebServiceRequest {
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
func (jbobject *ServicesEc2ModelModifyInstanceAttributeRequest) Clone2() (*JavaLangObject, error) {
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


