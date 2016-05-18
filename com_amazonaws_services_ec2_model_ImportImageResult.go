package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportImageResultInterface interface {
	JavaLangObjectInterface

	// public void setImportTaskId(java.lang.String)
	SetImportTaskId(a string) 

	// public java.lang.String getImportTaskId()
	GetImportTaskId() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withImportTaskId(java.lang.String)
	WithImportTaskId(a string) *ServicesEc2ModelImportImageResult

	// public void setArchitecture(java.lang.String)
	SetArchitecture(a string) 

	// public java.lang.String getArchitecture()
	GetArchitecture() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withArchitecture(java.lang.String)
	WithArchitecture(a string) *ServicesEc2ModelImportImageResult

	// public void setLicenseType(java.lang.String)
	SetLicenseType(a string) 

	// public java.lang.String getLicenseType()
	GetLicenseType() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withLicenseType(java.lang.String)
	WithLicenseType(a string) *ServicesEc2ModelImportImageResult

	// public void setPlatform(java.lang.String)
	SetPlatform(a string) 

	// public java.lang.String getPlatform()
	GetPlatform() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withPlatform(java.lang.String)
	WithPlatform(a string) *ServicesEc2ModelImportImageResult

	// public void setHypervisor(java.lang.String)
	SetHypervisor(a string) 

	// public java.lang.String getHypervisor()
	GetHypervisor() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withHypervisor(java.lang.String)
	WithHypervisor(a string) *ServicesEc2ModelImportImageResult

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImportImageResult

	// public java.util.List<com.amazonaws.services.ec2.model.SnapshotDetail> getSnapshotDetails()
	GetSnapshotDetails() []*ServicesEc2ModelSnapshotDetail

	// public void setSnapshotDetails(java.util.Collection<com.amazonaws.services.ec2.model.SnapshotDetail>)
	SetSnapshotDetails(a []*ServicesEc2ModelSnapshotDetail) 

	// public com.amazonaws.services.ec2.model.ImportImageResult withSnapshotDetails(com.amazonaws.services.ec2.model.SnapshotDetail...)
	WithSnapshotDetails(a ...*ServicesEc2ModelSnapshotDetail) *ServicesEc2ModelImportImageResult

	// public com.amazonaws.services.ec2.model.ImportImageResult withSnapshotDetails(java.util.Collection<com.amazonaws.services.ec2.model.SnapshotDetail>)
	WithSnapshotDetails2(a []*ServicesEc2ModelSnapshotDetail) *ServicesEc2ModelImportImageResult

	// public void setImageId(java.lang.String)
	SetImageId(a string) 

	// public java.lang.String getImageId()
	GetImageId() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withImageId(java.lang.String)
	WithImageId(a string) *ServicesEc2ModelImportImageResult

	// public void setProgress(java.lang.String)
	SetProgress(a string) 

	// public java.lang.String getProgress()
	GetProgress() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withProgress(java.lang.String)
	WithProgress(a string) *ServicesEc2ModelImportImageResult

	// public void setStatusMessage(java.lang.String)
	SetStatusMessage(a string) 

	// public java.lang.String getStatusMessage()
	GetStatusMessage() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withStatusMessage(java.lang.String)
	WithStatusMessage(a string) *ServicesEc2ModelImportImageResult

	// public void setStatus(java.lang.String)
	SetStatus(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.ImportImageResult withStatus(java.lang.String)
	WithStatus(a string) *ServicesEc2ModelImportImageResult

	// public com.amazonaws.services.ec2.model.ImportImageResult clone()
	Clone() *ServicesEc2ModelImportImageResult
}

type ServicesEc2ModelImportImageResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ImportImageResult()
func NewServicesEc2ModelImportImageResult() (*ServicesEc2ModelImportImageResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportImageResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportImageResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetImportTaskId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImportTaskId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImportTaskId()
func (jbobject *ServicesEc2ModelImportImageResult) GetImportTaskId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImportTaskId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportImageResult withImportTaskId(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithImportTaskId(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportTaskId", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetArchitecture(a string)  {
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
func (jbobject *ServicesEc2ModelImportImageResult) GetArchitecture() string {
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

// public com.amazonaws.services.ec2.model.ImportImageResult withArchitecture(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithArchitecture(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArchitecture", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLicenseType(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetLicenseType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLicenseType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getLicenseType()
func (jbobject *ServicesEc2ModelImportImageResult) GetLicenseType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLicenseType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportImageResult withLicenseType(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithLicenseType(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLicenseType", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetPlatform(a string)  {
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
func (jbobject *ServicesEc2ModelImportImageResult) GetPlatform() string {
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

// public com.amazonaws.services.ec2.model.ImportImageResult withPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithPlatform(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHypervisor(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetHypervisor(a string)  {
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
func (jbobject *ServicesEc2ModelImportImageResult) GetHypervisor() string {
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

// public com.amazonaws.services.ec2.model.ImportImageResult withHypervisor(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithHypervisor(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHypervisor", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelImportImageResult) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ImportImageResult withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithDescription(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.SnapshotDetail> getSnapshotDetails()
func (jbobject *ServicesEc2ModelImportImageResult) GetSnapshotDetails() []*ServicesEc2ModelSnapshotDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshotDetails", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelSnapshotDetail)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSnapshotDetails(java.util.Collection<com.amazonaws.services.ec2.model.SnapshotDetail>)
func (jbobject *ServicesEc2ModelImportImageResult) SetSnapshotDetails(a []*ServicesEc2ModelSnapshotDetail)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshotDetails", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImportImageResult withSnapshotDetails(com.amazonaws.services.ec2.model.SnapshotDetail...)
func (jbobject *ServicesEc2ModelImportImageResult) WithSnapshotDetails(a ...*ServicesEc2ModelSnapshotDetail) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/SnapshotDetail")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotDetails", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/SnapshotDetail"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportImageResult withSnapshotDetails(java.util.Collection<com.amazonaws.services.ec2.model.SnapshotDetail>)
func (jbobject *ServicesEc2ModelImportImageResult) WithSnapshotDetails2(a []*ServicesEc2ModelSnapshotDetail) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotDetails", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImageId(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetImageId(a string)  {
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
func (jbobject *ServicesEc2ModelImportImageResult) GetImageId() string {
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

// public com.amazonaws.services.ec2.model.ImportImageResult withImageId(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithImageId(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageId", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProgress(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetProgress(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProgress", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getProgress()
func (jbobject *ServicesEc2ModelImportImageResult) GetProgress() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProgress", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportImageResult withProgress(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithProgress(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProgress", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetStatusMessage(a string)  {
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
func (jbobject *ServicesEc2ModelImportImageResult) GetStatusMessage() string {
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

// public com.amazonaws.services.ec2.model.ImportImageResult withStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithStatusMessage(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatusMessage", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) SetStatus(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getStatus()
func (jbobject *ServicesEc2ModelImportImageResult) GetStatus() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStatus", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportImageResult withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelImportImageResult) WithStatus(a string) *ServicesEc2ModelImportImageResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/ImportImageResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImportImageResult) ToString() string {
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
func (jbobject *ServicesEc2ModelImportImageResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportImageResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportImageResult clone()
func (jbobject *ServicesEc2ModelImportImageResult) Clone() *ServicesEc2ModelImportImageResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportImageResult")
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
	unique_x := &ServicesEc2ModelImportImageResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelImportImageResult) Clone2() (*JavaLangObject, error) {
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


