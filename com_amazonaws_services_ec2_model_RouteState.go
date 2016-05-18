package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRouteStateInterface interface {

	// public static com.amazonaws.services.ec2.model.RouteState[] values()
	Values() []*ServicesEc2ModelRouteState

	// public static com.amazonaws.services.ec2.model.RouteState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelRouteState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.RouteState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelRouteState
}

type ServicesEc2ModelRouteState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.RouteState[] values()
func (jbobject *ServicesEc2ModelRouteState) Values() []*ServicesEc2ModelRouteState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RouteState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/RouteState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/RouteState")
	dst := new([]*ServicesEc2ModelRouteState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.RouteState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelRouteState) ValueOf(a string) *ServicesEc2ModelRouteState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RouteState", "valueOf", "com/amazonaws/services/ec2/model/RouteState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRouteState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRouteState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.RouteState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelRouteState) FromValue(a string) *ServicesEc2ModelRouteState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RouteState", "fromValue", "com/amazonaws/services/ec2/model/RouteState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRouteState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRouteState) Active() *ServicesEc2ModelRouteState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RouteState", "Active", "com/amazonaws/services/ec2/model/RouteState")
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
	unique_x := &ServicesEc2ModelRouteState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRouteState) SetFieldActive(val ServicesEc2ModelRouteStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RouteState", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelRouteState) Blackhole() *ServicesEc2ModelRouteState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RouteState", "Blackhole", "com/amazonaws/services/ec2/model/RouteState")
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
	unique_x := &ServicesEc2ModelRouteState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRouteState) SetFieldBlackhole(val ServicesEc2ModelRouteStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RouteState", "Blackhole", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


