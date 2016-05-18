package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelStatusNameInterface interface {

	// public static com.amazonaws.services.ec2.model.StatusName[] values()
	Values() []*ServicesEc2ModelStatusName

	// public static com.amazonaws.services.ec2.model.StatusName valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelStatusName

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.StatusName fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelStatusName
}

type ServicesEc2ModelStatusName struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.StatusName[] values()
func (jbobject *ServicesEc2ModelStatusName) Values() []*ServicesEc2ModelStatusName {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/StatusName", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/StatusName"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/StatusName")
	dst := new([]*ServicesEc2ModelStatusName)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.StatusName valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelStatusName) ValueOf(a string) *ServicesEc2ModelStatusName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/StatusName", "valueOf", "com/amazonaws/services/ec2/model/StatusName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelStatusName{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelStatusName) ToString() string {
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

// public static com.amazonaws.services.ec2.model.StatusName fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelStatusName) FromValue(a string) *ServicesEc2ModelStatusName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/StatusName", "fromValue", "com/amazonaws/services/ec2/model/StatusName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelStatusName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatusName) Reachability() *ServicesEc2ModelStatusName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/StatusName", "Reachability", "com/amazonaws/services/ec2/model/StatusName")
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
	unique_x := &ServicesEc2ModelStatusName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatusName) SetFieldReachability(val ServicesEc2ModelStatusNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/StatusName", "Reachability", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


