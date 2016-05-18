package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPermissionGroupInterface interface {

	// public static com.amazonaws.services.ec2.model.PermissionGroup[] values()
	Values() []*ServicesEc2ModelPermissionGroup

	// public static com.amazonaws.services.ec2.model.PermissionGroup valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelPermissionGroup

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.PermissionGroup fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelPermissionGroup
}

type ServicesEc2ModelPermissionGroup struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.PermissionGroup[] values()
func (jbobject *ServicesEc2ModelPermissionGroup) Values() []*ServicesEc2ModelPermissionGroup {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PermissionGroup", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/PermissionGroup"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/PermissionGroup")
	dst := new([]*ServicesEc2ModelPermissionGroup)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.PermissionGroup valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelPermissionGroup) ValueOf(a string) *ServicesEc2ModelPermissionGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PermissionGroup", "valueOf", "com/amazonaws/services/ec2/model/PermissionGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPermissionGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPermissionGroup) ToString() string {
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

// public static com.amazonaws.services.ec2.model.PermissionGroup fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelPermissionGroup) FromValue(a string) *ServicesEc2ModelPermissionGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PermissionGroup", "fromValue", "com/amazonaws/services/ec2/model/PermissionGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPermissionGroup{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPermissionGroup) All() *ServicesEc2ModelPermissionGroup {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/PermissionGroup", "All", "com/amazonaws/services/ec2/model/PermissionGroup")
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
	unique_x := &ServicesEc2ModelPermissionGroup{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPermissionGroup) SetFieldAll(val ServicesEc2ModelPermissionGroupInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/PermissionGroup", "All", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


