package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelExportToS3TaskSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setDiskImageFormat(java.lang.String)
	SetDiskImageFormat2(a string) 

	// public java.lang.String getDiskImageFormat()
	GetDiskImageFormat() string

	// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withDiskImageFormat(java.lang.String)
	WithDiskImageFormat2(a string) *ServicesEc2ModelExportToS3TaskSpecification

	// public void setDiskImageFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
	SetDiskImageFormat(a ServicesEc2ModelDiskImageFormatInterface) 

	// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withDiskImageFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
	WithDiskImageFormat(a ServicesEc2ModelDiskImageFormatInterface) *ServicesEc2ModelExportToS3TaskSpecification

	// public void setContainerFormat(java.lang.String)
	SetContainerFormat2(a string) 

	// public java.lang.String getContainerFormat()
	GetContainerFormat() string

	// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withContainerFormat(java.lang.String)
	WithContainerFormat2(a string) *ServicesEc2ModelExportToS3TaskSpecification

	// public void setContainerFormat(com.amazonaws.services.ec2.model.ContainerFormat)
	SetContainerFormat(a ServicesEc2ModelContainerFormatInterface) 

	// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withContainerFormat(com.amazonaws.services.ec2.model.ContainerFormat)
	WithContainerFormat(a ServicesEc2ModelContainerFormatInterface) *ServicesEc2ModelExportToS3TaskSpecification

	// public void setS3Bucket(java.lang.String)
	SetS3Bucket(a string) 

	// public java.lang.String getS3Bucket()
	GetS3Bucket() string

	// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withS3Bucket(java.lang.String)
	WithS3Bucket(a string) *ServicesEc2ModelExportToS3TaskSpecification

	// public void setS3Prefix(java.lang.String)
	SetS3Prefix(a string) 

	// public java.lang.String getS3Prefix()
	GetS3Prefix() string

	// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withS3Prefix(java.lang.String)
	WithS3Prefix(a string) *ServicesEc2ModelExportToS3TaskSpecification

	// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification clone()
	Clone() *ServicesEc2ModelExportToS3TaskSpecification
}

type ServicesEc2ModelExportToS3TaskSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification()
func NewServicesEc2ModelExportToS3TaskSpecification() (*ServicesEc2ModelExportToS3TaskSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ExportToS3TaskSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelExportToS3TaskSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDiskImageFormat(java.lang.String)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) SetDiskImageFormat2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDiskImageFormat", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDiskImageFormat()
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) GetDiskImageFormat() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDiskImageFormat", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withDiskImageFormat(java.lang.String)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) WithDiskImageFormat2(a string) *ServicesEc2ModelExportToS3TaskSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDiskImageFormat", "com/amazonaws/services/ec2/model/ExportToS3TaskSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportToS3TaskSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDiskImageFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) SetDiskImageFormat(a ServicesEc2ModelDiskImageFormatInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDiskImageFormat", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageFormat"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withDiskImageFormat(com.amazonaws.services.ec2.model.DiskImageFormat)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) WithDiskImageFormat(a ServicesEc2ModelDiskImageFormatInterface) *ServicesEc2ModelExportToS3TaskSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDiskImageFormat", "com/amazonaws/services/ec2/model/ExportToS3TaskSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageFormat"))
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
	unique_x := &ServicesEc2ModelExportToS3TaskSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setContainerFormat(java.lang.String)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) SetContainerFormat2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setContainerFormat", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getContainerFormat()
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) GetContainerFormat() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContainerFormat", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withContainerFormat(java.lang.String)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) WithContainerFormat2(a string) *ServicesEc2ModelExportToS3TaskSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withContainerFormat", "com/amazonaws/services/ec2/model/ExportToS3TaskSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportToS3TaskSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setContainerFormat(com.amazonaws.services.ec2.model.ContainerFormat)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) SetContainerFormat(a ServicesEc2ModelContainerFormatInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setContainerFormat", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ContainerFormat"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withContainerFormat(com.amazonaws.services.ec2.model.ContainerFormat)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) WithContainerFormat(a ServicesEc2ModelContainerFormatInterface) *ServicesEc2ModelExportToS3TaskSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withContainerFormat", "com/amazonaws/services/ec2/model/ExportToS3TaskSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ContainerFormat"))
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
	unique_x := &ServicesEc2ModelExportToS3TaskSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setS3Bucket(java.lang.String)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) SetS3Bucket(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setS3Bucket", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getS3Bucket()
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) GetS3Bucket() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getS3Bucket", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withS3Bucket(java.lang.String)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) WithS3Bucket(a string) *ServicesEc2ModelExportToS3TaskSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withS3Bucket", "com/amazonaws/services/ec2/model/ExportToS3TaskSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportToS3TaskSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setS3Prefix(java.lang.String)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) SetS3Prefix(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setS3Prefix", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getS3Prefix()
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) GetS3Prefix() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getS3Prefix", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification withS3Prefix(java.lang.String)
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) WithS3Prefix(a string) *ServicesEc2ModelExportToS3TaskSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withS3Prefix", "com/amazonaws/services/ec2/model/ExportToS3TaskSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportToS3TaskSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) ToString() string {
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
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ExportToS3TaskSpecification clone()
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) Clone() *ServicesEc2ModelExportToS3TaskSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ExportToS3TaskSpecification")
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
	unique_x := &ServicesEc2ModelExportToS3TaskSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelExportToS3TaskSpecification) Clone2() (*JavaLangObject, error) {
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


