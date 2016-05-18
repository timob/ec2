package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelExportTaskStateInterface interface {

	// public static com.amazonaws.services.ec2.model.ExportTaskState[] values()
	Values() []*ServicesEc2ModelExportTaskState

	// public static com.amazonaws.services.ec2.model.ExportTaskState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelExportTaskState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ExportTaskState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelExportTaskState
}

type ServicesEc2ModelExportTaskState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ExportTaskState[] values()
func (jbobject *ServicesEc2ModelExportTaskState) Values() []*ServicesEc2ModelExportTaskState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ExportTaskState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ExportTaskState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ExportTaskState")
	dst := new([]*ServicesEc2ModelExportTaskState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ExportTaskState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelExportTaskState) ValueOf(a string) *ServicesEc2ModelExportTaskState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ExportTaskState", "valueOf", "com/amazonaws/services/ec2/model/ExportTaskState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportTaskState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelExportTaskState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ExportTaskState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelExportTaskState) FromValue(a string) *ServicesEc2ModelExportTaskState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ExportTaskState", "fromValue", "com/amazonaws/services/ec2/model/ExportTaskState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExportTaskState) Active() *ServicesEc2ModelExportTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ExportTaskState", "Active", "com/amazonaws/services/ec2/model/ExportTaskState")
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
	unique_x := &ServicesEc2ModelExportTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExportTaskState) SetFieldActive(val ServicesEc2ModelExportTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ExportTaskState", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelExportTaskState) Cancelling() *ServicesEc2ModelExportTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ExportTaskState", "Cancelling", "com/amazonaws/services/ec2/model/ExportTaskState")
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
	unique_x := &ServicesEc2ModelExportTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExportTaskState) SetFieldCancelling(val ServicesEc2ModelExportTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ExportTaskState", "Cancelling", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelExportTaskState) Cancelled() *ServicesEc2ModelExportTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ExportTaskState", "Cancelled", "com/amazonaws/services/ec2/model/ExportTaskState")
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
	unique_x := &ServicesEc2ModelExportTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExportTaskState) SetFieldCancelled(val ServicesEc2ModelExportTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ExportTaskState", "Cancelled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelExportTaskState) Completed() *ServicesEc2ModelExportTaskState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ExportTaskState", "Completed", "com/amazonaws/services/ec2/model/ExportTaskState")
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
	unique_x := &ServicesEc2ModelExportTaskState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExportTaskState) SetFieldCompleted(val ServicesEc2ModelExportTaskStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ExportTaskState", "Completed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


