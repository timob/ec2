package ec2

import "github.com/timob/javabind"

type RegionsRegionInterface interface {
	JavaLangObjectInterface

	// public static com.amazonaws.regions.Region getRegion(com.amazonaws.regions.Regions)
	GetRegion(a RegionsRegionsInterface) *RegionsRegion

	// public java.lang.String getName()
	GetName() string

	// public java.lang.String getDomain()
	GetDomain() string

	// public java.lang.String getServiceEndpoint(java.lang.String)
	GetServiceEndpoint(a string) string

	// public boolean isServiceSupported(java.lang.String)
	IsServiceSupported(a string) bool

	// public boolean hasHttpsEndpoint(java.lang.String)
	HasHttpsEndpoint(a string) bool

	// public boolean hasHttpEndpoint(java.lang.String)
	HasHttpEndpoint(a string) bool

	// public java.util.Collection<java.lang.String> getAvailableEndpoints()
	GetAvailableEndpoints() []string
}

type RegionsRegion struct {
	JavaLangObject
}

// public com.amazonaws.regions.Region(com.amazonaws.regions.RegionImpl)
func NewRegionsRegion(a RegionsRegionImplInterface) (*RegionsRegion) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/regions/Region", conv_a.Value().Cast("com/amazonaws/regions/RegionImpl"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &RegionsRegion{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static com.amazonaws.regions.Region getRegion(com.amazonaws.regions.Regions)
func (jbobject *RegionsRegion) GetRegion(a RegionsRegionsInterface) *RegionsRegion {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/amazonaws/regions/Region", "getRegion", "com/amazonaws/regions/Region", conv_a.Value().Cast("com/amazonaws/regions/Regions"))
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
	unique_x := &RegionsRegion{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getName()
func (jbobject *RegionsRegion) GetName() string {
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

// public java.lang.String getDomain()
func (jbobject *RegionsRegion) GetDomain() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDomain", "java/lang/String")
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

// public java.lang.String getServiceEndpoint(java.lang.String)
func (jbobject *RegionsRegion) GetServiceEndpoint(a string) string {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getServiceEndpoint", "java/lang/String", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean isServiceSupported(java.lang.String)
func (jbobject *RegionsRegion) IsServiceSupported(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isServiceSupported", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean hasHttpsEndpoint(java.lang.String)
func (jbobject *RegionsRegion) HasHttpsEndpoint(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasHttpsEndpoint", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean hasHttpEndpoint(java.lang.String)
func (jbobject *RegionsRegion) HasHttpEndpoint(a string) bool {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasHttpEndpoint", javabind.Boolean, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public java.util.Collection<java.lang.String> getAvailableEndpoints()
func (jbobject *RegionsRegion) GetAvailableEndpoints() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAvailableEndpoints", "java/util/Collection")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCollection(javabind.NewJavaToGoString())
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean equals(java.lang.Object)
func (jbobject *RegionsRegion) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *RegionsRegion) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *RegionsRegion) ToString() string {
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


