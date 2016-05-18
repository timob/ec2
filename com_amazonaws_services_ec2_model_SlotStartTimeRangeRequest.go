package ec2

import "time"
import "github.com/timob/javabind"

type ServicesEc2ModelSlotStartTimeRangeRequestInterface interface {
	JavaLangObjectInterface

	// public void setEarliestTime(java.util.Date)
	SetEarliestTime(a time.Time) 

	// public java.util.Date getEarliestTime()
	GetEarliestTime() time.Time

	// public com.amazonaws.services.ec2.model.SlotStartTimeRangeRequest withEarliestTime(java.util.Date)
	WithEarliestTime(a time.Time) *ServicesEc2ModelSlotStartTimeRangeRequest

	// public void setLatestTime(java.util.Date)
	SetLatestTime(a time.Time) 

	// public java.util.Date getLatestTime()
	GetLatestTime() time.Time

	// public com.amazonaws.services.ec2.model.SlotStartTimeRangeRequest withLatestTime(java.util.Date)
	WithLatestTime(a time.Time) *ServicesEc2ModelSlotStartTimeRangeRequest

	// public com.amazonaws.services.ec2.model.SlotStartTimeRangeRequest clone()
	Clone() *ServicesEc2ModelSlotStartTimeRangeRequest
}

type ServicesEc2ModelSlotStartTimeRangeRequest struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SlotStartTimeRangeRequest()
func NewServicesEc2ModelSlotStartTimeRangeRequest() (*ServicesEc2ModelSlotStartTimeRangeRequest) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SlotStartTimeRangeRequest")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSlotStartTimeRangeRequest{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setEarliestTime(java.util.Date)
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) SetEarliestTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEarliestTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getEarliestTime()
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) GetEarliestTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEarliestTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.SlotStartTimeRangeRequest withEarliestTime(java.util.Date)
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) WithEarliestTime(a time.Time) *ServicesEc2ModelSlotStartTimeRangeRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEarliestTime", "com/amazonaws/services/ec2/model/SlotStartTimeRangeRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelSlotStartTimeRangeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public void setLatestTime(java.util.Date)
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) SetLatestTime(a time.Time)  {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLatestTime", javabind.Void, conv_a.Value().Cast("java/util/Date"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.Date getLatestTime()
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) GetLatestTime() time.Time {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLatestTime", "java/util/Date")
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

// public com.amazonaws.services.ec2.model.SlotStartTimeRangeRequest withLatestTime(java.util.Date)
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) WithLatestTime(a time.Time) *ServicesEc2ModelSlotStartTimeRangeRequest {
	conv_a := javabind.NewGoToJavaDate()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withLatestTime", "com/amazonaws/services/ec2/model/SlotStartTimeRangeRequest", conv_a.Value().Cast("java/util/Date"))
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
	unique_x := &ServicesEc2ModelSlotStartTimeRangeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) ToString() string {
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
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SlotStartTimeRangeRequest clone()
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) Clone() *ServicesEc2ModelSlotStartTimeRangeRequest {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SlotStartTimeRangeRequest")
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
	unique_x := &ServicesEc2ModelSlotStartTimeRangeRequest{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSlotStartTimeRangeRequest) Clone2() (*JavaLangObject, error) {
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


