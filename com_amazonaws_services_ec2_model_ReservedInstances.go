package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelReservedInstancesInterface interface {
	JavaLangObjectInterface

	// public void setReservedInstancesId(java.lang.String)
	SetReservedInstancesId(a string) 

	// public java.lang.String getReservedInstancesId()
	GetReservedInstancesId() string

	// public com.amazonaws.services.ec2.model.ReservedInstances withReservedInstancesId(java.lang.String)
	WithReservedInstancesId(a string) *ServicesEc2ModelReservedInstances

	// public void setInstanceType(java.lang.String)
	SetInstanceType2(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceType(java.lang.String)
	WithInstanceType2(a string) *ServicesEc2ModelReservedInstances

	// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	SetInstanceType(a ServicesEc2ModelInstanceTypeInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
	WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelReservedInstances

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.ReservedInstances withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelReservedInstances

	// public void setStart(java.util.Date)
	SetStart(a time.Time) 

	// public java.util.Date getStart()
	GetStart() time.Time

	// public com.amazonaws.services.ec2.model.ReservedInstances withStart(java.util.Date)
	WithStart(a time.Time) *ServicesEc2ModelReservedInstances

	// public void setEnd(java.util.Date)
	SetEnd(a time.Time) 

	// public java.util.Date getEnd()
	GetEnd() time.Time

	// public com.amazonaws.services.ec2.model.ReservedInstances withEnd(java.util.Date)
	WithEnd(a time.Time) *ServicesEc2ModelReservedInstances

	// public void setDuration(java.lang.Long)
	SetDuration(a int64) 

	// public java.lang.Long getDuration()
	GetDuration() int64

	// public com.amazonaws.services.ec2.model.ReservedInstances withDuration(java.lang.Long)
	WithDuration(a int64) *ServicesEc2ModelReservedInstances

	// public void setUsagePrice(java.lang.Float)
	SetUsagePrice(a float32) 

	// public java.lang.Float getUsagePrice()
	GetUsagePrice() float32

	// public com.amazonaws.services.ec2.model.ReservedInstances withUsagePrice(java.lang.Float)
	WithUsagePrice(a float32) *ServicesEc2ModelReservedInstances

	// public void setFixedPrice(java.lang.Float)
	SetFixedPrice(a float32) 

	// public java.lang.Float getFixedPrice()
	GetFixedPrice() float32

	// public com.amazonaws.services.ec2.model.ReservedInstances withFixedPrice(java.lang.Float)
	WithFixedPrice(a float32) *ServicesEc2ModelReservedInstances

	// public void setInstanceCount(java.lang.Integer)
	SetInstanceCount(a int) 

	// public java.lang.Integer getInstanceCount()
	GetInstanceCount() int

	// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceCount(java.lang.Integer)
	WithInstanceCount(a int) *ServicesEc2ModelReservedInstances

	// public void setProductDescription(java.lang.String)
	SetProductDescription2(a string) 

	// public java.lang.String getProductDescription()
	GetProductDescription() string

	// public com.amazonaws.services.ec2.model.ReservedInstances withProductDescription(java.lang.String)
	WithProductDescription2(a string) *ServicesEc2ModelReservedInstances

	// public void setProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
	SetProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstances withProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
	WithProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) *ServicesEc2ModelReservedInstances

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.ReservedInstances withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelReservedInstances

	// public void setState(com.amazonaws.services.ec2.model.ReservedInstanceState)
	SetState(a ServicesEc2ModelReservedInstanceStateInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstances withState(com.amazonaws.services.ec2.model.ReservedInstanceState)
	WithState(a ServicesEc2ModelReservedInstanceStateInterface) *ServicesEc2ModelReservedInstances

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.ReservedInstances withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelReservedInstances

	// public com.amazonaws.services.ec2.model.ReservedInstances withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelReservedInstances

	// public void setInstanceTenancy(java.lang.String)
	SetInstanceTenancy2(a string) 

	// public java.lang.String getInstanceTenancy()
	GetInstanceTenancy() string

	// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceTenancy(java.lang.String)
	WithInstanceTenancy2(a string) *ServicesEc2ModelReservedInstances

	// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	SetInstanceTenancy(a ServicesEc2ModelTenancyInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelReservedInstances

	// public void setCurrencyCode(java.lang.String)
	SetCurrencyCode2(a string) 

	// public java.lang.String getCurrencyCode()
	GetCurrencyCode() string

	// public com.amazonaws.services.ec2.model.ReservedInstances withCurrencyCode(java.lang.String)
	WithCurrencyCode2(a string) *ServicesEc2ModelReservedInstances

	// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstances withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
	WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelReservedInstances

	// public void setOfferingType(java.lang.String)
	SetOfferingType2(a string) 

	// public java.lang.String getOfferingType()
	GetOfferingType() string

	// public com.amazonaws.services.ec2.model.ReservedInstances withOfferingType(java.lang.String)
	WithOfferingType2(a string) *ServicesEc2ModelReservedInstances

	// public void setOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
	SetOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) 

	// public com.amazonaws.services.ec2.model.ReservedInstances withOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
	WithOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) *ServicesEc2ModelReservedInstances

	// public java.util.List<com.amazonaws.services.ec2.model.RecurringCharge> getRecurringCharges()
	GetRecurringCharges() []*ServicesEc2ModelRecurringCharge

	// public void setRecurringCharges(java.util.Collection<com.amazonaws.services.ec2.model.RecurringCharge>)
	SetRecurringCharges(a []*ServicesEc2ModelRecurringCharge) 

	// public com.amazonaws.services.ec2.model.ReservedInstances withRecurringCharges(com.amazonaws.services.ec2.model.RecurringCharge...)
	WithRecurringCharges(a ...*ServicesEc2ModelRecurringCharge) *ServicesEc2ModelReservedInstances

	// public com.amazonaws.services.ec2.model.ReservedInstances withRecurringCharges(java.util.Collection<com.amazonaws.services.ec2.model.RecurringCharge>)
	WithRecurringCharges2(a []*ServicesEc2ModelRecurringCharge) *ServicesEc2ModelReservedInstances

	// public com.amazonaws.services.ec2.model.ReservedInstances clone()
	Clone() *ServicesEc2ModelReservedInstances
}

