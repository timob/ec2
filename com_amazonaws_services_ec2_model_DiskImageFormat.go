package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDiskImageFormatInterface interface {

	// public static com.amazonaws.services.ec2.model.DiskImageFormat[] values()
	Values() []*ServicesEc2ModelDiskImageFormat

	// public static com.amazonaws.services.ec2.model.DiskImageFormat valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelDiskImageFormat

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.DiskImageFormat fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelDiskImageFormat
}

type ServicesEc2ModelDiskImageFormat struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.DiskImageFormat[] values()
func (jbobject *ServicesEc2ModelDiskImageFormat) Values() []*ServicesEc2ModelDiskImageFormat {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DiskImageFormat", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/DiskImageFormat"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/DiskImageFormat")
	dst := new([]*ServicesEc2ModelDiskImageFormat)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.DiskImageFormat valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageFormat) ValueOf(a string) *ServicesEc2ModelDiskImageFormat {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DiskImageFormat", "valueOf", "com/amazonaws/services/ec2/model/DiskImageFormat", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDiskImageFormat{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDiskImageFormat) ToString() string {
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

// public static com.amazonaws.services.ec2.model.DiskImageFormat fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelDiskImageFormat) FromValue(a string) *ServicesEc2ModelDiskImageFormat {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DiskImageFormat", "fromValue", "com/amazonaws/services/ec2/model/DiskImageFormat", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDiskImageFormat{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDiskImageFormat) VMDK() *ServicesEc2ModelDiskImageFormat {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/DiskImageFormat", "VMDK", "com/amazonaws/services/ec2/model/DiskImageFormat")
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
	unique_x := &ServicesEc2ModelDiskImageFormat{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDiskImageFormat) SetFieldVMDK(val ServicesEc2ModelDiskImageFormatInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/DiskImageFormat", "VMDK", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelDiskImageFormat) RAW() *ServicesEc2ModelDiskImageFormat {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/DiskImageFormat", "RAW", "com/amazonaws/services/ec2/model/DiskImageFormat")
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
	unique_x := &ServicesEc2ModelDiskImageFormat{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDiskImageFormat) SetFieldRAW(val ServicesEc2ModelDiskImageFormatInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/DiskImageFormat", "RAW", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelDiskImageFormat) VHD() *ServicesEc2ModelDiskImageFormat {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/DiskImageFormat", "VHD", "com/amazonaws/services/ec2/model/DiskImageFormat")
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
	unique_x := &ServicesEc2ModelDiskImageFormat{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDiskImageFormat) SetFieldVHD(val ServicesEc2ModelDiskImageFormatInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/DiskImageFormat", "VHD", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


