package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReservedInstancesOfferingInterface interface {
	JavaLangObjectInterface

	// public void setReservedInstancesOfferingId(java.lang.String)
	SetReservedInstancesOfferingId(a string) 

	// public java.lang.String getReservedInstancesOfferingId()
	GetReservedInstancesOfferingId() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withReservedInstancesOfferingId(java.lang.String)
	WithReservedInstancesOfferingId(a string) *ServicesEc2ModelReservedInstancesOffering

	// public void setInstanceType(java.lang.String)
	SetInstanceType2(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withInstanceType(java.lang.String)
	WithInstanceType2(a string) *ServicesEc2ModelReservedInstancesOffering

	// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	SetInstanceType(a ServicesEc2ModelInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelReservedInstancesOffering

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelReservedInstancesOffering

	// public void setDuration(java.lang.Long)
	SetDuration(a int64) 

	// public java.lang.Long getDuration()
	GetDuration() int64

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withDuration(java.lang.Long)
	WithDuration(a int64) *ServicesEc2ModelReservedInstancesOffering

	// public void setUsagePrice(java.lang.Float)
	SetUsagePrice(a float32) 

	// public java.lang.Float getUsagePrice()
	GetUsagePrice() float32

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withUsagePrice(java.lang.Float)
	WithUsagePrice(a float32) *ServicesEc2ModelReservedInstancesOffering

	// public void setFixedPrice(java.lang.Float)
	SetFixedPrice(a float32) 

	// public java.lang.Float getFixedPrice()
	GetFixedPrice() float32

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withFixedPrice(java.lang.Float)
	WithFixedPrice(a float32) *ServicesEc2ModelReservedInstancesOffering

	// public void setProductDescription(java.lang.String)
	SetProductDescription2(a string) 

	// public java.lang.String getProductDescription()
	GetProductDescription() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withProductDescription(java.lang.String)
	WithProductDescription2(a string) *ServicesEc2ModelReservedInstancesOffering

	// public void setProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
	SetProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
	WithProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) *ServicesEc2ModelReservedInstancesOffering

	// public void setInstanceTenancy(java.lang.String)
	SetInstanceTenancy2(a string) 

	// public java.lang.String getInstanceTenancy()
	GetInstanceTenancy() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withInstanceTenancy(java.lang.String)
	WithInstanceTenancy2(a string) *ServicesEc2ModelReservedInstancesOffering

	// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	SetInstanceTenancy(a ServicesEc2ModelTenancyInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelReservedInstancesOffering

	// public void setCurrencyCode(java.lang.String)
	SetCurrencyCode2(a string) 

	// public java.lang.String getCurrencyCode()
	GetCurrencyCode() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withCurrencyCode(java.lang.String)
	WithCurrencyCode2(a string) *ServicesEc2ModelReservedInstancesOffering

	// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelReservedInstancesOffering

	// public void setOfferingType(java.lang.String)
	SetOfferingType2(a string) 

	// public java.lang.String getOfferingType()
	GetOfferingType() string

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withOfferingType(java.lang.String)
	WithOfferingType2(a string) *ServicesEc2ModelReservedInstancesOffering

	// public void setOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
	SetOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
	WithOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) *ServicesEc2ModelReservedInstancesOffering

	// public java.util.List<com.amazonaws.services.ec2.model.RecurringCharge> getRecurringCharges()
	GetRecurringCharges() []*ServicesEc2ModelRecurringCharge

	// public void setRecurringCharges(java.util.Collection<com.amazonaws.services.ec2.model.RecurringCharge>)
	SetRecurringCharges(a []*ServicesEc2ModelRecurringCharge) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withRecurringCharges(com.amazonaws.services.ec2.model.RecurringCharge...)
	WithRecurringCharges(a ...*ServicesEc2ModelRecurringCharge) *ServicesEc2ModelReservedInstancesOffering

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withRecurringCharges(java.util.Collection<com.amazonaws.services.ec2.model.RecurringCharge>)
	WithRecurringCharges2(a []*ServicesEc2ModelRecurringCharge) *ServicesEc2ModelReservedInstancesOffering

	// public void setMarketplace(java.lang.Boolean)
	SetMarketplace(a bool) 

	// public java.lang.Boolean getMarketplace()
	GetMarketplace() bool

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withMarketplace(java.lang.Boolean)
	WithMarketplace(a bool) *ServicesEc2ModelReservedInstancesOffering

	// public java.lang.Boolean isMarketplace()
	IsMarketplace() bool

	// public java.util.List<com.amazonaws.services.ec2.model.PricingDetail> getPricingDetails()
	GetPricingDetails() []*ServicesEc2ModelPricingDetail

	// public void setPricingDetails(java.util.Collection<com.amazonaws.services.ec2.model.PricingDetail>)
	SetPricingDetails(a []*ServicesEc2ModelPricingDetail) 

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withPricingDetails(com.amazonaws.services.ec2.model.PricingDetail...)
	WithPricingDetails(a ...*ServicesEc2ModelPricingDetail) *ServicesEc2ModelReservedInstancesOffering

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withPricingDetails(java.util.Collection<com.amazonaws.services.ec2.model.PricingDetail>)
	WithPricingDetails2(a []*ServicesEc2ModelPricingDetail) *ServicesEc2ModelReservedInstancesOffering

	// public com.amazonaws.services.ec2.model.ReservedInstancesOffering clone()
	Clone() *ServicesEc2ModelReservedInstancesOffering
}

