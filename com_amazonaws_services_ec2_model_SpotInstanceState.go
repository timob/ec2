package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSpotInstanceStateInterface interface {

	// public static com.amazonaws.services.ec2.model.SpotInstanceState[] values()
	Values() []*ServicesEc2ModelSpotInstanceState

	// public static com.amazonaws.services.ec2.model.SpotInstanceState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelSpotInstanceState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.SpotInstanceState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelSpotInstanceState
}

type ServicesEc2ModelSpotInstanceState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.SpotInstanceState[] values()
func (jbobject *ServicesEc2ModelSpotInstanceState) Values() []*ServicesEc2ModelSpotInstanceState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SpotInstanceState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/SpotInstanceState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/SpotInstanceState")
	dst := new([]*ServicesEc2ModelSpotInstanceState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.SpotInstanceState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceState) ValueOf(a string) *ServicesEc2ModelSpotInstanceState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SpotInstanceState", "valueOf", "com/amazonaws/services/ec2/model/SpotInstanceState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSpotInstanceState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.SpotInstanceState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceState) FromValue(a string) *ServicesEc2ModelSpotInstanceState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SpotInstanceState", "fromValue", "com/amazonaws/services/ec2/model/SpotInstanceState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSpotInstanceState) Open() *ServicesEc2ModelSpotInstanceState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Open", "com/amazonaws/services/ec2/model/SpotInstanceState")
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
	unique_x := &ServicesEc2ModelSpotInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSpotInstanceState) SetFieldOpen(val ServicesEc2ModelSpotInstanceStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Open", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSpotInstanceState) Active() *ServicesEc2ModelSpotInstanceState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Active", "com/amazonaws/services/ec2/model/SpotInstanceState")
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
	unique_x := &ServicesEc2ModelSpotInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSpotInstanceState) SetFieldActive(val ServicesEc2ModelSpotInstanceStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSpotInstanceState) Closed() *ServicesEc2ModelSpotInstanceState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Closed", "com/amazonaws/services/ec2/model/SpotInstanceState")
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
	unique_x := &ServicesEc2ModelSpotInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSpotInstanceState) SetFieldClosed(val ServicesEc2ModelSpotInstanceStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Closed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSpotInstanceState) Cancelled() *ServicesEc2ModelSpotInstanceState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Cancelled", "com/amazonaws/services/ec2/model/SpotInstanceState")
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
	unique_x := &ServicesEc2ModelSpotInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSpotInstanceState) SetFieldCancelled(val ServicesEc2ModelSpotInstanceStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Cancelled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSpotInstanceState) Failed() *ServicesEc2ModelSpotInstanceState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Failed", "com/amazonaws/services/ec2/model/SpotInstanceState")
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
	unique_x := &ServicesEc2ModelSpotInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSpotInstanceState) SetFieldFailed(val ServicesEc2ModelSpotInstanceStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SpotInstanceState", "Failed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


