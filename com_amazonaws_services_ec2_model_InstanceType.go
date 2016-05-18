package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.InstanceType[] values()
	Values() []*ServicesEc2ModelInstanceType

	// public static com.amazonaws.services.ec2.model.InstanceType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelInstanceType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.InstanceType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelInstanceType
}

type ServicesEc2ModelInstanceType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.InstanceType[] values()
func (jbobject *ServicesEc2ModelInstanceType) Values() []*ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/InstanceType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/InstanceType")
	dst := new([]*ServicesEc2ModelInstanceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.InstanceType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceType) ValueOf(a string) *ServicesEc2ModelInstanceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceType", "valueOf", "com/amazonaws/services/ec2/model/InstanceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.InstanceType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceType) FromValue(a string) *ServicesEc2ModelInstanceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/InstanceType", "fromValue", "com/amazonaws/services/ec2/model/InstanceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) T1Micro() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T1Micro", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldT1Micro(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T1Micro", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M1Small() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M1Small", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM1Small(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M1Small", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M1Medium() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M1Medium", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM1Medium(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M1Medium", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M1Large() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M1Large", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM1Large(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M1Large", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M1Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M1Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM1Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M1Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M3Medium() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M3Medium", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM3Medium(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M3Medium", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M3Large() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M3Large", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM3Large(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M3Large", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M3Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M3Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM3Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M3Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M32xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M32xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM32xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M32xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M4Large() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M4Large", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM4Large(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M4Large", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M4Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M4Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM4Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M4Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M42xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M42xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM42xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M42xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M44xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M44xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM44xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M44xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M410xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M410xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM410xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M410xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) T2Nano() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Nano", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldT2Nano(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Nano", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) T2Micro() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Micro", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldT2Micro(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Micro", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) T2Small() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Small", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldT2Small(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Small", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) T2Medium() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Medium", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldT2Medium(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Medium", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) T2Large() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Large", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldT2Large(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "T2Large", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M2Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M2Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM2Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M2Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M22xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M22xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM22xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M22xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) M24xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M24xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldM24xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "M24xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) Cr18xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Cr18xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldCr18xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Cr18xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) I2Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "I2Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldI2Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "I2Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) I22xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "I22xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldI22xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "I22xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) I24xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "I24xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldI24xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "I24xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) I28xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "I28xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldI28xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "I28xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) Hi14xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Hi14xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldHi14xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Hi14xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) Hs18xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Hs18xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldHs18xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Hs18xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C1Medium() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C1Medium", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC1Medium(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C1Medium", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C1Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C1Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC1Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C1Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C3Large() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C3Large", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC3Large(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C3Large", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C3Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C3Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC3Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C3Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C32xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C32xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC32xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C32xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C34xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C34xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC34xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C34xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C38xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C38xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC38xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C38xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C4Large() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C4Large", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC4Large(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C4Large", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C4Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C4Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC4Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C4Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C42xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C42xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC42xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C42xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C44xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C44xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC44xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C44xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) C48xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C48xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldC48xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "C48xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) Cc14xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Cc14xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldCc14xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Cc14xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) Cc28xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Cc28xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldCc28xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Cc28xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) G22xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "G22xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldG22xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "G22xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) G28xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "G28xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldG28xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "G28xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) Cg14xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Cg14xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldCg14xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "Cg14xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) R3Large() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R3Large", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldR3Large(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R3Large", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) R3Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R3Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldR3Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R3Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) R32xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R32xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldR32xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R32xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) R34xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R34xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldR34xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R34xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) R38xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R38xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldR38xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "R38xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) D2Xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "D2Xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldD2Xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "D2Xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) D22xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "D22xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldD22xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "D22xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) D24xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "D24xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldD24xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "D24xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelInstanceType) D28xlarge() *ServicesEc2ModelInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/InstanceType", "D28xlarge", "com/amazonaws/services/ec2/model/InstanceType")
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
	unique_x := &ServicesEc2ModelInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelInstanceType) SetFieldD28xlarge(val ServicesEc2ModelInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/InstanceType", "D28xlarge", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


