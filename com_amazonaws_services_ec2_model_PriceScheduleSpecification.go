package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPriceScheduleSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setTerm(java.lang.Long)
	SetTerm(a int64) 

	// public java.lang.Long getTerm()
	GetTerm() int64

	// public com.amazonaws.services.ec2.model.PriceScheduleSpecification withTerm(java.lang.Long)
	WithTerm(a int64) *ServicesEc2ModelPriceScheduleSpecification

	// public void setPrice(java.lang.Double)
	SetPrice(a float64) 

	// public java.lang.Double getPrice()
	GetPrice() float64

	// public com.amazonaws.services.ec2.model.PriceScheduleSpecification withPrice(java.lang.Double)
	WithPrice(a float64) *ServicesEc2ModelPriceScheduleSpecification

	// public void setCurrencyCode(java.lang.String)
	SetCurrencyCode2(a string) 

	// public java.lang.String getCurrencyCode()
	GetCurrencyCode() string

	// public com.amazonaws.services.ec2.model.PriceScheduleSpecification withCurrencyCode(java.lang.String)
	WithCurrencyCode2(a string) *ServicesEc2ModelPriceScheduleSpecification

	// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) 

	// public com.amazonaws.services.ec2.model.PriceScheduleSpecification withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelPriceScheduleSpecification

	// public com.amazonaws.services.ec2.model.PriceScheduleSpecification clone()
	Clone() *ServicesEc2ModelPriceScheduleSpecification
}

type ServicesEc2ModelPriceScheduleSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PriceScheduleSpecification()
func NewServicesEc2ModelPriceScheduleSpecification() (*ServicesEc2ModelPriceScheduleSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PriceScheduleSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPriceScheduleSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setTerm(java.lang.Long)
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) SetTerm(a int64)  {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTerm", javabind.Void, conv_a.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Long getTerm()
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) GetTerm() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTerm", "java/lang/Long")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoLong()
	dst := new(int64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.PriceScheduleSpecification withTerm(java.lang.Long)
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) WithTerm(a int64) *ServicesEc2ModelPriceScheduleSpecification {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTerm", "com/amazonaws/services/ec2/model/PriceScheduleSpecification", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelPriceScheduleSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrice(java.lang.Double)
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) SetPrice(a float64)  {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrice", javabind.Void, conv_a.Value().Cast("java/lang/Double"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Double getPrice()
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) GetPrice() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrice", "java/lang/Double")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDouble()
	dst := new(float64)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.PriceScheduleSpecification withPrice(java.lang.Double)
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) WithPrice(a float64) *ServicesEc2ModelPriceScheduleSpecification {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrice", "com/amazonaws/services/ec2/model/PriceScheduleSpecification", conv_a.Value().Cast("java/lang/Double"))
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
	unique_x := &ServicesEc2ModelPriceScheduleSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) SetCurrencyCode2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCurrencyCode", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCurrencyCode()
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) GetCurrencyCode() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCurrencyCode", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PriceScheduleSpecification withCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) WithCurrencyCode2(a string) *ServicesEc2ModelPriceScheduleSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/PriceScheduleSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPriceScheduleSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCurrencyCode", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/CurrencyCodeValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.PriceScheduleSpecification withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelPriceScheduleSpecification {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/PriceScheduleSpecification", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CurrencyCodeValues"))
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
	unique_x := &ServicesEc2ModelPriceScheduleSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) ToString() string {
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

// public boolean equals(java.lang.Object)
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PriceScheduleSpecification clone()
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) Clone() *ServicesEc2ModelPriceScheduleSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PriceScheduleSpecification")
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
	unique_x := &ServicesEc2ModelPriceScheduleSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPriceScheduleSpecification) Clone2() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}


