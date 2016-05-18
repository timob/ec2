package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelConversionTaskStateInterface interface {

	// public static com.amazonaws.services.ec2.model.ConversionTaskState[] values()
	Values() []*ServicesEc2ModelConversionTaskState

	// public static com.amazonaws.services.ec2.model.ConversionTaskState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelConversionTaskState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ConversionTaskState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelConversionTaskState
}

type ServicesEc2ModelConversionTaskState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ConversionTaskState[] values()
func (jbobject *ServicesEc2ModelConversionTaskState) Values() []*ServicesEc2ModelConversionTaskState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ConversionTaskState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ConversionTaskState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ConversionTaskState")
	dst := new([]*ServicesEc2ModelConversionTaskState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ConversionTaskState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTaskState) ValueOf(a string) *ServicesEc2ModelConversionTaskState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ConversionTaskState", "valueOf", "com/amazonaws/services/ec2/model/ConversionTaskState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelConversionTaskState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelConversionTaskState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ConversionTaskState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelConversionTaskState) FromValue(a string) *ServicesEc2ModelConversionTaskState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ConversionTaskState", "fromValue", "com/amazonaws/services/ec2/model/ConversionTaskState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelConversionTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelConversionTaskState) Active() *ServicesEc2ModelConversionTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ConversionTaskState", "Active", "com/amazonaws/services/ec2/model/ConversionTaskState")
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
	unique_x := &ServicesEc2ModelConversionTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelConversionTaskState) SetFieldActive(val ServicesEc2ModelConversionTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ConversionTaskState", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelConversionTaskState) Cancelling() *ServicesEc2ModelConversionTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ConversionTaskState", "Cancelling", "com/amazonaws/services/ec2/model/ConversionTaskState")
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
	unique_x := &ServicesEc2ModelConversionTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelConversionTaskState) SetFieldCancelling(val ServicesEc2ModelConversionTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ConversionTaskState", "Cancelling", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelConversionTaskState) Cancelled() *ServicesEc2ModelConversionTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ConversionTaskState", "Cancelled", "com/amazonaws/services/ec2/model/ConversionTaskState")
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
	unique_x := &ServicesEc2ModelConversionTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelConversionTaskState) SetFieldCancelled(val ServicesEc2ModelConversionTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ConversionTaskState", "Cancelled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelConversionTaskState) Completed() *ServicesEc2ModelConversionTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ConversionTaskState", "Completed", "com/amazonaws/services/ec2/model/ConversionTaskState")
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
	unique_x := &ServicesEc2ModelConversionTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelConversionTaskState) SetFieldCompleted(val ServicesEc2ModelConversionTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ConversionTaskState", "Completed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


