package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelExcessCapacityTerminationPolicyInterface interface {

	// public static com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy[] values()
	Values() []*ServicesEc2ModelExcessCapacityTerminationPolicy

	// public static com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelExcessCapacityTerminationPolicy

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelExcessCapacityTerminationPolicy
}

type ServicesEc2ModelExcessCapacityTerminationPolicy struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy[] values()
func (jbobject *ServicesEc2ModelExcessCapacityTerminationPolicy) Values() []*ServicesEc2ModelExcessCapacityTerminationPolicy {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy")
	dst := new([]*ServicesEc2ModelExcessCapacityTerminationPolicy)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelExcessCapacityTerminationPolicy) ValueOf(a string) *ServicesEc2ModelExcessCapacityTerminationPolicy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy", "valueOf", "com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExcessCapacityTerminationPolicy{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelExcessCapacityTerminationPolicy) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ExcessCapacityTerminationPolicy fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelExcessCapacityTerminationPolicy) FromValue(a string) *ServicesEc2ModelExcessCapacityTerminationPolicy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy", "fromValue", "com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelExcessCapacityTerminationPolicy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExcessCapacityTerminationPolicy) NoTermination() *ServicesEc2ModelExcessCapacityTerminationPolicy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy", "NoTermination", "com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy")
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
	unique_x := &ServicesEc2ModelExcessCapacityTerminationPolicy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExcessCapacityTerminationPolicy) SetFieldNoTermination(val ServicesEc2ModelExcessCapacityTerminationPolicyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy", "NoTermination", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelExcessCapacityTerminationPolicy) Default() *ServicesEc2ModelExcessCapacityTerminationPolicy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy", "Default", "com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy")
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
	unique_x := &ServicesEc2ModelExcessCapacityTerminationPolicy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelExcessCapacityTerminationPolicy) SetFieldDefault(val ServicesEc2ModelExcessCapacityTerminationPolicyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ExcessCapacityTerminationPolicy", "Default", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


