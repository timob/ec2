package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelSpotInstanceRequestStateInterface interface {

	// public static com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState[] values()
	Values() []*ServicesEc2ModelCancelSpotInstanceRequestState

	// public static com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelCancelSpotInstanceRequestState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelCancelSpotInstanceRequestState
}

type ServicesEc2ModelCancelSpotInstanceRequestState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState[] values()
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) Values() []*ServicesEc2ModelCancelSpotInstanceRequestState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState")
	dst := new([]*ServicesEc2ModelCancelSpotInstanceRequestState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) ValueOf(a string) *ServicesEc2ModelCancelSpotInstanceRequestState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "valueOf", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.CancelSpotInstanceRequestState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) FromValue(a string) *ServicesEc2ModelCancelSpotInstanceRequestState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "fromValue", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) Active() *ServicesEc2ModelCancelSpotInstanceRequestState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Active", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState")
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) SetFieldActive(val ServicesEc2ModelCancelSpotInstanceRequestStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) Open() *ServicesEc2ModelCancelSpotInstanceRequestState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Open", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState")
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) SetFieldOpen(val ServicesEc2ModelCancelSpotInstanceRequestStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Open", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) Closed() *ServicesEc2ModelCancelSpotInstanceRequestState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Closed", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState")
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) SetFieldClosed(val ServicesEc2ModelCancelSpotInstanceRequestStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Closed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) Cancelled() *ServicesEc2ModelCancelSpotInstanceRequestState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Cancelled", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState")
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) SetFieldCancelled(val ServicesEc2ModelCancelSpotInstanceRequestStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Cancelled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) Completed() *ServicesEc2ModelCancelSpotInstanceRequestState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Completed", "com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState")
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
	unique_x := &ServicesEc2ModelCancelSpotInstanceRequestState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelSpotInstanceRequestState) SetFieldCompleted(val ServicesEc2ModelCancelSpotInstanceRequestStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CancelSpotInstanceRequestState", "Completed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


