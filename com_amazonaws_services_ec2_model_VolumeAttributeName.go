package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeAttributeNameInterface interface {

	// public static com.amazonaws.services.ec2.model.VolumeAttributeName[] values()
	Values() []*ServicesEc2ModelVolumeAttributeName

	// public static com.amazonaws.services.ec2.model.VolumeAttributeName valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVolumeAttributeName

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VolumeAttributeName fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVolumeAttributeName
}

type ServicesEc2ModelVolumeAttributeName struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VolumeAttributeName[] values()
func (jbobject *ServicesEc2ModelVolumeAttributeName) Values() []*ServicesEc2ModelVolumeAttributeName {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeAttributeName", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VolumeAttributeName"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VolumeAttributeName")
	dst := new([]*ServicesEc2ModelVolumeAttributeName)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VolumeAttributeName valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeAttributeName) ValueOf(a string) *ServicesEc2ModelVolumeAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeAttributeName", "valueOf", "com/amazonaws/services/ec2/model/VolumeAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeAttributeName) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VolumeAttributeName fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeAttributeName) FromValue(a string) *ServicesEc2ModelVolumeAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeAttributeName", "fromValue", "com/amazonaws/services/ec2/model/VolumeAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeAttributeName) AutoEnableIO() *ServicesEc2ModelVolumeAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeAttributeName", "AutoEnableIO", "com/amazonaws/services/ec2/model/VolumeAttributeName")
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
	unique_x := &ServicesEc2ModelVolumeAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeAttributeName) SetFieldAutoEnableIO(val ServicesEc2ModelVolumeAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeAttributeName", "AutoEnableIO", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeAttributeName) ProductCodes() *ServicesEc2ModelVolumeAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeAttributeName", "ProductCodes", "com/amazonaws/services/ec2/model/VolumeAttributeName")
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
	unique_x := &ServicesEc2ModelVolumeAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeAttributeName) SetFieldProductCodes(val ServicesEc2ModelVolumeAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeAttributeName", "ProductCodes", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


