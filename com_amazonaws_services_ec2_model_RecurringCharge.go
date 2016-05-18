package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRecurringChargeInterface interface {
	JavaLangObjectInterface

	// public void setFrequency(java.lang.String)
	SetFrequency2(a string) 

	// public java.lang.String getFrequency()
	GetFrequency() string

	// public com.amazonaws.services.ec2.model.RecurringCharge withFrequency(java.lang.String)
	WithFrequency2(a string) *ServicesEc2ModelRecurringCharge

	// public void setFrequency(com.amazonaws.services.ec2.model.RecurringChargeFrequency)
	SetFrequency(a ServicesEc2ModelRecurringChargeFrequencyInterface) 

	// public com.amazonaws.services.ec2.model.RecurringCharge withFrequency(com.amazonaws.services.ec2.model.RecurringChargeFrequency)
	WithFrequency(a ServicesEc2ModelRecurringChargeFrequencyInterface) *ServicesEc2ModelRecurringCharge

	// public void setAmount(java.lang.Double)
	SetAmount(a float64) 

	// public java.lang.Double getAmount()
	GetAmount() float64

	// public com.amazonaws.services.ec2.model.RecurringCharge withAmount(java.lang.Double)
	WithAmount(a float64) *ServicesEc2ModelRecurringCharge

	// public com.amazonaws.services.ec2.model.RecurringCharge clone()
	Clone() *ServicesEc2ModelRecurringCharge
}

type ServicesEc2ModelRecurringCharge struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.RecurringCharge()
func NewServicesEc2ModelRecurringCharge() (*ServicesEc2ModelRecurringCharge) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/RecurringCharge")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRecurringCharge{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setFrequency(java.lang.String)
func (jbobject *ServicesEc2ModelRecurringCharge) SetFrequency2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFrequency", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getFrequency()
func (jbobject *ServicesEc2ModelRecurringCharge) GetFrequency() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFrequency", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.RecurringCharge withFrequency(java.lang.String)
func (jbobject *ServicesEc2ModelRecurringCharge) WithFrequency2(a string) *ServicesEc2ModelRecurringCharge {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFrequency", "com/amazonaws/services/ec2/model/RecurringCharge", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRecurringCharge{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFrequency(com.amazonaws.services.ec2.model.RecurringChargeFrequency)
func (jbobject *ServicesEc2ModelRecurringCharge) SetFrequency(a ServicesEc2ModelRecurringChargeFrequencyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFrequency", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RecurringChargeFrequency"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.RecurringCharge withFrequency(com.amazonaws.services.ec2.model.RecurringChargeFrequency)
func (jbobject *ServicesEc2ModelRecurringCharge) WithFrequency(a ServicesEc2ModelRecurringChargeFrequencyInterface) *ServicesEc2ModelRecurringCharge {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFrequency", "com/amazonaws/services/ec2/model/RecurringCharge", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RecurringChargeFrequency"))
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
	unique_x := &ServicesEc2ModelRecurringCharge{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAmount(java.lang.Double)
func (jbobject *ServicesEc2ModelRecurringCharge) SetAmount(a float64)  {
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
func (jbobject *ServicesEc2ModelRecurringCharge) GetAmount() float64 {
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

// public com.amazonaws.services.ec2.model.RecurringCharge withAmount(java.lang.Double)
func (jbobject *ServicesEc2ModelRecurringCharge) WithAmount(a float64) *ServicesEc2ModelRecurringCharge {
	conv_a := javabind.NewGoToJavaDouble()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAmount", "com/amazonaws/services/ec2/model/RecurringCharge", conv_a.Value().Cast("java/lang/Double"))
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
	unique_x := &ServicesEc2ModelRecurringCharge{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRecurringCharge) ToString() string {
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
func (jbobject *ServicesEc2ModelRecurringCharge) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRecurringCharge) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.RecurringCharge clone()
func (jbobject *ServicesEc2ModelRecurringCharge) Clone() *ServicesEc2ModelRecurringCharge {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/RecurringCharge")
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
	unique_x := &ServicesEc2ModelRecurringCharge{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelRecurringCharge) Clone2() (*JavaLangObject, error) {
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


