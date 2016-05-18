package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPriceScheduleInterface interface {
	JavaLangObjectInterface

	// public void setTerm(java.lang.Long)
	SetTerm(a int64) 

	// public java.lang.Long getTerm()
	GetTerm() int64

	// public com.amazonaws.services.ec2.model.PriceSchedule withTerm(java.lang.Long)
	WithTerm(a int64) *ServicesEc2ModelPriceSchedule

	// public void setPrice(java.lang.Double)
	SetPrice(a float64) 

	// public java.lang.Double getPrice()
	GetPrice() float64

	// public com.amazonaws.services.ec2.model.PriceSchedule withPrice(java.lang.Double)
	WithPrice(a float64) *ServicesEc2ModelPriceSchedule

	// public void setCurrencyCode(java.lang.String)
	SetCurrencyCode2(a string) 

	// public java.lang.String getCurrencyCode()
	GetCurrencyCode() string

	// public com.amazonaws.services.ec2.model.PriceSchedule withCurrencyCode(java.lang.String)
	WithCurrencyCode2(a string) *ServicesEc2ModelPriceSchedule

	// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) 

	// public com.amazonaws.services.ec2.model.PriceSchedule withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelPriceSchedule

	// public void setActive(java.lang.Boolean)
	SetActive(a bool) 

	// public java.lang.Boolean getActive()
	GetActive() bool

	// public com.amazonaws.services.ec2.model.PriceSchedule withActive(java.lang.Boolean)
	WithActive(a bool) *ServicesEc2ModelPriceSchedule

	// public java.lang.Boolean isActive()
	IsActive() bool

	// public com.amazonaws.services.ec2.model.PriceSchedule clone()
	Clone() *ServicesEc2ModelPriceSchedule
}

type ServicesEc2ModelPriceSchedule struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PriceSchedule()
func NewServicesEc2ModelPriceSchedule() (*ServicesEc2ModelPriceSchedule) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PriceSchedule")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPriceSchedule{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setTerm(java.lang.Long)
func (jbobject *ServicesEc2ModelPriceSchedule) SetTerm(a int64)  {
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
func (jbobject *ServicesEc2ModelPriceSchedule) GetTerm() int64 {
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

// public com.amazonaws.services.ec2.model.PriceSchedule withTerm(java.lang.Long)
func (jbobject *ServicesEc2ModelPriceSchedule) WithTerm(a int64) *ServicesEc2ModelPriceSchedule {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTerm", "com/amazonaws/services/ec2/model/PriceSchedule", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelPriceSchedule{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrice(java.lang.Double)
func (jbobject *ServicesEc2ModelPriceSchedule) SetPrice(a float64)  {
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
func (jbobject *ServicesEc2ModelPriceSchedule) GetPrice() float64 {
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

// public com.amazonaws.services.ec2.model.PriceSchedule withPrice(java.lang.Double)
func (jbobject *ServicesEc2ModelPriceSchedule) WithPrice(a float64) *ServicesEc2ModelPriceSchedule {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrice", "com/amazonaws/services/ec2/model/PriceSchedule", conv_a.Value().Cast("java/lang/Double"))
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
	unique_x := &ServicesEc2ModelPriceSchedule{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelPriceSchedule) SetCurrencyCode2(a string)  {
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
func (jbobject *ServicesEc2ModelPriceSchedule) GetCurrencyCode() string {
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

// public com.amazonaws.services.ec2.model.PriceSchedule withCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelPriceSchedule) WithCurrencyCode2(a string) *ServicesEc2ModelPriceSchedule {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/PriceSchedule", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPriceSchedule{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelPriceSchedule) SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface)  {
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

// public com.amazonaws.services.ec2.model.PriceSchedule withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelPriceSchedule) WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelPriceSchedule {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/PriceSchedule", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CurrencyCodeValues"))
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
	unique_x := &ServicesEc2ModelPriceSchedule{}
	unique_x.Callable = dst
	return unique_x
}

// public void setActive(java.lang.Boolean)
func (jbobject *ServicesEc2ModelPriceSchedule) SetActive(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setActive", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getActive()
func (jbobject *ServicesEc2ModelPriceSchedule) GetActive() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getActive", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.PriceSchedule withActive(java.lang.Boolean)
func (jbobject *ServicesEc2ModelPriceSchedule) WithActive(a bool) *ServicesEc2ModelPriceSchedule {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withActive", "com/amazonaws/services/ec2/model/PriceSchedule", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelPriceSchedule{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isActive()
func (jbobject *ServicesEc2ModelPriceSchedule) IsActive() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isActive", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPriceSchedule) ToString() string {
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
func (jbobject *ServicesEc2ModelPriceSchedule) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPriceSchedule) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PriceSchedule clone()
func (jbobject *ServicesEc2ModelPriceSchedule) Clone() *ServicesEc2ModelPriceSchedule {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PriceSchedule")
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
	unique_x := &ServicesEc2ModelPriceSchedule{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPriceSchedule) Clone2() (*JavaLangObject, error) {
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