type ServicesEc2ModelReservedInstancesOffering struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering()
func NewServicesEc2ModelReservedInstancesOffering() (*ServicesEc2ModelReservedInstancesOffering) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReservedInstancesOffering")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReservedInstancesOffering{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesOfferingId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetReservedInstancesOfferingId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesOfferingId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservedInstancesOfferingId()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetReservedInstancesOfferingId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesOfferingId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withReservedInstancesOfferingId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithReservedInstancesOfferingId(a string) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesOfferingId", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetInstanceType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceType()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetInstanceType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithInstanceType2(a string) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetInstanceType(a ServicesEc2ModelInstanceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetAvailabilityZone(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailabilityZone", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAvailabilityZone()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetAvailabilityZone() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailabilityZone", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithAvailabilityZone(a string) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDuration(java.lang.Long)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetDuration(a int64)  {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDuration", javabind.Void, conv_a.Value().Cast("java/lang/Long"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Long getDuration()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetDuration() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDuration", "java/lang/Long")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withDuration(java.lang.Long)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithDuration(a int64) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDuration", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUsagePrice(java.lang.Float)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetUsagePrice(a float32)  {
	conv_a := javabind.NewGoToJavaFloat()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUsagePrice", javabind.Void, conv_a.Value().Cast("java/lang/Float"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Float getUsagePrice()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetUsagePrice() float32 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUsagePrice", "java/lang/Float")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoFloat()
	dst := new(float32)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withUsagePrice(java.lang.Float)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithUsagePrice(a float32) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaFloat()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUsagePrice", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/Float"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFixedPrice(java.lang.Float)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetFixedPrice(a float32)  {
	conv_a := javabind.NewGoToJavaFloat()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFixedPrice", javabind.Void, conv_a.Value().Cast("java/lang/Float"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Float getFixedPrice()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetFixedPrice() float32 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFixedPrice", "java/lang/Float")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoFloat()
	dst := new(float32)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withFixedPrice(java.lang.Float)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithFixedPrice(a float32) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaFloat()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFixedPrice", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/Float"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductDescription(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetProductDescription2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getProductDescription()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetProductDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withProductDescription(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithProductDescription2(a string) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescription", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetProductDescription(a ServicesEc2ModelRIProductDescriptionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductDescription", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RIProductDescription"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescription", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RIProductDescription"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetInstanceTenancy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceTenancy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceTenancy()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetInstanceTenancy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceTenancy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithInstanceTenancy2(a string) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetInstanceTenancy(a ServicesEc2ModelTenancyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceTenancy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetCurrencyCode2(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetCurrencyCode() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithCurrencyCode2(a string) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CurrencyCodeValues"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOfferingType(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetOfferingType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOfferingType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOfferingType()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetOfferingType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOfferingType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withOfferingType(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithOfferingType2(a string) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOfferingType", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOfferingType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/OfferingTypeValues"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOfferingType", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("com/amazonaws/services/ec2/model/OfferingTypeValues"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.RecurringCharge> getRecurringCharges()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetRecurringCharges() []*ServicesEc2ModelRecurringCharge {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRecurringCharges", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelRecurringCharge)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRecurringCharges(java.util.Collection<com.amazonaws.services.ec2.model.RecurringCharge>)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetRecurringCharges(a []*ServicesEc2ModelRecurringCharge)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRecurringCharges", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withRecurringCharges(com.amazonaws.services.ec2.model.RecurringCharge...)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithRecurringCharges(a ...*ServicesEc2ModelRecurringCharge) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/RecurringCharge")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRecurringCharges", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RecurringCharge"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withRecurringCharges(java.util.Collection<com.amazonaws.services.ec2.model.RecurringCharge>)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithRecurringCharges2(a []*ServicesEc2ModelRecurringCharge) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRecurringCharges", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMarketplace(java.lang.Boolean)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetMarketplace(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMarketplace", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getMarketplace()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetMarketplace() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMarketplace", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withMarketplace(java.lang.Boolean)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithMarketplace(a bool) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMarketplace", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isMarketplace()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) IsMarketplace() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isMarketplace", "java/lang/Boolean")
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

// public java.util.List<com.amazonaws.services.ec2.model.PricingDetail> getPricingDetails()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) GetPricingDetails() []*ServicesEc2ModelPricingDetail {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPricingDetails", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelPricingDetail)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPricingDetails(java.util.Collection<com.amazonaws.services.ec2.model.PricingDetail>)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) SetPricingDetails(a []*ServicesEc2ModelPricingDetail)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPricingDetails", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withPricingDetails(com.amazonaws.services.ec2.model.PricingDetail...)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithPricingDetails(a ...*ServicesEc2ModelPricingDetail) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/PricingDetail")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPricingDetails", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PricingDetail"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering withPricingDetails(java.util.Collection<com.amazonaws.services.ec2.model.PricingDetail>)
func (jbobject *ServicesEc2ModelReservedInstancesOffering) WithPricingDetails2(a []*ServicesEc2ModelPricingDetail) *ServicesEc2ModelReservedInstancesOffering {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPricingDetails", "com/amazonaws/services/ec2/model/ReservedInstancesOffering", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) ToString() string {
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
func (jbobject *ServicesEc2ModelReservedInstancesOffering) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReservedInstancesOffering) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReservedInstancesOffering clone()
func (jbobject *ServicesEc2ModelReservedInstancesOffering) Clone() *ServicesEc2ModelReservedInstancesOffering {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReservedInstancesOffering")
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
	unique_x := &ServicesEc2ModelReservedInstancesOffering{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelReservedInstancesOffering) Clone2() (*JavaLangObject, error) {
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


