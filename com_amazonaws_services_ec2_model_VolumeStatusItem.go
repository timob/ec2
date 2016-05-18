package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVolumeStatusItemInterface interface {
	JavaLangObjectInterface

	// public void setVolumeId(java.lang.String)
	SetVolumeId(a string) 

	// public java.lang.String getVolumeId()
	GetVolumeId() string

	// public com.amazonaws.services.ec2.model.VolumeStatusItem withVolumeId(java.lang.String)
	WithVolumeId(a string) *ServicesEc2ModelVolumeStatusItem

	// public void setAvailabilityZone(java.lang.String)
	SetAvailabilityZone(a string) 

	// public java.lang.String getAvailabilityZone()
	GetAvailabilityZone() string

	// public com.amazonaws.services.ec2.model.VolumeStatusItem withAvailabilityZone(java.lang.String)
	WithAvailabilityZone(a string) *ServicesEc2ModelVolumeStatusItem

	// public void setVolumeStatus(com.amazonaws.services.ec2.model.VolumeStatusInfo)
	SetVolumeStatus(a ServicesEc2ModelVolumeStatusInfoInterface) 

	// public com.amazonaws.services.ec2.model.VolumeStatusInfo getVolumeStatus()
	GetVolumeStatus() *ServicesEc2ModelVolumeStatusInfo

	// public com.amazonaws.services.ec2.model.VolumeStatusItem withVolumeStatus(com.amazonaws.services.ec2.model.VolumeStatusInfo)
	WithVolumeStatus(a ServicesEc2ModelVolumeStatusInfoInterface) *ServicesEc2ModelVolumeStatusItem

	// public java.util.List<com.amazonaws.services.ec2.model.VolumeStatusEvent> getEvents()
	GetEvents() []*ServicesEc2ModelVolumeStatusEvent

	// public void setEvents(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusEvent>)
	SetEvents(a []*ServicesEc2ModelVolumeStatusEvent) 

	// public com.amazonaws.services.ec2.model.VolumeStatusItem withEvents(com.amazonaws.services.ec2.model.VolumeStatusEvent...)
	WithEvents(a ...*ServicesEc2ModelVolumeStatusEvent) *ServicesEc2ModelVolumeStatusItem

	// public com.amazonaws.services.ec2.model.VolumeStatusItem withEvents(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusEvent>)
	WithEvents2(a []*ServicesEc2ModelVolumeStatusEvent) *ServicesEc2ModelVolumeStatusItem

	// public java.util.List<com.amazonaws.services.ec2.model.VolumeStatusAction> getActions()
	GetActions() []*ServicesEc2ModelVolumeStatusAction

	// public void setActions(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusAction>)
	SetActions(a []*ServicesEc2ModelVolumeStatusAction) 

	// public com.amazonaws.services.ec2.model.VolumeStatusItem withActions(com.amazonaws.services.ec2.model.VolumeStatusAction...)
	WithActions(a ...*ServicesEc2ModelVolumeStatusAction) *ServicesEc2ModelVolumeStatusItem

	// public com.amazonaws.services.ec2.model.VolumeStatusItem withActions(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusAction>)
	WithActions2(a []*ServicesEc2ModelVolumeStatusAction) *ServicesEc2ModelVolumeStatusItem

	// public com.amazonaws.services.ec2.model.VolumeStatusItem clone()
	Clone() *ServicesEc2ModelVolumeStatusItem
}

