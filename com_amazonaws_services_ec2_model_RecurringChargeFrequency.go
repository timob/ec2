package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRecurringChargeFrequencyInterface interface {

	// public static com.amazonaws.services.ec2.model.RecurringChargeFrequency[] values()
	Values() []*ServicesEc2ModelRecurringChargeFrequency

	// public static com.amazonaws.services.ec2.model.RecurringChargeFrequency valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelRecurringChargeFrequency

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.RecurringChargeFrequency fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelRecurringChargeFrequency
}

type ServicesEc2ModelRecurringChargeFrequency struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.RecurringChargeFrequency[] values()
func (jbobject *ServicesEc2ModelRecurringChargeFrequency) Values() []*ServicesEc2ModelRecurringChargeFrequency {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RecurringChargeFrequency", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/RecurringChargeFrequency"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/RecurringChargeFrequency")
	dst := new([]*ServicesEc2ModelRecurringChargeFrequency)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.RecurringChargeFrequency valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelRecurringChargeFrequency) ValueOf(a string) *ServicesEc2ModelRecurringChargeFrequency {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RecurringChargeFrequency", "valueOf", "com/amazonaws/services/ec2/model/RecurringChargeFrequency", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRecurringChargeFrequency{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRecurringChargeFrequency) ToString() string {
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

// public static com.amazonaws.services.ec2.model.RecurringChargeFrequency fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelRecurringChargeFrequency) FromValue(a string) *ServicesEc2ModelRecurringChargeFrequency {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/RecurringChargeFrequency", "fromValue", "com/amazonaws/services/ec2/model/RecurringChargeFrequency", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRecurringChargeFrequency{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRecurringChargeFrequency) Hourly() *ServicesEc2ModelRecurringChargeFrequency {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/RecurringChargeFrequency", "Hourly", "com/amazonaws/services/ec2/model/RecurringChargeFrequency")
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
	unique_x := &ServicesEc2ModelRecurringChargeFrequency{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelRecurringChargeFrequency) SetFieldHourly(val ServicesEc2ModelRecurringChargeFrequencyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/RecurringChargeFrequency", "Hourly", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


