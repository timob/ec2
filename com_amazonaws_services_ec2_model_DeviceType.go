package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDeviceTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.DeviceType[] values()
	Values() []*ServicesEc2ModelDeviceType

	// public static com.amazonaws.services.ec2.model.DeviceType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelDeviceType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.DeviceType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelDeviceType
}

type ServicesEc2ModelDeviceType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.DeviceType[] values()
func (jbobject *ServicesEc2ModelDeviceType) Values() []*ServicesEc2ModelDeviceType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DeviceType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/DeviceType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/DeviceType")
	dst := new([]*ServicesEc2ModelDeviceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.DeviceType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelDeviceType) ValueOf(a string) *ServicesEc2ModelDeviceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DeviceType", "valueOf", "com/amazonaws/services/ec2/model/DeviceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeviceType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDeviceType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.DeviceType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelDeviceType) FromValue(a string) *ServicesEc2ModelDeviceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DeviceType", "fromValue", "com/amazonaws/services/ec2/model/DeviceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDeviceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDeviceType) Ebs() *ServicesEc2ModelDeviceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/DeviceType", "Ebs", "com/amazonaws/services/ec2/model/DeviceType")
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
	unique_x := &ServicesEc2ModelDeviceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDeviceType) SetFieldEbs(val ServicesEc2ModelDeviceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/DeviceType", "Ebs", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelDeviceType) InstanceStore() *ServicesEc2ModelDeviceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/DeviceType", "InstanceStore", "com/amazonaws/services/ec2/model/DeviceType")
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
	unique_x := &ServicesEc2ModelDeviceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDeviceType) SetFieldInstanceStore(val ServicesEc2ModelDeviceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/DeviceType", "InstanceStore", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


