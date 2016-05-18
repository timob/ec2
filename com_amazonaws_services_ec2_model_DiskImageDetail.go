package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDiskImageDetailInterface interface {
	JavaLangObjectInterface

	// public void setFormat(java.lang.String)
	SetFormat2(a string) 

	// public java.lang.String getFormat()
	GetFormat() string

	// public com.amazonaws.services.ec2.model.DiskImageDetail withFormat(java.lang.String)
	WithFormat2(a string) *ServicesEc2ModelDiskImageDetail

	// public void setFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
	SetFormat(a ServicesEc2ModelDiskImageFormatInterface) 

	// public com.amazonaws.services.ec2.model.DiskImageDetail withFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
	WithFormat(a ServicesEc2ModelDiskImageFormatInterface) *ServicesEc2ModelDiskImageDetail

	// public void setBytes(java.lang.Long)
	SetBytes(a int64) 

	// public java.lang.Long getBytes()
	GetBytes() int64

	// public com.amazonaws.services.ec2.model.DiskImageDetail withBytes(java.lang.Long)
	WithBytes(a int64) *ServicesEc2ModelDiskImageDetail

	// public void setImportManifestUrl(java.lang.String)
	SetImportManifestUrl(a string) 

	// public java.lang.String getImportManifestUrl()
	GetImportManifestUrl() string

	// public com.amazonaws.services.ec2.model.DiskImageDetail withImportManifestUrl(java.lang.String)
	WithImportManifestUrl(a string) *ServicesEc2ModelDiskImageDetail

	// public com.amazonaws.services.ec2.model.DiskImageDetail clone()
	Clone() *ServicesEc2ModelDiskImageDetail
}

type ServicesEc2ModelDiskImageDetail struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DiskImageDetail()
func NewServicesEc2ModelDiskImageDetail() (*ServicesEc2ModelDiskImageDetail) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DiskImageDetail")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDiskImageDetail{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setFormat(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDetail) SetFormat2(a string)  {
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
func (jbobject *ServicesEc2ModelDiskImageDetail) GetFormat() string {
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

// public com.amazonaws.services.ec2.model.DiskImageDetail withFormat(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDetail) WithFormat2(a string) *ServicesEc2ModelDiskImageDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFormat", "com/amazonaws/services/ec2/model/DiskImageDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDiskImageDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
func (jbobject *ServicesEc2ModelDiskImageDetail) SetFormat(a ServicesEc2ModelDiskImageFormatInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFormat", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageFormat"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DiskImageDetail withFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
func (jbobject *ServicesEc2ModelDiskImageDetail) WithFormat(a ServicesEc2ModelDiskImageFormatInterface) *ServicesEc2ModelDiskImageDetail {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFormat", "com/amazonaws/services/ec2/model/DiskImageDetail", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageFormat"))
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
	unique_x := &ServicesEc2ModelDiskImageDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBytes(java.lang.Long)
func (jbobject *ServicesEc2ModelDiskImageDetail) SetBytes(a int64)  {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBytes", javabind.Void, conv_a.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Long getBytes()
func (jbobject *ServicesEc2ModelDiskImageDetail) GetBytes() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBytes", "java/lang/Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.DiskImageDetail withBytes(java.lang.Long)
func (jbobject *ServicesEc2ModelDiskImageDetail) WithBytes(a int64) *ServicesEc2ModelDiskImageDetail {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBytes", "com/amazonaws/services/ec2/model/DiskImageDetail", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelDiskImageDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImportManifestUrl(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDetail) SetImportManifestUrl(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImportManifestUrl", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getImportManifestUrl()
func (jbobject *ServicesEc2ModelDiskImageDetail) GetImportManifestUrl() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImportManifestUrl", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DiskImageDetail withImportManifestUrl(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDetail) WithImportManifestUrl(a string) *ServicesEc2ModelDiskImageDetail {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportManifestUrl", "com/amazonaws/services/ec2/model/DiskImageDetail", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDiskImageDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDiskImageDetail) ToString() string {
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
func (jbobject *ServicesEc2ModelDiskImageDetail) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDiskImageDetail) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DiskImageDetail clone()
func (jbobject *ServicesEc2ModelDiskImageDetail) Clone() *ServicesEc2ModelDiskImageDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DiskImageDetail")
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
	unique_x := &ServicesEc2ModelDiskImageDetail{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDiskImageDetail) Clone2() (*JavaLangObject, error) {
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


