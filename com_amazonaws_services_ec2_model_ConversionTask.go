package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelConversionTaskInterface interface {
	JavaLangObjectInterface

	// public void setConversionTaskId(java.lang.String)
	SetConversionTaskId(a string) 

	// public java.lang.String getConversionTaskId()
	GetConversionTaskId() string

	// public com.amazonaws.services.ec2.model.ConversionTask withConversionTaskId(java.lang.String)
	WithConversionTaskId(a string) *ServicesEc2ModelConversionTask

	// public void setExpirationTime(java.lang.String)
	SetExpirationTime(a string) 

	// public java.lang.String getExpirationTime()
	GetExpirationTime() string

	// public com.amazonaws.services.ec2.model.ConversionTask withExpirationTime(java.lang.String)
	WithExpirationTime(a string) *ServicesEc2ModelConversionTask

	// public void setImportInstance(com.amazonaws.services.ec2.model.ImportInstanceTaskDetails)
	SetImportInstance(a ServicesEc2ModelImportInstanceTaskDetailsInterface) 

	// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails getImportInstance()
	GetImportInstance() *ServicesEc2ModelImportInstanceTaskDetails

	// public com.amazonaws.services.ec2.model.ConversionTask withImportInstance(com.amazonaws.services.ec2.model.ImportInstanceTaskDetails)
	WithImportInstance(a ServicesEc2ModelImportInstanceTaskDetailsInterface) *ServicesEc2ModelConversionTask

	// public void setImportVolume(com.amazonaws.services.ec2.model.ImportVolumeTaskDetails)
	SetImportVolume(a ServicesEc2ModelImportVolumeTaskDetailsInterface) 

	// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails getImportVolume()
	GetImportVolume() *ServicesEc2ModelImportVolumeTaskDetails

	// public com.amazonaws.services.ec2.model.ConversionTask withImportVolume(com.amazonaws.services.ec2.model.ImportVolumeTaskDetails)
	WithImportVolume(a ServicesEc2ModelImportVolumeTaskDetailsInterface) *ServicesEc2ModelConversionTask

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.ConversionTask withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelConversionTask

	// public void setState(com.amazonaws.services.ec2.model.ConversionTaskState)
	SetState(a ServicesEc2ModelConversionTaskStateInterface) 

	// public com.amazonaws.services.ec2.model.ConversionTask withState(com.amazonaws.services.ec2.model.ConversionTaskState)
	WithState(a ServicesEc2ModelConversionTaskStateInterface) *ServicesEc2ModelConversionTask

	// public void setStatusMessage(java.lang.String)
	SetStatusMessage(a string) 

	// public java.lang.String getStatusMessage()
	GetStatusMessage() string

	// public com.amazonaws.services.ec2.model.ConversionTask withStatusMessage(java.lang.String)
	WithStatusMessage(a string) *ServicesEc2ModelConversionTask

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.ConversionTask withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelConversionTask

	// public com.amazonaws.services.ec2.model.ConversionTask withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelConversionTask

	// public com.amazonaws.services.ec2.model.ConversionTask clone()
	Clone() *ServicesEc2ModelConversionTask
}

