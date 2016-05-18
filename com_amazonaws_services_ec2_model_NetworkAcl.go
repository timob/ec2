package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNetworkAclInterface interface {
	JavaLangObjectInterface

	// public void setNetworkAclId(java.lang.String)
	SetNetworkAclId(a string) 

	// public java.lang.String getNetworkAclId()
	GetNetworkAclId() string

	// public com.amazonaws.services.ec2.model.NetworkAcl withNetworkAclId(java.lang.String)
	WithNetworkAclId(a string) *ServicesEc2ModelNetworkAcl

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.NetworkAcl withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelNetworkAcl

	// public void setIsDefault(java.lang.Boolean)
	SetIsDefault(a bool) 

	// public java.lang.Boolean getIsDefault()
	GetIsDefault() bool

	// public com.amazonaws.services.ec2.model.NetworkAcl withIsDefault(java.lang.Boolean)
	WithIsDefault(a bool) *ServicesEc2ModelNetworkAcl

	// public java.lang.Boolean isDefault()
	IsDefault() bool

	// public java.util.List<com.amazonaws.services.ec2.model.NetworkAclEntry> getEntries()
	GetEntries() []*ServicesEc2ModelNetworkAclEntry

	// public void setEntries(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAclEntry>)
	SetEntries(a []*ServicesEc2ModelNetworkAclEntry) 

	// public com.amazonaws.services.ec2.model.NetworkAcl withEntries(com.amazonaws.services.ec2.model.NetworkAclEntry...)
	WithEntries(a ...*ServicesEc2ModelNetworkAclEntry) *ServicesEc2ModelNetworkAcl

	// public com.amazonaws.services.ec2.model.NetworkAcl withEntries(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAclEntry>)
	WithEntries2(a []*ServicesEc2ModelNetworkAclEntry) *ServicesEc2ModelNetworkAcl

	// public java.util.List<com.amazonaws.services.ec2.model.NetworkAclAssociation> getAssociations()
	GetAssociations() []*ServicesEc2ModelNetworkAclAssociation

	// public void setAssociations(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAclAssociation>)
	SetAssociations(a []*ServicesEc2ModelNetworkAclAssociation) 

	// public com.amazonaws.services.ec2.model.NetworkAcl withAssociations(com.amazonaws.services.ec2.model.NetworkAclAssociation...)
	WithAssociations(a ...*ServicesEc2ModelNetworkAclAssociation) *ServicesEc2ModelNetworkAcl

	// public com.amazonaws.services.ec2.model.NetworkAcl withAssociations(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAclAssociation>)
	WithAssociations2(a []*ServicesEc2ModelNetworkAclAssociation) *ServicesEc2ModelNetworkAcl

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.NetworkAcl withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelNetworkAcl

	// public com.amazonaws.services.ec2.model.NetworkAcl withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelNetworkAcl

	// public com.amazonaws.services.ec2.model.NetworkAcl clone()
	Clone() *ServicesEc2ModelNetworkAcl
}

