package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelRouteInterface interface {
	JavaLangObjectInterface

	// public void setDestinationCidrBlock(java.lang.String)
	SetDestinationCidrBlock(a string) 

	// public java.lang.String getDestinationCidrBlock()
	GetDestinationCidrBlock() string

	// public com.amazonaws.services.ec2.model.Route withDestinationCidrBlock(java.lang.String)
	WithDestinationCidrBlock(a string) *ServicesEc2ModelRoute

	// public void setDestinationPrefixListId(java.lang.String)
	SetDestinationPrefixListId(a string) 

	// public java.lang.String getDestinationPrefixListId()
	GetDestinationPrefixListId() string

	// public com.amazonaws.services.ec2.model.Route withDestinationPrefixListId(java.lang.String)
	WithDestinationPrefixListId(a string) *ServicesEc2ModelRoute

	// public void setGatewayId(java.lang.String)
	SetGatewayId(a string) 

	// public java.lang.String getGatewayId()
	GetGatewayId() string

	// public com.amazonaws.services.ec2.model.Route withGatewayId(java.lang.String)
	WithGatewayId(a string) *ServicesEc2ModelRoute

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.Route withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelRoute

	// public void setInstanceOwnerId(java.lang.String)
	SetInstanceOwnerId(a string) 

	// public java.lang.String getInstanceOwnerId()
	GetInstanceOwnerId() string

	// public com.amazonaws.services.ec2.model.Route withInstanceOwnerId(java.lang.String)
	WithInstanceOwnerId(a string) *ServicesEc2ModelRoute

	// public void setNetworkInterfaceId(java.lang.String)
	SetNetworkInterfaceId(a string) 

	// public java.lang.String getNetworkInterfaceId()
	GetNetworkInterfaceId() string

	// public com.amazonaws.services.ec2.model.Route withNetworkInterfaceId(java.lang.String)
	WithNetworkInterfaceId(a string) *ServicesEc2ModelRoute

	// public void setVpcPeeringConnectionId(java.lang.String)
	SetVpcPeeringConnectionId(a string) 

	// public java.lang.String getVpcPeeringConnectionId()
	GetVpcPeeringConnectionId() string

	// public com.amazonaws.services.ec2.model.Route withVpcPeeringConnectionId(java.lang.String)
	WithVpcPeeringConnectionId(a string) *ServicesEc2ModelRoute

	// public void setNatGatewayId(java.lang.String)
	SetNatGatewayId(a string) 

	// public java.lang.String getNatGatewayId()
	GetNatGatewayId() string

	// public com.amazonaws.services.ec2.model.Route withNatGatewayId(java.lang.String)
	WithNatGatewayId(a string) *ServicesEc2ModelRoute

	// public void setState(java.lang.String)
	SetState2(a string) 

	// public java.lang.String getState()
	GetState() string

	// public com.amazonaws.services.ec2.model.Route withState(java.lang.String)
	WithState2(a string) *ServicesEc2ModelRoute

	// public void setState(com.amazonaws.services.ec2.model.RouteState)
	SetState(a ServicesEc2ModelRouteStateInterface) 

	// public com.amazonaws.services.ec2.model.Route withState(com.amazonaws.services.ec2.model.RouteState)
	WithState(a ServicesEc2ModelRouteStateInterface) *ServicesEc2ModelRoute

	// public void setOrigin(java.lang.String)
	SetOrigin2(a string) 

	// public java.lang.String getOrigin()
	GetOrigin() string

	// public com.amazonaws.services.ec2.model.Route withOrigin(java.lang.String)
	WithOrigin2(a string) *ServicesEc2ModelRoute

	// public void setOrigin(com.amazonaws.services.ec2.model.RouteOrigin)
	SetOrigin(a ServicesEc2ModelRouteOriginInterface) 

	// public com.amazonaws.services.ec2.model.Route withOrigin(com.amazonaws.services.ec2.model.RouteOrigin)
	WithOrigin(a ServicesEc2ModelRouteOriginInterface) *ServicesEc2ModelRoute

	// public com.amazonaws.services.ec2.model.Route clone()
	Clone() *ServicesEc2ModelRoute
}

