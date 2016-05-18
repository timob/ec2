package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImageTypeValuesInterface interface {

	// public static com.amazonaws.services.ec2.model.ImageTypeValues[] values()
	Values() []*ServicesEc2ModelImageTypeValues

	// public static com.amazonaws.services.ec2.model.ImageTypeValues valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelImageTypeValues

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ImageTypeValues fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelImageTypeValues
}

type ServicesEc2ModelImageTypeValues struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ImageTypeValues[] values()
func (jbobject *ServicesEc2ModelImageTypeValues) Values() []*ServicesEc2ModelImageTypeValues {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ImageTypeValues", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ImageTypeValues"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ImageTypeValues")
	dst := new([]*ServicesEc2ModelImageTypeValues)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ImageTypeValues valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelImageTypeValues) ValueOf(a string) *ServicesEc2ModelImageTypeValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ImageTypeValues", "valueOf", "com/amazonaws/services/ec2/model/ImageTypeValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImageTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImageTypeValues) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ImageTypeValues fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelImageTypeValues) FromValue(a string) *ServicesEc2ModelImageTypeValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ImageTypeValues", "fromValue", "com/amazonaws/services/ec2/model/ImageTypeValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImageTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageTypeValues) Machine() *ServicesEc2ModelImageTypeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageTypeValues", "Machine", "com/amazonaws/services/ec2/model/ImageTypeValues")
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
	unique_x := &ServicesEc2ModelImageTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageTypeValues) SetFieldMachine(val ServicesEc2ModelImageTypeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageTypeValues", "Machine", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelImageTypeValues) Kernel() *ServicesEc2ModelImageTypeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageTypeValues", "Kernel", "com/amazonaws/services/ec2/model/ImageTypeValues")
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
	unique_x := &ServicesEc2ModelImageTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageTypeValues) SetFieldKernel(val ServicesEc2ModelImageTypeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageTypeValues", "Kernel", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelImageTypeValues) Ramdisk() *ServicesEc2ModelImageTypeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageTypeValues", "Ramdisk", "com/amazonaws/services/ec2/model/ImageTypeValues")
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
	unique_x := &ServicesEc2ModelImageTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageTypeValues) SetFieldRamdisk(val ServicesEc2ModelImageTypeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageTypeValues", "Ramdisk", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


