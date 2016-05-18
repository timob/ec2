package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelScheduledInstanceInterface interface {
	JavaLangObjectInterface

	// public void setScheduledInstanceId(java.lang.String)
	SetScheduledInstanceId(a string) 

	// public java.lang.String getScheduledInstanceId()
	GetScheduledInstanceId() string

	// public com.amazonaws.services.ec2.model.ScheduledInstance withScheduledInstanceId(java.lang.String)
	WithScheduledInstanceId(a string) *ServicesEc2ModelScheduledInstance

	// public void setInstanceType(java.lang.String)
	SetInstanceType(a string) 

	// public java.lang.String getInstanceType()
	GetInstanceType() string

	// public com.amazonaws.services.ec2.model.ScheduledInstance withInstanceType(java.lang.String)
	WithInstanceType(a string) *ServicesEc2ModelScheduledInstance

	// public void setPlatform(java.lang.String)
	SetPlatform(a string) 

	// public java.lang.String getPlatform()
	GetPlatform() string

	// public com.amazonaws.services.ec2.model.ScheduledInstance withPlatform(java.lang.String)
	WithPlatform(a string) *ServicesEc2ModelScheduledInstance

	// public void setNetworkPlatform(java.lang.String)
	SetNetworkPlatform(a string) 

	// public java.lang.String getNetworkPlatform()
	GetNetworkPlatform() string

	// public com.amazonaws.services.ec2.model.ScheduledInstance withNetworkPlatform(java.lang.String)
	WithNetworkPlatform(a string) *ServicesEc2ModelScheduledInstance

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.ScheduledInstance withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelScheduledInstance

	// public void setSlotDurationInHours(java.lang.Integer)
	SetSlotDurationInHours(a int) 

	// public java.lang.Integer getSlotDurationInHours()
	GetSlotDurationInHours() int

	// public com.amazonaws.services.ec2.model.ScheduledInstance withSlotDurationInHours(java.lang.Integer)
	WithSlotDurationInHours(a int) *ServicesEc2ModelScheduledInstance

	// public void setRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence)
	SetRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceInterface) 

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence getRecurrence()
	GetRecurrence() *ServicesEc2ModelScheduledInstanceRecurrence

	// public com.amazonaws.services.ec2.model.ScheduledInstance withRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence)
	WithRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceInterface) *ServicesEc2ModelScheduledInstance

	// public void setPreviousSlotEndTime(java.util.Date)
	SetPreviousSlotEndTime(a time.Time) 

	// public java.util.Date getPreviousSlotEndTime()
	GetPreviousSlotEndTime() time.Time

	// public com.amazonaws.services.ec2.model.ScheduledInstance withPreviousSlotEndTime(java.util.Date)
	WithPreviousSlotEndTime(a time.Time) *ServicesEc2ModelScheduledInstance

	// public void setNextSlotStartTime(java.util.Date)
	SetNextSlotStartTime(a time.Time) 

	// public java.util.Date getNextSlotStartTime()
	GetNextSlotStartTime() time.Time

	// public com.amazonaws.services.ec2.model.ScheduledInstance withNextSlotStartTime(java.util.Date)
	WithNextSlotStartTime(a time.Time) *ServicesEc2ModelScheduledInstance

	// public void setHourlyPrice(java.lang.String)
	SetHourlyPrice(a string) 

	// public java.lang.String getHourlyPrice()
	GetHourlyPrice() string

	// public com.amazonaws.services.ec2.model.ScheduledInstance withHourlyPrice(java.lang.String)
	WithHourlyPrice(a string) *ServicesEc2ModelScheduledInstance

	// public void setTotalScheduledInstanceHours(java.lang.Integer)
	SetTotalScheduledInstanceHours(a int) 

	// public java.lang.Integer getTotalScheduledInstanceHours()
	GetTotalScheduledInstanceHours() int

	// public com.amazonaws.services.ec2.model.ScheduledInstance withTotalScheduledInstanceHours(java.lang.Integer)
	WithTotalScheduledInstanceHours(a int) *ServicesEc2ModelScheduledInstance

	// public void setInstanceCount(java.lang.Integer)
	SetInstanceCount(a int) 

	// public java.lang.Integer getInstanceCount()
	GetInstanceCount() int

	// public com.amazonaws.services.ec2.model.ScheduledInstance withInstanceCount(java.lang.Integer)
	WithInstanceCount(a int) *ServicesEc2ModelScheduledInstance

	// public void setTermStartDate(java.util.Date)
	SetTermStartDate(a time.Time) 

	// public java.util.Date getTermStartDate()
	GetTermStartDate() time.Time

	// public com.amazonaws.services.ec2.model.ScheduledInstance withTermStartDate(java.util.Date)
	WithTermStartDate(a time.Time) *ServicesEc2ModelScheduledInstance

	// public void setTermEndDate(java.util.Date)
	SetTermEndDate(a time.Time) 

	// public java.util.Date getTermEndDate()
	GetTermEndDate() time.Time

	// public com.amazonaws.services.ec2.model.ScheduledInstance withTermEndDate(java.util.Date)
	WithTermEndDate(a time.Time) *ServicesEc2ModelScheduledInstance

	// public void setCreateDate(java.util.Date)
	SetCreateDate(a time.Time) 

	// public java.util.Date getCreateDate()
	GetCreateDate() time.Time

	// public com.amazonaws.services.ec2.model.ScheduledInstance withCreateDate(java.util.Date)
	WithCreateDate(a time.Time) *ServicesEc2ModelScheduledInstance

	// public com.amazonaws.services.ec2.model.ScheduledInstance clone()
	Clone() *ServicesEc2ModelScheduledInstance
}

