package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpcAttributeNameInterface interface {

	// public static com.amazonaws.services.ec2.model.VpcAttributeName[] values()
	Values() []*ServicesEc2ModelVpcAttributeName

	// public static com.amazonaws.services.ec2.model.VpcAttributeName valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVpcAttributeName

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VpcAttributeName fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVpcAttributeName
}

type ServicesEc2ModelVpcAttributeName struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VpcAttributeName[] values()
func (jbobject *ServicesEc2ModelVpcAttributeName) Values() []*ServicesEc2ModelVpcAttributeName {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpcAttributeName", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VpcAttributeName"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VpcAttributeName")
	dst := new([]*ServicesEc2ModelVpcAttributeName)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VpcAttributeName valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVpcAttributeName) ValueOf(a string) *ServicesEc2ModelVpcAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpcAttributeName", "valueOf", "com/amazonaws/services/ec2/model/VpcAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpcAttributeName) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VpcAttributeName fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVpcAttributeName) FromValue(a string) *ServicesEc2ModelVpcAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpcAttributeName", "fromValue", "com/amazonaws/services/ec2/model/VpcAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcAttributeName) EnableDnsSupport() *ServicesEc2ModelVpcAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcAttributeName", "EnableDnsSupport", "com/amazonaws/services/ec2/model/VpcAttributeName")
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
	unique_x := &ServicesEc2ModelVpcAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcAttributeName) SetFieldEnableDnsSupport(val ServicesEc2ModelVpcAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcAttributeName", "EnableDnsSupport", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcAttributeName) EnableDnsHostnames() *ServicesEc2ModelVpcAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcAttributeName", "EnableDnsHostnames", "com/amazonaws/services/ec2/model/VpcAttributeName")
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
	unique_x := &ServicesEc2ModelVpcAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcAttributeName) SetFieldEnableDnsHostnames(val ServicesEc2ModelVpcAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcAttributeName", "EnableDnsHostnames", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


