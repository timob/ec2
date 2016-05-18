package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImportInstanceVolumeDetailItemInterface interface {
	JavaLangObjectInterface

	// public void setBytesConverted(java.lang.Long)
	SetBytesConverted(a int64) 

	// public java.lang.Long getBytesConverted()
	GetBytesConverted() int64

	// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withBytesConverted(java.lang.Long)
	WithBytesConverted(a int64) *ServicesEc2ModelImportInstanceVolumeDetailItem

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelImportInstanceVolumeDetailItem

	// public void setImage(com.amazonaws.services.ec2.model.DiskImageDescription)
	SetImage(a ServicesEc2ModelDiskImageDescriptionInterface) 

	// public com.amazonaws.services.ec2.model.DiskImageDescription getImage()
	GetImage() *ServicesEc2ModelDiskImageDescription

	// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withImage(com.amazonaws.services.ec2.model.DiskImageDescription)
	WithImage(a ServicesEc2ModelDiskImageDescriptionInterface) *ServicesEc2ModelImportInstanceVolumeDetailItem

	// public void setVolume(com.amazonaws.services.ec2.model.DiskImageVolumeDescription)
	SetVolume(a ServicesEc2ModelDiskImageVolumeDescriptionInterface) 

	// public com.amazonaws.services.ec2.model.DiskImageVolumeDescription getVolume()
	GetVolume() *ServicesEc2ModelDiskImageVolumeDescription

	// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withVolume(com.amazonaws.services.ec2.model.DiskImageVolumeDescription)
	WithVolume(a ServicesEc2ModelDiskImageVolumeDescriptionInterface) *ServicesEc2ModelImportInstanceVolumeDetailItem

	// public void setStatus(java.lang.String)
	SetStatus(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withStatus(java.lang.String)
	WithStatus(a string) *ServicesEc2ModelImportInstanceVolumeDetailItem

	// public void setStatusMessage(java.lang.String)
	SetStatusMessage(a string) 

	// public java.lang.String getStatusMessage()
	GetStatusMessage() string

	// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withStatusMessage(java.lang.String)
	WithStatusMessage(a string) *ServicesEc2ModelImportInstanceVolumeDetailItem

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelImportInstanceVolumeDetailItem

	// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem clone()
	Clone() *ServicesEc2ModelImportInstanceVolumeDetailItem
}

type ServicesEc2ModelImportInstanceVolumeDetailItem struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem()
func NewServicesEc2ModelImportInstanceVolumeDetailItem() (*ServicesEc2ModelImportInstanceVolumeDetailItem) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelImportInstanceVolumeDetailItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setBytesConverted(java.lang.Long)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) SetBytesConverted(a int64)  {
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
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) GetBytesConverted() int64 {
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

// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withBytesConverted(java.lang.Long)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) WithBytesConverted(a int64) *ServicesEc2ModelImportInstanceVolumeDetailItem {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withBytesConverted", "com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelImportInstanceVolumeDetailItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) WithAvailabilityZone(a string) *ServicesEc2ModelImportInstanceVolumeDetailItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceVolumeDetailItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setImage(com.amazonaws.services.ec2.model.DiskImageDescription)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) SetImage(a ServicesEc2ModelDiskImageDescriptionInterface)  {
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
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) GetImage() *ServicesEc2ModelDiskImageDescription {
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

// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withImage(com.amazonaws.services.ec2.model.DiskImageDescription)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) WithImage(a ServicesEc2ModelDiskImageDescriptionInterface) *ServicesEc2ModelImportInstanceVolumeDetailItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImage", "com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageDescription"))
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
	unique_x := &ServicesEc2ModelImportInstanceVolumeDetailItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolume(com.amazonaws.services.ec2.model.DiskImageVolumeDescription)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) SetVolume(a ServicesEc2ModelDiskImageVolumeDescriptionInterface)  {
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
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) GetVolume() *ServicesEc2ModelDiskImageVolumeDescription {
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

// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withVolume(com.amazonaws.services.ec2.model.DiskImageVolumeDescription)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) WithVolume(a ServicesEc2ModelDiskImageVolumeDescriptionInterface) *ServicesEc2ModelImportInstanceVolumeDetailItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolume", "com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DiskImageVolumeDescription"))
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
	unique_x := &ServicesEc2ModelImportInstanceVolumeDetailItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) SetStatus(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) GetStatus() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) WithStatus(a string) *ServicesEc2ModelImportInstanceVolumeDetailItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceVolumeDetailItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) SetStatusMessage(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) GetStatusMessage() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withStatusMessage(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) WithStatusMessage(a string) *ServicesEc2ModelImportInstanceVolumeDetailItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatusMessage", "com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceVolumeDetailItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) WithDescription(a string) *ServicesEc2ModelImportInstanceVolumeDetailItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImportInstanceVolumeDetailItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) ToString() string {
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
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ImportInstanceVolumeDetailItem clone()
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) Clone() *ServicesEc2ModelImportInstanceVolumeDetailItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ImportInstanceVolumeDetailItem")
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
	unique_x := &ServicesEc2ModelImportInstanceVolumeDetailItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelImportInstanceVolumeDetailItem) Clone2() (*JavaLangObject, error) {
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


