package ec2

import "github.com/timob/javabind"

type RegionsRegionImplInterface interface {

	// public abstract java.lang.String getName()
	GetName() string

	// public abstract java.lang.String getDomain()
	GetDomain() string

	// public abstract boolean isServiceSupported(java.lang.String)
	IsServiceSupported(a string) bool

	// public abstract java.lang.String getServiceEndpoint(java.lang.String)
	GetServiceEndpoint(a string) string

	// public abstract boolean hasHttpEndpoint(java.lang.String)
	HasHttpEndpoint(a string) bool

	// public abstract boolean hasHttpsEndpoint(java.lang.String)
	HasHttpsEndpoint(a string) bool

	// public abstract java.util.Collection<java.lang.String> getAvailableEndpoints()
	GetAvailableEndpoints() []string
}

type RegionsRegionImpl struct {
	JavaLangObject
}

// public abstract java.lang.String getName()
func (jbobject *RegionsRegionImpl) GetName() string {
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

// public abstract java.lang.String getDomain()
func (jbobject *RegionsRegionImpl) GetDomain() string {
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

// public abstract boolean isServiceSupported(java.lang.String)
func (jbobject *RegionsRegionImpl) IsServiceSupported(a string) bool {
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

// public abstract java.lang.String getServiceEndpoint(java.lang.String)
func (jbobject *RegionsRegionImpl) GetServiceEndpoint(a string) string {
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

// public abstract boolean hasHttpEndpoint(java.lang.String)
func (jbobject *RegionsRegionImpl) HasHttpEndpoint(a string) bool {
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

// public abstract boolean hasHttpsEndpoint(java.lang.String)
func (jbobject *RegionsRegionImpl) HasHttpsEndpoint(a string) bool {
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

// public abstract java.util.Collection<java.lang.String> getAvailableEndpoints()
func (jbobject *RegionsRegionImpl) GetAvailableEndpoints() []string {
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


