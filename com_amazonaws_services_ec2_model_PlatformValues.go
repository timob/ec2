package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPlatformValuesInterface interface {

	// public static com.amazonaws.services.ec2.model.PlatformValues[] values()
	Values() []*ServicesEc2ModelPlatformValues

	// public static com.amazonaws.services.ec2.model.PlatformValues valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelPlatformValues

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.PlatformValues fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelPlatformValues
}

type ServicesEc2ModelPlatformValues struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.PlatformValues[] values()
func (jbobject *ServicesEc2ModelPlatformValues) Values() []*ServicesEc2ModelPlatformValues {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PlatformValues", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/PlatformValues"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/PlatformValues")
	dst := new([]*ServicesEc2ModelPlatformValues)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.PlatformValues valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelPlatformValues) ValueOf(a string) *ServicesEc2ModelPlatformValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PlatformValues", "valueOf", "com/amazonaws/services/ec2/model/PlatformValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlatformValues{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPlatformValues) ToString() string {
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

// public static com.amazonaws.services.ec2.model.PlatformValues fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelPlatformValues) FromValue(a string) *ServicesEc2ModelPlatformValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PlatformValues", "fromValue", "com/amazonaws/services/ec2/model/PlatformValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlatformValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPlatformValues) Windows() *ServicesEc2ModelPlatformValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/PlatformValues", "Windows", "com/amazonaws/services/ec2/model/PlatformValues")
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
	unique_x := &ServicesEc2ModelPlatformValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPlatformValues) SetFieldWindows(val ServicesEc2ModelPlatformValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/PlatformValues", "Windows", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


