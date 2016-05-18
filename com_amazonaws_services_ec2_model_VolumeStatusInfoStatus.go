package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeStatusInfoStatusInterface interface {

	// public static com.amazonaws.services.ec2.model.VolumeStatusInfoStatus[] values()
	Values() []*ServicesEc2ModelVolumeStatusInfoStatus

	// public static com.amazonaws.services.ec2.model.VolumeStatusInfoStatus valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVolumeStatusInfoStatus

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VolumeStatusInfoStatus fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVolumeStatusInfoStatus
}

type ServicesEc2ModelVolumeStatusInfoStatus struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VolumeStatusInfoStatus[] values()
func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) Values() []*ServicesEc2ModelVolumeStatusInfoStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VolumeStatusInfoStatus")
	dst := new([]*ServicesEc2ModelVolumeStatusInfoStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VolumeStatusInfoStatus valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) ValueOf(a string) *ServicesEc2ModelVolumeStatusInfoStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", "valueOf", "com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusInfoStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VolumeStatusInfoStatus fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) FromValue(a string) *ServicesEc2ModelVolumeStatusInfoStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", "fromValue", "com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusInfoStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) Ok() *ServicesEc2ModelVolumeStatusInfoStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", "Ok", "com/amazonaws/services/ec2/model/VolumeStatusInfoStatus")
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
	unique_x := &ServicesEc2ModelVolumeStatusInfoStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) SetFieldOk(val ServicesEc2ModelVolumeStatusInfoStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", "Ok", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) Impaired() *ServicesEc2ModelVolumeStatusInfoStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", "Impaired", "com/amazonaws/services/ec2/model/VolumeStatusInfoStatus")
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
	unique_x := &ServicesEc2ModelVolumeStatusInfoStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) SetFieldImpaired(val ServicesEc2ModelVolumeStatusInfoStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", "Impaired", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) InsufficientData() *ServicesEc2ModelVolumeStatusInfoStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", "InsufficientData", "com/amazonaws/services/ec2/model/VolumeStatusInfoStatus")
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
	unique_x := &ServicesEc2ModelVolumeStatusInfoStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeStatusInfoStatus) SetFieldInsufficientData(val ServicesEc2ModelVolumeStatusInfoStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeStatusInfoStatus", "InsufficientData", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


