package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAffinityInterface interface {

	// public static com.amazonaws.services.ec2.model.Affinity[] values()
	Values() []*ServicesEc2ModelAffinity

	// public static com.amazonaws.services.ec2.model.Affinity valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelAffinity

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.Affinity fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelAffinity
}

type ServicesEc2ModelAffinity struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.Affinity[] values()
func (jbobject *ServicesEc2ModelAffinity) Values() []*ServicesEc2ModelAffinity {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/Affinity", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/Affinity"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/Affinity")
	dst := new([]*ServicesEc2ModelAffinity)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.Affinity valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelAffinity) ValueOf(a string) *ServicesEc2ModelAffinity {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/Affinity", "valueOf", "com/amazonaws/services/ec2/model/Affinity", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAffinity{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAffinity) ToString() string {
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

// public static com.amazonaws.services.ec2.model.Affinity fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelAffinity) FromValue(a string) *ServicesEc2ModelAffinity {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/Affinity", "fromValue", "com/amazonaws/services/ec2/model/Affinity", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAffinity{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAffinity) Default() *ServicesEc2ModelAffinity {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/Affinity", "Default", "com/amazonaws/services/ec2/model/Affinity")
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
	unique_x := &ServicesEc2ModelAffinity{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAffinity) SetFieldDefault(val ServicesEc2ModelAffinityInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/Affinity", "Default", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAffinity) Host() *ServicesEc2ModelAffinity {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/Affinity", "Host", "com/amazonaws/services/ec2/model/Affinity")
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
	unique_x := &ServicesEc2ModelAffinity{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAffinity) SetFieldHost(val ServicesEc2ModelAffinityInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/Affinity", "Host", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


