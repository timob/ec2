package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelScheduledInstancesPrivateIpAddressConfigInterface interface {
	JavaLangObjectInterface

	// public void setPrivateIpAddress(java.lang.String)
	SetPrivateIpAddress(a string) 

	// public java.lang.String getPrivateIpAddress()
	GetPrivateIpAddress() string

	// public com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig withPrivateIpAddress(java.lang.String)
	WithPrivateIpAddress(a string) *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig

	// public void setPrimary(java.lang.Boolean)
	SetPrimary(a bool) 

	// public java.lang.Boolean getPrimary()
	GetPrimary() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig withPrimary(java.lang.Boolean)
	WithPrimary(a bool) *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig

	// public java.lang.Boolean isPrimary()
	IsPrimary() bool

	// public com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig clone()
	Clone() *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig
}

type ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig()
func NewServicesEc2ModelScheduledInstancesPrivateIpAddressConfig() (*ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/ScheduledInstancesPrivateIpAddressConfig")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) SetPrivateIpAddress(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrivateIpAddress", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrivateIpAddress()
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) GetPrivateIpAddress() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrivateIpAddress", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig withPrivateIpAddress(java.lang.String)
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) WithPrivateIpAddress(a string) *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrivateIpAddress", "com/amazonaws/services/ec2/model/ScheduledInstancesPrivateIpAddressConfig", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public void setPrimary(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) SetPrimary(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrimary", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getPrimary()
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) GetPrimary() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrimary", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig withPrimary(java.lang.Boolean)
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) WithPrimary(a bool) *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrimary", "com/amazonaws/services/ec2/model/ScheduledInstancesPrivateIpAddressConfig", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isPrimary()
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) IsPrimary() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isPrimary", "java/lang/Boolean")
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
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) ToString() string {
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
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.ScheduledInstancesPrivateIpAddressConfig clone()
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) Clone() *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/ScheduledInstancesPrivateIpAddressConfig")
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
	unique_x := &ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelScheduledInstancesPrivateIpAddressConfig) Clone2() (*JavaLangObject, error) {
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


