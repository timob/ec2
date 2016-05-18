package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportInstanceLaunchSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setArchitecture(java.lang.String)
	SetArchitecture2(a string) 

	// public java.lang.String getArchitecture()
	GetArchitecture() string

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withArchitecture(java.lang.String)
	WithArchitecture2(a string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
	SetArchitecture(a ServicesEc2ModelArchitectureValuesInterface) 

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
	WithArchitecture(a ServicesEc2ModelArchitectureValuesInterface) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public java.util.List<java.lang.String> getGroupNames()
	GetGroupNames() []string

	// public void setGroupNames(java.util.Collection<java.lang.String>)
	SetGroupNames(a []string) 

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withGroupNames(java.lang.String...)
	WithGroupNames(a ...string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withGroupNames(java.util.Collection<java.lang.String>)
	WithGroupNames2(a []string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public java.util.List<java.lang.String> getGroupIds()
	GetGroupIds() []string

	// public void setGroupIds(java.util.Collection<java.lang.String>)
	SetGroupIds(a []string) 

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withGroupIds(java.lang.String...)
	WithGroupIds(a ...string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withGroupIds(java.util.Collection<java.lang.String>)
	WithGroupIds2(a []string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setAdditionalInfo(java.lang.String)
	SetAdditionalInfo(a string) 

	// public java.lang.String getAdditionalInfo()
	GetAdditionalInfo() string

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withAdditionalInfo(java.lang.String)
	WithAdditionalInfo(a string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setUserData(com.amazonaws.services.ec2.model.UserData)
	SetUserData(a ServicesEc2ModelUserDataInterface) 

	// public com.amazonaws.services.ec2.model.UserData getUserData()
	GetUserData() *ServicesEc2ModelUserData

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withUserData(com.amazonaws.services.ec2.model.UserData)
	WithUserData(a ServicesEc2ModelUserDataInterface) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setInstanceType(java.lang.String)
	SetInstanceType2(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withInstanceType(java.lang.String)
	WithInstanceType2(a string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	SetInstanceType(a ServicesEc2ModelInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setPlacement(com.amazonaws.services.ec2.model.Placement)
	SetPlacement(a ServicesEc2ModelPlacementInterface) 

	// public com.amazonaws.services.ec2.model.Placement getPlacement()
	GetPlacement() *ServicesEc2ModelPlacement

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withPlacement(com.amazonaws.services.ec2.model.Placement)
	WithPlacement(a ServicesEc2ModelPlacementInterface) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setMonitoring(java.lang.Boolean)
	SetMonitoring(a bool) 

	// public java.lang.Boolean getMonitoring()
	GetMonitoring() bool

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withMonitoring(java.lang.Boolean)
	WithMonitoring(a bool) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public java.lang.Boolean isMonitoring()
	IsMonitoring() bool

	// public void setSubnetId(java.lang.String)
	SetSubnetId(a string) 

	// public java.lang.String getSubnetId()
	GetSubnetId() string

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withSubnetId(java.lang.String)
	WithSubnetId(a string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setInstanceInitiatedShutdownBehavior(java.lang.String)
	SetInstanceInitiatedShutdownBehavior2(a string) 

	// public java.lang.String getInstanceInitiatedShutdownBehavior()
	GetInstanceInitiatedShutdownBehavior() string

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withInstanceInitiatedShutdownBehavior(java.lang.String)
	WithInstanceInitiatedShutdownBehavior2(a string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setInstanceInitiatedShutdownBehavior(com.amazonaws.services.ec2.model.ShutdownBehavior)
	SetInstanceInitiatedShutdownBehavior(a ServicesEc2ModelShutdownBehaviorInterface) 

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withInstanceInitiatedShutdownBehavior(com.amazonaws.services.ec2.model.ShutdownBehavior)
	WithInstanceInitiatedShutdownBehavior(a ServicesEc2ModelShutdownBehaviorInterface) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelImportInstanceLaunchSpecification

	// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification clone()
	Clone() *ServicesEc2ModelImportInstanceLaunchSpecification
}

type ServicesEc2ModelImportInstanceLaunchSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification()
func NewServicesEc2ModelImportInstanceLaunchSpecification() (*ServicesEc2ModelImportInstanceLaunchSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetArchitecture2(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetArchitecture() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithArchitecture2(a string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArchitecture", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetArchitecture(a ServicesEc2ModelArchitectureValuesInterface)  {
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

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withArchitecture(com.amazonaws.services.ec2.model.ArchitectureValues)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithArchitecture(a ServicesEc2ModelArchitectureValuesInterface) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArchitecture", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ArchitectureValues"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getGroupNames()
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetGroupNames() []string {
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
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetGroupNames(a []string)  {
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

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withGroupNames(java.lang.String...)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithGroupNames(a ...string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupNames", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withGroupNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithGroupNames2(a []string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupNames", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getGroupIds()
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetGroupIds() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupIds", "java/util/List")
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

// public void setGroupIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetGroupIds(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withGroupIds(java.lang.String...)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithGroupIds(a ...string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupIds", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withGroupIds(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithGroupIds2(a []string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupIds", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAdditionalInfo(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetAdditionalInfo(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAdditionalInfo", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAdditionalInfo()
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetAdditionalInfo() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAdditionalInfo", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withAdditionalInfo(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithAdditionalInfo(a string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAdditionalInfo", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserData(com.amazonaws.services.ec2.model.UserData)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetUserData(a ServicesEc2ModelUserDataInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserData", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/UserData"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.UserData getUserData()
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetUserData() *ServicesEc2ModelUserData {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserData", "com/amazonaws/services/ec2/model/UserData")
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
	unique_x := &ServicesEc2ModelUserData{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withUserData(com.amazonaws.services.ec2.model.UserData)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithUserData(a ServicesEc2ModelUserDataInterface) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserData", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/UserData"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetInstanceType2(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetInstanceType() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithInstanceType2(a string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetInstanceType(a ServicesEc2ModelInstanceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlacement(com.amazonaws.services.ec2.model.Placement)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetPlacement(a ServicesEc2ModelPlacementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlacement", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Placement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Placement getPlacement()
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetPlacement() *ServicesEc2ModelPlacement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPlacement", "com/amazonaws/services/ec2/model/Placement")
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
	unique_x := &ServicesEc2ModelPlacement{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withPlacement(com.amazonaws.services.ec2.model.Placement)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithPlacement(a ServicesEc2ModelPlacementInterface) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlacement", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Placement"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMonitoring(java.lang.Boolean)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetMonitoring(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMonitoring", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getMonitoring()
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetMonitoring() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonitoring", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withMonitoring(java.lang.Boolean)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithMonitoring(a bool) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMonitoring", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isMonitoring()
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) IsMonitoring() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isMonitoring", "java/lang/Boolean")
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

// public void setSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetSubnetId(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetSubnetId() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withSubnetId(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithSubnetId(a string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSubnetId", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceInitiatedShutdownBehavior(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetInstanceInitiatedShutdownBehavior2(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetInstanceInitiatedShutdownBehavior() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withInstanceInitiatedShutdownBehavior(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithInstanceInitiatedShutdownBehavior2(a string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceInitiatedShutdownBehavior", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceInitiatedShutdownBehavior(com.amazonaws.services.ec2.model.ShutdownBehavior)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetInstanceInitiatedShutdownBehavior(a ServicesEc2ModelShutdownBehaviorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceInitiatedShutdownBehavior", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ShutdownBehavior"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withInstanceInitiatedShutdownBehavior(com.amazonaws.services.ec2.model.ShutdownBehavior)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithInstanceInitiatedShutdownBehavior(a ServicesEc2ModelShutdownBehaviorInterface) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceInitiatedShutdownBehavior", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ShutdownBehavior"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) SetPrivateIpAddress(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) GetPrivateIpAddress() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) WithPrivateIpAddress(a string) *ServicesEc2ModelImportInstanceLaunchSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) ToString() string {
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
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportInstanceLaunchSpecification clone()
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) Clone() *ServicesEc2ModelImportInstanceLaunchSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportInstanceLaunchSpecification")
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
	unique_x := &ServicesEc2ModelImportInstanceLaunchSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelImportInstanceLaunchSpecification) Clone2() (*JavaLangObject, error) {
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


