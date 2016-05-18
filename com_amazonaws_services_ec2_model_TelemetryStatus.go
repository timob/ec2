package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelTelemetryStatusInterface interface {

	// public static com.amazonaws.services.ec2.model.TelemetryStatus[] values()
	Values() []*ServicesEc2ModelTelemetryStatus

	// public static com.amazonaws.services.ec2.model.TelemetryStatus valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelTelemetryStatus

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.TelemetryStatus fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelTelemetryStatus
}

type ServicesEc2ModelTelemetryStatus struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.TelemetryStatus[] values()
func (jbobject *ServicesEc2ModelTelemetryStatus) Values() []*ServicesEc2ModelTelemetryStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/TelemetryStatus", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/TelemetryStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/TelemetryStatus")
	dst := new([]*ServicesEc2ModelTelemetryStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.TelemetryStatus valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelTelemetryStatus) ValueOf(a string) *ServicesEc2ModelTelemetryStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/TelemetryStatus", "valueOf", "com/amazonaws/services/ec2/model/TelemetryStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTelemetryStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelTelemetryStatus) ToString() string {
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

// public static com.amazonaws.services.ec2.model.TelemetryStatus fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelTelemetryStatus) FromValue(a string) *ServicesEc2ModelTelemetryStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/TelemetryStatus", "fromValue", "com/amazonaws/services/ec2/model/TelemetryStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTelemetryStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTelemetryStatus) UP() *ServicesEc2ModelTelemetryStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/TelemetryStatus", "UP", "com/amazonaws/services/ec2/model/TelemetryStatus")
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
	unique_x := &ServicesEc2ModelTelemetryStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTelemetryStatus) SetFieldUP(val ServicesEc2ModelTelemetryStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/TelemetryStatus", "UP", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelTelemetryStatus) DOWN() *ServicesEc2ModelTelemetryStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/TelemetryStatus", "DOWN", "com/amazonaws/services/ec2/model/TelemetryStatus")
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
	unique_x := &ServicesEc2ModelTelemetryStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTelemetryStatus) SetFieldDOWN(val ServicesEc2ModelTelemetryStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/TelemetryStatus", "DOWN", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


