package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpnConnectionInterface interface {
	JavaLangObjectInterface

	// public void setVpnConnectionId(java.lang.String)
	SetVpnConnectionId(a string) 

	// public java.lang.String getVpnConnectionId()
	GetVpnConnectionId() string

	// public com.amazonaws.services.ec2.model.VpnConnection withVpnConnectionId(java.lang.String)
	WithVpnConnectionId(a string) *ServicesEc2ModelVpnConnection

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.VpnConnection withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelVpnConnection

	// public void setState(com.amazonaws.services.ec2.model.VpnState)
	SetState(a ServicesEc2ModelVpnStateInterface) 

	// public com.amazonaws.services.ec2.model.VpnConnection withState(com.amazonaws.services.ec2.model.VpnState)
	WithState(a ServicesEc2ModelVpnStateInterface) *ServicesEc2ModelVpnConnection

	// public void setCustomerGatewayConfiguration(java.lang.String)
	SetCustomerGatewayConfiguration(a string) 

	// public java.lang.String getCustomerGatewayConfiguration()
	GetCustomerGatewayConfiguration() string

	// public com.amazonaws.services.ec2.model.VpnConnection withCustomerGatewayConfiguration(java.lang.String)
	WithCustomerGatewayConfiguration(a string) *ServicesEc2ModelVpnConnection

	// public void setType(java.lang.String)
	SetType2(a string) 

	// public java.lang.String getType()
	GetType() string

	// public com.amazonaws.services.ec2.model.VpnConnection withType(java.lang.String)
	WithType2(a string) *ServicesEc2ModelVpnConnection

	// public void setType(com.amazonaws.services.ec2.model.GatewayType)
	SetType(a ServicesEc2ModelGatewayTypeInterface) 

	// public com.amazonaws.services.ec2.model.VpnConnection withType(com.amazonaws.services.ec2.model.GatewayType)
	WithType(a ServicesEc2ModelGatewayTypeInterface) *ServicesEc2ModelVpnConnection

	// public void setCustomerGatewayId(java.lang.String)
	SetCustomerGatewayId(a string) 

	// public java.lang.String getCustomerGatewayId()
	GetCustomerGatewayId() string

	// public com.amazonaws.services.ec2.model.VpnConnection withCustomerGatewayId(java.lang.String)
	WithCustomerGatewayId(a string) *ServicesEc2ModelVpnConnection

	// public void setVpnGatewayId(java.lang.String)
	SetVpnGatewayId(a string) 

	// public java.lang.String getVpnGatewayId()
	GetVpnGatewayId() string

	// public com.amazonaws.services.ec2.model.VpnConnection withVpnGatewayId(java.lang.String)
	WithVpnGatewayId(a string) *ServicesEc2ModelVpnConnection

	// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
	GetTags() []*ServicesEc2ModelTag

	// public void setTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	SetTags(a []*ServicesEc2ModelTag) 

	// public com.amazonaws.services.ec2.model.VpnConnection withTags(com.amazonaws.services.ec2.model.Tag...)
	WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelVpnConnection

	// public com.amazonaws.services.ec2.model.VpnConnection withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
	WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelVpnConnection

	// public java.util.List<com.amazonaws.services.ec2.model.VgwTelemetry> getVgwTelemetry()
	GetVgwTelemetry() []*ServicesEc2ModelVgwTelemetry

	// public void setVgwTelemetry(java.util.Collection<com.amazonaws.services.ec2.model.VgwTelemetry>)
	SetVgwTelemetry(a []*ServicesEc2ModelVgwTelemetry) 

	// public com.amazonaws.services.ec2.model.VpnConnection withVgwTelemetry(com.amazonaws.services.ec2.model.VgwTelemetry...)
	WithVgwTelemetry(a ...*ServicesEc2ModelVgwTelemetry) *ServicesEc2ModelVpnConnection

	// public com.amazonaws.services.ec2.model.VpnConnection withVgwTelemetry(java.util.Collection<com.amazonaws.services.ec2.model.VgwTelemetry>)
	WithVgwTelemetry2(a []*ServicesEc2ModelVgwTelemetry) *ServicesEc2ModelVpnConnection

	// public void setOptions(com.amazonaws.services.ec2.model.VpnConnectionOptions)
	SetOptions(a ServicesEc2ModelVpnConnectionOptionsInterface) 

	// public com.amazonaws.services.ec2.model.VpnConnectionOptions getOptions()
	GetOptions() *ServicesEc2ModelVpnConnectionOptions

	// public com.amazonaws.services.ec2.model.VpnConnection withOptions(com.amazonaws.services.ec2.model.VpnConnectionOptions)
	WithOptions(a ServicesEc2ModelVpnConnectionOptionsInterface) *ServicesEc2ModelVpnConnection

	// public java.util.List<com.amazonaws.services.ec2.model.VpnStaticRoute> getRoutes()
	GetRoutes() []*ServicesEc2ModelVpnStaticRoute

	// public void setRoutes(java.util.Collection<com.amazonaws.services.ec2.model.VpnStaticRoute>)
	SetRoutes(a []*ServicesEc2ModelVpnStaticRoute) 

	// public com.amazonaws.services.ec2.model.VpnConnection withRoutes(com.amazonaws.services.ec2.model.VpnStaticRoute...)
	WithRoutes(a ...*ServicesEc2ModelVpnStaticRoute) *ServicesEc2ModelVpnConnection

	// public com.amazonaws.services.ec2.model.VpnConnection withRoutes(java.util.Collection<com.amazonaws.services.ec2.model.VpnStaticRoute>)
	WithRoutes2(a []*ServicesEc2ModelVpnStaticRoute) *ServicesEc2ModelVpnConnection

	// public com.amazonaws.services.ec2.model.VpnConnection clone()
	Clone() *ServicesEc2ModelVpnConnection
}

