package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateVolumePermissionModificationsInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.CreateVolumePermission> getAdd()
	GetAdd() []*ServicesEc2ModelCreateVolumePermission

	// public void setAdd(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
	SetAdd(a []*ServicesEc2ModelCreateVolumePermission) 

	// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications withAdd(com.amazonaws.services.ec2.model.CreateVolumePermission...)
	WithAdd(a ...*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelCreateVolumePermissionModifications

	// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications withAdd(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
	WithAdd2(a []*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelCreateVolumePermissionModifications

	// public java.util.List<com.amazonaws.services.ec2.model.CreateVolumePermission> getRemove()
	GetRemove() []*ServicesEc2ModelCreateVolumePermission

	// public void setRemove(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
	SetRemove(a []*ServicesEc2ModelCreateVolumePermission) 

	// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications withRemove(com.amazonaws.services.ec2.model.CreateVolumePermission...)
	WithRemove(a ...*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelCreateVolumePermissionModifications

	// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications withRemove(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
	WithRemove2(a []*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelCreateVolumePermissionModifications

	// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications clone()
	Clone() *ServicesEc2ModelCreateVolumePermissionModifications
}

type ServicesEc2ModelCreateVolumePermissionModifications struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications()
func NewServicesEc2ModelCreateVolumePermissionModifications() (*ServicesEc2ModelCreateVolumePermissionModifications) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateVolumePermissionModifications")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateVolumePermissionModifications{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.CreateVolumePermission> getAdd()
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) GetAdd() []*ServicesEc2ModelCreateVolumePermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAdd", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelCreateVolumePermission)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAdd(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) SetAdd(a []*ServicesEc2ModelCreateVolumePermission)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAdd", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications withAdd(com.amazonaws.services.ec2.model.CreateVolumePermission...)
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) WithAdd(a ...*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelCreateVolumePermissionModifications {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/CreateVolumePermission")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAdd", "com/amazonaws/services/ec2/model/CreateVolumePermissionModifications", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVolumePermission"))
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
	unique_x := &ServicesEc2ModelCreateVolumePermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications withAdd(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) WithAdd2(a []*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelCreateVolumePermissionModifications {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAdd", "com/amazonaws/services/ec2/model/CreateVolumePermissionModifications", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateVolumePermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.CreateVolumePermission> getRemove()
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) GetRemove() []*ServicesEc2ModelCreateVolumePermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRemove", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelCreateVolumePermission)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRemove(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) SetRemove(a []*ServicesEc2ModelCreateVolumePermission)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRemove", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications withRemove(com.amazonaws.services.ec2.model.CreateVolumePermission...)
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) WithRemove(a ...*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelCreateVolumePermissionModifications {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/CreateVolumePermission")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRemove", "com/amazonaws/services/ec2/model/CreateVolumePermissionModifications", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVolumePermission"))
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
	unique_x := &ServicesEc2ModelCreateVolumePermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications withRemove(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) WithRemove2(a []*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelCreateVolumePermissionModifications {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRemove", "com/amazonaws/services/ec2/model/CreateVolumePermissionModifications", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelCreateVolumePermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateVolumePermissionModifications clone()
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) Clone() *ServicesEc2ModelCreateVolumePermissionModifications {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateVolumePermissionModifications")
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
	unique_x := &ServicesEc2ModelCreateVolumePermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateVolumePermissionModifications) Clone2() (*JavaLangObject, error) {
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


