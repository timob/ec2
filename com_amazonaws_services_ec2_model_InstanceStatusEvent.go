package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelInstanceStatusEventInterface interface {
	JavaLangObjectInterface

	// public void setCode(java.lang.String)
	SetCode2(a string) 

	// public java.lang.String getCode()
	GetCode() string

	// public com.amazonaws.services.ec2.model.InstanceStatusEvent withCode(java.lang.String)
	WithCode2(a string) *ServicesEc2ModelInstanceStatusEvent

	// public void setCode(com.amazonaws.services.ec2.model.EventCode)
	SetCode(a ServicesEc2ModelEventCodeInterface) 

	// public com.amazonaws.services.ec2.model.InstanceStatusEvent withCode(com.amazonaws.services.ec2.model.EventCode)
	WithCode(a ServicesEc2ModelEventCodeInterface) *ServicesEc2ModelInstanceStatusEvent

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.InstanceStatusEvent withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelInstanceStatusEvent

	// public void setNotBefore(java.util.Date)
	SetNotBefore(a time.Time) 

	// public java.util.Date getNotBefore()
	GetNotBefore() time.Time

	// public com.amazonaws.services.ec2.model.InstanceStatusEvent withNotBefore(java.util.Date)
	WithNotBefore(a time.Time) *ServicesEc2ModelInstanceStatusEvent

	// public void setNotAfter(java.util.Date)
	SetNotAfter(a time.Time) 

	// public java.util.Date getNotAfter()
	GetNotAfter() time.Time

	// public com.amazonaws.services.ec2.model.InstanceStatusEvent withNotAfter(java.util.Date)
	WithNotAfter(a time.Time) *ServicesEc2ModelInstanceStatusEvent

	// public com.amazonaws.services.ec2.model.InstanceStatusEvent clone()
	Clone() *ServicesEc2ModelInstanceStatusEvent
}

type ServicesEc2ModelInstanceStatusEvent struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceStatusEvent()
func NewServicesEc2ModelInstanceStatusEvent() (*ServicesEc2ModelInstanceStatusEvent) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceStatusEvent")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceStatusEvent{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setCode(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) SetCode2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCode", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCode()
func (jbobject *ServicesEc2ModelInstanceStatusEvent) GetCode() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCode", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceStatusEvent withCode(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) WithCode2(a string) *ServicesEc2ModelInstanceStatusEvent {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCode", "com/amazonaws/services/ec2/model/InstanceStatusEvent", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCode(com.amazonaws.services.ec2.model.EventCode)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) SetCode(a ServicesEc2ModelEventCodeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCode", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/EventCode"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceStatusEvent withCode(com.amazonaws.services.ec2.model.EventCode)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) WithCode(a ServicesEc2ModelEventCodeInterface) *ServicesEc2ModelInstanceStatusEvent {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCode", "com/amazonaws/services/ec2/model/InstanceStatusEvent", conv_a.Value().Cast("com/amazonaws/services/ec2/model/EventCode"))
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
	unique_x := &ServicesEc2ModelInstanceStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelInstanceStatusEvent) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.InstanceStatusEvent withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) WithDescription(a string) *ServicesEc2ModelInstanceStatusEvent {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/InstanceStatusEvent", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNotBefore(java.util.Date)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) SetNotBefore(a time.Time)  {
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
func (jbobject *ServicesEc2ModelInstanceStatusEvent) GetNotBefore() time.Time {
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

// public com.amazonaws.services.ec2.model.InstanceStatusEvent withNotBefore(java.util.Date)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) WithNotBefore(a time.Time) *ServicesEc2ModelInstanceStatusEvent {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNotBefore", "com/amazonaws/services/ec2/model/InstanceStatusEvent", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelInstanceStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNotAfter(java.util.Date)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) SetNotAfter(a time.Time)  {
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
func (jbobject *ServicesEc2ModelInstanceStatusEvent) GetNotAfter() time.Time {
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

// public com.amazonaws.services.ec2.model.InstanceStatusEvent withNotAfter(java.util.Date)
func (jbobject *ServicesEc2ModelInstanceStatusEvent) WithNotAfter(a time.Time) *ServicesEc2ModelInstanceStatusEvent {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNotAfter", "com/amazonaws/services/ec2/model/InstanceStatusEvent", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelInstanceStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceStatusEvent) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceStatusEvent) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceStatusEvent) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceStatusEvent clone()
func (jbobject *ServicesEc2ModelInstanceStatusEvent) Clone() *ServicesEc2ModelInstanceStatusEvent {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceStatusEvent")
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
	unique_x := &ServicesEc2ModelInstanceStatusEvent{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceStatusEvent) Clone2() (*JavaLangObject, error) {
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


