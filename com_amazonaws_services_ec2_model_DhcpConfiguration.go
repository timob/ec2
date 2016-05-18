package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDhcpConfigurationInterface interface {
	JavaLangObjectInterface

	// public void setKey(java.lang.String)
	SetKey(a string) 

	// public java.lang.String getKey()
	GetKey() string

	// public com.amazonaws.services.ec2.model.DhcpConfiguration withKey(java.lang.String)
	WithKey(a string) *ServicesEc2ModelDhcpConfiguration

	// public java.util.List<java.lang.String> getValues()
	GetValues() []string

	// public void setValues(java.util.Collection<java.lang.String>)
	SetValues(a []string) 

	// public com.amazonaws.services.ec2.model.DhcpConfiguration withValues(java.lang.String...)
	WithValues(a ...string) *ServicesEc2ModelDhcpConfiguration

	// public com.amazonaws.services.ec2.model.DhcpConfiguration withValues(java.util.Collection<java.lang.String>)
	WithValues2(a []string) *ServicesEc2ModelDhcpConfiguration

	// public com.amazonaws.services.ec2.model.DhcpConfiguration clone()
	Clone() *ServicesEc2ModelDhcpConfiguration
}

type ServicesEc2ModelDhcpConfiguration struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DhcpConfiguration()
func NewServicesEc2ModelDhcpConfiguration() (*ServicesEc2ModelDhcpConfiguration) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DhcpConfiguration")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDhcpConfiguration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setKey(java.lang.String)
func (jbobject *ServicesEc2ModelDhcpConfiguration) SetKey(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setKey", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getKey()
func (jbobject *ServicesEc2ModelDhcpConfiguration) GetKey() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getKey", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DhcpConfiguration withKey(java.lang.String)
func (jbobject *ServicesEc2ModelDhcpConfiguration) WithKey(a string) *ServicesEc2ModelDhcpConfiguration {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withKey", "com/amazonaws/services/ec2/model/DhcpConfiguration", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDhcpConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<java.lang.String> getValues()
func (jbobject *ServicesEc2ModelDhcpConfiguration) GetValues() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValues", "java/util/List")
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

// public void setValues(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDhcpConfiguration) SetValues(a []string)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValues", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DhcpConfiguration withValues(java.lang.String...)
func (jbobject *ServicesEc2ModelDhcpConfiguration) WithValues(a ...string) *ServicesEc2ModelDhcpConfiguration {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValues", "com/amazonaws/services/ec2/model/DhcpConfiguration", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDhcpConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DhcpConfiguration withValues(java.util.Collection<java.lang.String>)
func (jbobject *ServicesEc2ModelDhcpConfiguration) WithValues2(a []string) *ServicesEc2ModelDhcpConfiguration {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaString())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withValues", "com/amazonaws/services/ec2/model/DhcpConfiguration", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDhcpConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDhcpConfiguration) ToString() string {
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
func (jbobject *ServicesEc2ModelDhcpConfiguration) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDhcpConfiguration) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DhcpConfiguration clone()
func (jbobject *ServicesEc2ModelDhcpConfiguration) Clone() *ServicesEc2ModelDhcpConfiguration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DhcpConfiguration")
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
	unique_x := &ServicesEc2ModelDhcpConfiguration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDhcpConfiguration) Clone2() (*JavaLangObject, error) {
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


