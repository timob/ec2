package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelHostInterface interface {
	JavaLangObjectInterface

	// public void setHostId(java.lang.String)
	SetHostId(a string) 

	// public java.lang.String getHostId()
	GetHostId() string

	// public com.amazonaws.services.ec2.model.Host withHostId(java.lang.String)
	WithHostId(a string) *ServicesEc2ModelHost

	// public void setAutoPlacement(java.lang.String)
	SetAutoPlacement2(a string) 

	// public java.lang.String getAutoPlacement()
	GetAutoPlacement() string

	// public com.amazonaws.services.ec2.model.Host withAutoPlacement(java.lang.String)
	WithAutoPlacement2(a string) *ServicesEc2ModelHost

	// public void setAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
	SetAutoPlacement(a ServicesEc2ModelAutoPlacementInterface) 

	// public com.amazonaws.services.ec2.model.Host withAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
	WithAutoPlacement(a ServicesEc2ModelAutoPlacementInterface) *ServicesEc2ModelHost

	// public void setHostReservationId(java.lang.String)
	SetHostReservationId(a string) 

	// public java.lang.String getHostReservationId()
	GetHostReservationId() string

	// public com.amazonaws.services.ec2.model.Host withHostReservationId(java.lang.String)
	WithHostReservationId(a string) *ServicesEc2ModelHost

	// public void setClientToken(java.lang.String)
	SetClientToken(a string) 

	// public java.lang.String getClientToken()
	GetClientToken() string

	// public com.amazonaws.services.ec2.model.Host withClientToken(java.lang.String)
	WithClientToken(a string) *ServicesEc2ModelHost

	// public void setHostProperties(com.amazonaws.services.ec2.model.HostProperties)
	SetHostProperties(a ServicesEc2ModelHostPropertiesInterface) 

	// public com.amazonaws.services.ec2.model.HostProperties getHostProperties()
	GetHostProperties() *ServicesEc2ModelHostProperties

	// public com.amazonaws.services.ec2.model.Host withHostProperties(com.amazonaws.services.ec2.model.HostProperties)
	WithHostProperties(a ServicesEc2ModelHostPropertiesInterface) *ServicesEc2ModelHost

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.Host withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelHost

	// public void setState(com.amazonaws.services.ec2.model.AllocationState)
	SetState(a ServicesEc2ModelAllocationStateInterface) 

	// public com.amazonaws.services.ec2.model.Host withState(com.amazonaws.services.ec2.model.AllocationState)
	WithState(a ServicesEc2ModelAllocationStateInterface) *ServicesEc2ModelHost

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.Host withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelHost

	// public java.util.List<com.amazonaws.services.ec2.model.HostInstance> getInstances()
	GetInstances() []*ServicesEc2ModelHostInstance

	// public void setInstances(java.util.Collection<com.amazonaws.services.ec2.model.HostInstance>)
	SetInstances(a []*ServicesEc2ModelHostInstance) 

	// public com.amazonaws.services.ec2.model.Host withInstances(com.amazonaws.services.ec2.model.HostInstance...)
	WithInstances(a ...*ServicesEc2ModelHostInstance) *ServicesEc2ModelHost

	// public com.amazonaws.services.ec2.model.Host withInstances(java.util.Collection<com.amazonaws.services.ec2.model.HostInstance>)
	WithInstances2(a []*ServicesEc2ModelHostInstance) *ServicesEc2ModelHost

	// public void setAvailableCapacity(com.amazonaws.services.ec2.model.AvailableCapacity)
	SetAvailableCapacity(a ServicesEc2ModelAvailableCapacityInterface) 

	// public com.amazonaws.services.ec2.model.AvailableCapacity getAvailableCapacity()
	GetAvailableCapacity() *ServicesEc2ModelAvailableCapacity

	// public com.amazonaws.services.ec2.model.Host withAvailableCapacity(com.amazonaws.services.ec2.model.AvailableCapacity)
	WithAvailableCapacity(a ServicesEc2ModelAvailableCapacityInterface) *ServicesEc2ModelHost

	// public com.amazonaws.services.ec2.model.Host clone()
	Clone() *ServicesEc2ModelHost
}

