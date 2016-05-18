package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelContainerFormatInterface interface {

	// public static com.amazonaws.services.ec2.model.ContainerFormat[] values()
	Values() []*ServicesEc2ModelContainerFormat

	// public static com.amazonaws.services.ec2.model.ContainerFormat valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelContainerFormat

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ContainerFormat fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelContainerFormat
}

type ServicesEc2ModelContainerFormat struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ContainerFormat[] values()
func (jbobject *ServicesEc2ModelContainerFormat) Values() []*ServicesEc2ModelContainerFormat {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ContainerFormat", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ContainerFormat"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ContainerFormat")
	dst := new([]*ServicesEc2ModelContainerFormat)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ContainerFormat valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelContainerFormat) ValueOf(a string) *ServicesEc2ModelContainerFormat {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ContainerFormat", "valueOf", "com/amazonaws/services/ec2/model/ContainerFormat", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelContainerFormat{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelContainerFormat) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ContainerFormat fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelContainerFormat) FromValue(a string) *ServicesEc2ModelContainerFormat {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ContainerFormat", "fromValue", "com/amazonaws/services/ec2/model/ContainerFormat", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelContainerFormat{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelContainerFormat) Ova() *ServicesEc2ModelContainerFormat {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ContainerFormat", "Ova", "com/amazonaws/services/ec2/model/ContainerFormat")
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
	unique_x := &ServicesEc2ModelContainerFormat{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelContainerFormat) SetFieldOva(val ServicesEc2ModelContainerFormatInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ContainerFormat", "Ova", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


