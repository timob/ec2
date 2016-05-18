package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelEventTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.EventType[] values()
	Values() []*ServicesEc2ModelEventType

	// public static com.amazonaws.services.ec2.model.EventType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelEventType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.EventType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelEventType
}

type ServicesEc2ModelEventType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.EventType[] values()
func (jbobject *ServicesEc2ModelEventType) Values() []*ServicesEc2ModelEventType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/EventType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/EventType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/EventType")
	dst := new([]*ServicesEc2ModelEventType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.EventType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelEventType) ValueOf(a string) *ServicesEc2ModelEventType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/EventType", "valueOf", "com/amazonaws/services/ec2/model/EventType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEventType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelEventType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.EventType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelEventType) FromValue(a string) *ServicesEc2ModelEventType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/EventType", "fromValue", "com/amazonaws/services/ec2/model/EventType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventType) InstanceChange() *ServicesEc2ModelEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/EventType", "InstanceChange", "com/amazonaws/services/ec2/model/EventType")
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
	unique_x := &ServicesEc2ModelEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventType) SetFieldInstanceChange(val ServicesEc2ModelEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/EventType", "InstanceChange", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelEventType) FleetRequestChange() *ServicesEc2ModelEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/EventType", "FleetRequestChange", "com/amazonaws/services/ec2/model/EventType")
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
	unique_x := &ServicesEc2ModelEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventType) SetFieldFleetRequestChange(val ServicesEc2ModelEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/EventType", "FleetRequestChange", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelEventType) Error() *ServicesEc2ModelEventType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/EventType", "Error", "com/amazonaws/services/ec2/model/EventType")
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
	unique_x := &ServicesEc2ModelEventType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelEventType) SetFieldError(val ServicesEc2ModelEventTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/EventType", "Error", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


