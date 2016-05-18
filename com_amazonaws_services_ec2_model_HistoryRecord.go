package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelHistoryRecordInterface interface {
	JavaLangObjectInterface

	// public void setTimestamp(java.util.Date)
	SetTimestamp(a time.Time) 

	// public java.util.Date getTimestamp()
	GetTimestamp() time.Time

	// public com.amazonaws.services.ec2.model.HistoryRecord withTimestamp(java.util.Date)
	WithTimestamp(a time.Time) *ServicesEc2ModelHistoryRecord

	// public void setEventType(java.lang.String)
	SetEventType2(a string) 

	// public java.lang.String getEventType()
	GetEventType() string

	// public com.amazonaws.services.ec2.model.HistoryRecord withEventType(java.lang.String)
	WithEventType2(a string) *ServicesEc2ModelHistoryRecord

	// public void setEventType(com.amazonaws.services.ec2.model.EventType)
	SetEventType(a ServicesEc2ModelEventTypeInterface) 

	// public com.amazonaws.services.ec2.model.HistoryRecord withEventType(com.amazonaws.services.ec2.model.EventType)
	WithEventType(a ServicesEc2ModelEventTypeInterface) *ServicesEc2ModelHistoryRecord

	// public void setEventInformation(com.amazonaws.services.ec2.model.EventInformation)
	SetEventInformation(a ServicesEc2ModelEventInformationInterface) 

	// public com.amazonaws.services.ec2.model.EventInformation getEventInformation()
	GetEventInformation() *ServicesEc2ModelEventInformation

	// public com.amazonaws.services.ec2.model.HistoryRecord withEventInformation(com.amazonaws.services.ec2.model.EventInformation)
	WithEventInformation(a ServicesEc2ModelEventInformationInterface) *ServicesEc2ModelHistoryRecord

	// public com.amazonaws.services.ec2.model.HistoryRecord clone()
	Clone() *ServicesEc2ModelHistoryRecord
}

type ServicesEc2ModelHistoryRecord struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.HistoryRecord()
func NewServicesEc2ModelHistoryRecord() (*ServicesEc2ModelHistoryRecord) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/HistoryRecord")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelHistoryRecord{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setTimestamp(java.util.Date)
func (jbobject *ServicesEc2ModelHistoryRecord) SetTimestamp(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTimestamp", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getTimestamp()
func (jbobject *ServicesEc2ModelHistoryRecord) GetTimestamp() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTimestamp", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.HistoryRecord withTimestamp(java.util.Date)
func (jbobject *ServicesEc2ModelHistoryRecord) WithTimestamp(a time.Time) *ServicesEc2ModelHistoryRecord {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTimestamp", "com/amazonaws/services/ec2/model/HistoryRecord", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelHistoryRecord{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEventType(java.lang.String)
func (jbobject *ServicesEc2ModelHistoryRecord) SetEventType2(a string)  {
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
func (jbobject *ServicesEc2ModelHistoryRecord) GetEventType() string {
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

// public com.amazonaws.services.ec2.model.HistoryRecord withEventType(java.lang.String)
func (jbobject *ServicesEc2ModelHistoryRecord) WithEventType2(a string) *ServicesEc2ModelHistoryRecord {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEventType", "com/amazonaws/services/ec2/model/HistoryRecord", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHistoryRecord{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEventType(com.amazonaws.services.ec2.model.EventType)
func (jbobject *ServicesEc2ModelHistoryRecord) SetEventType(a ServicesEc2ModelEventTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEventType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/EventType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.HistoryRecord withEventType(com.amazonaws.services.ec2.model.EventType)
func (jbobject *ServicesEc2ModelHistoryRecord) WithEventType(a ServicesEc2ModelEventTypeInterface) *ServicesEc2ModelHistoryRecord {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEventType", "com/amazonaws/services/ec2/model/HistoryRecord", conv_a.Value().Cast("com/amazonaws/services/ec2/model/EventType"))
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
	unique_x := &ServicesEc2ModelHistoryRecord{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEventInformation(com.amazonaws.services.ec2.model.EventInformation)
func (jbobject *ServicesEc2ModelHistoryRecord) SetEventInformation(a ServicesEc2ModelEventInformationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEventInformation", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/EventInformation"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.EventInformation getEventInformation()
func (jbobject *ServicesEc2ModelHistoryRecord) GetEventInformation() *ServicesEc2ModelEventInformation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEventInformation", "com/amazonaws/services/ec2/model/EventInformation")
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
	unique_x := &ServicesEc2ModelEventInformation{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.HistoryRecord withEventInformation(com.amazonaws.services.ec2.model.EventInformation)
func (jbobject *ServicesEc2ModelHistoryRecord) WithEventInformation(a ServicesEc2ModelEventInformationInterface) *ServicesEc2ModelHistoryRecord {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEventInformation", "com/amazonaws/services/ec2/model/HistoryRecord", conv_a.Value().Cast("com/amazonaws/services/ec2/model/EventInformation"))
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
	unique_x := &ServicesEc2ModelHistoryRecord{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelHistoryRecord) ToString() string {
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
func (jbobject *ServicesEc2ModelHistoryRecord) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelHistoryRecord) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.HistoryRecord clone()
func (jbobject *ServicesEc2ModelHistoryRecord) Clone() *ServicesEc2ModelHistoryRecord {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/HistoryRecord")
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
	unique_x := &ServicesEc2ModelHistoryRecord{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelHistoryRecord) Clone2() (*JavaLangObject, error) {
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