type ServicesEc2ModelConversionTask struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ConversionTask()
func NewServicesEc2ModelConversionTask() (*ServicesEc2ModelConversionTask) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ConversionTask")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelConversionTask{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setConversionTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTask) SetConversionTaskId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setConversionTaskId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getConversionTaskId()
func (jbobject *ServicesEc2ModelConversionTask) GetConversionTaskId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getConversionTaskId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ConversionTask withConversionTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTask) WithConversionTaskId(a string) *ServicesEc2ModelConversionTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withConversionTaskId", "com/amazonaws/services/ec2/model/ConversionTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExpirationTime(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTask) SetExpirationTime(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpirationTime", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getExpirationTime()
func (jbobject *ServicesEc2ModelConversionTask) GetExpirationTime() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpirationTime", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ConversionTask withExpirationTime(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTask) WithExpirationTime(a string) *ServicesEc2ModelConversionTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withExpirationTime", "com/amazonaws/services/ec2/model/ConversionTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImportInstance(com.amazonaws.services.ec2.model.ImportInstanceTaskDetails)
func (jbobject *ServicesEc2ModelConversionTask) SetImportInstance(a ServicesEc2ModelImportInstanceTaskDetailsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImportInstance", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportInstanceTaskDetails"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportInstanceTaskDetails getImportInstance()
func (jbobject *ServicesEc2ModelConversionTask) GetImportInstance() *ServicesEc2ModelImportInstanceTaskDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImportInstance", "com/amazonaws/services/ec2/model/ImportInstanceTaskDetails")
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

// public com.amazonaws.services.ec2.model.ConversionTask withImportInstance(com.amazonaws.services.ec2.model.ImportInstanceTaskDetails)
func (jbobject *ServicesEc2ModelConversionTask) WithImportInstance(a ServicesEc2ModelImportInstanceTaskDetailsInterface) *ServicesEc2ModelConversionTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportInstance", "com/amazonaws/services/ec2/model/ConversionTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportInstanceTaskDetails"))
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImportVolume(com.amazonaws.services.ec2.model.ImportVolumeTaskDetails)
func (jbobject *ServicesEc2ModelConversionTask) SetImportVolume(a ServicesEc2ModelImportVolumeTaskDetailsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImportVolume", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportVolumeTaskDetails"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails getImportVolume()
func (jbobject *ServicesEc2ModelConversionTask) GetImportVolume() *ServicesEc2ModelImportVolumeTaskDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImportVolume", "com/amazonaws/services/ec2/model/ImportVolumeTaskDetails")
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
	unique_x := &ServicesEc2ModelImportVolumeTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ConversionTask withImportVolume(com.amazonaws.services.ec2.model.ImportVolumeTaskDetails)
func (jbobject *ServicesEc2ModelConversionTask) WithImportVolume(a ServicesEc2ModelImportVolumeTaskDetailsInterface) *ServicesEc2ModelConversionTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportVolume", "com/amazonaws/services/ec2/model/ConversionTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImportVolumeTaskDetails"))
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTask) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelConversionTask) GetState() string {
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

// public com.amazonaws.services.ec2.model.ConversionTask withState(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTask) WithState2(a string) *ServicesEc2ModelConversionTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/ConversionTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.ConversionTaskState)
func (jbobject *ServicesEc2ModelConversionTask) SetState(a ServicesEc2ModelConversionTaskStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ConversionTaskState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ConversionTask withState(com.amazonaws.services.ec2.model.ConversionTaskState)
func (jbobject *ServicesEc2ModelConversionTask) WithState(a ServicesEc2ModelConversionTaskStateInterface) *ServicesEc2ModelConversionTask {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/ConversionTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ConversionTaskState"))
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTask) SetStatusMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatusMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatusMessage()
func (jbobject *ServicesEc2ModelConversionTask) GetStatusMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatusMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ConversionTask withStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTask) WithStatusMessage(a string) *ServicesEc2ModelConversionTask {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatusMessage", "com/amazonaws/services/ec2/model/ConversionTask", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelConversionTask) GetTags() []*ServicesEc2ModelTag {
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
func (jbobject *ServicesEc2ModelConversionTask) SetTags(a []*ServicesEc2ModelTag)  {
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

// public com.amazonaws.services.ec2.model.ConversionTask withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelConversionTask) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelConversionTask {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/ConversionTask", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ConversionTask withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelConversionTask) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelConversionTask {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/ConversionTask", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelConversionTask) ToString() string {
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
func (jbobject *ServicesEc2ModelConversionTask) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelConversionTask) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ConversionTask clone()
func (jbobject *ServicesEc2ModelConversionTask) Clone() *ServicesEc2ModelConversionTask {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ConversionTask")
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
	unique_x := &ServicesEc2ModelConversionTask{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelConversionTask) Clone2() (*JavaLangObject, error) {
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


