package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVirtualizationTypeInterface interface {

	// public static com.amazonaws.services.ec2.model.VirtualizationType[] values()
	Values() []*ServicesEc2ModelVirtualizationType

	// public static com.amazonaws.services.ec2.model.VirtualizationType valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVirtualizationType

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VirtualizationType fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVirtualizationType
}

type ServicesEc2ModelVirtualizationType struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VirtualizationType[] values()
func (jbobject *ServicesEc2ModelVirtualizationType) Values() []*ServicesEc2ModelVirtualizationType {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VirtualizationType", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VirtualizationType"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VirtualizationType")
	dst := new([]*ServicesEc2ModelVirtualizationType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VirtualizationType valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVirtualizationType) ValueOf(a string) *ServicesEc2ModelVirtualizationType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VirtualizationType", "valueOf", "com/amazonaws/services/ec2/model/VirtualizationType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVirtualizationType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVirtualizationType) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VirtualizationType fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVirtualizationType) FromValue(a string) *ServicesEc2ModelVirtualizationType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VirtualizationType", "fromValue", "com/amazonaws/services/ec2/model/VirtualizationType", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVirtualizationType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVirtualizationType) Hvm() *ServicesEc2ModelVirtualizationType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VirtualizationType", "Hvm", "com/amazonaws/services/ec2/model/VirtualizationType")
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
	unique_x := &ServicesEc2ModelVirtualizationType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVirtualizationType) SetFieldHvm(val ServicesEc2ModelVirtualizationTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VirtualizationType", "Hvm", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVirtualizationType) Paravirtual() *ServicesEc2ModelVirtualizationType {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VirtualizationType", "Paravirtual", "com/amazonaws/services/ec2/model/VirtualizationType")
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
	unique_x := &ServicesEc2ModelVirtualizationType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVirtualizationType) SetFieldParavirtual(val ServicesEc2ModelVirtualizationTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VirtualizationType", "Paravirtual", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


