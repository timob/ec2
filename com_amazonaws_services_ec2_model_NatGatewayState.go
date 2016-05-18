package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNatGatewayStateInterface interface {

	// public static com.amazonaws.services.ec2.model.NatGatewayState[] values()
	Values() []*ServicesEc2ModelNatGatewayState

	// public static com.amazonaws.services.ec2.model.NatGatewayState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelNatGatewayState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.NatGatewayState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelNatGatewayState
}

type ServicesEc2ModelNatGatewayState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.NatGatewayState[] values()
func (jbobject *ServicesEc2ModelNatGatewayState) Values() []*ServicesEc2ModelNatGatewayState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NatGatewayState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/NatGatewayState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/NatGatewayState")
	dst := new([]*ServicesEc2ModelNatGatewayState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.NatGatewayState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelNatGatewayState) ValueOf(a string) *ServicesEc2ModelNatGatewayState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NatGatewayState", "valueOf", "com/amazonaws/services/ec2/model/NatGatewayState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNatGatewayState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNatGatewayState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.NatGatewayState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelNatGatewayState) FromValue(a string) *ServicesEc2ModelNatGatewayState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NatGatewayState", "fromValue", "com/amazonaws/services/ec2/model/NatGatewayState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNatGatewayState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNatGatewayState) Pending() *ServicesEc2ModelNatGatewayState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Pending", "com/amazonaws/services/ec2/model/NatGatewayState")
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
	unique_x := &ServicesEc2ModelNatGatewayState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNatGatewayState) SetFieldPending(val ServicesEc2ModelNatGatewayStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNatGatewayState) Failed() *ServicesEc2ModelNatGatewayState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Failed", "com/amazonaws/services/ec2/model/NatGatewayState")
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
	unique_x := &ServicesEc2ModelNatGatewayState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNatGatewayState) SetFieldFailed(val ServicesEc2ModelNatGatewayStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Failed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNatGatewayState) Available() *ServicesEc2ModelNatGatewayState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Available", "com/amazonaws/services/ec2/model/NatGatewayState")
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
	unique_x := &ServicesEc2ModelNatGatewayState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNatGatewayState) SetFieldAvailable(val ServicesEc2ModelNatGatewayStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNatGatewayState) Deleting() *ServicesEc2ModelNatGatewayState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Deleting", "com/amazonaws/services/ec2/model/NatGatewayState")
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
	unique_x := &ServicesEc2ModelNatGatewayState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNatGatewayState) SetFieldDeleting(val ServicesEc2ModelNatGatewayStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Deleting", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNatGatewayState) Deleted() *ServicesEc2ModelNatGatewayState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Deleted", "com/amazonaws/services/ec2/model/NatGatewayState")
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
	unique_x := &ServicesEc2ModelNatGatewayState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNatGatewayState) SetFieldDeleted(val ServicesEc2ModelNatGatewayStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NatGatewayState", "Deleted", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


