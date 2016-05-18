package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNetworkInterfaceTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.NetworkInterfaceType[] values()
	Values() []*ServicesEc2ModelNetworkInterfaceType

	// public static com.amazonaws.services.ec2.model.NetworkInterfaceType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelNetworkInterfaceType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.NetworkInterfaceType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelNetworkInterfaceType
}

type ServicesEc2ModelNetworkInterfaceType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.NetworkInterfaceType[] values()
func (jbobject *ServicesEc2ModelNetworkInterfaceType) Values() []*ServicesEc2ModelNetworkInterfaceType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NetworkInterfaceType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/NetworkInterfaceType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/NetworkInterfaceType")
	dst := new([]*ServicesEc2ModelNetworkInterfaceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.NetworkInterfaceType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceType) ValueOf(a string) *ServicesEc2ModelNetworkInterfaceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NetworkInterfaceType", "valueOf", "com/amazonaws/services/ec2/model/NetworkInterfaceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNetworkInterfaceType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.NetworkInterfaceType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceType) FromValue(a string) *ServicesEc2ModelNetworkInterfaceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NetworkInterfaceType", "fromValue", "com/amazonaws/services/ec2/model/NetworkInterfaceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceType) Interface() *ServicesEc2ModelNetworkInterfaceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceType", "Interface", "com/amazonaws/services/ec2/model/NetworkInterfaceType")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceType) SetFieldInterface(val ServicesEc2ModelNetworkInterfaceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceType", "Interface", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNetworkInterfaceType) NatGateway() *ServicesEc2ModelNetworkInterfaceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceType", "NatGateway", "com/amazonaws/services/ec2/model/NetworkInterfaceType")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceType) SetFieldNatGateway(val ServicesEc2ModelNetworkInterfaceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceType", "NatGateway", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