type ServicesEc2ModelVolumeStatusItem struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VolumeStatusItem()
func NewServicesEc2ModelVolumeStatusItem() (*ServicesEc2ModelVolumeStatusItem) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VolumeStatusItem")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVolumeStatusItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVolumeId(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusItem) SetVolumeId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVolumeId()
func (jbobject *ServicesEc2ModelVolumeStatusItem) GetVolumeId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VolumeStatusItem withVolumeId(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusItem) WithVolumeId(a string) *ServicesEc2ModelVolumeStatusItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeId", "com/amazonaws/services/ec2/model/VolumeStatusItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusItem) SetAvailabilityZone(a string)  {
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
func (jbobject *ServicesEc2ModelVolumeStatusItem) GetAvailabilityZone() string {
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

// public com.amazonaws.services.ec2.model.VolumeStatusItem withAvailabilityZone(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusItem) WithAvailabilityZone(a string) *ServicesEc2ModelVolumeStatusItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAvailabilityZone", "com/amazonaws/services/ec2/model/VolumeStatusItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVolumeStatus(com.amazonaws.services.ec2.model.VolumeStatusInfo)
func (jbobject *ServicesEc2ModelVolumeStatusItem) SetVolumeStatus(a ServicesEc2ModelVolumeStatusInfoInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeStatus", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusInfo"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VolumeStatusInfo getVolumeStatus()
func (jbobject *ServicesEc2ModelVolumeStatusItem) GetVolumeStatus() *ServicesEc2ModelVolumeStatusInfo {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeStatus", "com/amazonaws/services/ec2/model/VolumeStatusInfo")
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
	unique_x := &ServicesEc2ModelVolumeStatusInfo{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VolumeStatusItem withVolumeStatus(com.amazonaws.services.ec2.model.VolumeStatusInfo)
func (jbobject *ServicesEc2ModelVolumeStatusItem) WithVolumeStatus(a ServicesEc2ModelVolumeStatusInfoInterface) *ServicesEc2ModelVolumeStatusItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeStatus", "com/amazonaws/services/ec2/model/VolumeStatusItem", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusInfo"))
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
	unique_x := &ServicesEc2ModelVolumeStatusItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.VolumeStatusEvent> getEvents()
func (jbobject *ServicesEc2ModelVolumeStatusItem) GetEvents() []*ServicesEc2ModelVolumeStatusEvent {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEvents", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVolumeStatusEvent)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setEvents(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusEvent>)
func (jbobject *ServicesEc2ModelVolumeStatusItem) SetEvents(a []*ServicesEc2ModelVolumeStatusEvent)  {
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

// public com.amazonaws.services.ec2.model.VolumeStatusItem withEvents(com.amazonaws.services.ec2.model.VolumeStatusEvent...)
func (jbobject *ServicesEc2ModelVolumeStatusItem) WithEvents(a ...*ServicesEc2ModelVolumeStatusEvent) *ServicesEc2ModelVolumeStatusItem {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VolumeStatusEvent")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEvents", "com/amazonaws/services/ec2/model/VolumeStatusItem", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusEvent"))
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
	unique_x := &ServicesEc2ModelVolumeStatusItem{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VolumeStatusItem withEvents(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusEvent>)
func (jbobject *ServicesEc2ModelVolumeStatusItem) WithEvents2(a []*ServicesEc2ModelVolumeStatusEvent) *ServicesEc2ModelVolumeStatusItem {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEvents", "com/amazonaws/services/ec2/model/VolumeStatusItem", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVolumeStatusItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.VolumeStatusAction> getActions()
func (jbobject *ServicesEc2ModelVolumeStatusItem) GetActions() []*ServicesEc2ModelVolumeStatusAction {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getActions", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVolumeStatusAction)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setActions(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusAction>)
func (jbobject *ServicesEc2ModelVolumeStatusItem) SetActions(a []*ServicesEc2ModelVolumeStatusAction)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setActions", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VolumeStatusItem withActions(com.amazonaws.services.ec2.model.VolumeStatusAction...)
func (jbobject *ServicesEc2ModelVolumeStatusItem) WithActions(a ...*ServicesEc2ModelVolumeStatusAction) *ServicesEc2ModelVolumeStatusItem {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VolumeStatusAction")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withActions", "com/amazonaws/services/ec2/model/VolumeStatusItem", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VolumeStatusAction"))
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
	unique_x := &ServicesEc2ModelVolumeStatusItem{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VolumeStatusItem withActions(java.util.Collection<com.amazonaws.services.ec2.model.VolumeStatusAction>)
func (jbobject *ServicesEc2ModelVolumeStatusItem) WithActions2(a []*ServicesEc2ModelVolumeStatusAction) *ServicesEc2ModelVolumeStatusItem {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withActions", "com/amazonaws/services/ec2/model/VolumeStatusItem", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVolumeStatusItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeStatusItem) ToString() string {
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
func (jbobject *ServicesEc2ModelVolumeStatusItem) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVolumeStatusItem) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VolumeStatusItem clone()
func (jbobject *ServicesEc2ModelVolumeStatusItem) Clone() *ServicesEc2ModelVolumeStatusItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VolumeStatusItem")
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
	unique_x := &ServicesEc2ModelVolumeStatusItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVolumeStatusItem) Clone2() (*JavaLangObject, error) {
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


