package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelResourceTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.ResourceType[] values()
	Values() []*ServicesEc2ModelResourceType

	// public static com.amazonaws.services.ec2.model.ResourceType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelResourceType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ResourceType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelResourceType
}

type ServicesEc2ModelResourceType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ResourceType[] values()
func (jbobject *ServicesEc2ModelResourceType) Values() []*ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ResourceType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ResourceType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ResourceType")
	dst := new([]*ServicesEc2ModelResourceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ResourceType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelResourceType) ValueOf(a string) *ServicesEc2ModelResourceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ResourceType", "valueOf", "com/amazonaws/services/ec2/model/ResourceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelResourceType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ResourceType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelResourceType) FromValue(a string) *ServicesEc2ModelResourceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ResourceType", "fromValue", "com/amazonaws/services/ec2/model/ResourceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) CustomerGateway() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "CustomerGateway", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldCustomerGateway(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "CustomerGateway", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) DhcpOptions() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "DhcpOptions", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldDhcpOptions(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "DhcpOptions", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) Image() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Image", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldImage(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Image", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) Instance() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Instance", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldInstance(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Instance", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) InternetGateway() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "InternetGateway", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldInternetGateway(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "InternetGateway", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) NetworkAcl() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "NetworkAcl", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldNetworkAcl(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "NetworkAcl", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) NetworkInterface() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "NetworkInterface", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldNetworkInterface(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "NetworkInterface", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) ReservedInstances() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "ReservedInstances", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldReservedInstances(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "ReservedInstances", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) RouteTable() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "RouteTable", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldRouteTable(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "RouteTable", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) Snapshot() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Snapshot", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldSnapshot(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Snapshot", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) SpotInstancesRequest() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "SpotInstancesRequest", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldSpotInstancesRequest(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "SpotInstancesRequest", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) Subnet() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Subnet", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldSubnet(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Subnet", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) SecurityGroup() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "SecurityGroup", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldSecurityGroup(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "SecurityGroup", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) Volume() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Volume", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldVolume(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Volume", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) Vpc() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Vpc", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldVpc(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "Vpc", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) VpnConnection() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "VpnConnection", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldVpnConnection(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "VpnConnection", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelResourceType) VpnGateway() *ServicesEc2ModelResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ResourceType", "VpnGateway", "com/amazonaws/services/ec2/model/ResourceType")
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
	unique_x := &ServicesEc2ModelResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelResourceType) SetFieldVpnGateway(val ServicesEc2ModelResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ResourceType", "VpnGateway", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


