package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpnStaticRouteSourceInterface interface {

	// public static com.amazonaws.services.ec2.model.VpnStaticRouteSource[] values()
	Values() []*ServicesEc2ModelVpnStaticRouteSource

	// public static com.amazonaws.services.ec2.model.VpnStaticRouteSource valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVpnStaticRouteSource

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VpnStaticRouteSource fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVpnStaticRouteSource
}

type ServicesEc2ModelVpnStaticRouteSource struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VpnStaticRouteSource[] values()
func (jbobject *ServicesEc2ModelVpnStaticRouteSource) Values() []*ServicesEc2ModelVpnStaticRouteSource {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpnStaticRouteSource", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VpnStaticRouteSource"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VpnStaticRouteSource")
	dst := new([]*ServicesEc2ModelVpnStaticRouteSource)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VpnStaticRouteSource valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVpnStaticRouteSource) ValueOf(a string) *ServicesEc2ModelVpnStaticRouteSource {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpnStaticRouteSource", "valueOf", "com/amazonaws/services/ec2/model/VpnStaticRouteSource", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnStaticRouteSource{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpnStaticRouteSource) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VpnStaticRouteSource fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVpnStaticRouteSource) FromValue(a string) *ServicesEc2ModelVpnStaticRouteSource {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpnStaticRouteSource", "fromValue", "com/amazonaws/services/ec2/model/VpnStaticRouteSource", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnStaticRouteSource{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpnStaticRouteSource) Static() *ServicesEc2ModelVpnStaticRouteSource {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpnStaticRouteSource", "Static", "com/amazonaws/services/ec2/model/VpnStaticRouteSource")
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
	unique_x := &ServicesEc2ModelVpnStaticRouteSource{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpnStaticRouteSource) SetFieldStatic(val ServicesEc2ModelVpnStaticRouteSourceInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpnStaticRouteSource", "Static", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


