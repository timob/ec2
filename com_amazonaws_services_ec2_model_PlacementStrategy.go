package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPlacementStrategyInterface interface {

	// public static com.amazonaws.services.ec2.model.PlacementStrategy[] values()
	Values() []*ServicesEc2ModelPlacementStrategy

	// public static com.amazonaws.services.ec2.model.PlacementStrategy valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelPlacementStrategy

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.PlacementStrategy fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelPlacementStrategy
}

type ServicesEc2ModelPlacementStrategy struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.PlacementStrategy[] values()
func (jbobject *ServicesEc2ModelPlacementStrategy) Values() []*ServicesEc2ModelPlacementStrategy {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PlacementStrategy", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/PlacementStrategy"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/PlacementStrategy")
	dst := new([]*ServicesEc2ModelPlacementStrategy)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.PlacementStrategy valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementStrategy) ValueOf(a string) *ServicesEc2ModelPlacementStrategy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PlacementStrategy", "valueOf", "com/amazonaws/services/ec2/model/PlacementStrategy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacementStrategy{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPlacementStrategy) ToString() string {
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

// public static com.amazonaws.services.ec2.model.PlacementStrategy fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelPlacementStrategy) FromValue(a string) *ServicesEc2ModelPlacementStrategy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/PlacementStrategy", "fromValue", "com/amazonaws/services/ec2/model/PlacementStrategy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPlacementStrategy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPlacementStrategy) Cluster() *ServicesEc2ModelPlacementStrategy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/PlacementStrategy", "Cluster", "com/amazonaws/services/ec2/model/PlacementStrategy")
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
	unique_x := &ServicesEc2ModelPlacementStrategy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelPlacementStrategy) SetFieldCluster(val ServicesEc2ModelPlacementStrategyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/PlacementStrategy", "Cluster", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


