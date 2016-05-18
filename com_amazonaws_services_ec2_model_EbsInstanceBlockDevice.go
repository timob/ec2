package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelEbsInstanceBlockDeviceInterface interface {
	JavaLangObjectInterface

	// public void setVolumeId(java.lang.String)
	SetVolumeId(a string) 

	// public java.lang.String getVolumeId()
	GetVolumeId() string

	// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withVolumeId(java.lang.String)
	WithVolumeId(a string) *ServicesEc2ModelEbsInstanceBlockDevice

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelEbsInstanceBlockDevice

	// public void setStatus(com.amazonaws.services.ec2.model.AttachmentStatus)
	SetStatus(a ServicesEc2ModelAttachmentStatusInterface) 

	// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withStatus(com.amazonaws.services.ec2.model.AttachmentStatus)
	WithStatus(a ServicesEc2ModelAttachmentStatusInterface) *ServicesEc2ModelEbsInstanceBlockDevice

	// public void setAttachTime(java.util.Date)
	SetAttachTime(a time.Time) 

	// public java.util.Date getAttachTime()
	GetAttachTime() time.Time

	// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withAttachTime(java.util.Date)
	WithAttachTime(a time.Time) *ServicesEc2ModelEbsInstanceBlockDevice

	// public void setDeleteOnTermination(java.lang.Boolean)
	SetDeleteOnTermination(a bool) 

	// public java.lang.Boolean getDeleteOnTermination()
	GetDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withDeleteOnTermination(java.lang.Boolean)
	WithDeleteOnTermination(a bool) *ServicesEc2ModelEbsInstanceBlockDevice

	// public java.lang.Boolean isDeleteOnTermination()
	IsDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice clone()
	Clone() *ServicesEc2ModelEbsInstanceBlockDevice
}

type ServicesEc2ModelEbsInstanceBlockDevice struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice()
func NewServicesEc2ModelEbsInstanceBlockDevice() (*ServicesEc2ModelEbsInstanceBlockDevice) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/EbsInstanceBlockDevice")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelEbsInstanceBlockDevice{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVolumeId(java.lang.String)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) SetVolumeId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVolumeId()
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) GetVolumeId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withVolumeId(java.lang.String)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) WithVolumeId(a string) *ServicesEc2ModelEbsInstanceBlockDevice {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeId", "com/amazonaws/services/ec2/model/EbsInstanceBlockDevice", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEbsInstanceBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) SetStatus2(a string)  {
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
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) GetStatus() string {
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

// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) WithStatus2(a string) *ServicesEc2ModelEbsInstanceBlockDevice {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/EbsInstanceBlockDevice", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEbsInstanceBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.AttachmentStatus)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) SetStatus(a ServicesEc2ModelAttachmentStatusInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachmentStatus"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withStatus(com.amazonaws.services.ec2.model.AttachmentStatus)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) WithStatus(a ServicesEc2ModelAttachmentStatusInterface) *ServicesEc2ModelEbsInstanceBlockDevice {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/EbsInstanceBlockDevice", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachmentStatus"))
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
	unique_x := &ServicesEc2ModelEbsInstanceBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttachTime(java.util.Date)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) SetAttachTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getAttachTime()
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) GetAttachTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachTime", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withAttachTime(java.util.Date)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) WithAttachTime(a time.Time) *ServicesEc2ModelEbsInstanceBlockDevice {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachTime", "com/amazonaws/services/ec2/model/EbsInstanceBlockDevice", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelEbsInstanceBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) SetDeleteOnTermination(a bool)  {
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
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) GetDeleteOnTermination() bool {
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

// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice withDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) WithDeleteOnTermination(a bool) *ServicesEc2ModelEbsInstanceBlockDevice {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeleteOnTermination", "com/amazonaws/services/ec2/model/EbsInstanceBlockDevice", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelEbsInstanceBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDeleteOnTermination()
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) IsDeleteOnTermination() bool {
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

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) ToString() string {
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
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.EbsInstanceBlockDevice clone()
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) Clone() *ServicesEc2ModelEbsInstanceBlockDevice {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/EbsInstanceBlockDevice")
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
	unique_x := &ServicesEc2ModelEbsInstanceBlockDevice{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelEbsInstanceBlockDevice) Clone2() (*JavaLangObject, error) {
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


