package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpcInterface interface {
	JavaLangObjectInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.Vpc withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelVpc

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.Vpc withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelVpc

	// public void setState(com.amazonaws.services.ec2.model.VpcState)
	SetState(a ServicesEc2ModelVpcStateInterface) 

	// public com.amazonaws.services.ec2.model.Vpc withState(com.amazonaws.services.ec2.model.VpcState)
	WithState(a ServicesEc2ModelVpcStateInterface) *ServicesEc2ModelVpc

	// public void setCidrBlock(java.lang.String)
	SetCidrBlock(a string) 

	// public java.lang.String getCidrBlock()
	GetCidrBlock() string

	// public com.amazonaws.services.ec2.model.Vpc withCidrBlock(java.lang.String)
	WithCidrBlock(a string) *ServicesEc2ModelVpc

	// public void setDhcpOptionsId(java.lang.String)
	SetDhcpOptionsId(a string) 

	// public java.lang.String getDhcpOptionsId()
	GetDhcpOptionsId() string

	// public com.amazonaws.services.ec2.model.Vpc withDhcpOptionsId(java.lang.String)
	WithDhcpOptionsId(a string) *ServicesEc2ModelVpc

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.Vpc withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelVpc

	// public com.amazonaws.services.ec2.model.Vpc withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelVpc

	// public void setInstanceTenancy(java.lang.String)
	SetInstanceTenancy2(a string) 

	// public java.lang.String getInstanceTenancy()
	GetInstanceTenancy() string

	// public com.amazonaws.services.ec2.model.Vpc withInstanceTenancy(java.lang.String)
	WithInstanceTenancy2(a string) *ServicesEc2ModelVpc

	// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	SetInstanceTenancy(a ServicesEc2ModelTenancyInterface) 

	// public com.amazonaws.services.ec2.model.Vpc withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
	WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelVpc

	// public void setIsDefault(java.lang.Boolean)
	SetIsDefault(a bool) 

	// public java.lang.Boolean getIsDefault()
	GetIsDefault() bool

	// public com.amazonaws.services.ec2.model.Vpc withIsDefault(java.lang.Boolean)
	WithIsDefault(a bool) *ServicesEc2ModelVpc

	// public java.lang.Boolean isDefault()
	IsDefault() bool

	// public com.amazonaws.services.ec2.model.Vpc clone()
	Clone() *ServicesEc2ModelVpc
}

type ServicesEc2ModelVpc struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.Vpc()
func NewServicesEc2ModelVpc() (*ServicesEc2ModelVpc) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Vpc")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVpc{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) SetVpcId(a string)  {
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
func (jbobject *ServicesEc2ModelVpc) GetVpcId() string {
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

// public com.amazonaws.services.ec2.model.Vpc withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) WithVpcId(a string) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelVpc) GetState() string {
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

// public com.amazonaws.services.ec2.model.Vpc withState(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) WithState2(a string) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.VpcState)
func (jbobject *ServicesEc2ModelVpc) SetState(a ServicesEc2ModelVpcStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Vpc withState(com.amazonaws.services.ec2.model.VpcState)
func (jbobject *ServicesEc2ModelVpc) WithState(a ServicesEc2ModelVpcStateInterface) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpcState"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) SetCidrBlock(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCidrBlock", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCidrBlock()
func (jbobject *ServicesEc2ModelVpc) GetCidrBlock() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCidrBlock", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Vpc withCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) WithCidrBlock(a string) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCidrBlock", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDhcpOptionsId(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) SetDhcpOptionsId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDhcpOptionsId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDhcpOptionsId()
func (jbobject *ServicesEc2ModelVpc) GetDhcpOptionsId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDhcpOptionsId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Vpc withDhcpOptionsId(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) WithDhcpOptionsId(a string) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDhcpOptionsId", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelVpc) GetTags() []*ServicesEc2ModelTag {
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
func (jbobject *ServicesEc2ModelVpc) SetTags(a []*ServicesEc2ModelTag)  {
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

// public com.amazonaws.services.ec2.model.Vpc withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelVpc) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.Vpc withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelVpc) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) SetInstanceTenancy2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceTenancy", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceTenancy()
func (jbobject *ServicesEc2ModelVpc) GetInstanceTenancy() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceTenancy", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Vpc withInstanceTenancy(java.lang.String)
func (jbobject *ServicesEc2ModelVpc) WithInstanceTenancy2(a string) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelVpc) SetInstanceTenancy(a ServicesEc2ModelTenancyInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceTenancy", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Vpc withInstanceTenancy(com.amazonaws.services.ec2.model.Tenancy)
func (jbobject *ServicesEc2ModelVpc) WithInstanceTenancy(a ServicesEc2ModelTenancyInterface) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceTenancy", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tenancy"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public void setIsDefault(java.lang.Boolean)
func (jbobject *ServicesEc2ModelVpc) SetIsDefault(a bool)  {
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
func (jbobject *ServicesEc2ModelVpc) GetIsDefault() bool {
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

// public com.amazonaws.services.ec2.model.Vpc withIsDefault(java.lang.Boolean)
func (jbobject *ServicesEc2ModelVpc) WithIsDefault(a bool) *ServicesEc2ModelVpc {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withIsDefault", "com/amazonaws/services/ec2/model/Vpc", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isDefault()
func (jbobject *ServicesEc2ModelVpc) IsDefault() bool {
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

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpc) ToString() string {
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
func (jbobject *ServicesEc2ModelVpc) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVpc) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.Vpc clone()
func (jbobject *ServicesEc2ModelVpc) Clone() *ServicesEc2ModelVpc {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/Vpc")
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
	unique_x := &ServicesEc2ModelVpc{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVpc) Clone2() (*JavaLangObject, error) {
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


