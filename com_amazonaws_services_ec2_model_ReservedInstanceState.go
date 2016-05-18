package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReservedInstanceStateInterface interface {

	// public static com.amazonaws.services.ec2.model.ReservedInstanceState[] values()
	Values() []*ServicesEc2ModelReservedInstanceState

	// public static com.amazonaws.services.ec2.model.ReservedInstanceState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelReservedInstanceState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ReservedInstanceState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelReservedInstanceState
}

type ServicesEc2ModelReservedInstanceState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ReservedInstanceState[] values()
func (jbobject *ServicesEc2ModelReservedInstanceState) Values() []*ServicesEc2ModelReservedInstanceState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ReservedInstanceState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ReservedInstanceState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ReservedInstanceState")
	dst := new([]*ServicesEc2ModelReservedInstanceState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ReservedInstanceState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstanceState) ValueOf(a string) *ServicesEc2ModelReservedInstanceState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ReservedInstanceState", "valueOf", "com/amazonaws/services/ec2/model/ReservedInstanceState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReservedInstanceState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ReservedInstanceState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstanceState) FromValue(a string) *ServicesEc2ModelReservedInstanceState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ReservedInstanceState", "fromValue", "com/amazonaws/services/ec2/model/ReservedInstanceState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReservedInstanceState) PaymentPending() *ServicesEc2ModelReservedInstanceState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReservedInstanceState", "PaymentPending", "com/amazonaws/services/ec2/model/ReservedInstanceState")
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
	unique_x := &ServicesEc2ModelReservedInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReservedInstanceState) SetFieldPaymentPending(val ServicesEc2ModelReservedInstanceStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReservedInstanceState", "PaymentPending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReservedInstanceState) Active() *ServicesEc2ModelReservedInstanceState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReservedInstanceState", "Active", "com/amazonaws/services/ec2/model/ReservedInstanceState")
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
	unique_x := &ServicesEc2ModelReservedInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReservedInstanceState) SetFieldActive(val ServicesEc2ModelReservedInstanceStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReservedInstanceState", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReservedInstanceState) PaymentFailed() *ServicesEc2ModelReservedInstanceState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReservedInstanceState", "PaymentFailed", "com/amazonaws/services/ec2/model/ReservedInstanceState")
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
	unique_x := &ServicesEc2ModelReservedInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReservedInstanceState) SetFieldPaymentFailed(val ServicesEc2ModelReservedInstanceStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReservedInstanceState", "PaymentFailed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReservedInstanceState) Retired() *ServicesEc2ModelReservedInstanceState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReservedInstanceState", "Retired", "com/amazonaws/services/ec2/model/ReservedInstanceState")
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
	unique_x := &ServicesEc2ModelReservedInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReservedInstanceState) SetFieldRetired(val ServicesEc2ModelReservedInstanceStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReservedInstanceState", "Retired", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


