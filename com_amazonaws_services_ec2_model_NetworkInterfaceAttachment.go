package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelNetworkInterfaceAttachmentInterface interface {
	JavaLangObjectInterface

	// public void setAttachmentId(java.lang.String)
	SetAttachmentId(a string) 

	// public java.lang.String getAttachmentId()
	GetAttachmentId() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withAttachmentId(java.lang.String)
	WithAttachmentId(a string) *ServicesEc2ModelNetworkInterfaceAttachment

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelNetworkInterfaceAttachment

	// public void setInstanceOwnerId(java.lang.String)
	SetInstanceOwnerId(a string) 

	// public java.lang.String getInstanceOwnerId()
	GetInstanceOwnerId() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withInstanceOwnerId(java.lang.String)
	WithInstanceOwnerId(a string) *ServicesEc2ModelNetworkInterfaceAttachment

	// public void setDeviceIndex(java.lang.Integer)
	SetDeviceIndex(a int) 

	// public java.lang.Integer getDeviceIndex()
	GetDeviceIndex() int

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withDeviceIndex(java.lang.Integer)
	WithDeviceIndex(a int) *ServicesEc2ModelNetworkInterfaceAttachment

	// public void setStatus(java.lang.String)
	SetStatus2(a string) 

	// public java.lang.String getStatus()
	GetStatus() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withStatus(java.lang.String)
	WithStatus2(a string) *ServicesEc2ModelNetworkInterfaceAttachment

	// public void setStatus(com.amazonaws.services.ec2.model.AttachmentStatus)
	SetStatus(a ServicesEc2ModelAttachmentStatusInterface) 

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withStatus(com.amazonaws.services.ec2.model.AttachmentStatus)
	WithStatus(a ServicesEc2ModelAttachmentStatusInterface) *ServicesEc2ModelNetworkInterfaceAttachment

	// public void setAttachTime(java.util.Date)
	SetAttachTime(a time.Time) 

	// public java.util.Date getAttachTime()
	GetAttachTime() time.Time

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withAttachTime(java.util.Date)
	WithAttachTime(a time.Time) *ServicesEc2ModelNetworkInterfaceAttachment

	// public void setDeleteOnTermination(java.lang.Boolean)
	SetDeleteOnTermination(a bool) 

	// public java.lang.Boolean getDeleteOnTermination()
	GetDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withDeleteOnTermination(java.lang.Boolean)
	WithDeleteOnTermination(a bool) *ServicesEc2ModelNetworkInterfaceAttachment

	// public java.lang.Boolean isDeleteOnTermination()
	IsDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment clone()
	Clone() *ServicesEc2ModelNetworkInterfaceAttachment
}

type ServicesEc2ModelNetworkInterfaceAttachment struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment()
func NewServicesEc2ModelNetworkInterfaceAttachment() (*ServicesEc2ModelNetworkInterfaceAttachment) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/NetworkInterfaceAttachment")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAttachmentId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) SetAttachmentId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttachmentId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAttachmentId()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) GetAttachmentId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttachmentId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withAttachmentId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) WithAttachmentId(a string) *ServicesEc2ModelNetworkInterfaceAttachment {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachmentId", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) WithInstanceId(a string) *ServicesEc2ModelNetworkInterfaceAttachment {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) SetInstanceOwnerId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceOwnerId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceOwnerId()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) GetInstanceOwnerId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceOwnerId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withInstanceOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) WithInstanceOwnerId(a string) *ServicesEc2ModelNetworkInterfaceAttachment {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceOwnerId", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeviceIndex(java.lang.Integer)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) SetDeviceIndex(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDeviceIndex", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getDeviceIndex()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) GetDeviceIndex() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeviceIndex", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withDeviceIndex(java.lang.Integer)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) WithDeviceIndex(a int) *ServicesEc2ModelNetworkInterfaceAttachment {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeviceIndex", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) SetStatus2(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) GetStatus() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withStatus(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) WithStatus2(a string) *ServicesEc2ModelNetworkInterfaceAttachment {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStatus(com.amazonaws.services.ec2.model.AttachmentStatus)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) SetStatus(a ServicesEc2ModelAttachmentStatusInterface)  {
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withStatus(com.amazonaws.services.ec2.model.AttachmentStatus)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) WithStatus(a ServicesEc2ModelAttachmentStatusInterface) *ServicesEc2ModelNetworkInterfaceAttachment {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStatus", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AttachmentStatus"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAttachTime(java.util.Date)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) SetAttachTime(a time.Time)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) GetAttachTime() time.Time {
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withAttachTime(java.util.Date)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) WithAttachTime(a time.Time) *ServicesEc2ModelNetworkInterfaceAttachment {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachTime", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) SetDeleteOnTermination(a bool)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) GetDeleteOnTermination() bool {
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment withDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) WithDeleteOnTermination(a bool) *ServicesEc2ModelNetworkInterfaceAttachment {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeleteOnTermination", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDeleteOnTermination()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) IsDeleteOnTermination() bool {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) ToString() string {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachment clone()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) Clone() *ServicesEc2ModelNetworkInterfaceAttachment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachment")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachment{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachment) Clone2() (*JavaLangObject, error) {
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


