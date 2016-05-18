package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelGatewayTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.GatewayType[] values()
	Values() []*ServicesEc2ModelGatewayType

	// public static com.amazonaws.services.ec2.model.GatewayType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelGatewayType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.GatewayType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelGatewayType
}

type ServicesEc2ModelGatewayType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.GatewayType[] values()
func (jbobject *ServicesEc2ModelGatewayType) Values() []*ServicesEc2ModelGatewayType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/GatewayType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/GatewayType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/GatewayType")
	dst := new([]*ServicesEc2ModelGatewayType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.GatewayType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelGatewayType) ValueOf(a string) *ServicesEc2ModelGatewayType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/GatewayType", "valueOf", "com/amazonaws/services/ec2/model/GatewayType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelGatewayType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelGatewayType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.GatewayType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelGatewayType) FromValue(a string) *ServicesEc2ModelGatewayType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/GatewayType", "fromValue", "com/amazonaws/services/ec2/model/GatewayType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelGatewayType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelGatewayType) Ipsec1() *ServicesEc2ModelGatewayType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/GatewayType", "Ipsec1", "com/amazonaws/services/ec2/model/GatewayType")
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
	unique_x := &ServicesEc2ModelGatewayType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelGatewayType) SetFieldIpsec1(val ServicesEc2ModelGatewayTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/GatewayType", "Ipsec1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


