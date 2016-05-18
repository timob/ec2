package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelExportEnvironmentInterface interface {

	// public static com.amazonaws.services.ec2.model.ExportEnvironment[] values()
	Values() []*ServicesEc2ModelExportEnvironment

	// public static com.amazonaws.services.ec2.model.ExportEnvironment valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelExportEnvironment

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ExportEnvironment fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelExportEnvironment
}

type ServicesEc2ModelExportEnvironment struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ExportEnvironment[] values()
func (jbobject *ServicesEc2ModelExportEnvironment) Values() []*ServicesEc2ModelExportEnvironment {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ExportEnvironment", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ExportEnvironment"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ExportEnvironment")
	dst := new([]*ServicesEc2ModelExportEnvironment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ExportEnvironment valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelExportEnvironment) ValueOf(a string) *ServicesEc2ModelExportEnvironment {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ExportEnvironment", "valueOf", "com/amazonaws/services/ec2/model/ExportEnvironment", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportEnvironment{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelExportEnvironment) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ExportEnvironment fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelExportEnvironment) FromValue(a string) *ServicesEc2ModelExportEnvironment {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ExportEnvironment", "fromValue", "com/amazonaws/services/ec2/model/ExportEnvironment", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExportEnvironment{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExportEnvironment) Citrix() *ServicesEc2ModelExportEnvironment {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ExportEnvironment", "Citrix", "com/amazonaws/services/ec2/model/ExportEnvironment")
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
	unique_x := &ServicesEc2ModelExportEnvironment{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExportEnvironment) SetFieldCitrix(val ServicesEc2ModelExportEnvironmentInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ExportEnvironment", "Citrix", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelExportEnvironment) Vmware() *ServicesEc2ModelExportEnvironment {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ExportEnvironment", "Vmware", "com/amazonaws/services/ec2/model/ExportEnvironment")
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
	unique_x := &ServicesEc2ModelExportEnvironment{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExportEnvironment) SetFieldVmware(val ServicesEc2ModelExportEnvironmentInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ExportEnvironment", "Vmware", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelExportEnvironment) Microsoft() *ServicesEc2ModelExportEnvironment {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ExportEnvironment", "Microsoft", "com/amazonaws/services/ec2/model/ExportEnvironment")
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
	unique_x := &ServicesEc2ModelExportEnvironment{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExportEnvironment) SetFieldMicrosoft(val ServicesEc2ModelExportEnvironmentInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ExportEnvironment", "Microsoft", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


