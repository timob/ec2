package ec2

import "github.com/timob/javabind"

type RegionsRegionsInterface interface {

	// public static com.amazonaws.regions.Regions[] values()
	Values() []*RegionsRegions

	// public static com.amazonaws.regions.Regions valueOf(java.lang.String)
	ValueOf(a string) *RegionsRegions

	// public java.lang.String getName()
	GetName() string

	// public static com.amazonaws.regions.Regions fromName(java.lang.String)
	FromName(a string) *RegionsRegions

	// public static com.amazonaws.regions.Region getCurrentRegion()
	GetCurrentRegion() *RegionsRegion
}

type RegionsRegions struct {
	*javabind.Callable
}

// public static com.amazonaws.regions.Regions[] values()
func (jbobject *RegionsRegions) Values() []*RegionsRegions {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/regions/Regions", "values", javabind.ObjectArrayType("com/amazonaws/regions/Regions"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/amazonaws/regions/Regions")
	dst := new([]*RegionsRegions)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.amazonaws.regions.Regions valueOf(java.lang.String)
func (jbobject *RegionsRegions) ValueOf(a string) *RegionsRegions {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/regions/Regions", "valueOf", "com/amazonaws/regions/Regions", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getName()
func (jbobject *RegionsRegions) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public static com.amazonaws.regions.Regions fromName(java.lang.String)
func (jbobject *RegionsRegions) FromName(a string) *RegionsRegions {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/regions/Regions", "fromName", "com/amazonaws/regions/Regions", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.amazonaws.regions.Region getCurrentRegion()
func (jbobject *RegionsRegions) GetCurrentRegion() *RegionsRegion {
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/regions/Regions", "getCurrentRegion", "com/amazonaws/regions/Region")
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
	unique_x := &RegionsRegion{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) GovCloud() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "GovCloud", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldGovCloud(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "GovCloud", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) US_EAST_1() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "US_EAST_1", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldUS_EAST_1(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "US_EAST_1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) US_WEST_1() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "US_WEST_1", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldUS_WEST_1(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "US_WEST_1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) US_WEST_2() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "US_WEST_2", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldUS_WEST_2(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "US_WEST_2", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) EU_WEST_1() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "EU_WEST_1", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldEU_WEST_1(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "EU_WEST_1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) EU_CENTRAL_1() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "EU_CENTRAL_1", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldEU_CENTRAL_1(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "EU_CENTRAL_1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) AP_SOUTHEAST_1() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "AP_SOUTHEAST_1", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldAP_SOUTHEAST_1(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "AP_SOUTHEAST_1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) AP_SOUTHEAST_2() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "AP_SOUTHEAST_2", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldAP_SOUTHEAST_2(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "AP_SOUTHEAST_2", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) AP_NORTHEAST_1() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "AP_NORTHEAST_1", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldAP_NORTHEAST_1(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "AP_NORTHEAST_1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) AP_NORTHEAST_2() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "AP_NORTHEAST_2", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldAP_NORTHEAST_2(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "AP_NORTHEAST_2", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) SA_EAST_1() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "SA_EAST_1", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldSA_EAST_1(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "SA_EAST_1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) CN_NORTH_1() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "CN_NORTH_1", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldCN_NORTH_1(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "CN_NORTH_1", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *RegionsRegions) DEFAULT_REGION() *RegionsRegions {
	jret, err := javabind.GetEnv().GetStaticField("com/amazonaws/regions/Regions", "DEFAULT_REGION", "com/amazonaws/regions/Regions")
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
	unique_x := &RegionsRegions{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *RegionsRegions) SetFieldDEFAULT_REGION(val RegionsRegionsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/amazonaws/regions/Regions", "DEFAULT_REGION", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


