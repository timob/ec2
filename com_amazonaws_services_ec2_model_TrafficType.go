package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelTrafficTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.TrafficType[] values()
	Values() []*ServicesEc2ModelTrafficType

	// public static com.amazonaws.services.ec2.model.TrafficType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelTrafficType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.TrafficType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelTrafficType
}

type ServicesEc2ModelTrafficType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.TrafficType[] values()
func (jbobject *ServicesEc2ModelTrafficType) Values() []*ServicesEc2ModelTrafficType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/TrafficType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/TrafficType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/TrafficType")
	dst := new([]*ServicesEc2ModelTrafficType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.TrafficType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelTrafficType) ValueOf(a string) *ServicesEc2ModelTrafficType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/TrafficType", "valueOf", "com/amazonaws/services/ec2/model/TrafficType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTrafficType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelTrafficType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.TrafficType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelTrafficType) FromValue(a string) *ServicesEc2ModelTrafficType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/TrafficType", "fromValue", "com/amazonaws/services/ec2/model/TrafficType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelTrafficType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTrafficType) ACCEPT() *ServicesEc2ModelTrafficType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/TrafficType", "ACCEPT", "com/amazonaws/services/ec2/model/TrafficType")
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
	unique_x := &ServicesEc2ModelTrafficType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTrafficType) SetFieldACCEPT(val ServicesEc2ModelTrafficTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/TrafficType", "ACCEPT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelTrafficType) REJECT() *ServicesEc2ModelTrafficType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/TrafficType", "REJECT", "com/amazonaws/services/ec2/model/TrafficType")
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
	unique_x := &ServicesEc2ModelTrafficType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTrafficType) SetFieldREJECT(val ServicesEc2ModelTrafficTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/TrafficType", "REJECT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelTrafficType) ALL() *ServicesEc2ModelTrafficType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/TrafficType", "ALL", "com/amazonaws/services/ec2/model/TrafficType")
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
	unique_x := &ServicesEc2ModelTrafficType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelTrafficType) SetFieldALL(val ServicesEc2ModelTrafficTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/TrafficType", "ALL", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


