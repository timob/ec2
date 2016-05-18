package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRouteOriginInterface interface {

	// public static com.amazonaws.services.ec2.model.RouteOrigin[] values()
	Values() []*ServicesEc2ModelRouteOrigin

	// public static com.amazonaws.services.ec2.model.RouteOrigin valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelRouteOrigin

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.RouteOrigin fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelRouteOrigin
}

type ServicesEc2ModelRouteOrigin struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.RouteOrigin[] values()
func (jbobject *ServicesEc2ModelRouteOrigin) Values() []*ServicesEc2ModelRouteOrigin {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RouteOrigin", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/RouteOrigin"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/RouteOrigin")
	dst := new([]*ServicesEc2ModelRouteOrigin)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.RouteOrigin valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelRouteOrigin) ValueOf(a string) *ServicesEc2ModelRouteOrigin {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RouteOrigin", "valueOf", "com/amazonaws/services/ec2/model/RouteOrigin", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRouteOrigin{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRouteOrigin) ToString() string {
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

// public static com.amazonaws.services.ec2.model.RouteOrigin fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelRouteOrigin) FromValue(a string) *ServicesEc2ModelRouteOrigin {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RouteOrigin", "fromValue", "com/amazonaws/services/ec2/model/RouteOrigin", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRouteOrigin{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRouteOrigin) CreateRouteTable() *ServicesEc2ModelRouteOrigin {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RouteOrigin", "CreateRouteTable", "com/amazonaws/services/ec2/model/RouteOrigin")
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
	unique_x := &ServicesEc2ModelRouteOrigin{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRouteOrigin) SetFieldCreateRouteTable(val ServicesEc2ModelRouteOriginInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RouteOrigin", "CreateRouteTable", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelRouteOrigin) CreateRoute() *ServicesEc2ModelRouteOrigin {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RouteOrigin", "CreateRoute", "com/amazonaws/services/ec2/model/RouteOrigin")
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
	unique_x := &ServicesEc2ModelRouteOrigin{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRouteOrigin) SetFieldCreateRoute(val ServicesEc2ModelRouteOriginInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RouteOrigin", "CreateRoute", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelRouteOrigin) EnableVgwRoutePropagation() *ServicesEc2ModelRouteOrigin {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RouteOrigin", "EnableVgwRoutePropagation", "com/amazonaws/services/ec2/model/RouteOrigin")
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
	unique_x := &ServicesEc2ModelRouteOrigin{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRouteOrigin) SetFieldEnableVgwRoutePropagation(val ServicesEc2ModelRouteOriginInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RouteOrigin", "EnableVgwRoutePropagation", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


