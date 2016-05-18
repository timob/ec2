package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelImageStateInterface interface {

	// public static com.amazonaws.services.ec2.model.ImageState[] values()
	Values() []*ServicesEc2ModelImageState

	// public static com.amazonaws.services.ec2.model.ImageState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelImageState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ImageState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelImageState
}

type ServicesEc2ModelImageState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ImageState[] values()
func (jbobject *ServicesEc2ModelImageState) Values() []*ServicesEc2ModelImageState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ImageState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ImageState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ImageState")
	dst := new([]*ServicesEc2ModelImageState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ImageState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelImageState) ValueOf(a string) *ServicesEc2ModelImageState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ImageState", "valueOf", "com/amazonaws/services/ec2/model/ImageState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImageState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelImageState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ImageState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelImageState) FromValue(a string) *ServicesEc2ModelImageState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ImageState", "fromValue", "com/amazonaws/services/ec2/model/ImageState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelImageState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageState) Pending() *ServicesEc2ModelImageState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageState", "Pending", "com/amazonaws/services/ec2/model/ImageState")
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
	unique_x := &ServicesEc2ModelImageState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageState) SetFieldPending(val ServicesEc2ModelImageStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelImageState) Available() *ServicesEc2ModelImageState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageState", "Available", "com/amazonaws/services/ec2/model/ImageState")
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
	unique_x := &ServicesEc2ModelImageState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageState) SetFieldAvailable(val ServicesEc2ModelImageStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelImageState) Invalid() *ServicesEc2ModelImageState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageState", "Invalid", "com/amazonaws/services/ec2/model/ImageState")
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
	unique_x := &ServicesEc2ModelImageState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageState) SetFieldInvalid(val ServicesEc2ModelImageStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageState", "Invalid", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelImageState) Deregistered() *ServicesEc2ModelImageState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageState", "Deregistered", "com/amazonaws/services/ec2/model/ImageState")
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
	unique_x := &ServicesEc2ModelImageState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageState) SetFieldDeregistered(val ServicesEc2ModelImageStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageState", "Deregistered", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelImageState) Transient() *ServicesEc2ModelImageState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageState", "Transient", "com/amazonaws/services/ec2/model/ImageState")
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
	unique_x := &ServicesEc2ModelImageState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageState) SetFieldTransient(val ServicesEc2ModelImageStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageState", "Transient", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelImageState) Failed() *ServicesEc2ModelImageState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageState", "Failed", "com/amazonaws/services/ec2/model/ImageState")
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
	unique_x := &ServicesEc2ModelImageState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageState) SetFieldFailed(val ServicesEc2ModelImageStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageState", "Failed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelImageState) Error() *ServicesEc2ModelImageState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ImageState", "Error", "com/amazonaws/services/ec2/model/ImageState")
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
	unique_x := &ServicesEc2ModelImageState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelImageState) SetFieldError(val ServicesEc2ModelImageStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ImageState", "Error", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


