package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelHypervisorTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.HypervisorType[] values()
	Values() []*ServicesEc2ModelHypervisorType

	// public static com.amazonaws.services.ec2.model.HypervisorType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelHypervisorType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.HypervisorType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelHypervisorType
}

type ServicesEc2ModelHypervisorType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.HypervisorType[] values()
func (jbobject *ServicesEc2ModelHypervisorType) Values() []*ServicesEc2ModelHypervisorType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/HypervisorType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/HypervisorType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/HypervisorType")
	dst := new([]*ServicesEc2ModelHypervisorType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.HypervisorType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelHypervisorType) ValueOf(a string) *ServicesEc2ModelHypervisorType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/HypervisorType", "valueOf", "com/amazonaws/services/ec2/model/HypervisorType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHypervisorType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelHypervisorType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.HypervisorType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelHypervisorType) FromValue(a string) *ServicesEc2ModelHypervisorType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/HypervisorType", "fromValue", "com/amazonaws/services/ec2/model/HypervisorType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelHypervisorType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelHypervisorType) Ovm() *ServicesEc2ModelHypervisorType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/HypervisorType", "Ovm", "com/amazonaws/services/ec2/model/HypervisorType")
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
	unique_x := &ServicesEc2ModelHypervisorType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelHypervisorType) SetFieldOvm(val ServicesEc2ModelHypervisorTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/HypervisorType", "Ovm", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelHypervisorType) Xen() *ServicesEc2ModelHypervisorType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/HypervisorType", "Xen", "com/amazonaws/services/ec2/model/HypervisorType")
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
	unique_x := &ServicesEc2ModelHypervisorType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelHypervisorType) SetFieldXen(val ServicesEc2ModelHypervisorTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/HypervisorType", "Xen", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


