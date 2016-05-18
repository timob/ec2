package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelBatchStateInterface interface {

	// public static com.amazonaws.services.ec2.model.BatchState[] values()
	Values() []*ServicesEc2ModelBatchState

	// public static com.amazonaws.services.ec2.model.BatchState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelBatchState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.BatchState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelBatchState
}

type ServicesEc2ModelBatchState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.BatchState[] values()
func (jbobject *ServicesEc2ModelBatchState) Values() []*ServicesEc2ModelBatchState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/BatchState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/BatchState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/BatchState")
	dst := new([]*ServicesEc2ModelBatchState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.BatchState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelBatchState) ValueOf(a string) *ServicesEc2ModelBatchState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/BatchState", "valueOf", "com/amazonaws/services/ec2/model/BatchState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBatchState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelBatchState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.BatchState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelBatchState) FromValue(a string) *ServicesEc2ModelBatchState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/BatchState", "fromValue", "com/amazonaws/services/ec2/model/BatchState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBatchState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBatchState) Submitted() *ServicesEc2ModelBatchState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BatchState", "Submitted", "com/amazonaws/services/ec2/model/BatchState")
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
	unique_x := &ServicesEc2ModelBatchState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBatchState) SetFieldSubmitted(val ServicesEc2ModelBatchStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BatchState", "Submitted", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBatchState) Active() *ServicesEc2ModelBatchState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BatchState", "Active", "com/amazonaws/services/ec2/model/BatchState")
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
	unique_x := &ServicesEc2ModelBatchState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBatchState) SetFieldActive(val ServicesEc2ModelBatchStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BatchState", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBatchState) Cancelled() *ServicesEc2ModelBatchState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BatchState", "Cancelled", "com/amazonaws/services/ec2/model/BatchState")
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
	unique_x := &ServicesEc2ModelBatchState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBatchState) SetFieldCancelled(val ServicesEc2ModelBatchStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BatchState", "Cancelled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBatchState) Failed() *ServicesEc2ModelBatchState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BatchState", "Failed", "com/amazonaws/services/ec2/model/BatchState")
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
	unique_x := &ServicesEc2ModelBatchState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBatchState) SetFieldFailed(val ServicesEc2ModelBatchStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BatchState", "Failed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBatchState) Cancelled_running() *ServicesEc2ModelBatchState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BatchState", "Cancelled_running", "com/amazonaws/services/ec2/model/BatchState")
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
	unique_x := &ServicesEc2ModelBatchState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBatchState) SetFieldCancelled_running(val ServicesEc2ModelBatchStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BatchState", "Cancelled_running", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBatchState) Cancelled_terminating() *ServicesEc2ModelBatchState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BatchState", "Cancelled_terminating", "com/amazonaws/services/ec2/model/BatchState")
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
	unique_x := &ServicesEc2ModelBatchState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBatchState) SetFieldCancelled_terminating(val ServicesEc2ModelBatchStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BatchState", "Cancelled_terminating", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBatchState) Modifying() *ServicesEc2ModelBatchState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BatchState", "Modifying", "com/amazonaws/services/ec2/model/BatchState")
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
	unique_x := &ServicesEc2ModelBatchState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBatchState) SetFieldModifying(val ServicesEc2ModelBatchStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BatchState", "Modifying", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


