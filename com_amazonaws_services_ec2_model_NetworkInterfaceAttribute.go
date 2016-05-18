package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNetworkInterfaceAttributeInterface interface {

	// public static com.amazonaws.services.ec2.model.NetworkInterfaceAttribute[] values()
	Values() []*ServicesEc2ModelNetworkInterfaceAttribute

	// public static com.amazonaws.services.ec2.model.NetworkInterfaceAttribute valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelNetworkInterfaceAttribute

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.NetworkInterfaceAttribute fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelNetworkInterfaceAttribute
}

type ServicesEc2ModelNetworkInterfaceAttribute struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.NetworkInterfaceAttribute[] values()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) Values() []*ServicesEc2ModelNetworkInterfaceAttribute {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/NetworkInterfaceAttribute")
	dst := new([]*ServicesEc2ModelNetworkInterfaceAttribute)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.NetworkInterfaceAttribute valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) ValueOf(a string) *ServicesEc2ModelNetworkInterfaceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "valueOf", "com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) ToString() string {
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

// public static com.amazonaws.services.ec2.model.NetworkInterfaceAttribute fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) FromValue(a string) *ServicesEc2ModelNetworkInterfaceAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "fromValue", "com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) Description() *ServicesEc2ModelNetworkInterfaceAttribute {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "Description", "com/amazonaws/services/ec2/model/NetworkInterfaceAttribute")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) SetFieldDescription(val ServicesEc2ModelNetworkInterfaceAttributeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "Description", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) GroupSet() *ServicesEc2ModelNetworkInterfaceAttribute {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "GroupSet", "com/amazonaws/services/ec2/model/NetworkInterfaceAttribute")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) SetFieldGroupSet(val ServicesEc2ModelNetworkInterfaceAttributeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "GroupSet", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) SourceDestCheck() *ServicesEc2ModelNetworkInterfaceAttribute {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "SourceDestCheck", "com/amazonaws/services/ec2/model/NetworkInterfaceAttribute")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) SetFieldSourceDestCheck(val ServicesEc2ModelNetworkInterfaceAttributeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "SourceDestCheck", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) Attachment() *ServicesEc2ModelNetworkInterfaceAttribute {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "Attachment", "com/amazonaws/services/ec2/model/NetworkInterfaceAttribute")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceAttribute) SetFieldAttachment(val ServicesEc2ModelNetworkInterfaceAttributeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceAttribute", "Attachment", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


