package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReportInstanceReasonCodesInterface interface {

	// public static com.amazonaws.services.ec2.model.ReportInstanceReasonCodes[] values()
	Values() []*ServicesEc2ModelReportInstanceReasonCodes

	// public static com.amazonaws.services.ec2.model.ReportInstanceReasonCodes valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelReportInstanceReasonCodes

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ReportInstanceReasonCodes fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelReportInstanceReasonCodes
}

type ServicesEc2ModelReportInstanceReasonCodes struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ReportInstanceReasonCodes[] values()
func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) Values() []*ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
	dst := new([]*ServicesEc2ModelReportInstanceReasonCodes)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ReportInstanceReasonCodes valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) ValueOf(a string) *ServicesEc2ModelReportInstanceReasonCodes {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "valueOf", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ReportInstanceReasonCodes fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) FromValue(a string) *ServicesEc2ModelReportInstanceReasonCodes {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "fromValue", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) InstanceStuckInState() *ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "InstanceStuckInState", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) SetFieldInstanceStuckInState(val ServicesEc2ModelReportInstanceReasonCodesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "InstanceStuckInState", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) Unresponsive() *ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "Unresponsive", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) SetFieldUnresponsive(val ServicesEc2ModelReportInstanceReasonCodesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "Unresponsive", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) NotAcceptingCredentials() *ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "NotAcceptingCredentials", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) SetFieldNotAcceptingCredentials(val ServicesEc2ModelReportInstanceReasonCodesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "NotAcceptingCredentials", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) PasswordNotAvailable() *ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PasswordNotAvailable", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) SetFieldPasswordNotAvailable(val ServicesEc2ModelReportInstanceReasonCodesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PasswordNotAvailable", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) PerformanceNetwork() *ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PerformanceNetwork", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) SetFieldPerformanceNetwork(val ServicesEc2ModelReportInstanceReasonCodesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PerformanceNetwork", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) PerformanceInstanceStore() *ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PerformanceInstanceStore", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) SetFieldPerformanceInstanceStore(val ServicesEc2ModelReportInstanceReasonCodesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PerformanceInstanceStore", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) PerformanceEbsVolume() *ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PerformanceEbsVolume", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) SetFieldPerformanceEbsVolume(val ServicesEc2ModelReportInstanceReasonCodesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PerformanceEbsVolume", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) PerformanceOther() *ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PerformanceOther", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) SetFieldPerformanceOther(val ServicesEc2ModelReportInstanceReasonCodesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "PerformanceOther", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) Other() *ServicesEc2ModelReportInstanceReasonCodes {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "Other", "com/amazonaws/services/ec2/model/ReportInstanceReasonCodes")
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
	unique_x := &ServicesEc2ModelReportInstanceReasonCodes{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportInstanceReasonCodes) SetFieldOther(val ServicesEc2ModelReportInstanceReasonCodesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportInstanceReasonCodes", "Other", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


