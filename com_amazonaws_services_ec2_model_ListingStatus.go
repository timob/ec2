package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelListingStatusInterface interface {

	// public static com.amazonaws.services.ec2.model.ListingStatus[] values()
	Values() []*ServicesEc2ModelListingStatus

	// public static com.amazonaws.services.ec2.model.ListingStatus valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelListingStatus

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.ListingStatus fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelListingStatus
}

type ServicesEc2ModelListingStatus struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.ListingStatus[] values()
func (jbobject *ServicesEc2ModelListingStatus) Values() []*ServicesEc2ModelListingStatus {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ListingStatus", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/ListingStatus"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/ListingStatus")
	dst := new([]*ServicesEc2ModelListingStatus)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.ListingStatus valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelListingStatus) ValueOf(a string) *ServicesEc2ModelListingStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ListingStatus", "valueOf", "com/amazonaws/services/ec2/model/ListingStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelListingStatus{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelListingStatus) ToString() string {
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

// public static com.amazonaws.services.ec2.model.ListingStatus fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelListingStatus) FromValue(a string) *ServicesEc2ModelListingStatus {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/ListingStatus", "fromValue", "com/amazonaws/services/ec2/model/ListingStatus", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelListingStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingStatus) Active() *ServicesEc2ModelListingStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ListingStatus", "Active", "com/amazonaws/services/ec2/model/ListingStatus")
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
	unique_x := &ServicesEc2ModelListingStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingStatus) SetFieldActive(val ServicesEc2ModelListingStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ListingStatus", "Active", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelListingStatus) Pending() *ServicesEc2ModelListingStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ListingStatus", "Pending", "com/amazonaws/services/ec2/model/ListingStatus")
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
	unique_x := &ServicesEc2ModelListingStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingStatus) SetFieldPending(val ServicesEc2ModelListingStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ListingStatus", "Pending", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelListingStatus) Cancelled() *ServicesEc2ModelListingStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ListingStatus", "Cancelled", "com/amazonaws/services/ec2/model/ListingStatus")
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
	unique_x := &ServicesEc2ModelListingStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingStatus) SetFieldCancelled(val ServicesEc2ModelListingStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ListingStatus", "Cancelled", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelListingStatus) Closed() *ServicesEc2ModelListingStatus {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/ListingStatus", "Closed", "com/amazonaws/services/ec2/model/ListingStatus")
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
	unique_x := &ServicesEc2ModelListingStatus{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelListingStatus) SetFieldClosed(val ServicesEc2ModelListingStatusInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/ListingStatus", "Closed", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


