package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImageInterface interface {
	JavaLangObjectInterface

	// public void setImageId(java.lang.String)
	SetImageId(a string) 

	// public java.lang.String getImageId()
	GetImageId() string

	// public com.amazonaws.services.ec2.model.Image withImageId(java.lang.String)
	WithImageId(a string) *ServicesEc2ModelImage

	// public void setImageLocation(java.lang.String)
	SetImageLocation(a string) 

	// public java.lang.String getImageLocation()
	GetImageLocation() string

	// public com.amazonaws.services.ec2.model.Image withImageLocation(java.lang.String)
	WithImageLocation(a string) *ServicesEc2ModelImage

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.Image withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelImage

	// public void setState(com.amazonaws.services.ec2.model.ImageState)
	SetState(a ServicesEc2ModelImageStateInterface) 

	// public com.amazonaws.services.ec2.model.Image withState(com.amazonaws.services.ec2.model.ImageState)
	WithState(a ServicesEc2ModelImageStateInterface) *ServicesEc2ModelImage

	// public void setOwnerId(java.lang.String)
	SetOwnerId(a string) 

	// public java.lang.String getOwnerId()
	GetOwnerId() string

	// public com.amazonaws.services.ec2.model.Image withOwnerId(java.lang.String)
	WithOwnerId(a string) *ServicesEc2ModelImage

	// public void setCreationDate(java.lang.String)
	SetCreationDate(a string) 

	// public java.lang.String getCreationDate()
	GetCreationDate() string

	// public com.amazonaws.services.ec2.model.Image withCreationDate(java.lang.String)
	WithCreationDate(a string) *ServicesEc2ModelImage

	// public void setPublic(java.lang.Boolean)
	SetPublic(a bool) 

	// public java.lang.Boolean getPublic()
	GetPublic() bool

	// public com.amazonaws.services.ec2.model.Image withPublic(java.lang.Boolean)
	WithPublic(a bool) *ServicesEc2ModelImage

	// public java.lang.Boolean isPublic()
	IsPublic() bool

	// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
	GetProductCodes() []*ServicesEc2ModelProductCode

	// public void setProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	SetProductCodes(a []*ServicesEc2ModelProductCode) 

	// public com.amazonaws.services.ec2.model.Image withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
	WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelImage

	// public com.amazonaws.services.ec2.model.Image withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelImage

	// public void setArchitecture(java.lang.String)
	SetArchitecture2(a string) 

	// public java.lang.String getArchitecture()
	GetArchitecture() string

	// public com.amazonaws.services.ec2.model.Image withArchitecture(java.lang.String)
	WithArchitecture2(a string) *ServicesEc2ModelImage

	// public void setArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
	SetArchitecture(a ServicesEc2ModelArchitectureValuesInterface) 

	// public com.amazonaws.services.ec2.model.Image withArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
	WithArchitecture(a ServicesEc2ModelArchitectureValuesInterface) *ServicesEc2ModelImage

	// public void setImageType(java.lang.String)
	SetImageType2(a string) 

	// public java.lang.String getImageType()
	GetImageType() string

	// public com.amazonaws.services.ec2.model.Image withImageType(java.lang.String)
	WithImageType2(a string) *ServicesEc2ModelImage

	// public void setImageType(com.amazonaws.services.ec2.model.ImageTypeValues)
	SetImageType(a ServicesEc2ModelImageTypeValuesInterface) 

	// public com.amazonaws.services.ec2.model.Image withImageType(com.amazonaws.services.ec2.model.ImageTypeValues)
	WithImageType(a ServicesEc2ModelImageTypeValuesInterface) *ServicesEc2ModelImage

	// public void setKernelId(java.lang.String)
	SetKernelId(a string) 

	// public java.lang.String getKernelId()
	GetKernelId() string

	// public com.amazonaws.services.ec2.model.Image withKernelId(java.lang.String)
	WithKernelId(a string) *ServicesEc2ModelImage

	// public void setRamdiskId(java.lang.String)
	SetRamdiskId(a string) 

	// public java.lang.String getRamdiskId()
	GetRamdiskId() string

	// public com.amazonaws.services.ec2.model.Image withRamdiskId(java.lang.String)
	WithRamdiskId(a string) *ServicesEc2ModelImage

	// public void setPlatform(java.lang.String)
	SetPlatform2(a string) 

	// public java.lang.String getPlatform()
	GetPlatform() string

	// public com.amazonaws.services.ec2.model.Image withPlatform(java.lang.String)
	WithPlatform2(a string) *ServicesEc2ModelImage

	// public void setPlatform(com.amazonaws.services.ec2.model.PlatformValues)
	SetPlatform(a ServicesEc2ModelPlatformValuesInterface) 

	// public com.amazonaws.services.ec2.model.Image withPlatform(com.amazonaws.services.ec2.model.PlatformValues)
	WithPlatform(a ServicesEc2ModelPlatformValuesInterface) *ServicesEc2ModelImage

	// public void setSriovNetSupport(java.lang.String)
	SetSriovNetSupport(a string) 

	// public java.lang.String getSriovNetSupport()
	GetSriovNetSupport() string

	// public com.amazonaws.services.ec2.model.Image withSriovNetSupport(java.lang.String)
	WithSriovNetSupport(a string) *ServicesEc2ModelImage

	// public void setStateReason(com.amazonaws.services.ec2.model.StateReason)
	SetStateReason(a ServicesEc2ModelStateReasonInterface) 

	// public com.amazonaws.services.ec2.model.StateReason getStateReason()
	GetStateReason() *ServicesEc2ModelStateReason

	// public com.amazonaws.services.ec2.model.Image withStateReason(com.amazonaws.services.ec2.model.StateReason)
	WithStateReason(a ServicesEc2ModelStateReasonInterface) *ServicesEc2ModelImage

	// public void setImageOwnerAlias(java.lang.String)
	SetImageOwnerAlias(a string) 

	// public java.lang.String getImageOwnerAlias()
	GetImageOwnerAlias() string

	// public com.amazonaws.services.ec2.model.Image withImageOwnerAlias(java.lang.String)
	WithImageOwnerAlias(a string) *ServicesEc2ModelImage

	// public void setName(java.lang.String)
	SetName(a string) 

	// public java.lang.String getName()
	GetName() string

	// public com.amazonaws.services.ec2.model.Image withName(java.lang.String)
	WithName(a string) *ServicesEc2ModelImage

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.Image withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImage

	// public void setRootDeviceType(java.lang.String)
	SetRootDeviceType2(a string) 

	// public java.lang.String getRootDeviceType()
	GetRootDeviceType() string

	// public com.amazonaws.services.ec2.model.Image withRootDeviceType(java.lang.String)
	WithRootDeviceType2(a string) *ServicesEc2ModelImage

	// public void setRootDeviceType(com.amazonaws.services.ec2.model.DeviceType)
	SetRootDeviceType(a ServicesEc2ModelDeviceTypeInterface) 

	// public com.amazonaws.services.ec2.model.Image withRootDeviceType(com.amazonaws.services.ec2.model.DeviceType)
	WithRootDeviceType(a ServicesEc2ModelDeviceTypeInterface) *ServicesEc2ModelImage

	// public void setRootDeviceName(java.lang.String)
	SetRootDeviceName(a string) 

	// public java.lang.String getRootDeviceName()
	GetRootDeviceName() string

	// public com.amazonaws.services.ec2.model.Image withRootDeviceName(java.lang.String)
	WithRootDeviceName(a string) *ServicesEc2ModelImage

	// public java.util.List<com.amazonaws.services.ec2.model.BlockDeviceMapping> getBlockDeviceMappings()
	GetBlockDeviceMappings() []*ServicesEc2ModelBlockDeviceMapping

	// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
	SetBlockDeviceMappings(a []*ServicesEc2ModelBlockDeviceMapping) 

	// public com.amazonaws.services.ec2.model.Image withBlockDeviceMappings(com.amazonaws.services.ec2.model.BlockDeviceMapping...)
	WithBlockDeviceMappings(a ...*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelImage

	// public com.amazonaws.services.ec2.model.Image withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
	WithBlockDeviceMappings2(a []*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelImage

	// public void setVirtualizationType(java.lang.String)
	SetVirtualizationType2(a string) 

	// public java.lang.String getVirtualizationType()
	GetVirtualizationType() string

	// public com.amazonaws.services.ec2.model.Image withVirtualizationType(java.lang.String)
	WithVirtualizationType2(a string) *ServicesEc2ModelImage

	// public void setVirtualizationType(com.amazonaws.services.ec2.model.VirtualizationType)
	SetVirtualizationType(a ServicesEc2ModelVirtualizationTypeInterface) 

	// public com.amazonaws.services.ec2.model.Image withVirtualizationType(com.amazonaws.services.ec2.model.VirtualizationType)
	WithVirtualizationType(a ServicesEc2ModelVirtualizationTypeInterface) *ServicesEc2ModelImage

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.Image withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelImage

	// public com.amazonaws.services.ec2.model.Image withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelImage

	// public void setHypervisor(java.lang.String)
	SetHypervisor2(a string) 

	// public java.lang.String getHypervisor()
	GetHypervisor() string

	// public com.amazonaws.services.ec2.model.Image withHypervisor(java.lang.String)
	WithHypervisor2(a string) *ServicesEc2ModelImage

	// public void setHypervisor(com.amazonaws.services.ec2.model.HypervisorType)
	SetHypervisor(a ServicesEc2ModelHypervisorTypeInterface) 

	// public com.amazonaws.services.ec2.model.Image withHypervisor(com.amazonaws.services.ec2.model.HypervisorType)
	WithHypervisor(a ServicesEc2ModelHypervisorTypeInterface) *ServicesEc2ModelImage

	// public com.amazonaws.services.ec2.model.Image clone()
	Clone() *ServicesEc2ModelImage
}

type ServicesEc2ModelImage struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.Image()
func NewServicesEc2ModelImage() (*ServicesEc2ModelImage) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Image")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImage{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImageId(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetImageId(a string)  {
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
func (jbobject *ServicesEc2ModelImage) GetImageId() string {
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

// public com.amazonaws.services.ec2.model.Image withImageId(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithImageId(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageId", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImageLocation(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetImageLocation(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImageLocation", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImageLocation()
func (jbobject *ServicesEc2ModelImage) GetImageLocation() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageLocation", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withImageLocation(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithImageLocation(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageLocation", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getState()
func (jbobject *ServicesEc2ModelImage) GetState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withState(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithState2(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.ImageState)
func (jbobject *ServicesEc2ModelImage) SetState(a ServicesEc2ModelImageStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImageState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Image withState(com.amazonaws.services.ec2.model.ImageState)
func (jbobject *ServicesEc2ModelImage) WithState(a ServicesEc2ModelImageStateInterface) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImageState"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetOwnerId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOwnerId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOwnerId()
func (jbobject *ServicesEc2ModelImage) GetOwnerId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOwnerId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithOwnerId(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOwnerId", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreationDate(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetCreationDate(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreationDate", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCreationDate()
func (jbobject *ServicesEc2ModelImage) GetCreationDate() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreationDate", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withCreationDate(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithCreationDate(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreationDate", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPublic(java.lang.Boolean)
func (jbobject *ServicesEc2ModelImage) SetPublic(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPublic", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getPublic()
func (jbobject *ServicesEc2ModelImage) GetPublic() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPublic", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.Image withPublic(java.lang.Boolean)
func (jbobject *ServicesEc2ModelImage) WithPublic(a bool) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPublic", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isPublic()
func (jbobject *ServicesEc2ModelImage) IsPublic() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isPublic", "java/lang/Boolean")
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

// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
func (jbobject *ServicesEc2ModelImage) GetProductCodes() []*ServicesEc2ModelProductCode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductCodes", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelProductCode)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
func (jbobject *ServicesEc2ModelImage) SetProductCodes(a []*ServicesEc2ModelProductCode)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductCodes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Image withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
func (jbobject *ServicesEc2ModelImage) WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ProductCode")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ProductCode"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Image withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
func (jbobject *ServicesEc2ModelImage) WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetArchitecture2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArchitecture", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getArchitecture()
func (jbobject *ServicesEc2ModelImage) GetArchitecture() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArchitecture", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithArchitecture2(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArchitecture", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
func (jbobject *ServicesEc2ModelImage) SetArchitecture(a ServicesEc2ModelArchitectureValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArchitecture", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ArchitectureValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Image withArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
func (jbobject *ServicesEc2ModelImage) WithArchitecture(a ServicesEc2ModelArchitectureValuesInterface) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArchitecture", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ArchitectureValues"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImageType(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetImageType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImageType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImageType()
func (jbobject *ServicesEc2ModelImage) GetImageType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withImageType(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithImageType2(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageType", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImageType(com.amazonaws.services.ec2.model.ImageTypeValues)
func (jbobject *ServicesEc2ModelImage) SetImageType(a ServicesEc2ModelImageTypeValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImageType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImageTypeValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Image withImageType(com.amazonaws.services.ec2.model.ImageTypeValues)
func (jbobject *ServicesEc2ModelImage) WithImageType(a ServicesEc2ModelImageTypeValuesInterface) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageType", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImageTypeValues"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetKernelId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKernelId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKernelId()
func (jbobject *ServicesEc2ModelImage) GetKernelId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKernelId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withKernelId(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithKernelId(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKernelId", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetRamdiskId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRamdiskId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRamdiskId()
func (jbobject *ServicesEc2ModelImage) GetRamdiskId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRamdiskId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withRamdiskId(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithRamdiskId(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRamdiskId", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetPlatform2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlatform", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPlatform()
func (jbobject *ServicesEc2ModelImage) GetPlatform() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPlatform", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithPlatform2(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(com.amazonaws.services.ec2.model.PlatformValues)
func (jbobject *ServicesEc2ModelImage) SetPlatform(a ServicesEc2ModelPlatformValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlatform", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlatformValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Image withPlatform(com.amazonaws.services.ec2.model.PlatformValues)
func (jbobject *ServicesEc2ModelImage) WithPlatform(a ServicesEc2ModelPlatformValuesInterface) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlatformValues"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSriovNetSupport(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetSriovNetSupport(a string)  {
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
func (jbobject *ServicesEc2ModelImage) GetSriovNetSupport() string {
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

// public com.amazonaws.services.ec2.model.Image withSriovNetSupport(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithSriovNetSupport(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSriovNetSupport", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStateReason(com.amazonaws.services.ec2.model.StateReason)
func (jbobject *ServicesEc2ModelImage) SetStateReason(a ServicesEc2ModelStateReasonInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStateReason", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/StateReason"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.StateReason getStateReason()
func (jbobject *ServicesEc2ModelImage) GetStateReason() *ServicesEc2ModelStateReason {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStateReason", "com/amazonaws/services/ec2/model/StateReason")
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
	unique_x := &ServicesEc2ModelStateReason{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Image withStateReason(com.amazonaws.services.ec2.model.StateReason)
func (jbobject *ServicesEc2ModelImage) WithStateReason(a ServicesEc2ModelStateReasonInterface) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStateReason", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/StateReason"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImageOwnerAlias(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetImageOwnerAlias(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImageOwnerAlias", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImageOwnerAlias()
func (jbobject *ServicesEc2ModelImage) GetImageOwnerAlias() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageOwnerAlias", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withImageOwnerAlias(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithImageOwnerAlias(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageOwnerAlias", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getName()
func (jbobject *ServicesEc2ModelImage) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withName(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithName(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelImage) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.Image withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithDescription(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRootDeviceType(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetRootDeviceType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRootDeviceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRootDeviceType()
func (jbobject *ServicesEc2ModelImage) GetRootDeviceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRootDeviceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withRootDeviceType(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithRootDeviceType2(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRootDeviceType", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRootDeviceType(com.amazonaws.services.ec2.model.DeviceType)
func (jbobject *ServicesEc2ModelImage) SetRootDeviceType(a ServicesEc2ModelDeviceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRootDeviceType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeviceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Image withRootDeviceType(com.amazonaws.services.ec2.model.DeviceType)
func (jbobject *ServicesEc2ModelImage) WithRootDeviceType(a ServicesEc2ModelDeviceTypeInterface) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRootDeviceType", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DeviceType"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRootDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetRootDeviceName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRootDeviceName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRootDeviceName()
func (jbobject *ServicesEc2ModelImage) GetRootDeviceName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRootDeviceName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withRootDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithRootDeviceName(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRootDeviceName", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.BlockDeviceMapping> getBlockDeviceMappings()
func (jbobject *ServicesEc2ModelImage) GetBlockDeviceMappings() []*ServicesEc2ModelBlockDeviceMapping {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockDeviceMappings", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelBlockDeviceMapping)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
func (jbobject *ServicesEc2ModelImage) SetBlockDeviceMappings(a []*ServicesEc2ModelBlockDeviceMapping)  {
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

// public com.amazonaws.services.ec2.model.Image withBlockDeviceMappings(com.amazonaws.services.ec2.model.BlockDeviceMapping...)
func (jbobject *ServicesEc2ModelImage) WithBlockDeviceMappings(a ...*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/BlockDeviceMapping")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/BlockDeviceMapping"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Image withBlockDeviceMappings(java.util.Collection<com.amazonaws.services.ec2.model.BlockDeviceMapping>)
func (jbobject *ServicesEc2ModelImage) WithBlockDeviceMappings2(a []*ServicesEc2ModelBlockDeviceMapping) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBlockDeviceMappings", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVirtualizationType(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetVirtualizationType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVirtualizationType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVirtualizationType()
func (jbobject *ServicesEc2ModelImage) GetVirtualizationType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVirtualizationType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withVirtualizationType(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithVirtualizationType2(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVirtualizationType", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVirtualizationType(com.amazonaws.services.ec2.model.VirtualizationType)
func (jbobject *ServicesEc2ModelImage) SetVirtualizationType(a ServicesEc2ModelVirtualizationTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVirtualizationType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VirtualizationType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Image withVirtualizationType(com.amazonaws.services.ec2.model.VirtualizationType)
func (jbobject *ServicesEc2ModelImage) WithVirtualizationType(a ServicesEc2ModelVirtualizationTypeInterface) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVirtualizationType", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VirtualizationType"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelImage) GetTags() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTag)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelImage) SetTags(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Image withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelImage) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Image withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelImage) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHypervisor(java.lang.String)
func (jbobject *ServicesEc2ModelImage) SetHypervisor2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHypervisor", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getHypervisor()
func (jbobject *ServicesEc2ModelImage) GetHypervisor() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHypervisor", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Image withHypervisor(java.lang.String)
func (jbobject *ServicesEc2ModelImage) WithHypervisor2(a string) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHypervisor", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHypervisor(com.amazonaws.services.ec2.model.HypervisorType)
func (jbobject *ServicesEc2ModelImage) SetHypervisor(a ServicesEc2ModelHypervisorTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHypervisor", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/HypervisorType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Image withHypervisor(com.amazonaws.services.ec2.model.HypervisorType)
func (jbobject *ServicesEc2ModelImage) WithHypervisor(a ServicesEc2ModelHypervisorTypeInterface) *ServicesEc2ModelImage {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHypervisor", "com/amazonaws/services/ec2/model/Image", conv_a.Value().Cast("com/amazonaws/services/ec2/model/HypervisorType"))
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImage) ToString() string {
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
func (jbobject *ServicesEc2ModelImage) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImage) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.Image clone()
func (jbobject *ServicesEc2ModelImage) Clone() *ServicesEc2ModelImage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/Image")
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
	unique_x := &ServicesEc2ModelImage{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelImage) Clone2() (*JavaLangObject, error) {
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


