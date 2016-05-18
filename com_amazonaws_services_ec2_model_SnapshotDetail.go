package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSnapshotDetailInterface interface {
	JavaLangObjectInterface

	// public void setDiskImageSize(java.lang.Double)
	SetDiskImageSize(a float64) 

	// public java.lang.Double getDiskImageSize()
	GetDiskImageSize() float64

	// public com.amazonaws.services.ec2.model.SnapshotDetail withDiskImageSize(java.lang.Double)
	WithDiskImageSize(a float64) *ServicesEc2ModelSnapshotDetail

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.SnapshotDetail withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelSnapshotDetail

	// public void setFormat(java.lang.String)
	SetFormat(a string) 

	// public java.lang.String getFormat()
	GetFormat() string

	// public com.amazonaws.services.ec2.model.SnapshotDetail withFormat(java.lang.String)
	WithFormat(a string) *ServicesEc2ModelSnapshotDetail

	// public void setUrl(java.lang.String)
	SetUrl(a string) 

	// public java.lang.String getUrl()
	GetUrl() string

	// public com.amazonaws.services.ec2.model.SnapshotDetail withUrl(java.lang.String)
	WithUrl(a string) *ServicesEc2ModelSnapshotDetail

	// public void setUserBucket(com.amazonaws.services.ec2.model.UserBucketDetails)
	SetUserBucket(a ServicesEc2ModelUserBucketDetailsInterface) 

	// public com.amazonaws.services.ec2.model.UserBucketDetails getUserBucket()
	GetUserBucket() *ServicesEc2ModelUserBucketDetails

	// public com.amazonaws.services.ec2.model.SnapshotDetail withUserBucket(com.amazonaws.services.ec2.model.UserBucketDetails)
	WithUserBucket(a ServicesEc2ModelUserBucketDetailsInterface) *ServicesEc2ModelSnapshotDetail

	// public void setDeviceName(java.lang.String)
	SetDeviceName(a string) 

	// public java.lang.String getDeviceName()
	GetDeviceName() string

	// public com.amazonaws.services.ec2.model.SnapshotDetail withDeviceName(java.lang.String)
	WithDeviceName(a string) *ServicesEc2ModelSnapshotDetail

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.SnapshotDetail withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelSnapshotDetail

	// public void setProgress(java.lang.String)
	SetProgress(a string) 

	// public java.lang.String getProgress()
	GetProgress() string

	// public com.amazonaws.services.ec2.model.SnapshotDetail withProgress(java.lang.String)
	WithProgress(a string) *ServicesEc2ModelSnapshotDetail

	// public void setStatusMessage(java.lang.String)
	SetStatusMessage(a string) 

	// public java.lang.String getStatusMessage()
	GetStatusMessage() string

	// public com.amazonaws.services.ec2.model.SnapshotDetail withStatusMessage(java.lang.String)
	WithStatusMessage(a string) *ServicesEc2ModelSnapshotDetail

	// public void setStatus(java.lang.String)
	SetStatus(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.SnapshotDetail withStatus(java.lang.String)
	WithStatus(a string) *ServicesEc2ModelSnapshotDetail

	// public com.amazonaws.services.ec2.model.SnapshotDetail clone()
	Clone() *ServicesEc2ModelSnapshotDetail
}

type ServicesEc2ModelSnapshotDetail struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SnapshotDetail()
func NewServicesEc2ModelSnapshotDetail() (*ServicesEc2ModelSnapshotDetail) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SnapshotDetail")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSnapshotDetail{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDiskImageSize(java.lang.Double)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetDiskImageSize(a float64)  {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDiskImageSize", javabind.Void, conv_a.Value().Cast("java/lang/Double"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Double getDiskImageSize()
func (jbobject *ServicesEc2ModelSnapshotDetail) GetDiskImageSize() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDiskImageSize", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.SnapshotDetail withDiskImageSize(java.lang.Double)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithDiskImageSize(a float64) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDiskImageSize", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("java/lang/Double"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotDetail) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.SnapshotDetail withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithDescription(a string) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFormat(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetFormat(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFormat", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getFormat()
func (jbobject *ServicesEc2ModelSnapshotDetail) GetFormat() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFormat", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SnapshotDetail withFormat(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithFormat(a string) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFormat", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUrl(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetUrl(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUrl", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getUrl()
func (jbobject *ServicesEc2ModelSnapshotDetail) GetUrl() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUrl", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SnapshotDetail withUrl(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithUrl(a string) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUrl", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserBucket(com.amazonaws.services.ec2.model.UserBucketDetails)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetUserBucket(a ServicesEc2ModelUserBucketDetailsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserBucket", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/UserBucketDetails"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.UserBucketDetails getUserBucket()
func (jbobject *ServicesEc2ModelSnapshotDetail) GetUserBucket() *ServicesEc2ModelUserBucketDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserBucket", "com/amazonaws/services/ec2/model/UserBucketDetails")
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
	unique_x := &ServicesEc2ModelUserBucketDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SnapshotDetail withUserBucket(com.amazonaws.services.ec2.model.UserBucketDetails)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithUserBucket(a ServicesEc2ModelUserBucketDetailsInterface) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserBucket", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("com/amazonaws/services/ec2/model/UserBucketDetails"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetDeviceName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeviceName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDeviceName()
func (jbobject *ServicesEc2ModelSnapshotDetail) GetDeviceName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeviceName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SnapshotDetail withDeviceName(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithDeviceName(a string) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeviceName", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetSnapshotId(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotDetail) GetSnapshotId() string {
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

// public com.amazonaws.services.ec2.model.SnapshotDetail withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithSnapshotId(a string) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProgress(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetProgress(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotDetail) GetProgress() string {
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

// public com.amazonaws.services.ec2.model.SnapshotDetail withProgress(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithProgress(a string) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProgress", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetStatusMessage(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotDetail) GetStatusMessage() string {
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

// public com.amazonaws.services.ec2.model.SnapshotDetail withStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithStatusMessage(a string) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatusMessage", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) SetStatus(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotDetail) GetStatus() string {
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

// public com.amazonaws.services.ec2.model.SnapshotDetail withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDetail) WithStatus(a string) *ServicesEc2ModelSnapshotDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/SnapshotDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSnapshotDetail) ToString() string {
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
func (jbobject *ServicesEc2ModelSnapshotDetail) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSnapshotDetail) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SnapshotDetail clone()
func (jbobject *ServicesEc2ModelSnapshotDetail) Clone() *ServicesEc2ModelSnapshotDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SnapshotDetail")
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
	unique_x := &ServicesEc2ModelSnapshotDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSnapshotDetail) Clone2() (*JavaLangObject, error) {
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


