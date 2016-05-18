package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelHostTenancyInterface interface {

	// public static com.amazonaws.services.ec2.model.HostTenancy[] values()
	Values() []*ServicesEc2ModelHostTenancy

	// public static com.amazonaws.services.ec2.model.HostTenancy valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelHostTenancy

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.HostTenancy fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelHostTenancy
}

type ServicesEc2ModelHostTenancy struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.HostTenancy[] values()
func (jbobject *ServicesEc2ModelHostTenancy) Values() []*ServicesEc2ModelHostTenancy {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/HostTenancy", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/HostTenancy"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/HostTenancy")
	dst := new([]*ServicesEc2ModelHostTenancy)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.HostTenancy valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelHostTenancy) ValueOf(a string) *ServicesEc2ModelHostTenancy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/HostTenancy", "valueOf", "com/amazonaws/services/ec2/model/HostTenancy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHostTenancy{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelHostTenancy) ToString() string {
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

// public static com.amazonaws.services.ec2.model.HostTenancy fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelHostTenancy) FromValue(a string) *ServicesEc2ModelHostTenancy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/HostTenancy", "fromValue", "com/amazonaws/services/ec2/model/HostTenancy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHostTenancy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelHostTenancy) Dedicated() *ServicesEc2ModelHostTenancy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/HostTenancy", "Dedicated", "com/amazonaws/services/ec2/model/HostTenancy")
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
	unique_x := &ServicesEc2ModelHostTenancy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelHostTenancy) SetFieldDedicated(val ServicesEc2ModelHostTenancyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/HostTenancy", "Dedicated", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelHostTenancy) Host() *ServicesEc2ModelHostTenancy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/HostTenancy", "Host", "com/amazonaws/services/ec2/model/HostTenancy")
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
	unique_x := &ServicesEc2ModelHostTenancy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelHostTenancy) SetFieldHost(val ServicesEc2ModelHostTenancyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/HostTenancy", "Host", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


