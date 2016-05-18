package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpcStateInterface interface {

	// public static com.amazonaws.services.ec2.model.VpcState[] values()
	Values() []*ServicesEc2ModelVpcState

	// public static com.amazonaws.services.ec2.model.VpcState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVpcState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VpcState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVpcState
}

type ServicesEc2ModelVpcState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VpcState[] values()
func (jbobject *ServicesEc2ModelVpcState) Values() []*ServicesEc2ModelVpcState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpcState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VpcState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VpcState")
	dst := new([]*ServicesEc2ModelVpcState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VpcState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVpcState) ValueOf(a string) *ServicesEc2ModelVpcState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpcState", "valueOf", "com/amazonaws/services/ec2/model/VpcState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpcState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VpcState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVpcState) FromValue(a string) *ServicesEc2ModelVpcState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpcState", "fromValue", "com/amazonaws/services/ec2/model/VpcState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcState) Pending() *ServicesEc2ModelVpcState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcState", "Pending", "com/amazonaws/services/ec2/model/VpcState")
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
	unique_x := &ServicesEc2ModelVpcState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcState) SetFieldPending(val ServicesEc2ModelVpcStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcState) Available() *ServicesEc2ModelVpcState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcState", "Available", "com/amazonaws/services/ec2/model/VpcState")
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
	unique_x := &ServicesEc2ModelVpcState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcState) SetFieldAvailable(val ServicesEc2ModelVpcStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


