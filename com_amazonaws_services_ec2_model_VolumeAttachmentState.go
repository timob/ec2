package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeAttachmentStateInterface interface {

	// public static com.amazonaws.services.ec2.model.VolumeAttachmentState[] values()
	Values() []*ServicesEc2ModelVolumeAttachmentState

	// public static com.amazonaws.services.ec2.model.VolumeAttachmentState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVolumeAttachmentState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VolumeAttachmentState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVolumeAttachmentState
}

type ServicesEc2ModelVolumeAttachmentState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VolumeAttachmentState[] values()
func (jbobject *ServicesEc2ModelVolumeAttachmentState) Values() []*ServicesEc2ModelVolumeAttachmentState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeAttachmentState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VolumeAttachmentState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VolumeAttachmentState")
	dst := new([]*ServicesEc2ModelVolumeAttachmentState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VolumeAttachmentState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeAttachmentState) ValueOf(a string) *ServicesEc2ModelVolumeAttachmentState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeAttachmentState", "valueOf", "com/amazonaws/services/ec2/model/VolumeAttachmentState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeAttachmentState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeAttachmentState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VolumeAttachmentState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeAttachmentState) FromValue(a string) *ServicesEc2ModelVolumeAttachmentState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VolumeAttachmentState", "fromValue", "com/amazonaws/services/ec2/model/VolumeAttachmentState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeAttachmentState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeAttachmentState) Attaching() *ServicesEc2ModelVolumeAttachmentState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeAttachmentState", "Attaching", "com/amazonaws/services/ec2/model/VolumeAttachmentState")
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
	unique_x := &ServicesEc2ModelVolumeAttachmentState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeAttachmentState) SetFieldAttaching(val ServicesEc2ModelVolumeAttachmentStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeAttachmentState", "Attaching", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeAttachmentState) Attached() *ServicesEc2ModelVolumeAttachmentState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeAttachmentState", "Attached", "com/amazonaws/services/ec2/model/VolumeAttachmentState")
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
	unique_x := &ServicesEc2ModelVolumeAttachmentState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeAttachmentState) SetFieldAttached(val ServicesEc2ModelVolumeAttachmentStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeAttachmentState", "Attached", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeAttachmentState) Detaching() *ServicesEc2ModelVolumeAttachmentState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeAttachmentState", "Detaching", "com/amazonaws/services/ec2/model/VolumeAttachmentState")
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
	unique_x := &ServicesEc2ModelVolumeAttachmentState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeAttachmentState) SetFieldDetaching(val ServicesEc2ModelVolumeAttachmentStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeAttachmentState", "Detaching", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVolumeAttachmentState) Detached() *ServicesEc2ModelVolumeAttachmentState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VolumeAttachmentState", "Detached", "com/amazonaws/services/ec2/model/VolumeAttachmentState")
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
	unique_x := &ServicesEc2ModelVolumeAttachmentState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVolumeAttachmentState) SetFieldDetached(val ServicesEc2ModelVolumeAttachmentStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VolumeAttachmentState", "Detached", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


