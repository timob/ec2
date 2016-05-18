package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceStatusInterface interface {
	JavaLangObjectInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.InstanceStatus withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelInstanceStatus

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.InstanceStatus withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelInstanceStatus

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceStatusEvent> getEvents()
	GetEvents() []*ServicesEc2ModelInstanceStatusEvent

	// public void setEvents(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatusEvent>)
	SetEvents(a []*ServicesEc2ModelInstanceStatusEvent) 

	// public com.amazonaws.services.ec2.model.InstanceStatus withEvents(com.amazonaws.services.ec2.model.InstanceStatusEvent...)
	WithEvents(a ...*ServicesEc2ModelInstanceStatusEvent) *ServicesEc2ModelInstanceStatus

	// public com.amazonaws.services.ec2.model.InstanceStatus withEvents(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatusEvent>)
	WithEvents2(a []*ServicesEc2ModelInstanceStatusEvent) *ServicesEc2ModelInstanceStatus

	// public void setInstanceState(com.amazonaws.services.ec2.model.InstanceState)
	SetInstanceState(a ServicesEc2ModelInstanceStateInterface) 

	// public com.amazonaws.services.ec2.model.InstanceState getInstanceState()
	GetInstanceState() *ServicesEc2ModelInstanceState

	// public com.amazonaws.services.ec2.model.InstanceStatus withInstanceState(com.amazonaws.services.ec2.model.InstanceState)
	WithInstanceState(a ServicesEc2ModelInstanceStateInterface) *ServicesEc2ModelInstanceStatus

	// public void setSystemStatus(com.amazonaws.services.ec2.model.InstanceStatusSummary)
	SetSystemStatus(a ServicesEc2ModelInstanceStatusSummaryInterface) 

	// public com.amazonaws.services.ec2.model.InstanceStatusSummary getSystemStatus()
	GetSystemStatus() *ServicesEc2ModelInstanceStatusSummary

	// public com.amazonaws.services.ec2.model.InstanceStatus withSystemStatus(com.amazonaws.services.ec2.model.InstanceStatusSummary)
	WithSystemStatus(a ServicesEc2ModelInstanceStatusSummaryInterface) *ServicesEc2ModelInstanceStatus

	// public void setInstanceStatus(com.amazonaws.services.ec2.model.InstanceStatusSummary)
	SetInstanceStatus(a ServicesEc2ModelInstanceStatusSummaryInterface) 

	// public com.amazonaws.services.ec2.model.InstanceStatusSummary getInstanceStatus()
	GetInstanceStatus() *ServicesEc2ModelInstanceStatusSummary

	// public com.amazonaws.services.ec2.model.InstanceStatus withInstanceStatus(com.amazonaws.services.ec2.model.InstanceStatusSummary)
	WithInstanceStatus(a ServicesEc2ModelInstanceStatusSummaryInterface) *ServicesEc2ModelInstanceStatus

	// public com.amazonaws.services.ec2.model.InstanceStatus clone()
	Clone() *ServicesEc2ModelInstanceStatus
}

