package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelVolumeStatusEventInterface interface {
	JavaLangObjectInterface

	// public void setEventType(java.lang.String)
	SetEventType(a string) 

	// public java.lang.String getEventType()
	GetEventType() string

	// public com.amazonaws.services.ec2.model.VolumeStatusEvent withEventType(java.lang.String)
	WithEventType(a string) *ServicesEc2ModelVolumeStatusEvent

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.VolumeStatusEvent withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelVolumeStatusEvent

	// public void setNotBefore(java.util.Date)
	SetNotBefore(a time.Time) 

	// public java.util.Date getNotBefore()
	GetNotBefore() time.Time

	// public com.amazonaws.services.ec2.model.VolumeStatusEvent withNotBefore(java.util.Date)
	WithNotBefore(a time.Time) *ServicesEc2ModelVolumeStatusEvent

	// public void setNotAfter(java.util.Date)
	SetNotAfter(a time.Time) 

	// public java.util.Date getNotAfter()
	GetNotAfter() time.Time

	// public com.amazonaws.services.ec2.model.VolumeStatusEvent withNotAfter(java.util.Date)
	WithNotAfter(a time.Time) *ServicesEc2ModelVolumeStatusEvent

	// public void setEventId(java.lang.String)
	SetEventId(a string) 

	// public java.lang.String getEventId()
	GetEventId() string

	// public com.amazonaws.services.ec2.model.VolumeStatusEvent withEventId(java.lang.String)
	WithEventId(a string) *ServicesEc2ModelVolumeStatusEvent

	// public com.amazonaws.services.ec2.model.VolumeStatusEvent clone()
	Clone() *ServicesEc2ModelVolumeStatusEvent
}

type ServicesEc2ModelVolumeStatusEvent struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VolumeStatusEvent()
func NewServicesEc2ModelVolumeStatusEvent() (*ServicesEc2ModelVolumeStatusEvent) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VolumeStatusEvent")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVolumeStatusEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setEventType(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) SetEventType(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEventType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getEventType()
func (jbobject *ServicesEc2ModelVolumeStatusEvent) GetEventType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEventType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VolumeStatusEvent withEventType(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) WithEventType(a string) *ServicesEc2ModelVolumeStatusEvent {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEventType", "com/amazonaws/services/ec2/model/VolumeStatusEvent", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) SetDescription(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDescription", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDescription()
func (jbobject *ServicesEc2ModelVolumeStatusEvent) GetDescription() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDescription", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VolumeStatusEvent withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) WithDescription(a string) *ServicesEc2ModelVolumeStatusEvent {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/VolumeStatusEvent", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNotBefore(java.util.Date)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) SetNotBefore(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNotBefore", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getNotBefore()
func (jbobject *ServicesEc2ModelVolumeStatusEvent) GetNotBefore() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNotBefore", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.VolumeStatusEvent withNotBefore(java.util.Date)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) WithNotBefore(a time.Time) *ServicesEc2ModelVolumeStatusEvent {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNotBefore", "com/amazonaws/services/ec2/model/VolumeStatusEvent", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelVolumeStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNotAfter(java.util.Date)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) SetNotAfter(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNotAfter", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getNotAfter()
func (jbobject *ServicesEc2ModelVolumeStatusEvent) GetNotAfter() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNotAfter", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.VolumeStatusEvent withNotAfter(java.util.Date)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) WithNotAfter(a time.Time) *ServicesEc2ModelVolumeStatusEvent {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNotAfter", "com/amazonaws/services/ec2/model/VolumeStatusEvent", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelVolumeStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEventId(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) SetEventId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEventId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getEventId()
func (jbobject *ServicesEc2ModelVolumeStatusEvent) GetEventId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEventId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VolumeStatusEvent withEventId(java.lang.String)
func (jbobject *ServicesEc2ModelVolumeStatusEvent) WithEventId(a string) *ServicesEc2ModelVolumeStatusEvent {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEventId", "com/amazonaws/services/ec2/model/VolumeStatusEvent", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVolumeStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVolumeStatusEvent) ToString() string {
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
func (jbobject *ServicesEc2ModelVolumeStatusEvent) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVolumeStatusEvent) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VolumeStatusEvent clone()
func (jbobject *ServicesEc2ModelVolumeStatusEvent) Clone() *ServicesEc2ModelVolumeStatusEvent {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VolumeStatusEvent")
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
	unique_x := &ServicesEc2ModelVolumeStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVolumeStatusEvent) Clone2() (*JavaLangObject, error) {
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


