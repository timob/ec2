package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelStatusTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.StatusType[] values()
	Values() []*ServicesEc2ModelStatusType

	// public static com.amazonaws.services.ec2.model.StatusType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelStatusType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.StatusType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelStatusType
}

type ServicesEc2ModelStatusType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.StatusType[] values()
func (jbobject *ServicesEc2ModelStatusType) Values() []*ServicesEc2ModelStatusType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/StatusType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/StatusType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/StatusType")
	dst := new([]*ServicesEc2ModelStatusType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.StatusType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelStatusType) ValueOf(a string) *ServicesEc2ModelStatusType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/StatusType", "valueOf", "com/amazonaws/services/ec2/model/StatusType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelStatusType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelStatusType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.StatusType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelStatusType) FromValue(a string) *ServicesEc2ModelStatusType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/StatusType", "fromValue", "com/amazonaws/services/ec2/model/StatusType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelStatusType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatusType) Passed() *ServicesEc2ModelStatusType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/StatusType", "Passed", "com/amazonaws/services/ec2/model/StatusType")
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
	unique_x := &ServicesEc2ModelStatusType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatusType) SetFieldPassed(val ServicesEc2ModelStatusTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/StatusType", "Passed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelStatusType) Failed() *ServicesEc2ModelStatusType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/StatusType", "Failed", "com/amazonaws/services/ec2/model/StatusType")
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
	unique_x := &ServicesEc2ModelStatusType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatusType) SetFieldFailed(val ServicesEc2ModelStatusTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/StatusType", "Failed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelStatusType) InsufficientData() *ServicesEc2ModelStatusType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/StatusType", "InsufficientData", "com/amazonaws/services/ec2/model/StatusType")
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
	unique_x := &ServicesEc2ModelStatusType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatusType) SetFieldInsufficientData(val ServicesEc2ModelStatusTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/StatusType", "InsufficientData", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelStatusType) Initializing() *ServicesEc2ModelStatusType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/StatusType", "Initializing", "com/amazonaws/services/ec2/model/StatusType")
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
	unique_x := &ServicesEc2ModelStatusType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelStatusType) SetFieldInitializing(val ServicesEc2ModelStatusTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/StatusType", "Initializing", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


