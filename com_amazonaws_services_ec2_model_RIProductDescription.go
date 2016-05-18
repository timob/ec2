package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRIProductDescriptionInterface interface {

	// public static com.amazonaws.services.ec2.model.RIProductDescription[] values()
	Values() []*ServicesEc2ModelRIProductDescription

	// public static com.amazonaws.services.ec2.model.RIProductDescription valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelRIProductDescription

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.RIProductDescription fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelRIProductDescription
}

type ServicesEc2ModelRIProductDescription struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.RIProductDescription[] values()
func (jbobject *ServicesEc2ModelRIProductDescription) Values() []*ServicesEc2ModelRIProductDescription {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RIProductDescription", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/RIProductDescription"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/RIProductDescription")
	dst := new([]*ServicesEc2ModelRIProductDescription)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.RIProductDescription valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelRIProductDescription) ValueOf(a string) *ServicesEc2ModelRIProductDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RIProductDescription", "valueOf", "com/amazonaws/services/ec2/model/RIProductDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRIProductDescription{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRIProductDescription) ToString() string {
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

// public static com.amazonaws.services.ec2.model.RIProductDescription fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelRIProductDescription) FromValue(a string) *ServicesEc2ModelRIProductDescription {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RIProductDescription", "fromValue", "com/amazonaws/services/ec2/model/RIProductDescription", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRIProductDescription{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRIProductDescription) LinuxUNIX() *ServicesEc2ModelRIProductDescription {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RIProductDescription", "LinuxUNIX", "com/amazonaws/services/ec2/model/RIProductDescription")
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
	unique_x := &ServicesEc2ModelRIProductDescription{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRIProductDescription) SetFieldLinuxUNIX(val ServicesEc2ModelRIProductDescriptionInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RIProductDescription", "LinuxUNIX", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelRIProductDescription) LinuxUNIXAmazonVPC() *ServicesEc2ModelRIProductDescription {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RIProductDescription", "LinuxUNIXAmazonVPC", "com/amazonaws/services/ec2/model/RIProductDescription")
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
	unique_x := &ServicesEc2ModelRIProductDescription{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRIProductDescription) SetFieldLinuxUNIXAmazonVPC(val ServicesEc2ModelRIProductDescriptionInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RIProductDescription", "LinuxUNIXAmazonVPC", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelRIProductDescription) Windows() *ServicesEc2ModelRIProductDescription {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RIProductDescription", "Windows", "com/amazonaws/services/ec2/model/RIProductDescription")
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
	unique_x := &ServicesEc2ModelRIProductDescription{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRIProductDescription) SetFieldWindows(val ServicesEc2ModelRIProductDescriptionInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RIProductDescription", "Windows", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelRIProductDescription) WindowsAmazonVPC() *ServicesEc2ModelRIProductDescription {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RIProductDescription", "WindowsAmazonVPC", "com/amazonaws/services/ec2/model/RIProductDescription")
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
	unique_x := &ServicesEc2ModelRIProductDescription{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRIProductDescription) SetFieldWindowsAmazonVPC(val ServicesEc2ModelRIProductDescriptionInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RIProductDescription", "WindowsAmazonVPC", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


