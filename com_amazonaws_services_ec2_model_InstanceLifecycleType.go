package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceLifecycleTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.InstanceLifecycleType[] values()
	Values() []*ServicesEc2ModelInstanceLifecycleType

	// public static com.amazonaws.services.ec2.model.InstanceLifecycleType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelInstanceLifecycleType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.InstanceLifecycleType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelInstanceLifecycleType
}

type ServicesEc2ModelInstanceLifecycleType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.InstanceLifecycleType[] values()
func (jbobject *ServicesEc2ModelInstanceLifecycleType) Values() []*ServicesEc2ModelInstanceLifecycleType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceLifecycleType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/InstanceLifecycleType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/InstanceLifecycleType")
	dst := new([]*ServicesEc2ModelInstanceLifecycleType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.InstanceLifecycleType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceLifecycleType) ValueOf(a string) *ServicesEc2ModelInstanceLifecycleType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceLifecycleType", "valueOf", "com/amazonaws/services/ec2/model/InstanceLifecycleType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceLifecycleType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceLifecycleType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.InstanceLifecycleType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceLifecycleType) FromValue(a string) *ServicesEc2ModelInstanceLifecycleType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceLifecycleType", "fromValue", "com/amazonaws/services/ec2/model/InstanceLifecycleType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceLifecycleType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceLifecycleType) Spot() *ServicesEc2ModelInstanceLifecycleType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceLifecycleType", "Spot", "com/amazonaws/services/ec2/model/InstanceLifecycleType")
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
	unique_x := &ServicesEc2ModelInstanceLifecycleType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceLifecycleType) SetFieldSpot(val ServicesEc2ModelInstanceLifecycleTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceLifecycleType", "Spot", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceLifecycleType) Scheduled() *ServicesEc2ModelInstanceLifecycleType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceLifecycleType", "Scheduled", "com/amazonaws/services/ec2/model/InstanceLifecycleType")
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
	unique_x := &ServicesEc2ModelInstanceLifecycleType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceLifecycleType) SetFieldScheduled(val ServicesEc2ModelInstanceLifecycleTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceLifecycleType", "Scheduled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


