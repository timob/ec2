package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelTenancyInterface interface {

	// public static com.amazonaws.services.ec2.model.Tenancy[] values()
	Values() []*ServicesEc2ModelTenancy

	// public static com.amazonaws.services.ec2.model.Tenancy valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelTenancy

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.Tenancy fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelTenancy
}

type ServicesEc2ModelTenancy struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.Tenancy[] values()
func (jbobject *ServicesEc2ModelTenancy) Values() []*ServicesEc2ModelTenancy {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/Tenancy", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/Tenancy"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/Tenancy")
	dst := new([]*ServicesEc2ModelTenancy)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.Tenancy valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelTenancy) ValueOf(a string) *ServicesEc2ModelTenancy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/Tenancy", "valueOf", "com/amazonaws/services/ec2/model/Tenancy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTenancy{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelTenancy) ToString() string {
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

// public static com.amazonaws.services.ec2.model.Tenancy fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelTenancy) FromValue(a string) *ServicesEc2ModelTenancy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/Tenancy", "fromValue", "com/amazonaws/services/ec2/model/Tenancy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTenancy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTenancy) Default() *ServicesEc2ModelTenancy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/Tenancy", "Default", "com/amazonaws/services/ec2/model/Tenancy")
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
	unique_x := &ServicesEc2ModelTenancy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTenancy) SetFieldDefault(val ServicesEc2ModelTenancyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/Tenancy", "Default", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelTenancy) Dedicated() *ServicesEc2ModelTenancy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/Tenancy", "Dedicated", "com/amazonaws/services/ec2/model/Tenancy")
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
	unique_x := &ServicesEc2ModelTenancy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTenancy) SetFieldDedicated(val ServicesEc2ModelTenancyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/Tenancy", "Dedicated", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelTenancy) Host() *ServicesEc2ModelTenancy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/Tenancy", "Host", "com/amazonaws/services/ec2/model/Tenancy")
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
	unique_x := &ServicesEc2ModelTenancy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTenancy) SetFieldHost(val ServicesEc2ModelTenancyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/Tenancy", "Host", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


