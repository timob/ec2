package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface interface {

	// public static com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode[] values()
	Values() []*ServicesEc2ModelVpcPeeringConnectionStateReasonCode

	// public static com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelVpcPeeringConnectionStateReasonCode

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelVpcPeeringConnectionStateReasonCode
}

type ServicesEc2ModelVpcPeeringConnectionStateReasonCode struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode[] values()
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) Values() []*ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
	dst := new([]*ServicesEc2ModelVpcPeeringConnectionStateReasonCode)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) ValueOf(a string) *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "valueOf", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) ToString() string {
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

// public static com.amazonaws.services.ec2.model.VpcPeeringConnectionStateReasonCode fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) FromValue(a string) *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "fromValue", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) InitiatingRequest() *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "InitiatingRequest", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) SetFieldInitiatingRequest(val ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "InitiatingRequest", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) PendingAcceptance() *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "PendingAcceptance", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) SetFieldPendingAcceptance(val ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "PendingAcceptance", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) Active() *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Active", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) SetFieldActive(val ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) Deleted() *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Deleted", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) SetFieldDeleted(val ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Deleted", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) Rejected() *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Rejected", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) SetFieldRejected(val ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Rejected", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) Failed() *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Failed", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) SetFieldFailed(val ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Failed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) Expired() *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Expired", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) SetFieldExpired(val ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Expired", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) Provisioning() *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Provisioning", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) SetFieldProvisioning(val ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Provisioning", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) Deleting() *ServicesEc2ModelVpcPeeringConnectionStateReasonCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Deleting", "com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode")
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
	unique_x := &ServicesEc2ModelVpcPeeringConnectionStateReasonCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelVpcPeeringConnectionStateReasonCode) SetFieldDeleting(val ServicesEc2ModelVpcPeeringConnectionStateReasonCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/VpcPeeringConnectionStateReasonCode", "Deleting", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


