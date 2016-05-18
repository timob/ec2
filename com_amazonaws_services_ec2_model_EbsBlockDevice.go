package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelEbsBlockDeviceInterface interface {
	JavaLangObjectInterface

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.EbsBlockDevice withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelEbsBlockDevice

	// public void setVolumeSize(java.lang.Integer)
	SetVolumeSize(a int) 

	// public java.lang.Integer getVolumeSize()
	GetVolumeSize() int

	// public com.amazonaws.services.ec2.model.EbsBlockDevice withVolumeSize(java.lang.Integer)
	WithVolumeSize(a int) *ServicesEc2ModelEbsBlockDevice

	// public void setDeleteOnTermination(java.lang.Boolean)
	SetDeleteOnTermination(a bool) 

	// public java.lang.Boolean getDeleteOnTermination()
	GetDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.EbsBlockDevice withDeleteOnTermination(java.lang.Boolean)
	WithDeleteOnTermination(a bool) *ServicesEc2ModelEbsBlockDevice

	// public java.lang.Boolean isDeleteOnTermination()
	IsDeleteOnTermination() bool

	// public void setVolumeType(java.lang.String)
	SetVolumeType2(a string) 

	// public java.lang.String getVolumeType()
	GetVolumeType() string

	// public com.amazonaws.services.ec2.model.EbsBlockDevice withVolumeType(java.lang.String)
	WithVolumeType2(a string) *ServicesEc2ModelEbsBlockDevice

	// public void setVolumeType(com.amazonaws.services.ec2.model.VolumeType)
	SetVolumeType(a ServicesEc2ModelVolumeTypeInterface) 

	// public com.amazonaws.services.ec2.model.EbsBlockDevice withVolumeType(com.amazonaws.services.ec2.model.VolumeType)
	WithVolumeType(a ServicesEc2ModelVolumeTypeInterface) *ServicesEc2ModelEbsBlockDevice

	// public void setIops(java.lang.Integer)
	SetIops(a int) 

	// public java.lang.Integer getIops()
	GetIops() int

	// public com.amazonaws.services.ec2.model.EbsBlockDevice withIops(java.lang.Integer)
	WithIops(a int) *ServicesEc2ModelEbsBlockDevice

	// public void setEncrypted(java.lang.Boolean)
	SetEncrypted(a bool) 

	// public java.lang.Boolean getEncrypted()
	GetEncrypted() bool

	// public com.amazonaws.services.ec2.model.EbsBlockDevice withEncrypted(java.lang.Boolean)
	WithEncrypted(a bool) *ServicesEc2ModelEbsBlockDevice

	// public java.lang.Boolean isEncrypted()
	IsEncrypted() bool

	// public com.amazonaws.services.ec2.model.EbsBlockDevice clone()
	Clone() *ServicesEc2ModelEbsBlockDevice
}

type ServicesEc2ModelEbsBlockDevice struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.EbsBlockDevice()
func NewServicesEc2ModelEbsBlockDevice() (*ServicesEc2ModelEbsBlockDevice) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/EbsBlockDevice")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelEbsBlockDevice{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelEbsBlockDevice) SetSnapshotId(a string)  {
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
func (jbobject *ServicesEc2ModelEbsBlockDevice) GetSnapshotId() string {
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

// public com.amazonaws.services.ec2.model.EbsBlockDevice withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelEbsBlockDevice) WithSnapshotId(a string) *ServicesEc2ModelEbsBlockDevice {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/EbsBlockDevice", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEbsBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolumeSize(java.lang.Integer)
func (jbobject *ServicesEc2ModelEbsBlockDevice) SetVolumeSize(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeSize", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getVolumeSize()
func (jbobject *ServicesEc2ModelEbsBlockDevice) GetVolumeSize() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeSize", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.EbsBlockDevice withVolumeSize(java.lang.Integer)
func (jbobject *ServicesEc2ModelEbsBlockDevice) WithVolumeSize(a int) *ServicesEc2ModelEbsBlockDevice {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeSize", "com/amazonaws/services/ec2/model/EbsBlockDevice", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelEbsBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelEbsBlockDevice) SetDeleteOnTermination(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeleteOnTermination", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getDeleteOnTermination()
func (jbobject *ServicesEc2ModelEbsBlockDevice) GetDeleteOnTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeleteOnTermination", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.EbsBlockDevice withDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelEbsBlockDevice) WithDeleteOnTermination(a bool) *ServicesEc2ModelEbsBlockDevice {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeleteOnTermination", "com/amazonaws/services/ec2/model/EbsBlockDevice", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelEbsBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDeleteOnTermination()
func (jbobject *ServicesEc2ModelEbsBlockDevice) IsDeleteOnTermination() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDeleteOnTermination", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVolumeType(java.lang.String)
func (jbobject *ServicesEc2ModelEbsBlockDevice) SetVolumeType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVolumeType()
func (jbobject *ServicesEc2ModelEbsBlockDevice) GetVolumeType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.EbsBlockDevice withVolumeType(java.lang.String)
func (jbobject *ServicesEc2ModelEbsBlockDevice) WithVolumeType2(a string) *ServicesEc2ModelEbsBlockDevice {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeType", "com/amazonaws/services/ec2/model/EbsBlockDevice", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEbsBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolumeType(com.amazonaws.services.ec2.model.VolumeType)
func (jbobject *ServicesEc2ModelEbsBlockDevice) SetVolumeType(a ServicesEc2ModelVolumeTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.EbsBlockDevice withVolumeType(com.amazonaws.services.ec2.model.VolumeType)
func (jbobject *ServicesEc2ModelEbsBlockDevice) WithVolumeType(a ServicesEc2ModelVolumeTypeInterface) *ServicesEc2ModelEbsBlockDevice {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeType", "com/amazonaws/services/ec2/model/EbsBlockDevice", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeType"))
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
	unique_x := &ServicesEc2ModelEbsBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIops(java.lang.Integer)
func (jbobject *ServicesEc2ModelEbsBlockDevice) SetIops(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIops", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getIops()
func (jbobject *ServicesEc2ModelEbsBlockDevice) GetIops() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIops", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.EbsBlockDevice withIops(java.lang.Integer)
func (jbobject *ServicesEc2ModelEbsBlockDevice) WithIops(a int) *ServicesEc2ModelEbsBlockDevice {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIops", "com/amazonaws/services/ec2/model/EbsBlockDevice", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelEbsBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelEbsBlockDevice) SetEncrypted(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEncrypted", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEncrypted()
func (jbobject *ServicesEc2ModelEbsBlockDevice) GetEncrypted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEncrypted", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.EbsBlockDevice withEncrypted(java.lang.Boolean)
func (jbobject *ServicesEc2ModelEbsBlockDevice) WithEncrypted(a bool) *ServicesEc2ModelEbsBlockDevice {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEncrypted", "com/amazonaws/services/ec2/model/EbsBlockDevice", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelEbsBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEncrypted()
func (jbobject *ServicesEc2ModelEbsBlockDevice) IsEncrypted() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEncrypted", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelEbsBlockDevice) ToString() string {
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
func (jbobject *ServicesEc2ModelEbsBlockDevice) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelEbsBlockDevice) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.EbsBlockDevice clone()
func (jbobject *ServicesEc2ModelEbsBlockDevice) Clone() *ServicesEc2ModelEbsBlockDevice {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/EbsBlockDevice")
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
	unique_x := &ServicesEc2ModelEbsBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelEbsBlockDevice) Clone2() (*JavaLangObject, error) {
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


