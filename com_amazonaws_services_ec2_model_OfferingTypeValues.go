package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelOfferingTypeValuesInterface interface {

	// public static com.amazonaws.services.ec2.model.OfferingTypeValues[] values()
	Values() []*ServicesEc2ModelOfferingTypeValues

	// public static com.amazonaws.services.ec2.model.OfferingTypeValues valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelOfferingTypeValues

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.OfferingTypeValues fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelOfferingTypeValues
}

type ServicesEc2ModelOfferingTypeValues struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.OfferingTypeValues[] values()
func (jbobject *ServicesEc2ModelOfferingTypeValues) Values() []*ServicesEc2ModelOfferingTypeValues {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/OfferingTypeValues", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/OfferingTypeValues"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/OfferingTypeValues")
	dst := new([]*ServicesEc2ModelOfferingTypeValues)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.OfferingTypeValues valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelOfferingTypeValues) ValueOf(a string) *ServicesEc2ModelOfferingTypeValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/OfferingTypeValues", "valueOf", "com/amazonaws/services/ec2/model/OfferingTypeValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelOfferingTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelOfferingTypeValues) ToString() string {
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

// public static com.amazonaws.services.ec2.model.OfferingTypeValues fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelOfferingTypeValues) FromValue(a string) *ServicesEc2ModelOfferingTypeValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/OfferingTypeValues", "fromValue", "com/amazonaws/services/ec2/model/OfferingTypeValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelOfferingTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOfferingTypeValues) HeavyUtilization() *ServicesEc2ModelOfferingTypeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "HeavyUtilization", "com/amazonaws/services/ec2/model/OfferingTypeValues")
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
	unique_x := &ServicesEc2ModelOfferingTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOfferingTypeValues) SetFieldHeavyUtilization(val ServicesEc2ModelOfferingTypeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "HeavyUtilization", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelOfferingTypeValues) MediumUtilization() *ServicesEc2ModelOfferingTypeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "MediumUtilization", "com/amazonaws/services/ec2/model/OfferingTypeValues")
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
	unique_x := &ServicesEc2ModelOfferingTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOfferingTypeValues) SetFieldMediumUtilization(val ServicesEc2ModelOfferingTypeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "MediumUtilization", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelOfferingTypeValues) LightUtilization() *ServicesEc2ModelOfferingTypeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "LightUtilization", "com/amazonaws/services/ec2/model/OfferingTypeValues")
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
	unique_x := &ServicesEc2ModelOfferingTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOfferingTypeValues) SetFieldLightUtilization(val ServicesEc2ModelOfferingTypeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "LightUtilization", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelOfferingTypeValues) NoUpfront() *ServicesEc2ModelOfferingTypeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "NoUpfront", "com/amazonaws/services/ec2/model/OfferingTypeValues")
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
	unique_x := &ServicesEc2ModelOfferingTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOfferingTypeValues) SetFieldNoUpfront(val ServicesEc2ModelOfferingTypeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "NoUpfront", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelOfferingTypeValues) PartialUpfront() *ServicesEc2ModelOfferingTypeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "PartialUpfront", "com/amazonaws/services/ec2/model/OfferingTypeValues")
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
	unique_x := &ServicesEc2ModelOfferingTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOfferingTypeValues) SetFieldPartialUpfront(val ServicesEc2ModelOfferingTypeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "PartialUpfront", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelOfferingTypeValues) AllUpfront() *ServicesEc2ModelOfferingTypeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "AllUpfront", "com/amazonaws/services/ec2/model/OfferingTypeValues")
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
	unique_x := &ServicesEc2ModelOfferingTypeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelOfferingTypeValues) SetFieldAllUpfront(val ServicesEc2ModelOfferingTypeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/OfferingTypeValues", "AllUpfront", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


