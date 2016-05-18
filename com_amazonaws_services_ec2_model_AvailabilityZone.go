package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAvailabilityZoneInterface interface {
	JavaLangObjectInterface

	// public void setZoneName(java.lang.String)
	SetZoneName(a string) 

	// public java.lang.String getZoneName()
	GetZoneName() string

	// public com.amazonaws.services.ec2.model.AvailabilityZone withZoneName(java.lang.String)
	WithZoneName(a string) *ServicesEc2ModelAvailabilityZone

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.AvailabilityZone withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelAvailabilityZone

	// public void setState(com.amazonaws.services.ec2.model.AvailabilityZoneState)
	SetState(a ServicesEc2ModelAvailabilityZoneStateInterface) 

	// public com.amazonaws.services.ec2.model.AvailabilityZone withState(com.amazonaws.services.ec2.model.AvailabilityZoneState)
	WithState(a ServicesEc2ModelAvailabilityZoneStateInterface) *ServicesEc2ModelAvailabilityZone

	// public void setRegionName(java.lang.String)
	SetRegionName(a string) 

	// public java.lang.String getRegionName()
	GetRegionName() string

	// public com.amazonaws.services.ec2.model.AvailabilityZone withRegionName(java.lang.String)
	WithRegionName(a string) *ServicesEc2ModelAvailabilityZone

	// public java.util.List<com.amazonaws.services.ec2.model.AvailabilityZoneMessage> getMessages()
	GetMessages() []*ServicesEc2ModelAvailabilityZoneMessage

	// public void setMessages(java.util.Collection<com.amazonaws.services.ec2.model.AvailabilityZoneMessage>)
	SetMessages(a []*ServicesEc2ModelAvailabilityZoneMessage) 

	// public com.amazonaws.services.ec2.model.AvailabilityZone withMessages(com.amazonaws.services.ec2.model.AvailabilityZoneMessage...)
	WithMessages(a ...*ServicesEc2ModelAvailabilityZoneMessage) *ServicesEc2ModelAvailabilityZone

	// public com.amazonaws.services.ec2.model.AvailabilityZone withMessages(java.util.Collection<com.amazonaws.services.ec2.model.AvailabilityZoneMessage>)
	WithMessages2(a []*ServicesEc2ModelAvailabilityZoneMessage) *ServicesEc2ModelAvailabilityZone

	// public com.amazonaws.services.ec2.model.AvailabilityZone clone()
	Clone() *ServicesEc2ModelAvailabilityZone
}

type ServicesEc2ModelAvailabilityZone struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AvailabilityZone()
func NewServicesEc2ModelAvailabilityZone() (*ServicesEc2ModelAvailabilityZone) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AvailabilityZone")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAvailabilityZone{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setZoneName(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZone) SetZoneName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setZoneName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getZoneName()
func (jbobject *ServicesEc2ModelAvailabilityZone) GetZoneName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getZoneName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AvailabilityZone withZoneName(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZone) WithZoneName(a string) *ServicesEc2ModelAvailabilityZone {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withZoneName", "com/amazonaws/services/ec2/model/AvailabilityZone", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAvailabilityZone{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZone) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelAvailabilityZone) GetState() string {
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

// public com.amazonaws.services.ec2.model.AvailabilityZone withState(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZone) WithState2(a string) *ServicesEc2ModelAvailabilityZone {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/AvailabilityZone", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAvailabilityZone{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.AvailabilityZoneState)
func (jbobject *ServicesEc2ModelAvailabilityZone) SetState(a ServicesEc2ModelAvailabilityZoneStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/AvailabilityZoneState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AvailabilityZone withState(com.amazonaws.services.ec2.model.AvailabilityZoneState)
func (jbobject *ServicesEc2ModelAvailabilityZone) WithState(a ServicesEc2ModelAvailabilityZoneStateInterface) *ServicesEc2ModelAvailabilityZone {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/AvailabilityZone", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AvailabilityZoneState"))
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
	unique_x := &ServicesEc2ModelAvailabilityZone{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRegionName(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZone) SetRegionName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRegionName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRegionName()
func (jbobject *ServicesEc2ModelAvailabilityZone) GetRegionName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRegionName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AvailabilityZone withRegionName(java.lang.String)
func (jbobject *ServicesEc2ModelAvailabilityZone) WithRegionName(a string) *ServicesEc2ModelAvailabilityZone {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRegionName", "com/amazonaws/services/ec2/model/AvailabilityZone", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAvailabilityZone{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.AvailabilityZoneMessage> getMessages()
func (jbobject *ServicesEc2ModelAvailabilityZone) GetMessages() []*ServicesEc2ModelAvailabilityZoneMessage {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMessages", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelAvailabilityZoneMessage)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setMessages(java.util.Collection<com.amazonaws.services.ec2.model.AvailabilityZoneMessage>)
func (jbobject *ServicesEc2ModelAvailabilityZone) SetMessages(a []*ServicesEc2ModelAvailabilityZoneMessage)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMessages", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AvailabilityZone withMessages(com.amazonaws.services.ec2.model.AvailabilityZoneMessage...)
func (jbobject *ServicesEc2ModelAvailabilityZone) WithMessages(a ...*ServicesEc2ModelAvailabilityZoneMessage) *ServicesEc2ModelAvailabilityZone {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/AvailabilityZoneMessage")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMessages", "com/amazonaws/services/ec2/model/AvailabilityZone", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AvailabilityZoneMessage"))
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
	unique_x := &ServicesEc2ModelAvailabilityZone{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AvailabilityZone withMessages(java.util.Collection<com.amazonaws.services.ec2.model.AvailabilityZoneMessage>)
func (jbobject *ServicesEc2ModelAvailabilityZone) WithMessages2(a []*ServicesEc2ModelAvailabilityZoneMessage) *ServicesEc2ModelAvailabilityZone {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMessages", "com/amazonaws/services/ec2/model/AvailabilityZone", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelAvailabilityZone{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAvailabilityZone) ToString() string {
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
func (jbobject *ServicesEc2ModelAvailabilityZone) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAvailabilityZone) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AvailabilityZone clone()
func (jbobject *ServicesEc2ModelAvailabilityZone) Clone() *ServicesEc2ModelAvailabilityZone {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AvailabilityZone")
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
	unique_x := &ServicesEc2ModelAvailabilityZone{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAvailabilityZone) Clone2() (*JavaLangObject, error) {
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


