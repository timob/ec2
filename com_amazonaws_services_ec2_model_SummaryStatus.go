package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSummaryStatusInterface interface {

	// public static com.amazonaws.services.ec2.model.SummaryStatus[] values()
	Values() []*ServicesEc2ModelSummaryStatus

	// public static com.amazonaws.services.ec2.model.SummaryStatus valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelSummaryStatus

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.SummaryStatus fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelSummaryStatus
}

type ServicesEc2ModelSummaryStatus struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.SummaryStatus[] values()
func (jbobject *ServicesEc2ModelSummaryStatus) Values() []*ServicesEc2ModelSummaryStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SummaryStatus", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/SummaryStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/SummaryStatus")
	dst := new([]*ServicesEc2ModelSummaryStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.SummaryStatus valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelSummaryStatus) ValueOf(a string) *ServicesEc2ModelSummaryStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SummaryStatus", "valueOf", "com/amazonaws/services/ec2/model/SummaryStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSummaryStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSummaryStatus) ToString() string {
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

// public static com.amazonaws.services.ec2.model.SummaryStatus fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelSummaryStatus) FromValue(a string) *ServicesEc2ModelSummaryStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SummaryStatus", "fromValue", "com/amazonaws/services/ec2/model/SummaryStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSummaryStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSummaryStatus) Ok() *ServicesEc2ModelSummaryStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "Ok", "com/amazonaws/services/ec2/model/SummaryStatus")
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
	unique_x := &ServicesEc2ModelSummaryStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSummaryStatus) SetFieldOk(val ServicesEc2ModelSummaryStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "Ok", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSummaryStatus) Impaired() *ServicesEc2ModelSummaryStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "Impaired", "com/amazonaws/services/ec2/model/SummaryStatus")
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
	unique_x := &ServicesEc2ModelSummaryStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSummaryStatus) SetFieldImpaired(val ServicesEc2ModelSummaryStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "Impaired", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSummaryStatus) InsufficientData() *ServicesEc2ModelSummaryStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "InsufficientData", "com/amazonaws/services/ec2/model/SummaryStatus")
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
	unique_x := &ServicesEc2ModelSummaryStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSummaryStatus) SetFieldInsufficientData(val ServicesEc2ModelSummaryStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "InsufficientData", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSummaryStatus) NotApplicable() *ServicesEc2ModelSummaryStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "NotApplicable", "com/amazonaws/services/ec2/model/SummaryStatus")
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
	unique_x := &ServicesEc2ModelSummaryStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSummaryStatus) SetFieldNotApplicable(val ServicesEc2ModelSummaryStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "NotApplicable", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSummaryStatus) Initializing() *ServicesEc2ModelSummaryStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "Initializing", "com/amazonaws/services/ec2/model/SummaryStatus")
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
	unique_x := &ServicesEc2ModelSummaryStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSummaryStatus) SetFieldInitializing(val ServicesEc2ModelSummaryStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SummaryStatus", "Initializing", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


