package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSubnetStateInterface interface {

	// public static com.amazonaws.services.ec2.model.SubnetState[] values()
	Values() []*ServicesEc2ModelSubnetState

	// public static com.amazonaws.services.ec2.model.SubnetState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelSubnetState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.SubnetState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelSubnetState
}

type ServicesEc2ModelSubnetState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.SubnetState[] values()
func (jbobject *ServicesEc2ModelSubnetState) Values() []*ServicesEc2ModelSubnetState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SubnetState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/SubnetState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/SubnetState")
	dst := new([]*ServicesEc2ModelSubnetState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.SubnetState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelSubnetState) ValueOf(a string) *ServicesEc2ModelSubnetState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SubnetState", "valueOf", "com/amazonaws/services/ec2/model/SubnetState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSubnetState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSubnetState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.SubnetState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelSubnetState) FromValue(a string) *ServicesEc2ModelSubnetState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SubnetState", "fromValue", "com/amazonaws/services/ec2/model/SubnetState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSubnetState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSubnetState) Pending() *ServicesEc2ModelSubnetState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SubnetState", "Pending", "com/amazonaws/services/ec2/model/SubnetState")
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
	unique_x := &ServicesEc2ModelSubnetState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSubnetState) SetFieldPending(val ServicesEc2ModelSubnetStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SubnetState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSubnetState) Available() *ServicesEc2ModelSubnetState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SubnetState", "Available", "com/amazonaws/services/ec2/model/SubnetState")
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
	unique_x := &ServicesEc2ModelSubnetState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSubnetState) SetFieldAvailable(val ServicesEc2ModelSubnetStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SubnetState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