type ServicesEc2ModelNetworkAcl struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.NetworkAcl()
func NewServicesEc2ModelNetworkAcl() (*ServicesEc2ModelNetworkAcl) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/NetworkAcl")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelNetworkAcl{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setNetworkAclId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAcl) SetNetworkAclId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkAclId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNetworkAclId()
func (jbobject *ServicesEc2ModelNetworkAcl) GetNetworkAclId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkAclId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkAcl withNetworkAclId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAcl) WithNetworkAclId(a string) *ServicesEc2ModelNetworkAcl {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkAclId", "com/amazonaws/services/ec2/model/NetworkAcl", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAcl) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelNetworkAcl) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.NetworkAcl withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkAcl) WithVpcId(a string) *ServicesEc2ModelNetworkAcl {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/NetworkAcl", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIsDefault(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkAcl) SetIsDefault(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIsDefault", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getIsDefault()
func (jbobject *ServicesEc2ModelNetworkAcl) GetIsDefault() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIsDefault", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.NetworkAcl withIsDefault(java.lang.Boolean)
func (jbobject *ServicesEc2ModelNetworkAcl) WithIsDefault(a bool) *ServicesEc2ModelNetworkAcl {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIsDefault", "com/amazonaws/services/ec2/model/NetworkAcl", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDefault()
func (jbobject *ServicesEc2ModelNetworkAcl) IsDefault() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDefault", "java/lang/Boolean")
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

// public java.util.List<com.amazonaws.services.ec2.model.NetworkAclEntry> getEntries()
func (jbobject *ServicesEc2ModelNetworkAcl) GetEntries() []*ServicesEc2ModelNetworkAclEntry {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEntries", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelNetworkAclEntry)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setEntries(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAclEntry>)
func (jbobject *ServicesEc2ModelNetworkAcl) SetEntries(a []*ServicesEc2ModelNetworkAclEntry)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEntries", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkAcl withEntries(com.amazonaws.services.ec2.model.NetworkAclEntry...)
func (jbobject *ServicesEc2ModelNetworkAcl) WithEntries(a ...*ServicesEc2ModelNetworkAclEntry) *ServicesEc2ModelNetworkAcl {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/NetworkAclEntry")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEntries", "com/amazonaws/services/ec2/model/NetworkAcl", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkAclEntry"))
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.NetworkAcl withEntries(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAclEntry>)
func (jbobject *ServicesEc2ModelNetworkAcl) WithEntries2(a []*ServicesEc2ModelNetworkAclEntry) *ServicesEc2ModelNetworkAcl {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEntries", "com/amazonaws/services/ec2/model/NetworkAcl", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.NetworkAclAssociation> getAssociations()
func (jbobject *ServicesEc2ModelNetworkAcl) GetAssociations() []*ServicesEc2ModelNetworkAclAssociation {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAssociations", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelNetworkAclAssociation)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAssociations(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAclAssociation>)
func (jbobject *ServicesEc2ModelNetworkAcl) SetAssociations(a []*ServicesEc2ModelNetworkAclAssociation)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAssociations", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkAcl withAssociations(com.amazonaws.services.ec2.model.NetworkAclAssociation...)
func (jbobject *ServicesEc2ModelNetworkAcl) WithAssociations(a ...*ServicesEc2ModelNetworkAclAssociation) *ServicesEc2ModelNetworkAcl {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/NetworkAclAssociation")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociations", "com/amazonaws/services/ec2/model/NetworkAcl", conv_a.Value().Cast("com/amazonaws/services/ec2/model/NetworkAclAssociation"))
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.NetworkAcl withAssociations(java.util.Collection<com.amazonaws.services.ec2.model.NetworkAclAssociation>)
func (jbobject *ServicesEc2ModelNetworkAcl) WithAssociations2(a []*ServicesEc2ModelNetworkAclAssociation) *ServicesEc2ModelNetworkAcl {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociations", "com/amazonaws/services/ec2/model/NetworkAcl", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelNetworkAcl) GetTags() []*ServicesEc2ModelTag {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTags", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelTag)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelNetworkAcl) SetTags(a []*ServicesEc2ModelTag)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTags", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.NetworkAcl withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelNetworkAcl) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelNetworkAcl {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/NetworkAcl", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.NetworkAcl withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelNetworkAcl) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelNetworkAcl {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/NetworkAcl", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNetworkAcl) ToString() string {
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
func (jbobject *ServicesEc2ModelNetworkAcl) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelNetworkAcl) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.NetworkAcl clone()
func (jbobject *ServicesEc2ModelNetworkAcl) Clone() *ServicesEc2ModelNetworkAcl {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/NetworkAcl")
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
	unique_x := &ServicesEc2ModelNetworkAcl{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelNetworkAcl) Clone2() (*JavaLangObject, error) {
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


