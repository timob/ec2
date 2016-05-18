package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelEventCodeInterface interface {

	// public static com.amazonaws.services.ec2.model.EventCode[] values()
	Values() []*ServicesEc2ModelEventCode

	// public static com.amazonaws.services.ec2.model.EventCode valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelEventCode

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.EventCode fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelEventCode
}

type ServicesEc2ModelEventCode struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.EventCode[] values()
func (jbobject *ServicesEc2ModelEventCode) Values() []*ServicesEc2ModelEventCode {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/EventCode", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/EventCode"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/EventCode")
	dst := new([]*ServicesEc2ModelEventCode)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.EventCode valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelEventCode) ValueOf(a string) *ServicesEc2ModelEventCode {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/EventCode", "valueOf", "com/amazonaws/services/ec2/model/EventCode", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEventCode{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelEventCode) ToString() string {
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

// public static com.amazonaws.services.ec2.model.EventCode fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelEventCode) FromValue(a string) *ServicesEc2ModelEventCode {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/EventCode", "fromValue", "com/amazonaws/services/ec2/model/EventCode", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEventCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventCode) InstanceReboot() *ServicesEc2ModelEventCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/EventCode", "InstanceReboot", "com/amazonaws/services/ec2/model/EventCode")
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
	unique_x := &ServicesEc2ModelEventCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventCode) SetFieldInstanceReboot(val ServicesEc2ModelEventCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/EventCode", "InstanceReboot", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelEventCode) SystemReboot() *ServicesEc2ModelEventCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/EventCode", "SystemReboot", "com/amazonaws/services/ec2/model/EventCode")
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
	unique_x := &ServicesEc2ModelEventCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventCode) SetFieldSystemReboot(val ServicesEc2ModelEventCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/EventCode", "SystemReboot", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelEventCode) SystemMaintenance() *ServicesEc2ModelEventCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/EventCode", "SystemMaintenance", "com/amazonaws/services/ec2/model/EventCode")
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
	unique_x := &ServicesEc2ModelEventCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventCode) SetFieldSystemMaintenance(val ServicesEc2ModelEventCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/EventCode", "SystemMaintenance", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelEventCode) InstanceRetirement() *ServicesEc2ModelEventCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/EventCode", "InstanceRetirement", "com/amazonaws/services/ec2/model/EventCode")
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
	unique_x := &ServicesEc2ModelEventCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventCode) SetFieldInstanceRetirement(val ServicesEc2ModelEventCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/EventCode", "InstanceRetirement", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelEventCode) InstanceStop() *ServicesEc2ModelEventCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/EventCode", "InstanceStop", "com/amazonaws/services/ec2/model/EventCode")
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
	unique_x := &ServicesEc2ModelEventCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventCode) SetFieldInstanceStop(val ServicesEc2ModelEventCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/EventCode", "InstanceStop", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