type ServicesEc2ModelReservedInstances struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ReservedInstances()
func NewServicesEc2ModelReservedInstances() (*ServicesEc2ModelReservedInstances) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ReservedInstances")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReservedInstances{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) SetReservedInstancesId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservedInstancesId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservedInstancesId()
func (jbobject *ServicesEc2ModelReservedInstances) GetReservedInstancesId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservedInstancesId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstances withReservedInstancesId(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) WithReservedInstancesId(a string) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservedInstancesId", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) SetInstanceType2(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstances) GetInstanceType() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) WithInstanceType2(a string) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelReservedInstances) SetInstanceType(a ServicesEc2ModelInstanceTypeInterface)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceType(com.amazonaws.services.ec2.model.InstanceType)
func (jbobject *ServicesEc2ModelReservedInstances) WithInstanceType(a ServicesEc2ModelInstanceTypeInterface) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceType"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstances) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) WithAvailabilityZone(a string) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setStart(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstances) SetStart(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStart", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getStart()
func (jbobject *ServicesEc2ModelReservedInstances) GetStart() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStart", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.ReservedInstances withStart(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstances) WithStart(a time.Time) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withStart", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEnd(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstances) SetEnd(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnd", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getEnd()
func (jbobject *ServicesEc2ModelReservedInstances) GetEnd() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnd", "java/util/Date")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoDate()
	dst := new(time.Time)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.ReservedInstances withEnd(java.util.Date)
func (jbobject *ServicesEc2ModelReservedInstances) WithEnd(a time.Time) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEnd", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDuration(java.lang.Long)
func (jbobject *ServicesEc2ModelReservedInstances) SetDuration(a int64)  {
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
func (jbobject *ServicesEc2ModelReservedInstances) GetDuration() int64 {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withDuration(java.lang.Long)
func (jbobject *ServicesEc2ModelReservedInstances) WithDuration(a int64) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaLong()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDuration", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/Long"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setUsagePrice(java.lang.Float)
func (jbobject *ServicesEc2ModelReservedInstances) SetUsagePrice(a float32)  {
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
func (jbobject *ServicesEc2ModelReservedInstances) GetUsagePrice() float32 {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withUsagePrice(java.lang.Float)
func (jbobject *ServicesEc2ModelReservedInstances) WithUsagePrice(a float32) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaFloat()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUsagePrice", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/Float"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFixedPrice(java.lang.Float)
func (jbobject *ServicesEc2ModelReservedInstances) SetFixedPrice(a float32)  {
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
func (jbobject *ServicesEc2ModelReservedInstances) GetFixedPrice() float32 {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withFixedPrice(java.lang.Float)
func (jbobject *ServicesEc2ModelReservedInstances) WithFixedPrice(a float32) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaFloat()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFixedPrice", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/Float"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelReservedInstances) SetInstanceCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getInstanceCount()
func (jbobject *ServicesEc2ModelReservedInstances) GetInstanceCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceCount", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelReservedInstances) WithInstanceCount(a int) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceCount", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductDescription(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) SetProductDescription2(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstances) GetProductDescription() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withProductDescription(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) WithProductDescription2(a string) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescription", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
func (jbobject *ServicesEc2ModelReservedInstances) SetProductDescription(a ServicesEc2ModelRIProductDescriptionInterface)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withProductDescription(com.amazonaws.services.ec2.model.RIProductDescription)
func (jbobject *ServicesEc2ModelReservedInstances) WithProductDescription(a ServicesEc2ModelRIProductDescriptionInterface) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductDescription", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RIProductDescription"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) SetState2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getState()
func (jbobject *ServicesEc2ModelReservedInstances) GetState() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getState", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ReservedInstances withState(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) WithState2(a string) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.ReservedInstanceState)
func (jbobject *ServicesEc2ModelReservedInstances) SetState(a ServicesEc2ModelReservedInstanceStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstanceState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstances withState(com.amazonaws.services.ec2.model.ReservedInstanceState)
func (jbobject *ServicesEc2ModelReservedInstances) WithState(a ServicesEc2ModelReservedInstanceStateInterface) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ReservedInstanceState"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelReservedInstances) GetTags() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTag)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelReservedInstances) SetTags(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ReservedInstances withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelReservedInstances) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstances withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelReservedInstances) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) SetInstanceTenancy2(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstances) GetInstanceTenancy() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) WithInstanceTenancy2(a string) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelReservedInstances) SetInstanceTenancy(a ServicesEc2ModelTenancyInterface)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelReservedInstances) WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) SetCurrencyCode2(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstances) GetCurrencyCode() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withCurrencyCode(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) WithCurrencyCode2(a string) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelReservedInstances) SetCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withCurrencyCode(com.amazonaws.services.ec2.model.CurrencyCodeValues)
func (jbobject *ServicesEc2ModelReservedInstances) WithCurrencyCode(a ServicesEc2ModelCurrencyCodeValuesInterface) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCurrencyCode", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CurrencyCodeValues"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOfferingType(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) SetOfferingType2(a string)  {
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
func (jbobject *ServicesEc2ModelReservedInstances) GetOfferingType() string {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withOfferingType(java.lang.String)
func (jbobject *ServicesEc2ModelReservedInstances) WithOfferingType2(a string) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOfferingType", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
func (jbobject *ServicesEc2ModelReservedInstances) SetOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withOfferingType(com.amazonaws.services.ec2.model.OfferingTypeValues)
func (jbobject *ServicesEc2ModelReservedInstances) WithOfferingType(a ServicesEc2ModelOfferingTypeValuesInterface) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOfferingType", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("com/amazonaws/services/ec2/model/OfferingTypeValues"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.RecurringCharge> getRecurringCharges()
func (jbobject *ServicesEc2ModelReservedInstances) GetRecurringCharges() []*ServicesEc2ModelRecurringCharge {
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
func (jbobject *ServicesEc2ModelReservedInstances) SetRecurringCharges(a []*ServicesEc2ModelRecurringCharge)  {
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

// public com.amazonaws.services.ec2.model.ReservedInstances withRecurringCharges(com.amazonaws.services.ec2.model.RecurringCharge...)
func (jbobject *ServicesEc2ModelReservedInstances) WithRecurringCharges(a ...*ServicesEc2ModelRecurringCharge) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/RecurringCharge")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRecurringCharges", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RecurringCharge"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ReservedInstances withRecurringCharges(java.util.Collection<com.amazonaws.services.ec2.model.RecurringCharge>)
func (jbobject *ServicesEc2ModelReservedInstances) WithRecurringCharges2(a []*ServicesEc2ModelRecurringCharge) *ServicesEc2ModelReservedInstances {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRecurringCharges", "com/amazonaws/services/ec2/model/ReservedInstances", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReservedInstances) ToString() string {
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
func (jbobject *ServicesEc2ModelReservedInstances) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReservedInstances) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ReservedInstances clone()
func (jbobject *ServicesEc2ModelReservedInstances) Clone() *ServicesEc2ModelReservedInstances {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ReservedInstances")
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
	unique_x := &ServicesEc2ModelReservedInstances{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelReservedInstances) Clone2() (*JavaLangObject, error) {
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


