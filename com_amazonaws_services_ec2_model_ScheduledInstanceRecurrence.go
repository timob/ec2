package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelScheduledInstanceRecurrenceInterface interface {
	JavaLangObjectInterface

	// public void setFrequency(java.lang.String)
	SetFrequency(a string) 

	// public java.lang.String getFrequency()
	GetFrequency() string

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withFrequency(java.lang.String)
	WithFrequency(a string) *ServicesEc2ModelScheduledInstanceRecurrence

	// public void setInterval(java.lang.Integer)
	SetInterval(a int) 

	// public java.lang.Integer getInterval()
	GetInterval() int

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withInterval(java.lang.Integer)
	WithInterval(a int) *ServicesEc2ModelScheduledInstanceRecurrence

	// public java.util.List<java.lang.Integer> getOccurrenceDaySet()
	GetOccurrenceDaySet() []int

	// public void setOccurrenceDaySet(java.util.Collection<java.lang.Integer>)
	SetOccurrenceDaySet(a []int) 

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withOccurrenceDaySet(java.lang.Integer...)
	WithOccurrenceDaySet(a ...int) *ServicesEc2ModelScheduledInstanceRecurrence

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withOccurrenceDaySet(java.util.Collection<java.lang.Integer>)
	WithOccurrenceDaySet2(a []int) *ServicesEc2ModelScheduledInstanceRecurrence

	// public void setOccurrenceRelativeToEnd(java.lang.Boolean)
	SetOccurrenceRelativeToEnd(a bool) 

	// public java.lang.Boolean getOccurrenceRelativeToEnd()
	GetOccurrenceRelativeToEnd() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withOccurrenceRelativeToEnd(java.lang.Boolean)
	WithOccurrenceRelativeToEnd(a bool) *ServicesEc2ModelScheduledInstanceRecurrence

	// public java.lang.Boolean isOccurrenceRelativeToEnd()
	IsOccurrenceRelativeToEnd() bool

	// public void setOccurrenceUnit(java.lang.String)
	SetOccurrenceUnit(a string) 

	// public java.lang.String getOccurrenceUnit()
	GetOccurrenceUnit() string

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withOccurrenceUnit(java.lang.String)
	WithOccurrenceUnit(a string) *ServicesEc2ModelScheduledInstanceRecurrence

	// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence clone()
	Clone() *ServicesEc2ModelScheduledInstanceRecurrence
}

type ServicesEc2ModelScheduledInstanceRecurrence struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence()
func NewServicesEc2ModelScheduledInstanceRecurrence() (*ServicesEc2ModelScheduledInstanceRecurrence) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelScheduledInstanceRecurrence{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setFrequency(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) SetFrequency(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFrequency", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getFrequency()
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) GetFrequency() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFrequency", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withFrequency(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) WithFrequency(a string) *ServicesEc2ModelScheduledInstanceRecurrence {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFrequency", "com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceRecurrence{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInterval(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) SetInterval(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInterval", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getInterval()
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) GetInterval() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInterval", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withInterval(java.lang.Integer)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) WithInterval(a int) *ServicesEc2ModelScheduledInstanceRecurrence {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInterval", "com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceRecurrence{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.Integer> getOccurrenceDaySet()
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) GetOccurrenceDaySet() []int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOccurrenceDaySet", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoInteger())
	dst := new([]int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setOccurrenceDaySet(java.util.Collection<java.lang.Integer>)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) SetOccurrenceDaySet(a []int)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaInteger())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOccurrenceDaySet", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withOccurrenceDaySet(java.lang.Integer...)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) WithOccurrenceDaySet(a ...int) *ServicesEc2ModelScheduledInstanceRecurrence {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaInteger(), "java/lang/Integer")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOccurrenceDaySet", "com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceRecurrence{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withOccurrenceDaySet(java.util.Collection<java.lang.Integer>)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) WithOccurrenceDaySet2(a []int) *ServicesEc2ModelScheduledInstanceRecurrence {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaInteger())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOccurrenceDaySet", "com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceRecurrence{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOccurrenceRelativeToEnd(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) SetOccurrenceRelativeToEnd(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOccurrenceRelativeToEnd", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getOccurrenceRelativeToEnd()
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) GetOccurrenceRelativeToEnd() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOccurrenceRelativeToEnd", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withOccurrenceRelativeToEnd(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) WithOccurrenceRelativeToEnd(a bool) *ServicesEc2ModelScheduledInstanceRecurrence {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOccurrenceRelativeToEnd", "com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceRecurrence{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isOccurrenceRelativeToEnd()
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) IsOccurrenceRelativeToEnd() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isOccurrenceRelativeToEnd", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setOccurrenceUnit(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) SetOccurrenceUnit(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOccurrenceUnit", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOccurrenceUnit()
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) GetOccurrenceUnit() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOccurrenceUnit", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence withOccurrenceUnit(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) WithOccurrenceUnit(a string) *ServicesEc2ModelScheduledInstanceRecurrence {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOccurrenceUnit", "com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstanceRecurrence{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) ToString() string {
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
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ScheduledInstanceRecurrence clone()
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) Clone() *ServicesEc2ModelScheduledInstanceRecurrence {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ScheduledInstanceRecurrence")
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

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelScheduledInstanceRecurrence) Clone2() (*JavaLangObject, error) {
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


