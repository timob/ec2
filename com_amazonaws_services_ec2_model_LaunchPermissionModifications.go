package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelLaunchPermissionModificationsInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.LaunchPermission> getAdd()
	GetAdd() []*ServicesEc2ModelLaunchPermission

	// public void setAdd(java.util.Collection<com.amazonaws.services.ec2.model.LaunchPermission>)
	SetAdd(a []*ServicesEc2ModelLaunchPermission) 

	// public com.amazonaws.services.ec2.model.LaunchPermissionModifications withAdd(com.amazonaws.services.ec2.model.LaunchPermission...)
	WithAdd(a ...*ServicesEc2ModelLaunchPermission) *ServicesEc2ModelLaunchPermissionModifications

	// public com.amazonaws.services.ec2.model.LaunchPermissionModifications withAdd(java.util.Collection<com.amazonaws.services.ec2.model.LaunchPermission>)
	WithAdd2(a []*ServicesEc2ModelLaunchPermission) *ServicesEc2ModelLaunchPermissionModifications

	// public java.util.List<com.amazonaws.services.ec2.model.LaunchPermission> getRemove()
	GetRemove() []*ServicesEc2ModelLaunchPermission

	// public void setRemove(java.util.Collection<com.amazonaws.services.ec2.model.LaunchPermission>)
	SetRemove(a []*ServicesEc2ModelLaunchPermission) 

	// public com.amazonaws.services.ec2.model.LaunchPermissionModifications withRemove(com.amazonaws.services.ec2.model.LaunchPermission...)
	WithRemove(a ...*ServicesEc2ModelLaunchPermission) *ServicesEc2ModelLaunchPermissionModifications

	// public com.amazonaws.services.ec2.model.LaunchPermissionModifications withRemove(java.util.Collection<com.amazonaws.services.ec2.model.LaunchPermission>)
	WithRemove2(a []*ServicesEc2ModelLaunchPermission) *ServicesEc2ModelLaunchPermissionModifications

	// public com.amazonaws.services.ec2.model.LaunchPermissionModifications clone()
	Clone() *ServicesEc2ModelLaunchPermissionModifications
}

type ServicesEc2ModelLaunchPermissionModifications struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.LaunchPermissionModifications()
func NewServicesEc2ModelLaunchPermissionModifications() (*ServicesEc2ModelLaunchPermissionModifications) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/LaunchPermissionModifications")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelLaunchPermissionModifications{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.LaunchPermission> getAdd()
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) GetAdd() []*ServicesEc2ModelLaunchPermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAdd", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelLaunchPermission)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAdd(java.util.Collection<com.amazonaws.services.ec2.model.LaunchPermission>)
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) SetAdd(a []*ServicesEc2ModelLaunchPermission)  {
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

// public com.amazonaws.services.ec2.model.LaunchPermissionModifications withAdd(com.amazonaws.services.ec2.model.LaunchPermission...)
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) WithAdd(a ...*ServicesEc2ModelLaunchPermission) *ServicesEc2ModelLaunchPermissionModifications {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/LaunchPermission")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAdd", "com/amazonaws/services/ec2/model/LaunchPermissionModifications", conv_a.Value().Cast("com/amazonaws/services/ec2/model/LaunchPermission"))
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
	unique_x := &ServicesEc2ModelLaunchPermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.LaunchPermissionModifications withAdd(java.util.Collection<com.amazonaws.services.ec2.model.LaunchPermission>)
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) WithAdd2(a []*ServicesEc2ModelLaunchPermission) *ServicesEc2ModelLaunchPermissionModifications {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAdd", "com/amazonaws/services/ec2/model/LaunchPermissionModifications", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelLaunchPermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.LaunchPermission> getRemove()
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) GetRemove() []*ServicesEc2ModelLaunchPermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRemove", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelLaunchPermission)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRemove(java.util.Collection<com.amazonaws.services.ec2.model.LaunchPermission>)
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) SetRemove(a []*ServicesEc2ModelLaunchPermission)  {
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

// public com.amazonaws.services.ec2.model.LaunchPermissionModifications withRemove(com.amazonaws.services.ec2.model.LaunchPermission...)
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) WithRemove(a ...*ServicesEc2ModelLaunchPermission) *ServicesEc2ModelLaunchPermissionModifications {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/LaunchPermission")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRemove", "com/amazonaws/services/ec2/model/LaunchPermissionModifications", conv_a.Value().Cast("com/amazonaws/services/ec2/model/LaunchPermission"))
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
	unique_x := &ServicesEc2ModelLaunchPermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.LaunchPermissionModifications withRemove(java.util.Collection<com.amazonaws.services.ec2.model.LaunchPermission>)
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) WithRemove2(a []*ServicesEc2ModelLaunchPermission) *ServicesEc2ModelLaunchPermissionModifications {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRemove", "com/amazonaws/services/ec2/model/LaunchPermissionModifications", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelLaunchPermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) ToString() string {
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
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.LaunchPermissionModifications clone()
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) Clone() *ServicesEc2ModelLaunchPermissionModifications {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/LaunchPermissionModifications")
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
	unique_x := &ServicesEc2ModelLaunchPermissionModifications{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelLaunchPermissionModifications) Clone2() (*JavaLangObject, error) {
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


