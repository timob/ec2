package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelListingStateInterface interface {

	// public static com.amazonaws.services.ec2.model.ListingState[] values()
	Values() []*ServicesEc2ModelListingState

	// public static com.amazonaws.services.ec2.model.ListingState valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelListingState

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ListingState fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelListingState
}

type ServicesEc2ModelListingState struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ListingState[] values()
func (jbobject *ServicesEc2ModelListingState) Values() []*ServicesEc2ModelListingState {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ListingState", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ListingState"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ListingState")
	dst := new([]*ServicesEc2ModelListingState)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ListingState valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelListingState) ValueOf(a string) *ServicesEc2ModelListingState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ListingState", "valueOf", "com/amazonaws/services/ec2/model/ListingState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelListingState{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelListingState) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ListingState fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelListingState) FromValue(a string) *ServicesEc2ModelListingState {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ListingState", "fromValue", "com/amazonaws/services/ec2/model/ListingState", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelListingState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingState) Available() *ServicesEc2ModelListingState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ListingState", "Available", "com/amazonaws/services/ec2/model/ListingState")
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
	unique_x := &ServicesEc2ModelListingState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingState) SetFieldAvailable(val ServicesEc2ModelListingStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ListingState", "Available", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelListingState) Sold() *ServicesEc2ModelListingState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ListingState", "Sold", "com/amazonaws/services/ec2/model/ListingState")
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
	unique_x := &ServicesEc2ModelListingState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingState) SetFieldSold(val ServicesEc2ModelListingStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ListingState", "Sold", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelListingState) Cancelled() *ServicesEc2ModelListingState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ListingState", "Cancelled", "com/amazonaws/services/ec2/model/ListingState")
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
	unique_x := &ServicesEc2ModelListingState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingState) SetFieldCancelled(val ServicesEc2ModelListingStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ListingState", "Cancelled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelListingState) Pending() *ServicesEc2ModelListingState {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ListingState", "Pending", "com/amazonaws/services/ec2/model/ListingState")
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
	unique_x := &ServicesEc2ModelListingState{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingState) SetFieldPending(val ServicesEc2ModelListingStateInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ListingState", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


