package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceStateNameInterface interface {

	// public static com.amazonaws.services.ec2.model.InstanceStateName[] values()
	Values() []*ServicesEc2ModelInstanceStateName

	// public static com.amazonaws.services.ec2.model.InstanceStateName valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelInstanceStateName

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.InstanceStateName fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelInstanceStateName
}

type ServicesEc2ModelInstanceStateName struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.InstanceStateName[] values()
func (jbobject *ServicesEc2ModelInstanceStateName) Values() []*ServicesEc2ModelInstanceStateName {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceStateName", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/InstanceStateName"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/InstanceStateName")
	dst := new([]*ServicesEc2ModelInstanceStateName)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.InstanceStateName valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStateName) ValueOf(a string) *ServicesEc2ModelInstanceStateName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceStateName", "valueOf", "com/amazonaws/services/ec2/model/InstanceStateName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStateName{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceStateName) ToString() string {
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

// public static com.amazonaws.services.ec2.model.InstanceStateName fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStateName) FromValue(a string) *ServicesEc2ModelInstanceStateName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceStateName", "fromValue", "com/amazonaws/services/ec2/model/InstanceStateName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStateName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceStateName) Pending() *ServicesEc2ModelInstanceStateName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Pending", "com/amazonaws/services/ec2/model/InstanceStateName")
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
	unique_x := &ServicesEc2ModelInstanceStateName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceStateName) SetFieldPending(val ServicesEc2ModelInstanceStateNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceStateName) Running() *ServicesEc2ModelInstanceStateName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Running", "com/amazonaws/services/ec2/model/InstanceStateName")
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
	unique_x := &ServicesEc2ModelInstanceStateName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceStateName) SetFieldRunning(val ServicesEc2ModelInstanceStateNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Running", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceStateName) ShuttingDown() *ServicesEc2ModelInstanceStateName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "ShuttingDown", "com/amazonaws/services/ec2/model/InstanceStateName")
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
	unique_x := &ServicesEc2ModelInstanceStateName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceStateName) SetFieldShuttingDown(val ServicesEc2ModelInstanceStateNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "ShuttingDown", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceStateName) Terminated() *ServicesEc2ModelInstanceStateName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Terminated", "com/amazonaws/services/ec2/model/InstanceStateName")
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
	unique_x := &ServicesEc2ModelInstanceStateName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceStateName) SetFieldTerminated(val ServicesEc2ModelInstanceStateNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Terminated", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceStateName) Stopping() *ServicesEc2ModelInstanceStateName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Stopping", "com/amazonaws/services/ec2/model/InstanceStateName")
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
	unique_x := &ServicesEc2ModelInstanceStateName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceStateName) SetFieldStopping(val ServicesEc2ModelInstanceStateNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Stopping", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceStateName) Stopped() *ServicesEc2ModelInstanceStateName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Stopped", "com/amazonaws/services/ec2/model/InstanceStateName")
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
	unique_x := &ServicesEc2ModelInstanceStateName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceStateName) SetFieldStopped(val ServicesEc2ModelInstanceStateNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceStateName", "Stopped", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


