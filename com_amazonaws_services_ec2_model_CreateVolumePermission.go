package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVolumePermissionInterface interface {
	JavaLangObjectInterface

	// public void setUserId(java.lang.String)
	SetUserId(a string) 

	// public java.lang.String getUserId()
	GetUserId() string

	// public com.amazonaws.services.ec2.model.CreateVolumePermission withUserId(java.lang.String)
	WithUserId(a string) *ServicesEc2ModelCreateVolumePermission

	// public void setGroup(java.lang.String)
	SetGroup2(a string) 

	// public java.lang.String getGroup()
	GetGroup() string

	// public com.amazonaws.services.ec2.model.CreateVolumePermission withGroup(java.lang.String)
	WithGroup2(a string) *ServicesEc2ModelCreateVolumePermission

	// public void setGroup(com.amazonaws.services.ec2.model.PermissionGroup)
	SetGroup(a ServicesEc2ModelPermissionGroupInterface) 

	// public com.amazonaws.services.ec2.model.CreateVolumePermission withGroup(com.amazonaws.services.ec2.model.PermissionGroup)
	WithGroup(a ServicesEc2ModelPermissionGroupInterface) *ServicesEc2ModelCreateVolumePermission

	// public com.amazonaws.services.ec2.model.CreateVolumePermission clone()
	Clone() *ServicesEc2ModelCreateVolumePermission
}

type ServicesEc2ModelCreateVolumePermission struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateVolumePermission()
func NewServicesEc2ModelCreateVolumePermission() (*ServicesEc2ModelCreateVolumePermission) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVolumePermission")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVolumePermission{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setUserId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumePermission) SetUserId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getUserId()
func (jbobject *ServicesEc2ModelCreateVolumePermission) GetUserId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVolumePermission withUserId(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumePermission) WithUserId(a string) *ServicesEc2ModelCreateVolumePermission {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserId", "com/amazonaws/services/ec2/model/CreateVolumePermission", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVolumePermission{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGroup(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumePermission) SetGroup2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroup", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGroup()
func (jbobject *ServicesEc2ModelCreateVolumePermission) GetGroup() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroup", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CreateVolumePermission withGroup(java.lang.String)
func (jbobject *ServicesEc2ModelCreateVolumePermission) WithGroup2(a string) *ServicesEc2ModelCreateVolumePermission {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroup", "com/amazonaws/services/ec2/model/CreateVolumePermission", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCreateVolumePermission{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGroup(com.amazonaws.services.ec2.model.PermissionGroup)
func (jbobject *ServicesEc2ModelCreateVolumePermission) SetGroup(a ServicesEc2ModelPermissionGroupInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroup", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/PermissionGroup"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateVolumePermission withGroup(com.amazonaws.services.ec2.model.PermissionGroup)
func (jbobject *ServicesEc2ModelCreateVolumePermission) WithGroup(a ServicesEc2ModelPermissionGroupInterface) *ServicesEc2ModelCreateVolumePermission {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroup", "com/amazonaws/services/ec2/model/CreateVolumePermission", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PermissionGroup"))
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
	unique_x := &ServicesEc2ModelCreateVolumePermission{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateVolumePermission) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVolumePermission) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVolumePermission) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVolumePermission clone()
func (jbobject *ServicesEc2ModelCreateVolumePermission) Clone() *ServicesEc2ModelCreateVolumePermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVolumePermission")
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
	unique_x := &ServicesEc2ModelCreateVolumePermission{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateVolumePermission) Clone2() (*JavaLangObject, error) {
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


