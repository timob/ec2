package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAllocationStrategyInterface interface {

	// public static com.amazonaws.services.ec2.model.AllocationStrategy[] values()
	Values() []*ServicesEc2ModelAllocationStrategy

	// public static com.amazonaws.services.ec2.model.AllocationStrategy valueOf(java.lang.String)
	ValueOf(a string) *ServicesEc2ModelAllocationStrategy

	// public java.lang.String toString()
	ToString() string

	// public static com.amazonaws.services.ec2.model.AllocationStrategy fromValue(java.lang.String)
	FromValue(a string) *ServicesEc2ModelAllocationStrategy
}

type ServicesEc2ModelAllocationStrategy struct {
	*javabind.Callable
}

// public static com.amazonaws.services.ec2.model.AllocationStrategy[] values()
func (jbobject *ServicesEc2ModelAllocationStrategy) Values() []*ServicesEc2ModelAllocationStrategy {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AllocationStrategy", "values", javabind.ObjectArrayType("com/amazonaws/services/ec2/model/AllocationStrategy"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/services/ec2/model/AllocationStrategy")
	dst := new([]*ServicesEc2ModelAllocationStrategy)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.services.ec2.model.AllocationStrategy valueOf(java.lang.String)
func (jbobject *ServicesEc2ModelAllocationStrategy) ValueOf(a string) *ServicesEc2ModelAllocationStrategy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AllocationStrategy", "valueOf", "com/amazonaws/services/ec2/model/AllocationStrategy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocationStrategy{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAllocationStrategy) ToString() string {
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

// public static com.amazonaws.services.ec2.model.AllocationStrategy fromValue(java.lang.String)
func (jbobject *ServicesEc2ModelAllocationStrategy) FromValue(a string) *ServicesEc2ModelAllocationStrategy {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/services/ec2/model/AllocationStrategy", "fromValue", "com/amazonaws/services/ec2/model/AllocationStrategy", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAllocationStrategy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAllocationStrategy) LowestPrice() *ServicesEc2ModelAllocationStrategy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AllocationStrategy", "LowestPrice", "com/amazonaws/services/ec2/model/AllocationStrategy")
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
	unique_x := &ServicesEc2ModelAllocationStrategy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAllocationStrategy) SetFieldLowestPrice(val ServicesEc2ModelAllocationStrategyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AllocationStrategy", "LowestPrice", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ServicesEc2ModelAllocationStrategy) Diversified() *ServicesEc2ModelAllocationStrategy {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/services/ec2/model/AllocationStrategy", "Diversified", "com/amazonaws/services/ec2/model/AllocationStrategy")
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
	unique_x := &ServicesEc2ModelAllocationStrategy{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ServicesEc2ModelAllocationStrategy) SetFieldDiversified(val ServicesEc2ModelAllocationStrategyInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/services/ec2/model/AllocationStrategy", "Diversified", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


