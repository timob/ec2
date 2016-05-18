package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSnapshotStateInterface interface {

	// public static com.amazonaws.services.ec2.model.SnapshotState[] values()
	Values() []*ServicesEc2ModelSnapshotState

	// public static com.amazonaws.services.ec2.model.SnapshotState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelSnapshotState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.SnapshotState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelSnapshotState
}

type ServicesEc2ModelSnapshotState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.SnapshotState[] values()
func (jbobject *ServicesEc2ModelSnapshotState) Values() []*ServicesEc2ModelSnapshotState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SnapshotState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/SnapshotState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/SnapshotState")
	dst := new([]*ServicesEc2ModelSnapshotState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.SnapshotState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotState) ValueOf(a string) *ServicesEc2ModelSnapshotState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SnapshotState", "valueOf", "com/amazonaws/services/ec2/model/SnapshotState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSnapshotState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.SnapshotState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotState) FromValue(a string) *ServicesEc2ModelSnapshotState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SnapshotState", "fromValue", "com/amazonaws/services/ec2/model/SnapshotState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSnapshotState) Pending() *ServicesEc2ModelSnapshotState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SnapshotState", "Pending", "com/amazonaws/services/ec2/model/SnapshotState")
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
	unique_x := &ServicesEc2ModelSnapshotState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSnapshotState) SetFieldPending(val ServicesEc2ModelSnapshotStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SnapshotState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSnapshotState) Completed() *ServicesEc2ModelSnapshotState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SnapshotState", "Completed", "com/amazonaws/services/ec2/model/SnapshotState")
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
	unique_x := &ServicesEc2ModelSnapshotState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSnapshotState) SetFieldCompleted(val ServicesEc2ModelSnapshotStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SnapshotState", "Completed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSnapshotState) Error() *ServicesEc2ModelSnapshotState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SnapshotState", "Error", "com/amazonaws/services/ec2/model/SnapshotState")
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
	unique_x := &ServicesEc2ModelSnapshotState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSnapshotState) SetFieldError(val ServicesEc2ModelSnapshotStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SnapshotState", "Error", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


