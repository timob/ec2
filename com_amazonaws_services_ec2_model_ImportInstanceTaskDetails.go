package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportInstanceTaskDetailsInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem> getVolumes()
	GetVolumes() []*ServicesEc2ModelImportInstanceVolumeDetailItem

	// public void setVolumes(java.util.Collection<com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem>)
	SetVolumes(a []*ServicesEc2ModelImportInstanceVolumeDetailItem) 

	// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withVolumes(com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem...)
	WithVolumes(a ...*ServicesEc2ModelImportInstanceVolumeDetailItem) *ServicesEc2ModelImportInstanceTaskDetails

	// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withVolumes(java.util.Collection<com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem>)
	WithVolumes2(a []*ServicesEc2ModelImportInstanceVolumeDetailItem) *ServicesEc2ModelImportInstanceTaskDetails

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelImportInstanceTaskDetails

	// public void setPlatform(java.lang.String)
	SetPlatform2(a string) 

	// public java.lang.String getPlatform()
	GetPlatform() string

	// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withPlatform(java.lang.String)
	WithPlatform2(a string) *ServicesEc2ModelImportInstanceTaskDetails

	// public void setPlatform(com.amazonaws.services.ec2.model.PlatformValues)
	SetPlatform(a ServicesEc2ModelPlatformValuesInterface) 

	// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withPlatform(com.amazonaws.services.ec2.model.PlatformValues)
	WithPlatform(a ServicesEc2ModelPlatformValuesInterface) *ServicesEc2ModelImportInstanceTaskDetails

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImportInstanceTaskDetails

	// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails clone()
	Clone() *ServicesEc2ModelImportInstanceTaskDetails
}

type ServicesEc2ModelImportInstanceTaskDetails struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails()
func NewServicesEc2ModelImportInstanceTaskDetails() (*ServicesEc2ModelImportInstanceTaskDetails) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportInstanceTaskDetails")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportInstanceTaskDetails{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem> getVolumes()
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) GetVolumes() []*ServicesEc2ModelImportInstanceVolumeDetailItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumes", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelImportInstanceVolumeDetailItem)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVolumes(java.util.Collection<com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem>)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) SetVolumes(a []*ServicesEc2ModelImportInstanceVolumeDetailItem)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withVolumes(com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem...)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) WithVolumes(a ...*ServicesEc2ModelImportInstanceVolumeDetailItem) *ServicesEc2ModelImportInstanceTaskDetails {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumes", "com/amazonaws/services/ec2/model/ImportInstanceTaskDetails", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem"))
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
	unique_x := &ServicesEc2ModelImportInstanceTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withVolumes(java.util.Collection<com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem>)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) WithVolumes2(a []*ServicesEc2ModelImportInstanceVolumeDetailItem) *ServicesEc2ModelImportInstanceTaskDetails {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumes", "com/amazonaws/services/ec2/model/ImportInstanceTaskDetails", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelImportInstanceTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) SetInstanceId(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) GetInstanceId() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) WithInstanceId(a string) *ServicesEc2ModelImportInstanceTaskDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/ImportInstanceTaskDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) SetPlatform2(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) GetPlatform() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) WithPlatform2(a string) *ServicesEc2ModelImportInstanceTaskDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/ImportInstanceTaskDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(com.amazonaws.services.ec2.model.PlatformValues)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) SetPlatform(a ServicesEc2ModelPlatformValuesInterface)  {
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

// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withPlatform(com.amazonaws.services.ec2.model.PlatformValues)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) WithPlatform(a ServicesEc2ModelPlatformValuesInterface) *ServicesEc2ModelImportInstanceTaskDetails {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/ImportInstanceTaskDetails", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PlatformValues"))
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
	unique_x := &ServicesEc2ModelImportInstanceTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) WithDescription(a string) *ServicesEc2ModelImportInstanceTaskDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ImportInstanceTaskDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) ToString() string {
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
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails clone()
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) Clone() *ServicesEc2ModelImportInstanceTaskDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportInstanceTaskDetails")
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
	unique_x := &ServicesEc2ModelImportInstanceTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelImportInstanceTaskDetails) Clone2() (*JavaLangObject, error) {
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


