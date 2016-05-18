package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSecurityGroupInterface interface {
	JavaLangObjectInterface

	// public void setOwnerId(java.lang.String)
	SetOwnerId(a string) 

	// public java.lang.String getOwnerId()
	GetOwnerId() string

	// public com.amazonaws.services.ec2.model.SecurityGroup withOwnerId(java.lang.String)
	WithOwnerId(a string) *ServicesEc2ModelSecurityGroup

	// public void setGroupName(java.lang.String)
	SetGroupName(a string) 

	// public java.lang.String getGroupName()
	GetGroupName() string

	// public com.amazonaws.services.ec2.model.SecurityGroup withGroupName(java.lang.String)
	WithGroupName(a string) *ServicesEc2ModelSecurityGroup

	// public void setGroupId(java.lang.String)
	SetGroupId(a string) 

	// public java.lang.String getGroupId()
	GetGroupId() string

	// public com.amazonaws.services.ec2.model.SecurityGroup withGroupId(java.lang.String)
	WithGroupId(a string) *ServicesEc2ModelSecurityGroup

	// public void setDescription(java.lang.String)
	SetDescription(a string) 

	// public java.lang.String getDescription()
	GetDescription() string

	// public com.amazonaws.services.ec2.model.SecurityGroup withDescription(java.lang.String)
	WithDescription(a string) *ServicesEc2ModelSecurityGroup

	// public java.util.List<com.amazonaws.services.ec2.model.IpPermission> getIpPermissions()
	GetIpPermissions() []*ServicesEc2ModelIpPermission

	// public void setIpPermissions(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
	SetIpPermissions(a []*ServicesEc2ModelIpPermission) 

	// public com.amazonaws.services.ec2.model.SecurityGroup withIpPermissions(com.amazonaws.services.ec2.model.IpPermission...)
	WithIpPermissions(a ...*ServicesEc2ModelIpPermission) *ServicesEc2ModelSecurityGroup

	// public com.amazonaws.services.ec2.model.SecurityGroup withIpPermissions(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
	WithIpPermissions2(a []*ServicesEc2ModelIpPermission) *ServicesEc2ModelSecurityGroup

	// public java.util.List<com.amazonaws.services.ec2.model.IpPermission> getIpPermissionsEgress()
	GetIpPermissionsEgress() []*ServicesEc2ModelIpPermission

	// public void setIpPermissionsEgress(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
	SetIpPermissionsEgress(a []*ServicesEc2ModelIpPermission) 

	// public com.amazonaws.services.ec2.model.SecurityGroup withIpPermissionsEgress(com.amazonaws.services.ec2.model.IpPermission...)
	WithIpPermissionsEgress(a ...*ServicesEc2ModelIpPermission) *ServicesEc2ModelSecurityGroup

	// public com.amazonaws.services.ec2.model.SecurityGroup withIpPermissionsEgress(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
	WithIpPermissionsEgress2(a []*ServicesEc2ModelIpPermission) *ServicesEc2ModelSecurityGroup

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.SecurityGroup withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelSecurityGroup

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.SecurityGroup withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelSecurityGroup

	// public com.amazonaws.services.ec2.model.SecurityGroup withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelSecurityGroup

	// public com.amazonaws.services.ec2.model.SecurityGroup clone()
	Clone() *ServicesEc2ModelSecurityGroup
}

type ServicesEc2ModelSecurityGroup struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SecurityGroup()
func NewServicesEc2ModelSecurityGroup() (*ServicesEc2ModelSecurityGroup) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SecurityGroup")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSecurityGroup{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) SetOwnerId(a string)  {
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
func (jbobject *ServicesEc2ModelSecurityGroup) GetOwnerId() string {
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

// public com.amazonaws.services.ec2.model.SecurityGroup withOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) WithOwnerId(a string) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOwnerId", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) SetGroupName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGroupName()
func (jbobject *ServicesEc2ModelSecurityGroup) GetGroupName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SecurityGroup withGroupName(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) WithGroupName(a string) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupName", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGroupId(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) SetGroupId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGroupId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGroupId()
func (jbobject *ServicesEc2ModelSecurityGroup) GetGroupId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGroupId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.SecurityGroup withGroupId(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) WithGroupId(a string) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGroupId", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) SetDescription(a string)  {
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
func (jbobject *ServicesEc2ModelSecurityGroup) GetDescription() string {
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

// public com.amazonaws.services.ec2.model.SecurityGroup withDescription(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) WithDescription(a string) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDescription", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.IpPermission> getIpPermissions()
func (jbobject *ServicesEc2ModelSecurityGroup) GetIpPermissions() []*ServicesEc2ModelIpPermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIpPermissions", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelIpPermission)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setIpPermissions(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
func (jbobject *ServicesEc2ModelSecurityGroup) SetIpPermissions(a []*ServicesEc2ModelIpPermission)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIpPermissions", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SecurityGroup withIpPermissions(com.amazonaws.services.ec2.model.IpPermission...)
func (jbobject *ServicesEc2ModelSecurityGroup) WithIpPermissions(a ...*ServicesEc2ModelIpPermission) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/IpPermission")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpPermissions", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("com/amazonaws/services/ec2/model/IpPermission"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SecurityGroup withIpPermissions(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
func (jbobject *ServicesEc2ModelSecurityGroup) WithIpPermissions2(a []*ServicesEc2ModelIpPermission) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpPermissions", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.IpPermission> getIpPermissionsEgress()
func (jbobject *ServicesEc2ModelSecurityGroup) GetIpPermissionsEgress() []*ServicesEc2ModelIpPermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIpPermissionsEgress", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelIpPermission)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setIpPermissionsEgress(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
func (jbobject *ServicesEc2ModelSecurityGroup) SetIpPermissionsEgress(a []*ServicesEc2ModelIpPermission)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIpPermissionsEgress", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.SecurityGroup withIpPermissionsEgress(com.amazonaws.services.ec2.model.IpPermission...)
func (jbobject *ServicesEc2ModelSecurityGroup) WithIpPermissionsEgress(a ...*ServicesEc2ModelIpPermission) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/IpPermission")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpPermissionsEgress", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("com/amazonaws/services/ec2/model/IpPermission"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SecurityGroup withIpPermissionsEgress(java.util.Collection<com.amazonaws.services.ec2.model.IpPermission>)
func (jbobject *ServicesEc2ModelSecurityGroup) WithIpPermissionsEgress2(a []*ServicesEc2ModelIpPermission) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIpPermissionsEgress", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) SetVpcId(a string)  {
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
func (jbobject *ServicesEc2ModelSecurityGroup) GetVpcId() string {
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

// public com.amazonaws.services.ec2.model.SecurityGroup withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelSecurityGroup) WithVpcId(a string) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelSecurityGroup) GetTags() []*ServicesEc2ModelTag {
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
func (jbobject *ServicesEc2ModelSecurityGroup) SetTags(a []*ServicesEc2ModelTag)  {
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

// public com.amazonaws.services.ec2.model.SecurityGroup withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelSecurityGroup) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.SecurityGroup withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelSecurityGroup) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelSecurityGroup {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/SecurityGroup", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSecurityGroup) ToString() string {
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
func (jbobject *ServicesEc2ModelSecurityGroup) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSecurityGroup) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SecurityGroup clone()
func (jbobject *ServicesEc2ModelSecurityGroup) Clone() *ServicesEc2ModelSecurityGroup {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SecurityGroup")
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
	unique_x := &ServicesEc2ModelSecurityGroup{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSecurityGroup) Clone2() (*JavaLangObject, error) {
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


