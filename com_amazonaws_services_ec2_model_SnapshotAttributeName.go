package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSnapshotAttributeNameInterface interface {

	// public static com.amazonaws.services.ec2.model.SnapshotAttributeName[] values()
	Values() []*ServicesEc2ModelSnapshotAttributeName

	// public static com.amazonaws.services.ec2.model.SnapshotAttributeName valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelSnapshotAttributeName

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.SnapshotAttributeName fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelSnapshotAttributeName
}

type ServicesEc2ModelSnapshotAttributeName struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.SnapshotAttributeName[] values()
func (jbobject *ServicesEc2ModelSnapshotAttributeName) Values() []*ServicesEc2ModelSnapshotAttributeName {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SnapshotAttributeName", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/SnapshotAttributeName"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/SnapshotAttributeName")
	dst := new([]*ServicesEc2ModelSnapshotAttributeName)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.SnapshotAttributeName valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotAttributeName) ValueOf(a string) *ServicesEc2ModelSnapshotAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SnapshotAttributeName", "valueOf", "com/amazonaws/services/ec2/model/SnapshotAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSnapshotAttributeName) ToString() string {
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

// public static com.amazonaws.services.ec2.model.SnapshotAttributeName fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelSnapshotAttributeName) FromValue(a string) *ServicesEc2ModelSnapshotAttributeName {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/SnapshotAttributeName", "fromValue", "com/amazonaws/services/ec2/model/SnapshotAttributeName", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelSnapshotAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSnapshotAttributeName) ProductCodes() *ServicesEc2ModelSnapshotAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SnapshotAttributeName", "ProductCodes", "com/amazonaws/services/ec2/model/SnapshotAttributeName")
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
	unique_x := &ServicesEc2ModelSnapshotAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSnapshotAttributeName) SetFieldProductCodes(val ServicesEc2ModelSnapshotAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SnapshotAttributeName", "ProductCodes", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelSnapshotAttributeName) CreateVolumePermission() *ServicesEc2ModelSnapshotAttributeName {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/SnapshotAttributeName", "CreateVolumePermission", "com/amazonaws/services/ec2/model/SnapshotAttributeName")
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
	unique_x := &ServicesEc2ModelSnapshotAttributeName{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelSnapshotAttributeName) SetFieldCreateVolumePermission(val ServicesEc2ModelSnapshotAttributeNameInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/SnapshotAttributeName", "CreateVolumePermission", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


