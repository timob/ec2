package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCurrencyCodeValuesInterface interface {

	// public static com.amazonaws.services.ec2.model.CurrencyCodeValues[] values()
	Values() []*ServicesEc2ModelCurrencyCodeValues

	// public static com.amazonaws.services.ec2.model.CurrencyCodeValues valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelCurrencyCodeValues

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.CurrencyCodeValues fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelCurrencyCodeValues
}

type ServicesEc2ModelCurrencyCodeValues struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.CurrencyCodeValues[] values()
func (jbobject *ServicesEc2ModelCurrencyCodeValues) Values() []*ServicesEc2ModelCurrencyCodeValues {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/CurrencyCodeValues", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/CurrencyCodeValues"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/CurrencyCodeValues")
	dst := new([]*ServicesEc2ModelCurrencyCodeValues)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.CurrencyCodeValues valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelCurrencyCodeValues) ValueOf(a string) *ServicesEc2ModelCurrencyCodeValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/CurrencyCodeValues", "valueOf", "com/amazonaws/services/ec2/model/CurrencyCodeValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCurrencyCodeValues{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCurrencyCodeValues) ToString() string {
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

// public static com.amazonaws.services.ec2.model.CurrencyCodeValues fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelCurrencyCodeValues) FromValue(a string) *ServicesEc2ModelCurrencyCodeValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/CurrencyCodeValues", "fromValue", "com/amazonaws/services/ec2/model/CurrencyCodeValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCurrencyCodeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCurrencyCodeValues) USD() *ServicesEc2ModelCurrencyCodeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CurrencyCodeValues", "USD", "com/amazonaws/services/ec2/model/CurrencyCodeValues")
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
	unique_x := &ServicesEc2ModelCurrencyCodeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCurrencyCodeValues) SetFieldUSD(val ServicesEc2ModelCurrencyCodeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CurrencyCodeValues", "USD", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


