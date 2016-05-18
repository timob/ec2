package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelReservationInterface interface {
	JavaLangObjectInterface

	// public void setReservationId(java.lang.String)
	SetReservationId(a string) 

	// public java.lang.String getReservationId()
	GetReservationId() string

	// public com.amazonaws.services.ec2.model.Reservation withReservationId(java.lang.String)
	WithReservationId(a string) *ServicesEc2ModelReservation

	// public void setOwnerId(java.lang.String)
	SetOwnerId(a string) 

	// public java.lang.String getOwnerId()
	GetOwnerId() string

	// public com.amazonaws.services.ec2.model.Reservation withOwnerId(java.lang.String)
	WithOwnerId(a string) *ServicesEc2ModelReservation

	// public void setRequesterId(java.lang.String)
	SetRequesterId(a string) 

	// public java.lang.String getRequesterId()
	GetRequesterId() string

	// public com.amazonaws.services.ec2.model.Reservation withRequesterId(java.lang.String)
	WithRequesterId(a string) *ServicesEc2ModelReservation

	// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getGroups()
	GetGroups() []*ServicesEc2ModelGroupIdentifier

	// public void setGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	SetGroups(a []*ServicesEc2ModelGroupIdentifier) 

	// public com.amazonaws.services.ec2.model.Reservation withGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
	WithGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelReservation

	// public com.amazonaws.services.ec2.model.Reservation withGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
	WithGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelReservation

	// public java.util.List<com.amazonaws.services.ec2.model.Instance> getInstances()
	GetInstances() []*ServicesEc2ModelInstance

	// public void setInstances(java.util.Collection<com.amazonaws.services.ec2.model.Instance>)
	SetInstances(a []*ServicesEc2ModelInstance) 

	// public com.amazonaws.services.ec2.model.Reservation withInstances(com.amazonaws.services.ec2.model.Instance...)
	WithInstances(a ...*ServicesEc2ModelInstance) *ServicesEc2ModelReservation

	// public com.amazonaws.services.ec2.model.Reservation withInstances(java.util.Collection<com.amazonaws.services.ec2.model.Instance>)
	WithInstances2(a []*ServicesEc2ModelInstance) *ServicesEc2ModelReservation

	// public java.util.List<java.lang.String> getGroupNames()
	GetGroupNames() []string

	// public void setGroupNames(java.util.Collection<java.lang.String>)
	SetGroupNames(a []string) 

	// public com.amazonaws.services.ec2.model.Reservation withGroupNames(java.lang.String...)
	WithGroupNames(a ...string) *ServicesEc2ModelReservation

	// public com.amazonaws.services.ec2.model.Reservation withGroupNames(java.util.Collection<java.lang.String>)
	WithGroupNames2(a []string) *ServicesEc2ModelReservation

	// public com.amazonaws.services.ec2.model.Reservation clone()
	Clone() *ServicesEc2ModelReservation
}

type ServicesEc2ModelReservation struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.Reservation()
func NewServicesEc2ModelReservation() (*ServicesEc2ModelReservation) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Reservation")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelReservation{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setReservationId(java.lang.String)
func (jbobject *ServicesEc2ModelReservation) SetReservationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setReservationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getReservationId()
func (jbobject *ServicesEc2ModelReservation) GetReservationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getReservationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Reservation withReservationId(java.lang.String)
func (jbobject *ServicesEc2ModelReservation) WithReservationId(a string) *ServicesEc2ModelReservation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withReservationId", "com/amazonaws/services/ec2/model/Reservation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelReservation) SetOwnerId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOwnerId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOwnerId()
func (jbobject *ServicesEc2ModelReservation) GetOwnerId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOwnerId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Reservation withOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelReservation) WithOwnerId(a string) *ServicesEc2ModelReservation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOwnerId", "com/amazonaws/services/ec2/model/Reservation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public void setRequesterId(java.lang.String)
func (jbobject *ServicesEc2ModelReservation) SetRequesterId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRequesterId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getRequesterId()
func (jbobject *ServicesEc2ModelReservation) GetRequesterId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRequesterId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Reservation withRequesterId(java.lang.String)
func (jbobject *ServicesEc2ModelReservation) WithRequesterId(a string) *ServicesEc2ModelReservation {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRequesterId", "com/amazonaws/services/ec2/model/Reservation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.GroupIdentifier> getGroups()
func (jbobject *ServicesEc2ModelReservation) GetGroups() []*ServicesEc2ModelGroupIdentifier {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroups", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelGroupIdentifier)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelReservation) SetGroups(a []*ServicesEc2ModelGroupIdentifier)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroups", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Reservation withGroups(com.amazonaws.services.ec2.model.GroupIdentifier...)
func (jbobject *ServicesEc2ModelReservation) WithGroups(a ...*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelReservation {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/GroupIdentifier")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/Reservation", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GroupIdentifier"))
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Reservation withGroups(java.util.Collection<com.amazonaws.services.ec2.model.GroupIdentifier>)
func (jbobject *ServicesEc2ModelReservation) WithGroups2(a []*ServicesEc2ModelGroupIdentifier) *ServicesEc2ModelReservation {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroups", "com/amazonaws/services/ec2/model/Reservation", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Instance> getInstances()
func (jbobject *ServicesEc2ModelReservation) GetInstances() []*ServicesEc2ModelInstance {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstances", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstance)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setInstances(java.util.Collection<com.amazonaws.services.ec2.model.Instance>)
func (jbobject *ServicesEc2ModelReservation) SetInstances(a []*ServicesEc2ModelInstance)  {
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

// public com.amazonaws.services.ec2.model.Reservation withInstances(com.amazonaws.services.ec2.model.Instance...)
func (jbobject *ServicesEc2ModelReservation) WithInstances(a ...*ServicesEc2ModelInstance) *ServicesEc2ModelReservation {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Instance")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstances", "com/amazonaws/services/ec2/model/Reservation", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Instance"))
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Reservation withInstances(java.util.Collection<com.amazonaws.services.ec2.model.Instance>)
func (jbobject *ServicesEc2ModelReservation) WithInstances2(a []*ServicesEc2ModelInstance) *ServicesEc2ModelReservation {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstances", "com/amazonaws/services/ec2/model/Reservation", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getGroupNames()
func (jbobject *ServicesEc2ModelReservation) GetGroupNames() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupNames", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setGroupNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelReservation) SetGroupNames(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupNames", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Reservation withGroupNames(java.lang.String...)
func (jbobject *ServicesEc2ModelReservation) WithGroupNames(a ...string) *ServicesEc2ModelReservation {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupNames", "com/amazonaws/services/ec2/model/Reservation", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Reservation withGroupNames(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelReservation) WithGroupNames2(a []string) *ServicesEc2ModelReservation {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupNames", "com/amazonaws/services/ec2/model/Reservation", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelReservation) ToString() string {
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
func (jbobject *ServicesEc2ModelReservation) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelReservation) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.Reservation clone()
func (jbobject *ServicesEc2ModelReservation) Clone() *ServicesEc2ModelReservation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/Reservation")
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
	unique_x := &ServicesEc2ModelReservation{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelReservation) Clone2() (*JavaLangObject, error) {
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


