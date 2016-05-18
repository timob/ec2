package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeStatusNameInterface interface {

	// public static com.amazonaws.services.ec2.model.VolumeStatusName[] values()
	Values() []*ServicesEc2ModelVolumeStatusName

	// public static com.amazonaws.services.ec2.model.VolumeStatusName valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVolumeStatusName

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VolumeStatusName fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVolumeStatusName
}

type ServicesEc2ModelVolumeStatusName struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VolumeStatusName[] values()
func (jbobject *ServicesEc2ModelVolumeStatusName) Values() []*ServicesEc2ModelVolumeStatusName {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeStatusName", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VolumeStatusName"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VolumeStatusName")
	dst := new([]*ServicesEc2ModelVolumeStatusName)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VolumeStatusName valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusName) ValueOf(a string) *ServicesEc2ModelVolumeStatusName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeStatusName", "valueOf", "com/amazonaws/services/ec2/model/VolumeStatusName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusName{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeStatusName) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VolumeStatusName fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusName) FromValue(a string) *ServicesEc2ModelVolumeStatusName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeStatusName", "fromValue", "com/amazonaws/services/ec2/model/VolumeStatusName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeStatusName) IoEnabled() *ServicesEc2ModelVolumeStatusName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeStatusName", "IoEnabled", "com/amazonaws/services/ec2/model/VolumeStatusName")
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
	unique_x := &ServicesEc2ModelVolumeStatusName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeStatusName) SetFieldIoEnabled(val ServicesEc2ModelVolumeStatusNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeStatusName", "IoEnabled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeStatusName) IoPerformance() *ServicesEc2ModelVolumeStatusName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeStatusName", "IoPerformance", "com/amazonaws/services/ec2/model/VolumeStatusName")
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
	unique_x := &ServicesEc2ModelVolumeStatusName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeStatusName) SetFieldIoPerformance(val ServicesEc2ModelVolumeStatusNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeStatusName", "IoPerformance", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


