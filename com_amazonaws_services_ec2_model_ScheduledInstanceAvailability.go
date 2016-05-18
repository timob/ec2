package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelScheduledInstanceAvailabilityInterface interface {
	JavaLangObjectInterface

	// public void setInstanceType(java.lang.String)
	SetInstanceType(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withInstanceType(java.lang.String)
	WithInstanceType(a string) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setPlatform(java.lang.String)
	SetPlatform(a string) 

	// public java.lang.String getPlatform()
	GetPlatform() string

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withPlatform(java.lang.String)
	WithPlatform(a string) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setNetworkPlatform(java.lang.String)
	SetNetworkPlatform(a string) 

	// public java.lang.String getNetworkPlatform()
	GetNetworkPlatform() string

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withNetworkPlatform(java.lang.String)
	WithNetworkPlatform(a string) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setPurchaseToken(java.lang.String)
	SetPurchaseToken(a string) 

	// public java.lang.String getPurchaseToken()
	GetPurchaseToken() string

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withPurchaseToken(java.lang.String)
	WithPurchaseToken(a string) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setSlotDurationInHours(java.lang.Integer)
	SetSlotDurationInHours(a int) 

	// public java.lang.Integer getSlotDurationInHours()
	GetSlotDurationInHours() int

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withSlotDurationInHours(java.lang.Integer)
	WithSlotDurationInHours(a int) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence)
	SetRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceInterface) 

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence getRecurrence()
	GetRecurrence() *ServicesEc2ModelScheduledInstanceRecurrence

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence)
	WithRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceInterface) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setFirstSlotStartTime(java.util.Date)
	SetFirstSlotStartTime(a time.Time) 

	// public java.util.Date getFirstSlotStartTime()
	GetFirstSlotStartTime() time.Time

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withFirstSlotStartTime(java.util.Date)
	WithFirstSlotStartTime(a time.Time) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setHourlyPrice(java.lang.String)
	SetHourlyPrice(a string) 

	// public java.lang.String getHourlyPrice()
	GetHourlyPrice() string

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withHourlyPrice(java.lang.String)
	WithHourlyPrice(a string) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setTotalScheduledInstanceHours(java.lang.Integer)
	SetTotalScheduledInstanceHours(a int) 

	// public java.lang.Integer getTotalScheduledInstanceHours()
	GetTotalScheduledInstanceHours() int

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withTotalScheduledInstanceHours(java.lang.Integer)
	WithTotalScheduledInstanceHours(a int) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setAvailableInstanceCount(java.lang.Integer)
	SetAvailableInstanceCount(a int) 

	// public java.lang.Integer getAvailableInstanceCount()
	GetAvailableInstanceCount() int

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withAvailableInstanceCount(java.lang.Integer)
	WithAvailableInstanceCount(a int) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setMinTermDurationInDays(java.lang.Integer)
	SetMinTermDurationInDays(a int) 

	// public java.lang.Integer getMinTermDurationInDays()
	GetMinTermDurationInDays() int

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withMinTermDurationInDays(java.lang.Integer)
	WithMinTermDurationInDays(a int) *ServicesEc2ModelScheduledInstanceAvailability

	// public void setMaxTermDurationInDays(java.lang.Integer)
	SetMaxTermDurationInDays(a int) 

	// public java.lang.Integer getMaxTermDurationInDays()
	GetMaxTermDurationInDays() int

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withMaxTermDurationInDays(java.lang.Integer)
	WithMaxTermDurationInDays(a int) *ServicesEc2ModelScheduledInstanceAvailability

	// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability clone()
	Clone() *ServicesEc2ModelScheduledInstanceAvailability
}

type ServicesEc2ModelScheduledInstanceAvailability struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability()
func NewServicesEc2ModelScheduledInstanceAvailability() (*ServicesEc2ModelScheduledInstanceAvailability) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ScheduledInstanceAvailability")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelScheduledInstanceAvailability{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetInstanceType(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetInstanceType() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithInstanceType(a string) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetPlatform(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPlatform", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPlatform()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetPlatform() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPlatform", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithPlatform(a string) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNetworkPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetNetworkPlatform(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkPlatform", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNetworkPlatform()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetNetworkPlatform() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkPlatform", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withNetworkPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithNetworkPlatform(a string) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkPlatform", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithAvailabilityZone(a string) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPurchaseToken(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetPurchaseToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPurchaseToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPurchaseToken()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetPurchaseToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPurchaseToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withPurchaseToken(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithPurchaseToken(a string) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPurchaseToken", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSlotDurationInHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetSlotDurationInHours(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSlotDurationInHours", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getSlotDurationInHours()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetSlotDurationInHours() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSlotDurationInHours", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withSlotDurationInHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithSlotDurationInHours(a int) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSlotDurationInHours", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRecurrence", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence getRecurrence()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetRecurrence() *ServicesEc2ModelScheduledInstanceRecurrence {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRecurrence", "com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence")
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
	unique_x := &ServicesEc2ModelScheduledInstanceRecurrence{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceInterface) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRecurrence", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFirstSlotStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetFirstSlotStartTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFirstSlotStartTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getFirstSlotStartTime()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetFirstSlotStartTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFirstSlotStartTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withFirstSlotStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithFirstSlotStartTime(a time.Time) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFirstSlotStartTime", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHourlyPrice(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetHourlyPrice(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHourlyPrice", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getHourlyPrice()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetHourlyPrice() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHourlyPrice", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withHourlyPrice(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithHourlyPrice(a string) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHourlyPrice", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTotalScheduledInstanceHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetTotalScheduledInstanceHours(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTotalScheduledInstanceHours", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getTotalScheduledInstanceHours()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetTotalScheduledInstanceHours() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTotalScheduledInstanceHours", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withTotalScheduledInstanceHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithTotalScheduledInstanceHours(a int) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTotalScheduledInstanceHours", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailableInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetAvailableInstanceCount(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailableInstanceCount", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getAvailableInstanceCount()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetAvailableInstanceCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailableInstanceCount", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withAvailableInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithAvailableInstanceCount(a int) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailableInstanceCount", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMinTermDurationInDays(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetMinTermDurationInDays(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMinTermDurationInDays", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getMinTermDurationInDays()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetMinTermDurationInDays() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMinTermDurationInDays", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withMinTermDurationInDays(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithMinTermDurationInDays(a int) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMinTermDurationInDays", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMaxTermDurationInDays(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) SetMaxTermDurationInDays(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMaxTermDurationInDays", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getMaxTermDurationInDays()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) GetMaxTermDurationInDays() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMaxTermDurationInDays", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability withMaxTermDurationInDays(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) WithMaxTermDurationInDays(a int) *ServicesEc2ModelScheduledInstanceAvailability {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMaxTermDurationInDays", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) ToString() string {
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
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ScheduledInstanceAvailability clone()
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) Clone() *ServicesEc2ModelScheduledInstanceAvailability {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ScheduledInstanceAvailability")
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
	unique_x := &ServicesEc2ModelScheduledInstanceAvailability{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelScheduledInstanceAvailability) Clone2() (*JavaLangObject, error) {
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


