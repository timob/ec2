package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportVolumeTaskDetailsInterface interface {
	JavaLangObjectInterface

	// public void setBytesConverted(java.lang.Long)
	SetBytesConverted(a int64) 

	// public java.lang.Long getBytesConverted()
	GetBytesConverted() int64

	// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withBytesConverted(java.lang.Long)
	WithBytesConverted(a int64) *ServicesEc2ModelImportVolumeTaskDetails

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelImportVolumeTaskDetails

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImportVolumeTaskDetails

	// public void setImage(com.amazonaws.services.ec2.model.DiskImageDescription)
	SetImage(a ServicesEc2ModelDiskImageDescriptionInterface) 

	// public com.amazonaws.services.ec2.model.DiskImageDescription getImage()
	GetImage() *ServicesEc2ModelDiskImageDescription

	// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withImage(com.amazonaws.services.ec2.model.DiskImageDescription)
	WithImage(a ServicesEc2ModelDiskImageDescriptionInterface) *ServicesEc2ModelImportVolumeTaskDetails

	// public void setVolume(com.amazonaws.services.ec2.model.DiskImageVolumeDescription)
	SetVolume(a ServicesEc2ModelDiskImageVolumeDescriptionInterface) 

	// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription getVolume()
	GetVolume() *ServicesEc2ModelDiskImageVolumeDescription

	// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withVolume(com.amazonaws.services.ec2.model.DiskImageVolumeDescription)
	WithVolume(a ServicesEc2ModelDiskImageVolumeDescriptionInterface) *ServicesEc2ModelImportVolumeTaskDetails

	// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails clone()
	Clone() *ServicesEc2ModelImportVolumeTaskDetails
}

type ServicesEc2ModelImportVolumeTaskDetails struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails()
func NewServicesEc2ModelImportVolumeTaskDetails() (*ServicesEc2ModelImportVolumeTaskDetails) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportVolumeTaskDetails")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportVolumeTaskDetails{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setBytesConverted(java.lang.Long)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) SetBytesConverted(a int64)  {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBytesConverted", javabind.Void, conv_a.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Long getBytesConverted()
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) GetBytesConverted() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBytesConverted", "java/lang/Long")
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

// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withBytesConverted(java.lang.Long)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) WithBytesConverted(a int64) *ServicesEc2ModelImportVolumeTaskDetails {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBytesConverted", "com/amazonaws/services/ec2/model/ImportVolumeTaskDetails", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelImportVolumeTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) SetAvailabilityZone(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailabilityZone", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAvailabilityZone()
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) GetAvailabilityZone() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailabilityZone", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) WithAvailabilityZone(a string) *ServicesEc2ModelImportVolumeTaskDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/ImportVolumeTaskDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportVolumeTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) WithDescription(a string) *ServicesEc2ModelImportVolumeTaskDetails {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ImportVolumeTaskDetails", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportVolumeTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImage(com.amazonaws.services.ec2.model.DiskImageDescription)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) SetImage(a ServicesEc2ModelDiskImageDescriptionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImage", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageDescription"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DiskImageDescription getImage()
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) GetImage() *ServicesEc2ModelDiskImageDescription {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImage", "com/amazonaws/services/ec2/model/DiskImageDescription")
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

// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withImage(com.amazonaws.services.ec2.model.DiskImageDescription)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) WithImage(a ServicesEc2ModelDiskImageDescriptionInterface) *ServicesEc2ModelImportVolumeTaskDetails {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImage", "com/amazonaws/services/ec2/model/ImportVolumeTaskDetails", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageDescription"))
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
	unique_x := &ServicesEc2ModelImportVolumeTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolume(com.amazonaws.services.ec2.model.DiskImageVolumeDescription)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) SetVolume(a ServicesEc2ModelDiskImageVolumeDescriptionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolume", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageVolumeDescription"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription getVolume()
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) GetVolume() *ServicesEc2ModelDiskImageVolumeDescription {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolume", "com/amazonaws/services/ec2/model/DiskImageVolumeDescription")
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
	unique_x := &ServicesEc2ModelDiskImageVolumeDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails withVolume(com.amazonaws.services.ec2.model.DiskImageVolumeDescription)
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) WithVolume(a ServicesEc2ModelDiskImageVolumeDescriptionInterface) *ServicesEc2ModelImportVolumeTaskDetails {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolume", "com/amazonaws/services/ec2/model/ImportVolumeTaskDetails", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageVolumeDescription"))
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
	unique_x := &ServicesEc2ModelImportVolumeTaskDetails{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) ToString() string {
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
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportVolumeTaskDetails clone()
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) Clone() *ServicesEc2ModelImportVolumeTaskDetails {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportVolumeTaskDetails")
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

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelImportVolumeTaskDetails) Clone2() (*JavaLangObject, error) {
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