type ServicesEc2ModelRoute struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.Route()
func NewServicesEc2ModelRoute() (*ServicesEc2ModelRoute) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/Route")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelRoute{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDestinationCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetDestinationCidrBlock(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDestinationCidrBlock", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDestinationCidrBlock()
func (jbobject *ServicesEc2ModelRoute) GetDestinationCidrBlock() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDestinationCidrBlock", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Route withDestinationCidrBlock(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithDestinationCidrBlock(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDestinationCidrBlock", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setDestinationPrefixListId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetDestinationPrefixListId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDestinationPrefixListId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getDestinationPrefixListId()
func (jbobject *ServicesEc2ModelRoute) GetDestinationPrefixListId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDestinationPrefixListId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Route withDestinationPrefixListId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithDestinationPrefixListId(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withDestinationPrefixListId", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getGatewayId()
func (jbobject *ServicesEc2ModelRoute) GetGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Route withGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithGatewayId(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withGatewayId", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelRoute) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Route withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithInstanceId(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setInstanceOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetInstanceOwnerId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceOwnerId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceOwnerId()
func (jbobject *ServicesEc2ModelRoute) GetInstanceOwnerId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceOwnerId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Route withInstanceOwnerId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithInstanceOwnerId(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceOwnerId", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetNetworkInterfaceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNetworkInterfaceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNetworkInterfaceId()
func (jbobject *ServicesEc2ModelRoute) GetNetworkInterfaceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNetworkInterfaceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Route withNetworkInterfaceId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithNetworkInterfaceId(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNetworkInterfaceId", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetVpcPeeringConnectionId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcPeeringConnectionId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcPeeringConnectionId()
func (jbobject *ServicesEc2ModelRoute) GetVpcPeeringConnectionId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcPeeringConnectionId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Route withVpcPeeringConnectionId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithVpcPeeringConnectionId(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcPeeringConnectionId", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNatGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetNatGatewayId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNatGatewayId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNatGatewayId()
func (jbobject *ServicesEc2ModelRoute) GetNatGatewayId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNatGatewayId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Route withNatGatewayId(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithNatGatewayId(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNatGatewayId", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetState2(a string)  {
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
func (jbobject *ServicesEc2ModelRoute) GetState() string {
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

// public com.amazonaws.services.ec2.model.Route withState(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithState2(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setState(com.amazonaws.services.ec2.model.RouteState)
func (jbobject *ServicesEc2ModelRoute) SetState(a ServicesEc2ModelRouteStateInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setState", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RouteState"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Route withState(com.amazonaws.services.ec2.model.RouteState)
func (jbobject *ServicesEc2ModelRoute) WithState(a ServicesEc2ModelRouteStateInterface) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withState", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RouteState"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOrigin(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) SetOrigin2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrigin", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getOrigin()
func (jbobject *ServicesEc2ModelRoute) GetOrigin() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrigin", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.Route withOrigin(java.lang.String)
func (jbobject *ServicesEc2ModelRoute) WithOrigin2(a string) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOrigin", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public void setOrigin(com.amazonaws.services.ec2.model.RouteOrigin)
func (jbobject *ServicesEc2ModelRoute) SetOrigin(a ServicesEc2ModelRouteOriginInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setOrigin", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/RouteOrigin"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Route withOrigin(com.amazonaws.services.ec2.model.RouteOrigin)
func (jbobject *ServicesEc2ModelRoute) WithOrigin(a ServicesEc2ModelRouteOriginInterface) *ServicesEc2ModelRoute {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withOrigin", "com/amazonaws/services/ec2/model/Route", conv_a.Value().Cast("com/amazonaws/services/ec2/model/RouteOrigin"))
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelRoute) ToString() string {
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
func (jbobject *ServicesEc2ModelRoute) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelRoute) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.Route clone()
func (jbobject *ServicesEc2ModelRoute) Clone() *ServicesEc2ModelRoute {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/Route")
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
	unique_x := &ServicesEc2ModelRoute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelRoute) Clone2() (*JavaLangObject, error) {
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