type ServicesEc2ModelScheduledInstance struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ScheduledInstance()
func NewServicesEc2ModelScheduledInstance() (*ServicesEc2ModelScheduledInstance) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ScheduledInstance")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelScheduledInstance{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setScheduledInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) SetScheduledInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setScheduledInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getScheduledInstanceId()
func (jbobject *ServicesEc2ModelScheduledInstance) GetScheduledInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScheduledInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withScheduledInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) WithScheduledInstanceId(a string) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withScheduledInstanceId", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) SetInstanceType(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstance) GetInstanceType() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withInstanceType(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) WithInstanceType(a string) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceType", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) SetPlatform(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstance) GetPlatform() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) WithPlatform(a string) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPlatform", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNetworkPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) SetNetworkPlatform(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstance) GetNetworkPlatform() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withNetworkPlatform(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) WithNetworkPlatform(a string) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkPlatform", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstance) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) WithAvailabilityZone(a string) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSlotDurationInHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstance) SetSlotDurationInHours(a int)  {
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
func (jbobject *ServicesEc2ModelScheduledInstance) GetSlotDurationInHours() int {
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withSlotDurationInHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstance) WithSlotDurationInHours(a int) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSlotDurationInHours", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence)
func (jbobject *ServicesEc2ModelScheduledInstance) SetRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceInterface)  {
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
func (jbobject *ServicesEc2ModelScheduledInstance) GetRecurrence() *ServicesEc2ModelScheduledInstanceRecurrence {
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withRecurrence(com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence)
func (jbobject *ServicesEc2ModelScheduledInstance) WithRecurrence(a ServicesEc2ModelScheduledInstanceRecurrenceInterface) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRecurrence", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPreviousSlotEndTime(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) SetPreviousSlotEndTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPreviousSlotEndTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getPreviousSlotEndTime()
func (jbobject *ServicesEc2ModelScheduledInstance) GetPreviousSlotEndTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPreviousSlotEndTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withPreviousSlotEndTime(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) WithPreviousSlotEndTime(a time.Time) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPreviousSlotEndTime", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextSlotStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) SetNextSlotStartTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNextSlotStartTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getNextSlotStartTime()
func (jbobject *ServicesEc2ModelScheduledInstance) GetNextSlotStartTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNextSlotStartTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withNextSlotStartTime(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) WithNextSlotStartTime(a time.Time) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextSlotStartTime", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHourlyPrice(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) SetHourlyPrice(a string)  {
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
func (jbobject *ServicesEc2ModelScheduledInstance) GetHourlyPrice() string {
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withHourlyPrice(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstance) WithHourlyPrice(a string) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHourlyPrice", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTotalScheduledInstanceHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstance) SetTotalScheduledInstanceHours(a int)  {
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
func (jbobject *ServicesEc2ModelScheduledInstance) GetTotalScheduledInstanceHours() int {
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withTotalScheduledInstanceHours(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstance) WithTotalScheduledInstanceHours(a int) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTotalScheduledInstanceHours", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstance) SetInstanceCount(a int)  {
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
func (jbobject *ServicesEc2ModelScheduledInstance) GetInstanceCount() int {
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withInstanceCount(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstance) WithInstanceCount(a int) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceCount", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTermStartDate(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) SetTermStartDate(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTermStartDate", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getTermStartDate()
func (jbobject *ServicesEc2ModelScheduledInstance) GetTermStartDate() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTermStartDate", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withTermStartDate(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) WithTermStartDate(a time.Time) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTermStartDate", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTermEndDate(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) SetTermEndDate(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTermEndDate", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getTermEndDate()
func (jbobject *ServicesEc2ModelScheduledInstance) GetTermEndDate() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTermEndDate", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withTermEndDate(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) WithTermEndDate(a time.Time) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTermEndDate", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCreateDate(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) SetCreateDate(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreateDate", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getCreateDate()
func (jbobject *ServicesEc2ModelScheduledInstance) GetCreateDate() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreateDate", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.ScheduledInstance withCreateDate(java.util.Date)
func (jbobject *ServicesEc2ModelScheduledInstance) WithCreateDate(a time.Time) *ServicesEc2ModelScheduledInstance {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateDate", "com/amazonaws/services/ec2/model/ScheduledInstance", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelScheduledInstance) ToString() string {
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
func (jbobject *ServicesEc2ModelScheduledInstance) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelScheduledInstance) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ScheduledInstance clone()
func (jbobject *ServicesEc2ModelScheduledInstance) Clone() *ServicesEc2ModelScheduledInstance {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ScheduledInstance")
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
	unique_x := &ServicesEc2ModelScheduledInstance{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelScheduledInstance) Clone2() (*JavaLangObject, error) {
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


