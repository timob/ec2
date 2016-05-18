package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReportStatusTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.ReportStatusType[] values()
	Values() []*ServicesEc2ModelReportStatusType

	// public static com.amazonaws.services.ec2.model.ReportStatusType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelReportStatusType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ReportStatusType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelReportStatusType
}

type ServicesEc2ModelReportStatusType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ReportStatusType[] values()
func (jbobject *ServicesEc2ModelReportStatusType) Values() []*ServicesEc2ModelReportStatusType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ReportStatusType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ReportStatusType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ReportStatusType")
	dst := new([]*ServicesEc2ModelReportStatusType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ReportStatusType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelReportStatusType) ValueOf(a string) *ServicesEc2ModelReportStatusType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ReportStatusType", "valueOf", "com/amazonaws/services/ec2/model/ReportStatusType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReportStatusType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReportStatusType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ReportStatusType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelReportStatusType) FromValue(a string) *ServicesEc2ModelReportStatusType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ReportStatusType", "fromValue", "com/amazonaws/services/ec2/model/ReportStatusType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReportStatusType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportStatusType) Ok() *ServicesEc2ModelReportStatusType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportStatusType", "Ok", "com/amazonaws/services/ec2/model/ReportStatusType")
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
	unique_x := &ServicesEc2ModelReportStatusType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportStatusType) SetFieldOk(val ServicesEc2ModelReportStatusTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportStatusType", "Ok", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelReportStatusType) Impaired() *ServicesEc2ModelReportStatusType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ReportStatusType", "Impaired", "com/amazonaws/services/ec2/model/ReportStatusType")
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
	unique_x := &ServicesEc2ModelReportStatusType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelReportStatusType) SetFieldImpaired(val ServicesEc2ModelReportStatusTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ReportStatusType", "Impaired", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