type ServicesEc2ModelHost struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.Host()
func NewServicesEc2ModelHost() (*ServicesEc2ModelHost) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Host")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelHost{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setHostId(java.lang.String)
func (jbobject *ServicesEc2ModelHost) SetHostId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHostId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getHostId()
func (jbobject *ServicesEc2ModelHost) GetHostId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHostId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Host withHostId(java.lang.String)
func (jbobject *ServicesEc2ModelHost) WithHostId(a string) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHostId", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAutoPlacement(java.lang.String)
func (jbobject *ServicesEc2ModelHost) SetAutoPlacement2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAutoPlacement", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAutoPlacement()
func (jbobject *ServicesEc2ModelHost) GetAutoPlacement() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAutoPlacement", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Host withAutoPlacement(java.lang.String)
func (jbobject *ServicesEc2ModelHost) WithAutoPlacement2(a string) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAutoPlacement", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
func (jbobject *ServicesEc2ModelHost) SetAutoPlacement(a ServicesEc2ModelAutoPlacementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAutoPlacement", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AutoPlacement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Host withAutoPlacement(com.amazonaws.services.ec2.model.AutoPlacement)
func (jbobject *ServicesEc2ModelHost) WithAutoPlacement(a ServicesEc2ModelAutoPlacementInterface) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAutoPlacement", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AutoPlacement"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHostReservationId(java.lang.String)
func (jbobject *ServicesEc2ModelHost) SetHostReservationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHostReservationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getHostReservationId()
func (jbobject *ServicesEc2ModelHost) GetHostReservationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHostReservationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Host withHostReservationId(java.lang.String)
func (jbobject *ServicesEc2ModelHost) WithHostReservationId(a string) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHostReservationId", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public void setClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelHost) SetClientToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClientToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getClientToken()
func (jbobject *ServicesEc2ModelHost) GetClientToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClientToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Host withClientToken(java.lang.String)
func (jbobject *ServicesEc2ModelHost) WithClientToken(a string) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withClientToken", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public void setHostProperties(com.amazonaws.services.ec2.model.HostProperties)
func (jbobject *ServicesEc2ModelHost) SetHostProperties(a ServicesEc2ModelHostPropertiesInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setHostProperties", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/HostProperties"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.HostProperties getHostProperties()
func (jbobject *ServicesEc2ModelHost) GetHostProperties() *ServicesEc2ModelHostProperties {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getHostProperties", "com/amazonaws/services/ec2/model/HostProperties")
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
	unique_x := &ServicesEc2ModelHostProperties{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Host withHostProperties(com.amazonaws.services.ec2.model.HostProperties)
func (jbobject *ServicesEc2ModelHost) WithHostProperties(a ServicesEc2ModelHostPropertiesInterface) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withHostProperties", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("com/amazonaws/services/ec2/model/HostProperties"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelHost) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelHost) GetState() string {
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

// public com.amazonaws.services.ec2.model.Host withState(java.lang.String)
func (jbobject *ServicesEc2ModelHost) WithState2(a string) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.AllocationState)
func (jbobject *ServicesEc2ModelHost) SetState(a ServicesEc2ModelAllocationStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AllocationState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Host withState(com.amazonaws.services.ec2.model.AllocationState)
func (jbobject *ServicesEc2ModelHost) WithState(a ServicesEc2ModelAllocationStateInterface) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AllocationState"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelHost) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelHost) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.Host withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelHost) WithAvailabilityZone(a string) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.HostInstance> getInstances()
func (jbobject *ServicesEc2ModelHost) GetInstances() []*ServicesEc2ModelHostInstance {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstances", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelHostInstance)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setInstances(java.util.Collection<com.amazonaws.services.ec2.model.HostInstance>)
func (jbobject *ServicesEc2ModelHost) SetInstances(a []*ServicesEc2ModelHostInstance)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstances", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Host withInstances(com.amazonaws.services.ec2.model.HostInstance...)
func (jbobject *ServicesEc2ModelHost) WithInstances(a ...*ServicesEc2ModelHostInstance) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/HostInstance")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstances", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("com/amazonaws/services/ec2/model/HostInstance"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Host withInstances(java.util.Collection<com.amazonaws.services.ec2.model.HostInstance>)
func (jbobject *ServicesEc2ModelHost) WithInstances2(a []*ServicesEc2ModelHostInstance) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstances", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailableCapacity(com.amazonaws.services.ec2.model.AvailableCapacity)
func (jbobject *ServicesEc2ModelHost) SetAvailableCapacity(a ServicesEc2ModelAvailableCapacityInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAvailableCapacity", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AvailableCapacity"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AvailableCapacity getAvailableCapacity()
func (jbobject *ServicesEc2ModelHost) GetAvailableCapacity() *ServicesEc2ModelAvailableCapacity {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailableCapacity", "com/amazonaws/services/ec2/model/AvailableCapacity")
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
	unique_x := &ServicesEc2ModelAvailableCapacity{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Host withAvailableCapacity(com.amazonaws.services.ec2.model.AvailableCapacity)
func (jbobject *ServicesEc2ModelHost) WithAvailableCapacity(a ServicesEc2ModelAvailableCapacityInterface) *ServicesEc2ModelHost {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailableCapacity", "com/amazonaws/services/ec2/model/Host", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AvailableCapacity"))
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelHost) ToString() string {
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
func (jbobject *ServicesEc2ModelHost) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelHost) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.Host clone()
func (jbobject *ServicesEc2ModelHost) Clone() *ServicesEc2ModelHost {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/Host")
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
	unique_x := &ServicesEc2ModelHost{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelHost) Clone2() (*JavaLangObject, error) {
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


