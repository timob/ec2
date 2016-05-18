package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpnStateInterface interface {

	// public static com.amazonaws.services.ec2.model.VpnState[] values()
	Values() []*ServicesEc2ModelVpnState

	// public static com.amazonaws.services.ec2.model.VpnState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVpnState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VpnState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVpnState
}

type ServicesEc2ModelVpnState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VpnState[] values()
func (jbobject *ServicesEc2ModelVpnState) Values() []*ServicesEc2ModelVpnState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpnState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VpnState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VpnState")
	dst := new([]*ServicesEc2ModelVpnState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VpnState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVpnState) ValueOf(a string) *ServicesEc2ModelVpnState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpnState", "valueOf", "com/amazonaws/services/ec2/model/VpnState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpnState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VpnState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVpnState) FromValue(a string) *ServicesEc2ModelVpnState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpnState", "fromValue", "com/amazonaws/services/ec2/model/VpnState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpnState) Pending() *ServicesEc2ModelVpnState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpnState", "Pending", "com/amazonaws/services/ec2/model/VpnState")
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
	unique_x := &ServicesEc2ModelVpnState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpnState) SetFieldPending(val ServicesEc2ModelVpnStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpnState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpnState) Available() *ServicesEc2ModelVpnState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpnState", "Available", "com/amazonaws/services/ec2/model/VpnState")
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
	unique_x := &ServicesEc2ModelVpnState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpnState) SetFieldAvailable(val ServicesEc2ModelVpnStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpnState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpnState) Deleting() *ServicesEc2ModelVpnState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpnState", "Deleting", "com/amazonaws/services/ec2/model/VpnState")
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
	unique_x := &ServicesEc2ModelVpnState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpnState) SetFieldDeleting(val ServicesEc2ModelVpnStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpnState", "Deleting", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpnState) Deleted() *ServicesEc2ModelVpnState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpnState", "Deleted", "com/amazonaws/services/ec2/model/VpnState")
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
	unique_x := &ServicesEc2ModelVpnState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpnState) SetFieldDeleted(val ServicesEc2ModelVpnStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpnState", "Deleted", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


