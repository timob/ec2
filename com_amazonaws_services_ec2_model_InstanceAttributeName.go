package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceAttributeNameInterface interface {

	// public static com.amazonaws.services.ec2.model.InstanceAttributeName[] values()
	Values() []*ServicesEc2ModelInstanceAttributeName

	// public static com.amazonaws.services.ec2.model.InstanceAttributeName valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelInstanceAttributeName

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.InstanceAttributeName fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelInstanceAttributeName
}

type ServicesEc2ModelInstanceAttributeName struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.InstanceAttributeName[] values()
func (jbobject *ServicesEc2ModelInstanceAttributeName) Values() []*ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceAttributeName", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/InstanceAttributeName"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/InstanceAttributeName")
	dst := new([]*ServicesEc2ModelInstanceAttributeName)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.InstanceAttributeName valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttributeName) ValueOf(a string) *ServicesEc2ModelInstanceAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceAttributeName", "valueOf", "com/amazonaws/services/ec2/model/InstanceAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceAttributeName) ToString() string {
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

// public static com.amazonaws.services.ec2.model.InstanceAttributeName fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceAttributeName) FromValue(a string) *ServicesEc2ModelInstanceAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceAttributeName", "fromValue", "com/amazonaws/services/ec2/model/InstanceAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) InstanceType() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "InstanceType", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldInstanceType(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "InstanceType", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) Kernel() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "Kernel", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldKernel(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "Kernel", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) Ramdisk() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "Ramdisk", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldRamdisk(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "Ramdisk", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) UserData() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "UserData", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldUserData(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "UserData", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) DisableApiTermination() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "DisableApiTermination", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldDisableApiTermination(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "DisableApiTermination", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) InstanceInitiatedShutdownBehavior() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "InstanceInitiatedShutdownBehavior", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldInstanceInitiatedShutdownBehavior(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "InstanceInitiatedShutdownBehavior", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) RootDeviceName() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "RootDeviceName", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldRootDeviceName(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "RootDeviceName", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) BlockDeviceMapping() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "BlockDeviceMapping", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldBlockDeviceMapping(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "BlockDeviceMapping", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) ProductCodes() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "ProductCodes", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldProductCodes(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "ProductCodes", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SourceDestCheck() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "SourceDestCheck", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldSourceDestCheck(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "SourceDestCheck", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) GroupSet() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "GroupSet", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldGroupSet(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "GroupSet", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) EbsOptimized() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "EbsOptimized", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldEbsOptimized(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "EbsOptimized", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SriovNetSupport() *ServicesEc2ModelInstanceAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "SriovNetSupport", "com/amazonaws/services/ec2/model/InstanceAttributeName")
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
	unique_x := &ServicesEc2ModelInstanceAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceAttributeName) SetFieldSriovNetSupport(val ServicesEc2ModelInstanceAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceAttributeName", "SriovNetSupport", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


