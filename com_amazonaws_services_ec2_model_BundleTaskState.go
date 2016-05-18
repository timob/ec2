package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelBundleTaskStateInterface interface {

	// public static com.amazonaws.services.ec2.model.BundleTaskState[] values()
	Values() []*ServicesEc2ModelBundleTaskState

	// public static com.amazonaws.services.ec2.model.BundleTaskState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelBundleTaskState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.BundleTaskState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelBundleTaskState
}

type ServicesEc2ModelBundleTaskState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.BundleTaskState[] values()
func (jbobject *ServicesEc2ModelBundleTaskState) Values() []*ServicesEc2ModelBundleTaskState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/BundleTaskState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/BundleTaskState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/BundleTaskState")
	dst := new([]*ServicesEc2ModelBundleTaskState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.BundleTaskState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTaskState) ValueOf(a string) *ServicesEc2ModelBundleTaskState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/BundleTaskState", "valueOf", "com/amazonaws/services/ec2/model/BundleTaskState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBundleTaskState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelBundleTaskState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.BundleTaskState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTaskState) FromValue(a string) *ServicesEc2ModelBundleTaskState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/BundleTaskState", "fromValue", "com/amazonaws/services/ec2/model/BundleTaskState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBundleTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBundleTaskState) Pending() *ServicesEc2ModelBundleTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Pending", "com/amazonaws/services/ec2/model/BundleTaskState")
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
	unique_x := &ServicesEc2ModelBundleTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBundleTaskState) SetFieldPending(val ServicesEc2ModelBundleTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBundleTaskState) WaitingForShutdown() *ServicesEc2ModelBundleTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "WaitingForShutdown", "com/amazonaws/services/ec2/model/BundleTaskState")
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
	unique_x := &ServicesEc2ModelBundleTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBundleTaskState) SetFieldWaitingForShutdown(val ServicesEc2ModelBundleTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "WaitingForShutdown", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBundleTaskState) Bundling() *ServicesEc2ModelBundleTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Bundling", "com/amazonaws/services/ec2/model/BundleTaskState")
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
	unique_x := &ServicesEc2ModelBundleTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBundleTaskState) SetFieldBundling(val ServicesEc2ModelBundleTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Bundling", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBundleTaskState) Storing() *ServicesEc2ModelBundleTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Storing", "com/amazonaws/services/ec2/model/BundleTaskState")
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
	unique_x := &ServicesEc2ModelBundleTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBundleTaskState) SetFieldStoring(val ServicesEc2ModelBundleTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Storing", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBundleTaskState) Cancelling() *ServicesEc2ModelBundleTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Cancelling", "com/amazonaws/services/ec2/model/BundleTaskState")
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
	unique_x := &ServicesEc2ModelBundleTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBundleTaskState) SetFieldCancelling(val ServicesEc2ModelBundleTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Cancelling", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBundleTaskState) Complete() *ServicesEc2ModelBundleTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Complete", "com/amazonaws/services/ec2/model/BundleTaskState")
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
	unique_x := &ServicesEc2ModelBundleTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBundleTaskState) SetFieldComplete(val ServicesEc2ModelBundleTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Complete", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelBundleTaskState) Failed() *ServicesEc2ModelBundleTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Failed", "com/amazonaws/services/ec2/model/BundleTaskState")
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
	unique_x := &ServicesEc2ModelBundleTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelBundleTaskState) SetFieldFailed(val ServicesEc2ModelBundleTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/BundleTaskState", "Failed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


