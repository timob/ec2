package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReservedInstanceLimitPriceInterface interface {
	JavaLangObjectInterface

	// public void setAmount(java.lang.Double)
	SetAmount(a float64) 

	// public java.lang.Double getAmount()
	GetAmount() float64

	// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice withAmount(java.lang.Double)
	WithAmount(a float64) *ServicesEc2ModelReservedInstanceLimitPrice

	// public void setCurrencyCode(java.lang.String)
	SetCurrencyCode2(a string) 

	// public java.lang.String getCurrencyCode()
	GetCurrencyCode() string

	// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice withCurrencyCode(java.lang.String)
	WithCurrencyCode2(a string) *ServicesEc2ModelReservedInstanceLimitPrice

	// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelReservedInstanceLimitPrice

	// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice clone()
	Clone() *ServicesEc2ModelReservedInstanceLimitPrice
}

type ServicesEc2ModelReservedInstanceLimitPrice struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice()
func NewServicesEc2ModelReservedInstanceLimitPrice() (*ServicesEc2ModelReservedInstanceLimitPrice) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReservedInstanceLimitPrice")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReservedInstanceLimitPrice{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAmount(java.lang.Double)
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) SetAmount(a float64)  {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAmount", javabind.Void, conv_a.Value().Cast("java/lang/Double"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Double getAmount()
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) GetAmount() float64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAmount", "java/lang/Double")
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

// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice withAmount(java.lang.Double)
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) WithAmount(a float64) *ServicesEc2ModelReservedInstanceLimitPrice {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAmount", "com/amazonaws/services/ec2/model/ReservedInstanceLimitPrice", conv_a.Value().Cast("java/lang/Double"))
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
	unique_x := &ServicesEc2ModelReservedInstanceLimitPrice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) SetCurrencyCode2(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) GetCurrencyCode() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice withCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) WithCurrencyCode2(a string) *ServicesEc2ModelReservedInstanceLimitPrice {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/ReservedInstanceLimitPrice", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstanceLimitPrice{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelReservedInstanceLimitPrice {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/ReservedInstanceLimitPrice", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CurrencyCodeValues"))
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
	unique_x := &ServicesEc2ModelReservedInstanceLimitPrice{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) ToString() string {
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
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReservedInstanceLimitPrice clone()
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) Clone() *ServicesEc2ModelReservedInstanceLimitPrice {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReservedInstanceLimitPrice")
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
	unique_x := &ServicesEc2ModelReservedInstanceLimitPrice{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelReservedInstanceLimitPrice) Clone2() (*JavaLangObject, error) {
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


