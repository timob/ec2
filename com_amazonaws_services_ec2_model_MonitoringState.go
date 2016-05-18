package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelMonitoringStateInterface interface {

	// public static com.amazonaws.services.ec2.model.MonitoringState[] values()
	Values() []*ServicesEc2ModelMonitoringState

	// public static com.amazonaws.services.ec2.model.MonitoringState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelMonitoringState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.MonitoringState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelMonitoringState
}

type ServicesEc2ModelMonitoringState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.MonitoringState[] values()
func (jbobject *ServicesEc2ModelMonitoringState) Values() []*ServicesEc2ModelMonitoringState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/MonitoringState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/MonitoringState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/MonitoringState")
	dst := new([]*ServicesEc2ModelMonitoringState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.MonitoringState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelMonitoringState) ValueOf(a string) *ServicesEc2ModelMonitoringState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/MonitoringState", "valueOf", "com/amazonaws/services/ec2/model/MonitoringState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelMonitoringState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelMonitoringState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.MonitoringState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelMonitoringState) FromValue(a string) *ServicesEc2ModelMonitoringState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/MonitoringState", "fromValue", "com/amazonaws/services/ec2/model/MonitoringState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelMonitoringState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelMonitoringState) Disabled() *ServicesEc2ModelMonitoringState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/MonitoringState", "Disabled", "com/amazonaws/services/ec2/model/MonitoringState")
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
	unique_x := &ServicesEc2ModelMonitoringState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelMonitoringState) SetFieldDisabled(val ServicesEc2ModelMonitoringStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/MonitoringState", "Disabled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelMonitoringState) Disabling() *ServicesEc2ModelMonitoringState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/MonitoringState", "Disabling", "com/amazonaws/services/ec2/model/MonitoringState")
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
	unique_x := &ServicesEc2ModelMonitoringState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelMonitoringState) SetFieldDisabling(val ServicesEc2ModelMonitoringStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/MonitoringState", "Disabling", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelMonitoringState) Enabled() *ServicesEc2ModelMonitoringState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/MonitoringState", "Enabled", "com/amazonaws/services/ec2/model/MonitoringState")
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
	unique_x := &ServicesEc2ModelMonitoringState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelMonitoringState) SetFieldEnabled(val ServicesEc2ModelMonitoringStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/MonitoringState", "Enabled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelMonitoringState) Pending() *ServicesEc2ModelMonitoringState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/MonitoringState", "Pending", "com/amazonaws/services/ec2/model/MonitoringState")
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
	unique_x := &ServicesEc2ModelMonitoringState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelMonitoringState) SetFieldPending(val ServicesEc2ModelMonitoringStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/MonitoringState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


