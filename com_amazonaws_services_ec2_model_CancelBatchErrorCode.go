package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCancelBatchErrorCodeInterface interface {

	// public static com.amazonaws.services.ec2.model.CancelBatchErrorCode[] values()
	Values() []*ServicesEc2ModelCancelBatchErrorCode

	// public static com.amazonaws.services.ec2.model.CancelBatchErrorCode valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelCancelBatchErrorCode

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.CancelBatchErrorCode fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelCancelBatchErrorCode
}

type ServicesEc2ModelCancelBatchErrorCode struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.CancelBatchErrorCode[] values()
func (jbobject *ServicesEc2ModelCancelBatchErrorCode) Values() []*ServicesEc2ModelCancelBatchErrorCode {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/CancelBatchErrorCode"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/CancelBatchErrorCode")
	dst := new([]*ServicesEc2ModelCancelBatchErrorCode)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.CancelBatchErrorCode valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelCancelBatchErrorCode) ValueOf(a string) *ServicesEc2ModelCancelBatchErrorCode {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "valueOf", "com/amazonaws/services/ec2/model/CancelBatchErrorCode", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelBatchErrorCode{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCancelBatchErrorCode) ToString() string {
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

// public static com.amazonaws.services.ec2.model.CancelBatchErrorCode fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelCancelBatchErrorCode) FromValue(a string) *ServicesEc2ModelCancelBatchErrorCode {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "fromValue", "com/amazonaws/services/ec2/model/CancelBatchErrorCode", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCancelBatchErrorCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelBatchErrorCode) FleetRequestIdDoesNotExist() *ServicesEc2ModelCancelBatchErrorCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "FleetRequestIdDoesNotExist", "com/amazonaws/services/ec2/model/CancelBatchErrorCode")
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
	unique_x := &ServicesEc2ModelCancelBatchErrorCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelBatchErrorCode) SetFieldFleetRequestIdDoesNotExist(val ServicesEc2ModelCancelBatchErrorCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "FleetRequestIdDoesNotExist", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelCancelBatchErrorCode) FleetRequestIdMalformed() *ServicesEc2ModelCancelBatchErrorCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "FleetRequestIdMalformed", "com/amazonaws/services/ec2/model/CancelBatchErrorCode")
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
	unique_x := &ServicesEc2ModelCancelBatchErrorCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelBatchErrorCode) SetFieldFleetRequestIdMalformed(val ServicesEc2ModelCancelBatchErrorCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "FleetRequestIdMalformed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelCancelBatchErrorCode) FleetRequestNotInCancellableState() *ServicesEc2ModelCancelBatchErrorCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "FleetRequestNotInCancellableState", "com/amazonaws/services/ec2/model/CancelBatchErrorCode")
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
	unique_x := &ServicesEc2ModelCancelBatchErrorCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelBatchErrorCode) SetFieldFleetRequestNotInCancellableState(val ServicesEc2ModelCancelBatchErrorCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "FleetRequestNotInCancellableState", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelCancelBatchErrorCode) UnexpectedError() *ServicesEc2ModelCancelBatchErrorCode {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "UnexpectedError", "com/amazonaws/services/ec2/model/CancelBatchErrorCode")
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
	unique_x := &ServicesEc2ModelCancelBatchErrorCode{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelCancelBatchErrorCode) SetFieldUnexpectedError(val ServicesEc2ModelCancelBatchErrorCodeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/CancelBatchErrorCode", "UnexpectedError", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


