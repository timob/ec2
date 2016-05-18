package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelIpPermissionInterface interface {
	JavaLangObjectInterface

	// public void setIpProtocol(java.lang.String)
	SetIpProtocol(a string) 

	// public java.lang.String getIpProtocol()
	GetIpProtocol() string

	// public com.amazonaws.services.ec2.model.IpPermission withIpProtocol(java.lang.String)
	WithIpProtocol(a string) *ServicesEc2ModelIpPermission

	// public void setFromPort(java.lang.Integer)
	SetFromPort(a int) 

	// public java.lang.Integer getFromPort()
	GetFromPort() int

	// public com.amazonaws.services.ec2.model.IpPermission withFromPort(java.lang.Integer)
	WithFromPort(a int) *ServicesEc2ModelIpPermission

	// public void setToPort(java.lang.Integer)
	SetToPort(a int) 

	// public java.lang.Integer getToPort()
	GetToPort() int

	// public com.amazonaws.services.ec2.model.IpPermission withToPort(java.lang.Integer)
	WithToPort(a int) *ServicesEc2ModelIpPermission

	// public java.util.List<com.amazonaws.services.ec2.model.UserIdGroupPair> getUserIdGroupPairs()
	GetUserIdGroupPairs() []*ServicesEc2ModelUserIdGroupPair

	// public void setUserIdGroupPairs(java.util.Collection<com.amazonaws.services.ec2.model.UserIdGroupPair>)
	SetUserIdGroupPairs(a []*ServicesEc2ModelUserIdGroupPair) 

	// public com.amazonaws.services.ec2.model.IpPermission withUserIdGroupPairs(com.amazonaws.services.ec2.model.UserIdGroupPair...)
	WithUserIdGroupPairs(a ...*ServicesEc2ModelUserIdGroupPair) *ServicesEc2ModelIpPermission

	// public com.amazonaws.services.ec2.model.IpPermission withUserIdGroupPairs(java.util.Collection<com.amazonaws.services.ec2.model.UserIdGroupPair>)
	WithUserIdGroupPairs2(a []*ServicesEc2ModelUserIdGroupPair) *ServicesEc2ModelIpPermission

	// public java.util.List<java.lang.String> getIpRanges()
	GetIpRanges() []string

	// public void setIpRanges(java.util.Collection<java.lang.String>)
	SetIpRanges(a []string) 

	// public com.amazonaws.services.ec2.model.IpPermission withIpRanges(java.lang.String...)
	WithIpRanges(a ...string) *ServicesEc2ModelIpPermission

	// public com.amazonaws.services.ec2.model.IpPermission withIpRanges(java.util.Collection<java.lang.String>)
	WithIpRanges2(a []string) *ServicesEc2ModelIpPermission

	// public java.util.List<com.amazonaws.services.ec2.model.PrefixListId> getPrefixListIds()
	GetPrefixListIds() []*ServicesEc2ModelPrefixListId

	// public void setPrefixListIds(java.util.Collection<com.amazonaws.services.ec2.model.PrefixListId>)
	SetPrefixListIds(a []*ServicesEc2ModelPrefixListId) 

	// public com.amazonaws.services.ec2.model.IpPermission withPrefixListIds(com.amazonaws.services.ec2.model.PrefixListId...)
	WithPrefixListIds(a ...*ServicesEc2ModelPrefixListId) *ServicesEc2ModelIpPermission

	// public com.amazonaws.services.ec2.model.IpPermission withPrefixListIds(java.util.Collection<com.amazonaws.services.ec2.model.PrefixListId>)
	WithPrefixListIds2(a []*ServicesEc2ModelPrefixListId) *ServicesEc2ModelIpPermission

	// public com.amazonaws.services.ec2.model.IpPermission clone()
	Clone() *ServicesEc2ModelIpPermission
}

