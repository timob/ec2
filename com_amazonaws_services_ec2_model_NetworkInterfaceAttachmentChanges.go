package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNetworkInterfaceAttachmentChangesInterface interface {
	JavaLangObjectInterface

	// public void setAttachmentId(java.lang.String)
	SetAttachmentId(a string) 

	// public java.lang.String getAttachmentId()
	GetAttachmentId() string

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges withAttachmentId(java.lang.String)
	WithAttachmentId(a string) *ServicesEc2ModelNetworkInterfaceAttachmentChanges

	// public void setDeleteOnTermination(java.lang.Boolean)
	SetDeleteOnTermination(a bool) 

	// public java.lang.Boolean getDeleteOnTermination()
	GetDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges withDeleteOnTermination(java.lang.Boolean)
	WithDeleteOnTermination(a bool) *ServicesEc2ModelNetworkInterfaceAttachmentChanges

	// public java.lang.Boolean isDeleteOnTermination()
	IsDeleteOnTermination() bool

	// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges clone()
	Clone() *ServicesEc2ModelNetworkInterfaceAttachmentChanges
}

type ServicesEc2ModelNetworkInterfaceAttachmentChanges struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges()
func NewServicesEc2ModelNetworkInterfaceAttachmentChanges() (*ServicesEc2ModelNetworkInterfaceAttachmentChanges) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/NetworkInterfaceAttachmentChanges")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelNetworkInterfaceAttachmentChanges{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAttachmentId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) SetAttachmentId(a string)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) GetAttachmentId() string {
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges withAttachmentId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) WithAttachmentId(a string) *ServicesEc2ModelNetworkInterfaceAttachmentChanges {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttachmentId", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachmentChanges", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachmentChanges{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) SetDeleteOnTermination(a bool)  {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) GetDeleteOnTermination() bool {
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

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges withDeleteOnTermination(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) WithDeleteOnTermination(a bool) *ServicesEc2ModelNetworkInterfaceAttachmentChanges {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDeleteOnTermination", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachmentChanges", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachmentChanges{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDeleteOnTermination()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) IsDeleteOnTermination() bool {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) ToString() string {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.NetworkInterfaceAttachmentChanges clone()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) Clone() *ServicesEc2ModelNetworkInterfaceAttachmentChanges {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/NetworkInterfaceAttachmentChanges")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttachmentChanges{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelNetworkInterfaceAttachmentChanges) Clone2() (*JavaLangObject, error) {
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


