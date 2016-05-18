package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.VolumeType[] values()
	Values() []*ServicesEc2ModelVolumeType

	// public static com.amazonaws.services.ec2.model.VolumeType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVolumeType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VolumeType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVolumeType
}

type ServicesEc2ModelVolumeType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VolumeType[] values()
func (jbobject *ServicesEc2ModelVolumeType) Values() []*ServicesEc2ModelVolumeType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VolumeType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VolumeType")
	dst := new([]*ServicesEc2ModelVolumeType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VolumeType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeType) ValueOf(a string) *ServicesEc2ModelVolumeType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeType", "valueOf", "com/amazonaws/services/ec2/model/VolumeType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VolumeType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeType) FromValue(a string) *ServicesEc2ModelVolumeType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeType", "fromValue", "com/amazonaws/services/ec2/model/VolumeType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeType) Standard() *ServicesEc2ModelVolumeType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeType", "Standard", "com/amazonaws/services/ec2/model/VolumeType")
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
	unique_x := &ServicesEc2ModelVolumeType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeType) SetFieldStandard(val ServicesEc2ModelVolumeTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeType", "Standard", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeType) Io1() *ServicesEc2ModelVolumeType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeType", "Io1", "com/amazonaws/services/ec2/model/VolumeType")
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
	unique_x := &ServicesEc2ModelVolumeType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeType) SetFieldIo1(val ServicesEc2ModelVolumeTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeType", "Io1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeType) Gp2() *ServicesEc2ModelVolumeType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeType", "Gp2", "com/amazonaws/services/ec2/model/VolumeType")
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
	unique_x := &ServicesEc2ModelVolumeType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeType) SetFieldGp2(val ServicesEc2ModelVolumeTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeType", "Gp2", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


