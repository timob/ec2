package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSpotInstanceTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.SpotInstanceType[] values()
	Values() []*ServicesEc2ModelSpotInstanceType

	// public static com.amazonaws.services.ec2.model.SpotInstanceType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelSpotInstanceType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.SpotInstanceType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelSpotInstanceType
}

type ServicesEc2ModelSpotInstanceType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.SpotInstanceType[] values()
func (jbobject *ServicesEc2ModelSpotInstanceType) Values() []*ServicesEc2ModelSpotInstanceType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SpotInstanceType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/SpotInstanceType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/SpotInstanceType")
	dst := new([]*ServicesEc2ModelSpotInstanceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.SpotInstanceType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceType) ValueOf(a string) *ServicesEc2ModelSpotInstanceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SpotInstanceType", "valueOf", "com/amazonaws/services/ec2/model/SpotInstanceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSpotInstanceType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.SpotInstanceType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelSpotInstanceType) FromValue(a string) *ServicesEc2ModelSpotInstanceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SpotInstanceType", "fromValue", "com/amazonaws/services/ec2/model/SpotInstanceType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSpotInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSpotInstanceType) OneTime() *ServicesEc2ModelSpotInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SpotInstanceType", "OneTime", "com/amazonaws/services/ec2/model/SpotInstanceType")
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
	unique_x := &ServicesEc2ModelSpotInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSpotInstanceType) SetFieldOneTime(val ServicesEc2ModelSpotInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SpotInstanceType", "OneTime", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSpotInstanceType) Persistent() *ServicesEc2ModelSpotInstanceType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SpotInstanceType", "Persistent", "com/amazonaws/services/ec2/model/SpotInstanceType")
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
	unique_x := &ServicesEc2ModelSpotInstanceType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSpotInstanceType) SetFieldPersistent(val ServicesEc2ModelSpotInstanceTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SpotInstanceType", "Persistent", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


