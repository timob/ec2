package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSnapshotTaskDetailInterface interface {
	JavaLangObjectInterface

	// public void setDiskImageSize(java.lang.Double)
	SetDiskImageSize(a float64) 

	// public java.lang.Double getDiskImageSize()
	GetDiskImageSize() float64

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withDiskImageSize(java.lang.Double)
	WithDiskImageSize(a float64) *ServicesEc2ModelSnapshotTaskDetail

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelSnapshotTaskDetail

	// public void setFormat(java.lang.String)
	SetFormat(a string) 

	// public java.lang.String getFormat()
	GetFormat() string

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withFormat(java.lang.String)
	WithFormat(a string) *ServicesEc2ModelSnapshotTaskDetail

	// public void setUrl(java.lang.String)
	SetUrl(a string) 

	// public java.lang.String getUrl()
	GetUrl() string

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withUrl(java.lang.String)
	WithUrl(a string) *ServicesEc2ModelSnapshotTaskDetail

	// public void setUserBucket(com.amazonaws.services.ec2.model.UserBucketDetails)
	SetUserBucket(a ServicesEc2ModelUserBucketDetailsInterface) 

	// public com.amazonaws.services.ec2.model.UserBucketDetails getUserBucket()
	GetUserBucket() *ServicesEc2ModelUserBucketDetails

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withUserBucket(com.amazonaws.services.ec2.model.UserBucketDetails)
	WithUserBucket(a ServicesEc2ModelUserBucketDetailsInterface) *ServicesEc2ModelSnapshotTaskDetail

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelSnapshotTaskDetail

	// public void setProgress(java.lang.String)
	SetProgress(a string) 

	// public java.lang.String getProgress()
	GetProgress() string

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withProgress(java.lang.String)
	WithProgress(a string) *ServicesEc2ModelSnapshotTaskDetail

	// public void setStatusMessage(java.lang.String)
	SetStatusMessage(a string) 

	// public java.lang.String getStatusMessage()
	GetStatusMessage() string

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withStatusMessage(java.lang.String)
	WithStatusMessage(a string) *ServicesEc2ModelSnapshotTaskDetail

	// public void setStatus(java.lang.String)
	SetStatus(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withStatus(java.lang.String)
	WithStatus(a string) *ServicesEc2ModelSnapshotTaskDetail

	// public com.amazonaws.services.ec2.model.SnapshotTaskDetail clone()
	Clone() *ServicesEc2ModelSnapshotTaskDetail
}

type ServicesEc2ModelSnapshotTaskDetail struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail()
func NewServicesEc2ModelSnapshotTaskDetail() (*ServicesEc2ModelSnapshotTaskDetail) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SnapshotTaskDetail")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSnapshotTaskDetail{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDiskImageSize(java.lang.Double)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) SetDiskImageSize(a float64)  {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) GetDiskImageSize() float64 {
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

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withDiskImageSize(java.lang.Double)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) WithDiskImageSize(a float64) *ServicesEc2ModelSnapshotTaskDetail {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDiskImageSize", "com/amazonaws/services/ec2/model/SnapshotTaskDetail", conv_a.Value().Cast("java/lang/Double"))
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) WithDescription(a string) *ServicesEc2ModelSnapshotTaskDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/SnapshotTaskDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFormat(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) SetFormat(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) GetFormat() string {
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

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withFormat(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) WithFormat(a string) *ServicesEc2ModelSnapshotTaskDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFormat", "com/amazonaws/services/ec2/model/SnapshotTaskDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUrl(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) SetUrl(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) GetUrl() string {
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

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withUrl(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) WithUrl(a string) *ServicesEc2ModelSnapshotTaskDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUrl", "com/amazonaws/services/ec2/model/SnapshotTaskDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserBucket(com.amazonaws.services.ec2.model.UserBucketDetails)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) SetUserBucket(a ServicesEc2ModelUserBucketDetailsInterface)  {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) GetUserBucket() *ServicesEc2ModelUserBucketDetails {
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

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withUserBucket(com.amazonaws.services.ec2.model.UserBucketDetails)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) WithUserBucket(a ServicesEc2ModelUserBucketDetailsInterface) *ServicesEc2ModelSnapshotTaskDetail {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserBucket", "com/amazonaws/services/ec2/model/SnapshotTaskDetail", conv_a.Value().Cast("com/amazonaws/services/ec2/model/UserBucketDetails"))
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) SetSnapshotId(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) GetSnapshotId() string {
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

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) WithSnapshotId(a string) *ServicesEc2ModelSnapshotTaskDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/SnapshotTaskDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProgress(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) SetProgress(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) GetProgress() string {
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

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withProgress(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) WithProgress(a string) *ServicesEc2ModelSnapshotTaskDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProgress", "com/amazonaws/services/ec2/model/SnapshotTaskDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) SetStatusMessage(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) GetStatusMessage() string {
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

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) WithStatusMessage(a string) *ServicesEc2ModelSnapshotTaskDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatusMessage", "com/amazonaws/services/ec2/model/SnapshotTaskDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) SetStatus(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) GetStatus() string {
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

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) WithStatus(a string) *ServicesEc2ModelSnapshotTaskDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/SnapshotTaskDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) ToString() string {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SnapshotTaskDetail clone()
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) Clone() *ServicesEc2ModelSnapshotTaskDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SnapshotTaskDetail")
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
	unique_x := &ServicesEc2ModelSnapshotTaskDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSnapshotTaskDetail) Clone2() (*JavaLangObject, error) {
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


