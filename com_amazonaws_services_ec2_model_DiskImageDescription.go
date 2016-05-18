package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDiskImageDescriptionInterface interface {
	JavaLangObjectInterface

	// public void setFormat(java.lang.String)
	SetFormat2(a string) 

	// public java.lang.String getFormat()
	GetFormat() string

	// public com.amazonaws.services.ec2.model.DiskImageDescription withFormat(java.lang.String)
	WithFormat2(a string) *ServicesEc2ModelDiskImageDescription

	// public void setFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
	SetFormat(a ServicesEc2ModelDiskImageFormatInterface) 

	// public com.amazonaws.services.ec2.model.DiskImageDescription withFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
	WithFormat(a ServicesEc2ModelDiskImageFormatInterface) *ServicesEc2ModelDiskImageDescription

	// public void setSize(java.lang.Long)
	SetSize(a int64) 

	// public java.lang.Long getSize()
	GetSize() int64

	// public com.amazonaws.services.ec2.model.DiskImageDescription withSize(java.lang.Long)
	WithSize(a int64) *ServicesEc2ModelDiskImageDescription

	// public void setImportManifestUrl(java.lang.String)
	SetImportManifestUrl(a string) 

	// public java.lang.String getImportManifestUrl()
	GetImportManifestUrl() string

	// public com.amazonaws.services.ec2.model.DiskImageDescription withImportManifestUrl(java.lang.String)
	WithImportManifestUrl(a string) *ServicesEc2ModelDiskImageDescription

	// public void setChecksum(java.lang.String)
	SetChecksum(a string) 

	// public java.lang.String getChecksum()
	GetChecksum() string

	// public com.amazonaws.services.ec2.model.DiskImageDescription withChecksum(java.lang.String)
	WithChecksum(a string) *ServicesEc2ModelDiskImageDescription

	// public com.amazonaws.services.ec2.model.DiskImageDescription clone()
	Clone() *ServicesEc2ModelDiskImageDescription
}

type ServicesEc2ModelDiskImageDescription struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DiskImageDescription()
func NewServicesEc2ModelDiskImageDescription() (*ServicesEc2ModelDiskImageDescription) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DiskImageDescription")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDiskImageDescription{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setFormat(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDescription) SetFormat2(a string)  {
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
func (jbobject *ServicesEc2ModelDiskImageDescription) GetFormat() string {
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

// public com.amazonaws.services.ec2.model.DiskImageDescription withFormat(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDescription) WithFormat2(a string) *ServicesEc2ModelDiskImageDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFormat", "com/amazonaws/services/ec2/model/DiskImageDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDiskImageDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
func (jbobject *ServicesEc2ModelDiskImageDescription) SetFormat(a ServicesEc2ModelDiskImageFormatInterface)  {
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

// public com.amazonaws.services.ec2.model.DiskImageDescription withFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
func (jbobject *ServicesEc2ModelDiskImageDescription) WithFormat(a ServicesEc2ModelDiskImageFormatInterface) *ServicesEc2ModelDiskImageDescription {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFormat", "com/amazonaws/services/ec2/model/DiskImageDescription", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageFormat"))
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
	unique_x := &ServicesEc2ModelDiskImageDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSize(java.lang.Long)
func (jbobject *ServicesEc2ModelDiskImageDescription) SetSize(a int64)  {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSize", javabind.Void, conv_a.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Long getSize()
func (jbobject *ServicesEc2ModelDiskImageDescription) GetSize() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSize", "java/lang/Long")
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

// public com.amazonaws.services.ec2.model.DiskImageDescription withSize(java.lang.Long)
func (jbobject *ServicesEc2ModelDiskImageDescription) WithSize(a int64) *ServicesEc2ModelDiskImageDescription {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSize", "com/amazonaws/services/ec2/model/DiskImageDescription", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelDiskImageDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImportManifestUrl(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDescription) SetImportManifestUrl(a string)  {
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
func (jbobject *ServicesEc2ModelDiskImageDescription) GetImportManifestUrl() string {
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

// public com.amazonaws.services.ec2.model.DiskImageDescription withImportManifestUrl(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDescription) WithImportManifestUrl(a string) *ServicesEc2ModelDiskImageDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImportManifestUrl", "com/amazonaws/services/ec2/model/DiskImageDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDiskImageDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public void setChecksum(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDescription) SetChecksum(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setChecksum", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getChecksum()
func (jbobject *ServicesEc2ModelDiskImageDescription) GetChecksum() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getChecksum", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DiskImageDescription withChecksum(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageDescription) WithChecksum(a string) *ServicesEc2ModelDiskImageDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withChecksum", "com/amazonaws/services/ec2/model/DiskImageDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDiskImageDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDiskImageDescription) ToString() string {
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
func (jbobject *ServicesEc2ModelDiskImageDescription) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDiskImageDescription) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DiskImageDescription clone()
func (jbobject *ServicesEc2ModelDiskImageDescription) Clone() *ServicesEc2ModelDiskImageDescription {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DiskImageDescription")
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
	unique_x := &ServicesEc2ModelDiskImageDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDiskImageDescription) Clone2() (*JavaLangObject, error) {
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


