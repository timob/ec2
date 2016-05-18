package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDhcpOptionsInterface interface {
	JavaLangObjectInterface

	// public void setDhcpOptionsId(java.lang.String)
	SetDhcpOptionsId(a string) 

	// public java.lang.String getDhcpOptionsId()
	GetDhcpOptionsId() string

	// public com.amazonaws.services.ec2.model.DhcpOptions withDhcpOptionsId(java.lang.String)
	WithDhcpOptionsId(a string) *ServicesEc2ModelDhcpOptions

	// public java.util.List<com.amazonaws.services.ec2.model.DhcpConfiguration> getDhcpConfigurations()
	GetDhcpConfigurations() []*ServicesEc2ModelDhcpConfiguration

	// public void setDhcpConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.DhcpConfiguration>)
	SetDhcpConfigurations(a []*ServicesEc2ModelDhcpConfiguration) 

	// public com.amazonaws.services.ec2.model.DhcpOptions withDhcpConfigurations(com.amazonaws.services.ec2.model.DhcpConfiguration...)
	WithDhcpConfigurations(a ...*ServicesEc2ModelDhcpConfiguration) *ServicesEc2ModelDhcpOptions

	// public com.amazonaws.services.ec2.model.DhcpOptions withDhcpConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.DhcpConfiguration>)
	WithDhcpConfigurations2(a []*ServicesEc2ModelDhcpConfiguration) *ServicesEc2ModelDhcpOptions

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.DhcpOptions withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelDhcpOptions

	// public com.amazonaws.services.ec2.model.DhcpOptions withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelDhcpOptions

	// public com.amazonaws.services.ec2.model.DhcpOptions clone()
	Clone() *ServicesEc2ModelDhcpOptions
}

type ServicesEc2ModelDhcpOptions struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DhcpOptions()
func NewServicesEc2ModelDhcpOptions() (*ServicesEc2ModelDhcpOptions) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DhcpOptions")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDhcpOptions{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDhcpOptionsId(java.lang.String)
func (jbobject *ServicesEc2ModelDhcpOptions) SetDhcpOptionsId(a string)  {
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
func (jbobject *ServicesEc2ModelDhcpOptions) GetDhcpOptionsId() string {
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

// public com.amazonaws.services.ec2.model.DhcpOptions withDhcpOptionsId(java.lang.String)
func (jbobject *ServicesEc2ModelDhcpOptions) WithDhcpOptionsId(a string) *ServicesEc2ModelDhcpOptions {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDhcpOptionsId", "com/amazonaws/services/ec2/model/DhcpOptions", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDhcpOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.DhcpConfiguration> getDhcpConfigurations()
func (jbobject *ServicesEc2ModelDhcpOptions) GetDhcpConfigurations() []*ServicesEc2ModelDhcpConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDhcpConfigurations", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelDhcpConfiguration)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setDhcpConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.DhcpConfiguration>)
func (jbobject *ServicesEc2ModelDhcpOptions) SetDhcpConfigurations(a []*ServicesEc2ModelDhcpConfiguration)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDhcpConfigurations", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DhcpOptions withDhcpConfigurations(com.amazonaws.services.ec2.model.DhcpConfiguration...)
func (jbobject *ServicesEc2ModelDhcpOptions) WithDhcpConfigurations(a ...*ServicesEc2ModelDhcpConfiguration) *ServicesEc2ModelDhcpOptions {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/DhcpConfiguration")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDhcpConfigurations", "com/amazonaws/services/ec2/model/DhcpOptions", conv_a.Value().Cast("com/amazonaws/services/ec2/model/DhcpConfiguration"))
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
	unique_x := &ServicesEc2ModelDhcpOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DhcpOptions withDhcpConfigurations(java.util.Collection<com.amazonaws.services.ec2.model.DhcpConfiguration>)
func (jbobject *ServicesEc2ModelDhcpOptions) WithDhcpConfigurations2(a []*ServicesEc2ModelDhcpConfiguration) *ServicesEc2ModelDhcpOptions {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDhcpConfigurations", "com/amazonaws/services/ec2/model/DhcpOptions", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDhcpOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelDhcpOptions) GetTags() []*ServicesEc2ModelTag {
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
func (jbobject *ServicesEc2ModelDhcpOptions) SetTags(a []*ServicesEc2ModelTag)  {
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

// public com.amazonaws.services.ec2.model.DhcpOptions withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelDhcpOptions) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelDhcpOptions {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/DhcpOptions", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelDhcpOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DhcpOptions withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelDhcpOptions) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelDhcpOptions {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/DhcpOptions", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDhcpOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDhcpOptions) ToString() string {
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
func (jbobject *ServicesEc2ModelDhcpOptions) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDhcpOptions) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DhcpOptions clone()
func (jbobject *ServicesEc2ModelDhcpOptions) Clone() *ServicesEc2ModelDhcpOptions {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DhcpOptions")
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
	unique_x := &ServicesEc2ModelDhcpOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDhcpOptions) Clone2() (*JavaLangObject, error) {
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