type ServicesEc2ModelVpnConnection struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.VpnConnection()
func NewServicesEc2ModelVpnConnection() (*ServicesEc2ModelVpnConnection) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/VpnConnection")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelVpnConnection{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpnConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) SetVpnConnectionId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpnConnectionId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpnConnectionId()
func (jbobject *ServicesEc2ModelVpnConnection) GetVpnConnectionId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpnConnectionId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpnConnection withVpnConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) WithVpnConnectionId(a string) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnConnectionId", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelVpnConnection) GetState() string {
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

// public com.amazonaws.services.ec2.model.VpnConnection withState(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) WithState2(a string) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.VpnState)
func (jbobject *ServicesEc2ModelVpnConnection) SetState(a ServicesEc2ModelVpnStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnConnection withState(com.amazonaws.services.ec2.model.VpnState)
func (jbobject *ServicesEc2ModelVpnConnection) WithState(a ServicesEc2ModelVpnStateInterface) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnState"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCustomerGatewayConfiguration(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) SetCustomerGatewayConfiguration(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCustomerGatewayConfiguration", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCustomerGatewayConfiguration()
func (jbobject *ServicesEc2ModelVpnConnection) GetCustomerGatewayConfiguration() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCustomerGatewayConfiguration", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpnConnection withCustomerGatewayConfiguration(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) WithCustomerGatewayConfiguration(a string) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCustomerGatewayConfiguration", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setType(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) SetType2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getType()
func (jbobject *ServicesEc2ModelVpnConnection) GetType() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpnConnection withType(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) WithType2(a string) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setType(com.amazonaws.services.ec2.model.GatewayType)
func (jbobject *ServicesEc2ModelVpnConnection) SetType(a ServicesEc2ModelGatewayTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/GatewayType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnConnection withType(com.amazonaws.services.ec2.model.GatewayType)
func (jbobject *ServicesEc2ModelVpnConnection) WithType(a ServicesEc2ModelGatewayTypeInterface) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/GatewayType"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCustomerGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) SetCustomerGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCustomerGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCustomerGatewayId()
func (jbobject *ServicesEc2ModelVpnConnection) GetCustomerGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCustomerGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpnConnection withCustomerGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) WithCustomerGatewayId(a string) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCustomerGatewayId", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpnGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) SetVpnGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpnGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpnGatewayId()
func (jbobject *ServicesEc2ModelVpnConnection) GetVpnGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpnGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.VpnConnection withVpnGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelVpnConnection) WithVpnGatewayId(a string) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpnGatewayId", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.Tag> getTags()
func (jbobject *ServicesEc2ModelVpnConnection) GetTags() []*ServicesEc2ModelTag {
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
func (jbobject *ServicesEc2ModelVpnConnection) SetTags(a []*ServicesEc2ModelTag)  {
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

// public com.amazonaws.services.ec2.model.VpnConnection withTags(com.amazonaws.services.ec2.model.Tag...)
func (jbobject *ServicesEc2ModelVpnConnection) WithTags(a ...*ServicesEc2ModelTag) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Tag")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Tag"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VpnConnection withTags(java.util.Collection<com.amazonaws.services.ec2.model.Tag>)
func (jbobject *ServicesEc2ModelVpnConnection) WithTags2(a []*ServicesEc2ModelTag) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withTags", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.VgwTelemetry> getVgwTelemetry()
func (jbobject *ServicesEc2ModelVpnConnection) GetVgwTelemetry() []*ServicesEc2ModelVgwTelemetry {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVgwTelemetry", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVgwTelemetry)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setVgwTelemetry(java.util.Collection<com.amazonaws.services.ec2.model.VgwTelemetry>)
func (jbobject *ServicesEc2ModelVpnConnection) SetVgwTelemetry(a []*ServicesEc2ModelVgwTelemetry)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVgwTelemetry", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnConnection withVgwTelemetry(com.amazonaws.services.ec2.model.VgwTelemetry...)
func (jbobject *ServicesEc2ModelVpnConnection) WithVgwTelemetry(a ...*ServicesEc2ModelVgwTelemetry) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VgwTelemetry")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVgwTelemetry", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VgwTelemetry"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VpnConnection withVgwTelemetry(java.util.Collection<com.amazonaws.services.ec2.model.VgwTelemetry>)
func (jbobject *ServicesEc2ModelVpnConnection) WithVgwTelemetry2(a []*ServicesEc2ModelVgwTelemetry) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVgwTelemetry", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOptions(com.amazonaws.services.ec2.model.VpnConnectionOptions)
func (jbobject *ServicesEc2ModelVpnConnection) SetOptions(a ServicesEc2ModelVpnConnectionOptionsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOptions", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnConnectionOptions"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnConnectionOptions getOptions()
func (jbobject *ServicesEc2ModelVpnConnection) GetOptions() *ServicesEc2ModelVpnConnectionOptions {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOptions", "com/amazonaws/services/ec2/model/VpnConnectionOptions")
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
	unique_x := &ServicesEc2ModelVpnConnectionOptions{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VpnConnection withOptions(com.amazonaws.services.ec2.model.VpnConnectionOptions)
func (jbobject *ServicesEc2ModelVpnConnection) WithOptions(a ServicesEc2ModelVpnConnectionOptionsInterface) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOptions", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnConnectionOptions"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.VpnStaticRoute> getRoutes()
func (jbobject *ServicesEc2ModelVpnConnection) GetRoutes() []*ServicesEc2ModelVpnStaticRoute {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRoutes", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelVpnStaticRoute)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setRoutes(java.util.Collection<com.amazonaws.services.ec2.model.VpnStaticRoute>)
func (jbobject *ServicesEc2ModelVpnConnection) SetRoutes(a []*ServicesEc2ModelVpnStaticRoute)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRoutes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.VpnConnection withRoutes(com.amazonaws.services.ec2.model.VpnStaticRoute...)
func (jbobject *ServicesEc2ModelVpnConnection) WithRoutes(a ...*ServicesEc2ModelVpnStaticRoute) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/VpnStaticRoute")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRoutes", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("com/amazonaws/services/ec2/model/VpnStaticRoute"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.VpnConnection withRoutes(java.util.Collection<com.amazonaws.services.ec2.model.VpnStaticRoute>)
func (jbobject *ServicesEc2ModelVpnConnection) WithRoutes2(a []*ServicesEc2ModelVpnStaticRoute) *ServicesEc2ModelVpnConnection {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withRoutes", "com/amazonaws/services/ec2/model/VpnConnection", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpnConnection) ToString() string {
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
func (jbobject *ServicesEc2ModelVpnConnection) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelVpnConnection) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.VpnConnection clone()
func (jbobject *ServicesEc2ModelVpnConnection) Clone() *ServicesEc2ModelVpnConnection {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/VpnConnection")
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
	unique_x := &ServicesEc2ModelVpnConnection{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelVpnConnection) Clone2() (*JavaLangObject, error) {
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


