package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAccountAttributeNameInterface interface {

	// public static com.amazonaws.services.ec2.model.AccountAttributeName[] values()
	Values() []*ServicesEc2ModelAccountAttributeName

	// public static com.amazonaws.services.ec2.model.AccountAttributeName valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelAccountAttributeName

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.AccountAttributeName fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelAccountAttributeName
}

type ServicesEc2ModelAccountAttributeName struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.AccountAttributeName[] values()
func (jbobject *ServicesEc2ModelAccountAttributeName) Values() []*ServicesEc2ModelAccountAttributeName {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AccountAttributeName", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/AccountAttributeName"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/AccountAttributeName")
	dst := new([]*ServicesEc2ModelAccountAttributeName)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.AccountAttributeName valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelAccountAttributeName) ValueOf(a string) *ServicesEc2ModelAccountAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AccountAttributeName", "valueOf", "com/amazonaws/services/ec2/model/AccountAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAccountAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAccountAttributeName) ToString() string {
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

// public static com.amazonaws.services.ec2.model.AccountAttributeName fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelAccountAttributeName) FromValue(a string) *ServicesEc2ModelAccountAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AccountAttributeName", "fromValue", "com/amazonaws/services/ec2/model/AccountAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAccountAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAccountAttributeName) SupportedPlatforms() *ServicesEc2ModelAccountAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AccountAttributeName", "SupportedPlatforms", "com/amazonaws/services/ec2/model/AccountAttributeName")
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
	unique_x := &ServicesEc2ModelAccountAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAccountAttributeName) SetFieldSupportedPlatforms(val ServicesEc2ModelAccountAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AccountAttributeName", "SupportedPlatforms", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAccountAttributeName) DefaultVpc() *ServicesEc2ModelAccountAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AccountAttributeName", "DefaultVpc", "com/amazonaws/services/ec2/model/AccountAttributeName")
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
	unique_x := &ServicesEc2ModelAccountAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAccountAttributeName) SetFieldDefaultVpc(val ServicesEc2ModelAccountAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AccountAttributeName", "DefaultVpc", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


