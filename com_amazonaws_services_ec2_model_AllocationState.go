package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAllocationStateInterface interface {

	// public static com.amazonaws.services.ec2.model.AllocationState[] values()
	Values() []*ServicesEc2ModelAllocationState

	// public static com.amazonaws.services.ec2.model.AllocationState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelAllocationState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.AllocationState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelAllocationState
}

type ServicesEc2ModelAllocationState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.AllocationState[] values()
func (jbobject *ServicesEc2ModelAllocationState) Values() []*ServicesEc2ModelAllocationState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AllocationState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/AllocationState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/AllocationState")
	dst := new([]*ServicesEc2ModelAllocationState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.AllocationState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelAllocationState) ValueOf(a string) *ServicesEc2ModelAllocationState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AllocationState", "valueOf", "com/amazonaws/services/ec2/model/AllocationState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocationState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAllocationState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.AllocationState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelAllocationState) FromValue(a string) *ServicesEc2ModelAllocationState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AllocationState", "fromValue", "com/amazonaws/services/ec2/model/AllocationState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocationState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAllocationState) Available() *ServicesEc2ModelAllocationState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AllocationState", "Available", "com/amazonaws/services/ec2/model/AllocationState")
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
	unique_x := &ServicesEc2ModelAllocationState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAllocationState) SetFieldAvailable(val ServicesEc2ModelAllocationStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AllocationState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAllocationState) UnderAssessment() *ServicesEc2ModelAllocationState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AllocationState", "UnderAssessment", "com/amazonaws/services/ec2/model/AllocationState")
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
	unique_x := &ServicesEc2ModelAllocationState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAllocationState) SetFieldUnderAssessment(val ServicesEc2ModelAllocationStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AllocationState", "UnderAssessment", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAllocationState) PermanentFailure() *ServicesEc2ModelAllocationState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AllocationState", "PermanentFailure", "com/amazonaws/services/ec2/model/AllocationState")
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
	unique_x := &ServicesEc2ModelAllocationState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAllocationState) SetFieldPermanentFailure(val ServicesEc2ModelAllocationStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AllocationState", "PermanentFailure", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAllocationState) Released() *ServicesEc2ModelAllocationState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AllocationState", "Released", "com/amazonaws/services/ec2/model/AllocationState")
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
	unique_x := &ServicesEc2ModelAllocationState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAllocationState) SetFieldReleased(val ServicesEc2ModelAllocationStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AllocationState", "Released", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAllocationState) ReleasedPermanentFailure() *ServicesEc2ModelAllocationState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AllocationState", "ReleasedPermanentFailure", "com/amazonaws/services/ec2/model/AllocationState")
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
	unique_x := &ServicesEc2ModelAllocationState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAllocationState) SetFieldReleasedPermanentFailure(val ServicesEc2ModelAllocationStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AllocationState", "ReleasedPermanentFailure", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


