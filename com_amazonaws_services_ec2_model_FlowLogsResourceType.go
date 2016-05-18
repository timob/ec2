package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelFlowLogsResourceTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.FlowLogsResourceType[] values()
	Values() []*ServicesEc2ModelFlowLogsResourceType

	// public static com.amazonaws.services.ec2.model.FlowLogsResourceType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelFlowLogsResourceType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.FlowLogsResourceType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelFlowLogsResourceType
}

type ServicesEc2ModelFlowLogsResourceType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.FlowLogsResourceType[] values()
func (jbobject *ServicesEc2ModelFlowLogsResourceType) Values() []*ServicesEc2ModelFlowLogsResourceType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/FlowLogsResourceType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/FlowLogsResourceType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/FlowLogsResourceType")
	dst := new([]*ServicesEc2ModelFlowLogsResourceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.FlowLogsResourceType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLogsResourceType) ValueOf(a string) *ServicesEc2ModelFlowLogsResourceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/FlowLogsResourceType", "valueOf", "com/amazonaws/services/ec2/model/FlowLogsResourceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLogsResourceType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelFlowLogsResourceType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.FlowLogsResourceType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelFlowLogsResourceType) FromValue(a string) *ServicesEc2ModelFlowLogsResourceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/FlowLogsResourceType", "fromValue", "com/amazonaws/services/ec2/model/FlowLogsResourceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelFlowLogsResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelFlowLogsResourceType) VPC() *ServicesEc2ModelFlowLogsResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/FlowLogsResourceType", "VPC", "com/amazonaws/services/ec2/model/FlowLogsResourceType")
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
	unique_x := &ServicesEc2ModelFlowLogsResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelFlowLogsResourceType) SetFieldVPC(val ServicesEc2ModelFlowLogsResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/FlowLogsResourceType", "VPC", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelFlowLogsResourceType) Subnet() *ServicesEc2ModelFlowLogsResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/FlowLogsResourceType", "Subnet", "com/amazonaws/services/ec2/model/FlowLogsResourceType")
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
	unique_x := &ServicesEc2ModelFlowLogsResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelFlowLogsResourceType) SetFieldSubnet(val ServicesEc2ModelFlowLogsResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/FlowLogsResourceType", "Subnet", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelFlowLogsResourceType) NetworkInterface() *ServicesEc2ModelFlowLogsResourceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/FlowLogsResourceType", "NetworkInterface", "com/amazonaws/services/ec2/model/FlowLogsResourceType")
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
	unique_x := &ServicesEc2ModelFlowLogsResourceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelFlowLogsResourceType) SetFieldNetworkInterface(val ServicesEc2ModelFlowLogsResourceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/FlowLogsResourceType", "NetworkInterface", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