type ServicesEc2ModelInstanceStatus struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceStatus()
func NewServicesEc2ModelInstanceStatus() (*ServicesEc2ModelInstanceStatus) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceStatus")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceStatus{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatus) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelInstanceStatus) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceStatus withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatus) WithInstanceId(a string) *ServicesEc2ModelInstanceStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/InstanceStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatus) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceStatus) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.InstanceStatus withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatus) WithAvailabilityZone(a string) *ServicesEc2ModelInstanceStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/InstanceStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceStatusEvent> getEvents()
func (jbobject *ServicesEc2ModelInstanceStatus) GetEvents() []*ServicesEc2ModelInstanceStatusEvent {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEvents", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceStatusEvent)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setEvents(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatusEvent>)
func (jbobject *ServicesEc2ModelInstanceStatus) SetEvents(a []*ServicesEc2ModelInstanceStatusEvent)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEvents", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceStatus withEvents(com.amazonaws.services.ec2.model.InstanceStatusEvent...)
func (jbobject *ServicesEc2ModelInstanceStatus) WithEvents(a ...*ServicesEc2ModelInstanceStatusEvent) *ServicesEc2ModelInstanceStatus {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceStatusEvent")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEvents", "com/amazonaws/services/ec2/model/InstanceStatus", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStatusEvent"))
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
	unique_x := &ServicesEc2ModelInstanceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceStatus withEvents(java.util.Collection<com.amazonaws.services.ec2.model.InstanceStatusEvent>)
func (jbobject *ServicesEc2ModelInstanceStatus) WithEvents2(a []*ServicesEc2ModelInstanceStatusEvent) *ServicesEc2ModelInstanceStatus {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEvents", "com/amazonaws/services/ec2/model/InstanceStatus", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelInstanceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceState(com.amazonaws.services.ec2.model.InstanceState)
func (jbobject *ServicesEc2ModelInstanceStatus) SetInstanceState(a ServicesEc2ModelInstanceStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceState getInstanceState()
func (jbobject *ServicesEc2ModelInstanceStatus) GetInstanceState() *ServicesEc2ModelInstanceState {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceState", "com/amazonaws/services/ec2/model/InstanceState")
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
	unique_x := &ServicesEc2ModelInstanceState{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceStatus withInstanceState(com.amazonaws.services.ec2.model.InstanceState)
func (jbobject *ServicesEc2ModelInstanceStatus) WithInstanceState(a ServicesEc2ModelInstanceStateInterface) *ServicesEc2ModelInstanceStatus {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceState", "com/amazonaws/services/ec2/model/InstanceStatus", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceState"))
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
	unique_x := &ServicesEc2ModelInstanceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public void setSystemStatus(com.amazonaws.services.ec2.model.InstanceStatusSummary)
func (jbobject *ServicesEc2ModelInstanceStatus) SetSystemStatus(a ServicesEc2ModelInstanceStatusSummaryInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSystemStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStatusSummary"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceStatusSummary getSystemStatus()
func (jbobject *ServicesEc2ModelInstanceStatus) GetSystemStatus() *ServicesEc2ModelInstanceStatusSummary {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSystemStatus", "com/amazonaws/services/ec2/model/InstanceStatusSummary")
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
	unique_x := &ServicesEc2ModelInstanceStatusSummary{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceStatus withSystemStatus(com.amazonaws.services.ec2.model.InstanceStatusSummary)
func (jbobject *ServicesEc2ModelInstanceStatus) WithSystemStatus(a ServicesEc2ModelInstanceStatusSummaryInterface) *ServicesEc2ModelInstanceStatus {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSystemStatus", "com/amazonaws/services/ec2/model/InstanceStatus", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStatusSummary"))
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
	unique_x := &ServicesEc2ModelInstanceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceStatus(com.amazonaws.services.ec2.model.InstanceStatusSummary)
func (jbobject *ServicesEc2ModelInstanceStatus) SetInstanceStatus(a ServicesEc2ModelInstanceStatusSummaryInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStatusSummary"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceStatusSummary getInstanceStatus()
func (jbobject *ServicesEc2ModelInstanceStatus) GetInstanceStatus() *ServicesEc2ModelInstanceStatusSummary {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceStatus", "com/amazonaws/services/ec2/model/InstanceStatusSummary")
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
	unique_x := &ServicesEc2ModelInstanceStatusSummary{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceStatus withInstanceStatus(com.amazonaws.services.ec2.model.InstanceStatusSummary)
func (jbobject *ServicesEc2ModelInstanceStatus) WithInstanceStatus(a ServicesEc2ModelInstanceStatusSummaryInterface) *ServicesEc2ModelInstanceStatus {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceStatus", "com/amazonaws/services/ec2/model/InstanceStatus", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceStatusSummary"))
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
	unique_x := &ServicesEc2ModelInstanceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceStatus) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceStatus) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceStatus) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceStatus clone()
func (jbobject *ServicesEc2ModelInstanceStatus) Clone() *ServicesEc2ModelInstanceStatus {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceStatus")
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
	unique_x := &ServicesEc2ModelInstanceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceStatus) Clone2() (*JavaLangObject, error) {
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


