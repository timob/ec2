package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDatafeedSubscriptionStateInterface interface {

	// public static com.amazonaws.services.ec2.model.DatafeedSubscriptionState[] values()
	Values() []*ServicesEc2ModelDatafeedSubscriptionState

	// public static com.amazonaws.services.ec2.model.DatafeedSubscriptionState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelDatafeedSubscriptionState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.DatafeedSubscriptionState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelDatafeedSubscriptionState
}

type ServicesEc2ModelDatafeedSubscriptionState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.DatafeedSubscriptionState[] values()
func (jbobject *ServicesEc2ModelDatafeedSubscriptionState) Values() []*ServicesEc2ModelDatafeedSubscriptionState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DatafeedSubscriptionState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/DatafeedSubscriptionState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/DatafeedSubscriptionState")
	dst := new([]*ServicesEc2ModelDatafeedSubscriptionState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.DatafeedSubscriptionState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelDatafeedSubscriptionState) ValueOf(a string) *ServicesEc2ModelDatafeedSubscriptionState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DatafeedSubscriptionState", "valueOf", "com/amazonaws/services/ec2/model/DatafeedSubscriptionState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDatafeedSubscriptionState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDatafeedSubscriptionState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.DatafeedSubscriptionState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelDatafeedSubscriptionState) FromValue(a string) *ServicesEc2ModelDatafeedSubscriptionState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/DatafeedSubscriptionState", "fromValue", "com/amazonaws/services/ec2/model/DatafeedSubscriptionState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDatafeedSubscriptionState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDatafeedSubscriptionState) Active() *ServicesEc2ModelDatafeedSubscriptionState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/DatafeedSubscriptionState", "Active", "com/amazonaws/services/ec2/model/DatafeedSubscriptionState")
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
	unique_x := &ServicesEc2ModelDatafeedSubscriptionState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDatafeedSubscriptionState) SetFieldActive(val ServicesEc2ModelDatafeedSubscriptionStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/DatafeedSubscriptionState", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelDatafeedSubscriptionState) Inactive() *ServicesEc2ModelDatafeedSubscriptionState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/DatafeedSubscriptionState", "Inactive", "com/amazonaws/services/ec2/model/DatafeedSubscriptionState")
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
	unique_x := &ServicesEc2ModelDatafeedSubscriptionState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelDatafeedSubscriptionState) SetFieldInactive(val ServicesEc2ModelDatafeedSubscriptionStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/DatafeedSubscriptionState", "Inactive", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


