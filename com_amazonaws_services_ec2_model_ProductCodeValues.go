package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelProductCodeValuesInterface interface {

	// public static com.amazonaws.services.ec2.model.ProductCodeValues[] values()
	Values() []*ServicesEc2ModelProductCodeValues

	// public static com.amazonaws.services.ec2.model.ProductCodeValues valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelProductCodeValues

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ProductCodeValues fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelProductCodeValues
}

type ServicesEc2ModelProductCodeValues struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ProductCodeValues[] values()
func (jbobject *ServicesEc2ModelProductCodeValues) Values() []*ServicesEc2ModelProductCodeValues {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ProductCodeValues", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ProductCodeValues"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ProductCodeValues")
	dst := new([]*ServicesEc2ModelProductCodeValues)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ProductCodeValues valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelProductCodeValues) ValueOf(a string) *ServicesEc2ModelProductCodeValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ProductCodeValues", "valueOf", "com/amazonaws/services/ec2/model/ProductCodeValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelProductCodeValues{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelProductCodeValues) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ProductCodeValues fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelProductCodeValues) FromValue(a string) *ServicesEc2ModelProductCodeValues {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ProductCodeValues", "fromValue", "com/amazonaws/services/ec2/model/ProductCodeValues", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelProductCodeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelProductCodeValues) Devpay() *ServicesEc2ModelProductCodeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ProductCodeValues", "Devpay", "com/amazonaws/services/ec2/model/ProductCodeValues")
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
	unique_x := &ServicesEc2ModelProductCodeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelProductCodeValues) SetFieldDevpay(val ServicesEc2ModelProductCodeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ProductCodeValues", "Devpay", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelProductCodeValues) Marketplace() *ServicesEc2ModelProductCodeValues {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ProductCodeValues", "Marketplace", "com/amazonaws/services/ec2/model/ProductCodeValues")
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
	unique_x := &ServicesEc2ModelProductCodeValues{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelProductCodeValues) SetFieldMarketplace(val ServicesEc2ModelProductCodeValuesInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ProductCodeValues", "Marketplace", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


