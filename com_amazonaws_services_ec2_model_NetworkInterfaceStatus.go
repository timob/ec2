package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelNetworkInterfaceStatusInterface interface {

	// public static com.amazonaws.services.ec2.model.NetworkInterfaceStatus[] values()
	Values() []*ServicesEc2ModelNetworkInterfaceStatus

	// public static com.amazonaws.services.ec2.model.NetworkInterfaceStatus valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelNetworkInterfaceStatus

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.NetworkInterfaceStatus fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelNetworkInterfaceStatus
}

type ServicesEc2ModelNetworkInterfaceStatus struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.NetworkInterfaceStatus[] values()
func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) Values() []*ServicesEc2ModelNetworkInterfaceStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/NetworkInterfaceStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/NetworkInterfaceStatus")
	dst := new([]*ServicesEc2ModelNetworkInterfaceStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.NetworkInterfaceStatus valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) ValueOf(a string) *ServicesEc2ModelNetworkInterfaceStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "valueOf", "com/amazonaws/services/ec2/model/NetworkInterfaceStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) ToString() string {
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

// public static com.amazonaws.services.ec2.model.NetworkInterfaceStatus fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) FromValue(a string) *ServicesEc2ModelNetworkInterfaceStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "fromValue", "com/amazonaws/services/ec2/model/NetworkInterfaceStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelNetworkInterfaceStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) Available() *ServicesEc2ModelNetworkInterfaceStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "Available", "com/amazonaws/services/ec2/model/NetworkInterfaceStatus")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) SetFieldAvailable(val ServicesEc2ModelNetworkInterfaceStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) Attaching() *ServicesEc2ModelNetworkInterfaceStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "Attaching", "com/amazonaws/services/ec2/model/NetworkInterfaceStatus")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) SetFieldAttaching(val ServicesEc2ModelNetworkInterfaceStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "Attaching", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) InUse() *ServicesEc2ModelNetworkInterfaceStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "InUse", "com/amazonaws/services/ec2/model/NetworkInterfaceStatus")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) SetFieldInUse(val ServicesEc2ModelNetworkInterfaceStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "InUse", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) Detaching() *ServicesEc2ModelNetworkInterfaceStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "Detaching", "com/amazonaws/services/ec2/model/NetworkInterfaceStatus")
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
	unique_x := &ServicesEc2ModelNetworkInterfaceStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelNetworkInterfaceStatus) SetFieldDetaching(val ServicesEc2ModelNetworkInterfaceStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/NetworkInterfaceStatus", "Detaching", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


