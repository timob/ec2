package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelShutdownBehaviorInterface interface {

	// public static com.amazonaws.services.ec2.model.ShutdownBehavior[] values()
	Values() []*ServicesEc2ModelShutdownBehavior

	// public static com.amazonaws.services.ec2.model.ShutdownBehavior valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelShutdownBehavior

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ShutdownBehavior fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelShutdownBehavior
}

type ServicesEc2ModelShutdownBehavior struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ShutdownBehavior[] values()
func (jbobject *ServicesEc2ModelShutdownBehavior) Values() []*ServicesEc2ModelShutdownBehavior {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ShutdownBehavior", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ShutdownBehavior"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ShutdownBehavior")
	dst := new([]*ServicesEc2ModelShutdownBehavior)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ShutdownBehavior valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelShutdownBehavior) ValueOf(a string) *ServicesEc2ModelShutdownBehavior {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ShutdownBehavior", "valueOf", "com/amazonaws/services/ec2/model/ShutdownBehavior", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelShutdownBehavior{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelShutdownBehavior) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ShutdownBehavior fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelShutdownBehavior) FromValue(a string) *ServicesEc2ModelShutdownBehavior {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ShutdownBehavior", "fromValue", "com/amazonaws/services/ec2/model/ShutdownBehavior", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelShutdownBehavior{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelShutdownBehavior) Stop() *ServicesEc2ModelShutdownBehavior {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ShutdownBehavior", "Stop", "com/amazonaws/services/ec2/model/ShutdownBehavior")
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
	unique_x := &ServicesEc2ModelShutdownBehavior{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelShutdownBehavior) SetFieldStop(val ServicesEc2ModelShutdownBehaviorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ShutdownBehavior", "Stop", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelShutdownBehavior) Terminate() *ServicesEc2ModelShutdownBehavior {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ShutdownBehavior", "Terminate", "com/amazonaws/services/ec2/model/ShutdownBehavior")
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
	unique_x := &ServicesEc2ModelShutdownBehavior{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelShutdownBehavior) SetFieldTerminate(val ServicesEc2ModelShutdownBehaviorInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ShutdownBehavior", "Terminate", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


