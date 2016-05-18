package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeStateInterface interface {

	// public static com.amazonaws.services.ec2.model.VolumeState[] values()
	Values() []*ServicesEc2ModelVolumeState

	// public static com.amazonaws.services.ec2.model.VolumeState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVolumeState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VolumeState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVolumeState
}

type ServicesEc2ModelVolumeState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VolumeState[] values()
func (jbobject *ServicesEc2ModelVolumeState) Values() []*ServicesEc2ModelVolumeState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VolumeState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VolumeState")
	dst := new([]*ServicesEc2ModelVolumeState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VolumeState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeState) ValueOf(a string) *ServicesEc2ModelVolumeState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeState", "valueOf", "com/amazonaws/services/ec2/model/VolumeState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VolumeState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeState) FromValue(a string) *ServicesEc2ModelVolumeState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeState", "fromValue", "com/amazonaws/services/ec2/model/VolumeState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeState) Creating() *ServicesEc2ModelVolumeState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Creating", "com/amazonaws/services/ec2/model/VolumeState")
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
	unique_x := &ServicesEc2ModelVolumeState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeState) SetFieldCreating(val ServicesEc2ModelVolumeStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Creating", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeState) Available() *ServicesEc2ModelVolumeState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Available", "com/amazonaws/services/ec2/model/VolumeState")
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
	unique_x := &ServicesEc2ModelVolumeState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeState) SetFieldAvailable(val ServicesEc2ModelVolumeStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeState) InUse() *ServicesEc2ModelVolumeState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeState", "InUse", "com/amazonaws/services/ec2/model/VolumeState")
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
	unique_x := &ServicesEc2ModelVolumeState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeState) SetFieldInUse(val ServicesEc2ModelVolumeStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeState", "InUse", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeState) Deleting() *ServicesEc2ModelVolumeState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Deleting", "com/amazonaws/services/ec2/model/VolumeState")
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
	unique_x := &ServicesEc2ModelVolumeState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeState) SetFieldDeleting(val ServicesEc2ModelVolumeStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Deleting", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeState) Deleted() *ServicesEc2ModelVolumeState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Deleted", "com/amazonaws/services/ec2/model/VolumeState")
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
	unique_x := &ServicesEc2ModelVolumeState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeState) SetFieldDeleted(val ServicesEc2ModelVolumeStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Deleted", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeState) Error() *ServicesEc2ModelVolumeState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Error", "com/amazonaws/services/ec2/model/VolumeState")
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
	unique_x := &ServicesEc2ModelVolumeState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeState) SetFieldError(val ServicesEc2ModelVolumeStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeState", "Error", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


