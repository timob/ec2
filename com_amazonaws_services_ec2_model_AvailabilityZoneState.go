package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAvailabilityZoneStateInterface interface {

	// public static com.amazonaws.services.ec2.model.AvailabilityZoneState[] values()
	Values() []*ServicesEc2ModelAvailabilityZoneState

	// public static com.amazonaws.services.ec2.model.AvailabilityZoneState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelAvailabilityZoneState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.AvailabilityZoneState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelAvailabilityZoneState
}

type ServicesEc2ModelAvailabilityZoneState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.AvailabilityZoneState[] values()
func (jbobject *ServicesEc2ModelAvailabilityZoneState) Values() []*ServicesEc2ModelAvailabilityZoneState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AvailabilityZoneState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/AvailabilityZoneState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/AvailabilityZoneState")
	dst := new([]*ServicesEc2ModelAvailabilityZoneState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.AvailabilityZoneState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZoneState) ValueOf(a string) *ServicesEc2ModelAvailabilityZoneState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AvailabilityZoneState", "valueOf", "com/amazonaws/services/ec2/model/AvailabilityZoneState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAvailabilityZoneState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAvailabilityZoneState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.AvailabilityZoneState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZoneState) FromValue(a string) *ServicesEc2ModelAvailabilityZoneState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AvailabilityZoneState", "fromValue", "com/amazonaws/services/ec2/model/AvailabilityZoneState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAvailabilityZoneState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAvailabilityZoneState) Available() *ServicesEc2ModelAvailabilityZoneState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AvailabilityZoneState", "Available", "com/amazonaws/services/ec2/model/AvailabilityZoneState")
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
	unique_x := &ServicesEc2ModelAvailabilityZoneState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAvailabilityZoneState) SetFieldAvailable(val ServicesEc2ModelAvailabilityZoneStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AvailabilityZoneState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAvailabilityZoneState) Information() *ServicesEc2ModelAvailabilityZoneState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AvailabilityZoneState", "Information", "com/amazonaws/services/ec2/model/AvailabilityZoneState")
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
	unique_x := &ServicesEc2ModelAvailabilityZoneState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAvailabilityZoneState) SetFieldInformation(val ServicesEc2ModelAvailabilityZoneStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AvailabilityZoneState", "Information", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAvailabilityZoneState) Impaired() *ServicesEc2ModelAvailabilityZoneState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AvailabilityZoneState", "Impaired", "com/amazonaws/services/ec2/model/AvailabilityZoneState")
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
	unique_x := &ServicesEc2ModelAvailabilityZoneState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAvailabilityZoneState) SetFieldImpaired(val ServicesEc2ModelAvailabilityZoneStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AvailabilityZoneState", "Impaired", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAvailabilityZoneState) Unavailable() *ServicesEc2ModelAvailabilityZoneState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AvailabilityZoneState", "Unavailable", "com/amazonaws/services/ec2/model/AvailabilityZoneState")
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
	unique_x := &ServicesEc2ModelAvailabilityZoneState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAvailabilityZoneState) SetFieldUnavailable(val ServicesEc2ModelAvailabilityZoneStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AvailabilityZoneState", "Unavailable", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