type ServicesEc2ModelIpPermission struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.IpPermission()
func NewServicesEc2ModelIpPermission() (*ServicesEc2ModelIpPermission) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/IpPermission")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelIpPermission{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setIpProtocol(java.lang.String)
func (jbobject *ServicesEc2ModelIpPermission) SetIpProtocol(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIpProtocol", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getIpProtocol()
func (jbobject *ServicesEc2ModelIpPermission) GetIpProtocol() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIpProtocol", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.IpPermission withIpProtocol(java.lang.String)
func (jbobject *ServicesEc2ModelIpPermission) WithIpProtocol(a string) *ServicesEc2ModelIpPermission {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpProtocol", "com/amazonaws/services/ec2/model/IpPermission", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public void setFromPort(java.lang.Integer)
func (jbobject *ServicesEc2ModelIpPermission) SetFromPort(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFromPort", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getFromPort()
func (jbobject *ServicesEc2ModelIpPermission) GetFromPort() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFromPort", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.IpPermission withFromPort(java.lang.Integer)
func (jbobject *ServicesEc2ModelIpPermission) WithFromPort(a int) *ServicesEc2ModelIpPermission {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withFromPort", "com/amazonaws/services/ec2/model/IpPermission", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public void setToPort(java.lang.Integer)
func (jbobject *ServicesEc2ModelIpPermission) SetToPort(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setToPort", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getToPort()
func (jbobject *ServicesEc2ModelIpPermission) GetToPort() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getToPort", "java/lang/Integer")
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

// public com.amazonaws.services.ec2.model.IpPermission withToPort(java.lang.Integer)
func (jbobject *ServicesEc2ModelIpPermission) WithToPort(a int) *ServicesEc2ModelIpPermission {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withToPort", "com/amazonaws/services/ec2/model/IpPermission", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.UserIdGroupPair> getUserIdGroupPairs()
func (jbobject *ServicesEc2ModelIpPermission) GetUserIdGroupPairs() []*ServicesEc2ModelUserIdGroupPair {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUserIdGroupPairs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelUserIdGroupPair)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setUserIdGroupPairs(java.util.Collection<com.amazonaws.services.ec2.model.UserIdGroupPair>)
func (jbobject *ServicesEc2ModelIpPermission) SetUserIdGroupPairs(a []*ServicesEc2ModelUserIdGroupPair)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUserIdGroupPairs", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.IpPermission withUserIdGroupPairs(com.amazonaws.services.ec2.model.UserIdGroupPair...)
func (jbobject *ServicesEc2ModelIpPermission) WithUserIdGroupPairs(a ...*ServicesEc2ModelUserIdGroupPair) *ServicesEc2ModelIpPermission {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/UserIdGroupPair")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserIdGroupPairs", "com/amazonaws/services/ec2/model/IpPermission", conv_a.Value().Cast("com/amazonaws/services/ec2/model/UserIdGroupPair"))
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.IpPermission withUserIdGroupPairs(java.util.Collection<com.amazonaws.services.ec2.model.UserIdGroupPair>)
func (jbobject *ServicesEc2ModelIpPermission) WithUserIdGroupPairs2(a []*ServicesEc2ModelUserIdGroupPair) *ServicesEc2ModelIpPermission {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withUserIdGroupPairs", "com/amazonaws/services/ec2/model/IpPermission", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getIpRanges()
func (jbobject *ServicesEc2ModelIpPermission) GetIpRanges() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIpRanges", "java/util/List")
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

// public void setIpRanges(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelIpPermission) SetIpRanges(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIpRanges", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.IpPermission withIpRanges(java.lang.String...)
func (jbobject *ServicesEc2ModelIpPermission) WithIpRanges(a ...string) *ServicesEc2ModelIpPermission {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpRanges", "com/amazonaws/services/ec2/model/IpPermission", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.IpPermission withIpRanges(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelIpPermission) WithIpRanges2(a []string) *ServicesEc2ModelIpPermission {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpRanges", "com/amazonaws/services/ec2/model/IpPermission", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.PrefixListId> getPrefixListIds()
func (jbobject *ServicesEc2ModelIpPermission) GetPrefixListIds() []*ServicesEc2ModelPrefixListId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrefixListIds", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelPrefixListId)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPrefixListIds(java.util.Collection<com.amazonaws.services.ec2.model.PrefixListId>)
func (jbobject *ServicesEc2ModelIpPermission) SetPrefixListIds(a []*ServicesEc2ModelPrefixListId)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrefixListIds", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.IpPermission withPrefixListIds(com.amazonaws.services.ec2.model.PrefixListId...)
func (jbobject *ServicesEc2ModelIpPermission) WithPrefixListIds(a ...*ServicesEc2ModelPrefixListId) *ServicesEc2ModelIpPermission {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/PrefixListId")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefixListIds", "com/amazonaws/services/ec2/model/IpPermission", conv_a.Value().Cast("com/amazonaws/services/ec2/model/PrefixListId"))
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.IpPermission withPrefixListIds(java.util.Collection<com.amazonaws.services.ec2.model.PrefixListId>)
func (jbobject *ServicesEc2ModelIpPermission) WithPrefixListIds2(a []*ServicesEc2ModelPrefixListId) *ServicesEc2ModelIpPermission {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefixListIds", "com/amazonaws/services/ec2/model/IpPermission", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelIpPermission) ToString() string {
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
func (jbobject *ServicesEc2ModelIpPermission) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelIpPermission) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.IpPermission clone()
func (jbobject *ServicesEc2ModelIpPermission) Clone() *ServicesEc2ModelIpPermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/IpPermission")
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
	unique_x := &ServicesEc2ModelIpPermission{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelIpPermission) Clone2() (*JavaLangObject, error) {
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


