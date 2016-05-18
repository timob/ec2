package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelStatusInterface interface {

	// public static com.amazonaws.services.ec2.model.Status[] values()
	Values() []*ServicesEc2ModelStatus

	// public static com.amazonaws.services.ec2.model.Status valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelStatus

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.Status fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelStatus
}

type ServicesEc2ModelStatus struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.Status[] values()
func (jbobject *ServicesEc2ModelStatus) Values() []*ServicesEc2ModelStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/Status", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/Status"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/Status")
	dst := new([]*ServicesEc2ModelStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.Status valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelStatus) ValueOf(a string) *ServicesEc2ModelStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/Status", "valueOf", "com/amazonaws/services/ec2/model/Status", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelStatus) ToString() string {
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

// public static com.amazonaws.services.ec2.model.Status fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelStatus) FromValue(a string) *ServicesEc2ModelStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/Status", "fromValue", "com/amazonaws/services/ec2/model/Status", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatus) MoveInProgress() *ServicesEc2ModelStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/Status", "MoveInProgress", "com/amazonaws/services/ec2/model/Status")
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
	unique_x := &ServicesEc2ModelStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatus) SetFieldMoveInProgress(val ServicesEc2ModelStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/Status", "MoveInProgress", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelStatus) InVpc() *ServicesEc2ModelStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/Status", "InVpc", "com/amazonaws/services/ec2/model/Status")
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
	unique_x := &ServicesEc2ModelStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatus) SetFieldInVpc(val ServicesEc2ModelStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/Status", "InVpc", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelStatus) InClassic() *ServicesEc2ModelStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/Status", "InClassic", "com/amazonaws/services/ec2/model/Status")
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
	unique_x := &ServicesEc2ModelStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatus) SetFieldInClassic(val ServicesEc2ModelStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/Status", "InClassic", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


