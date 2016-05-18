package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSnapshotDiskContainerInterface interface {
	JavaLangObjectInterface

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.SnapshotDiskContainer withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelSnapshotDiskContainer

	// public void setFormat(java.lang.String)
	SetFormat(a string) 

	// public java.lang.String getFormat()
	GetFormat() string

	// public com.amazonaws.services.ec2.model.SnapshotDiskContainer withFormat(java.lang.String)
	WithFormat(a string) *ServicesEc2ModelSnapshotDiskContainer

	// public void setUrl(java.lang.String)
	SetUrl(a string) 

	// public java.lang.String getUrl()
	GetUrl() string

	// public com.amazonaws.services.ec2.model.SnapshotDiskContainer withUrl(java.lang.String)
	WithUrl(a string) *ServicesEc2ModelSnapshotDiskContainer

	// public void setUserBucket(com.amazonaws.services.ec2.model.UserBucket)
	SetUserBucket(a ServicesEc2ModelUserBucketInterface) 

	// public com.amazonaws.services.ec2.model.UserBucket getUserBucket()
	GetUserBucket() *ServicesEc2ModelUserBucket

	// public com.amazonaws.services.ec2.model.SnapshotDiskContainer withUserBucket(com.amazonaws.services.ec2.model.UserBucket)
	WithUserBucket(a ServicesEc2ModelUserBucketInterface) *ServicesEc2ModelSnapshotDiskContainer

	// public com.amazonaws.services.ec2.model.SnapshotDiskContainer clone()
	Clone() *ServicesEc2ModelSnapshotDiskContainer
}

type ServicesEc2ModelSnapshotDiskContainer struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SnapshotDiskContainer()
func NewServicesEc2ModelSnapshotDiskContainer() (*ServicesEc2ModelSnapshotDiskContainer) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SnapshotDiskContainer")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSnapshotDiskContainer{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.SnapshotDiskContainer withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) WithDescription(a string) *ServicesEc2ModelSnapshotDiskContainer {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/SnapshotDiskContainer", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDiskContainer{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFormat(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) SetFormat(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) GetFormat() string {
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

// public com.amazonaws.services.ec2.model.SnapshotDiskContainer withFormat(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) WithFormat(a string) *ServicesEc2ModelSnapshotDiskContainer {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFormat", "com/amazonaws/services/ec2/model/SnapshotDiskContainer", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDiskContainer{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUrl(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) SetUrl(a string)  {
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
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) GetUrl() string {
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

// public com.amazonaws.services.ec2.model.SnapshotDiskContainer withUrl(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) WithUrl(a string) *ServicesEc2ModelSnapshotDiskContainer {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUrl", "com/amazonaws/services/ec2/model/SnapshotDiskContainer", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotDiskContainer{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUserBucket(com.amazonaws.services.ec2.model.UserBucket)
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) SetUserBucket(a ServicesEc2ModelUserBucketInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserBucket", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/UserBucket"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.UserBucket getUserBucket()
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) GetUserBucket() *ServicesEc2ModelUserBucket {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserBucket", "com/amazonaws/services/ec2/model/UserBucket")
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
	unique_x := &ServicesEc2ModelUserBucket{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SnapshotDiskContainer withUserBucket(com.amazonaws.services.ec2.model.UserBucket)
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) WithUserBucket(a ServicesEc2ModelUserBucketInterface) *ServicesEc2ModelSnapshotDiskContainer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserBucket", "com/amazonaws/services/ec2/model/SnapshotDiskContainer", conv_a.Value().Cast("com/amazonaws/services/ec2/model/UserBucket"))
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
	unique_x := &ServicesEc2ModelSnapshotDiskContainer{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) ToString() string {
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
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SnapshotDiskContainer clone()
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) Clone() *ServicesEc2ModelSnapshotDiskContainer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SnapshotDiskContainer")
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
	unique_x := &ServicesEc2ModelSnapshotDiskContainer{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSnapshotDiskContainer) Clone2() (*JavaLangObject, error) {
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


