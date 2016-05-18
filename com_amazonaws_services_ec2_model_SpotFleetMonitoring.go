package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelSpotFleetMonitoringInterface interface {
	JavaLangObjectInterface

	// public void setEnabled(java.lang.Boolean)
	SetEnabled(a bool) 

	// public java.lang.Boolean getEnabled()
	GetEnabled() bool

	// public com.amazonaws.services.ec2.model.SpotFleetMonitoring withEnabled(java.lang.Boolean)
	WithEnabled(a bool) *ServicesEc2ModelSpotFleetMonitoring

	// public java.lang.Boolean isEnabled()
	IsEnabled() bool

	// public com.amazonaws.services.ec2.model.SpotFleetMonitoring clone()
	Clone() *ServicesEc2ModelSpotFleetMonitoring
}

type ServicesEc2ModelSpotFleetMonitoring struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.SpotFleetMonitoring()
func NewServicesEc2ModelSpotFleetMonitoring() (*ServicesEc2ModelSpotFleetMonitoring) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/SpotFleetMonitoring")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelSpotFleetMonitoring{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setEnabled(java.lang.Boolean)
func (jbobject *ServicesEc2ModelSpotFleetMonitoring) SetEnabled(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnabled", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEnabled()
func (jbobject *ServicesEc2ModelSpotFleetMonitoring) GetEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnabled", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.SpotFleetMonitoring withEnabled(java.lang.Boolean)
func (jbobject *ServicesEc2ModelSpotFleetMonitoring) WithEnabled(a bool) *ServicesEc2ModelSpotFleetMonitoring {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEnabled", "com/amazonaws/services/ec2/model/SpotFleetMonitoring", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelSpotFleetMonitoring{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEnabled()
func (jbobject *ServicesEc2ModelSpotFleetMonitoring) IsEnabled() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnabled", "java/lang/Boolean")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelSpotFleetMonitoring) ToString() string {
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

// public boolean equals(java.lang.Object)
func (jbobject *ServicesEc2ModelSpotFleetMonitoring) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelSpotFleetMonitoring) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.SpotFleetMonitoring clone()
func (jbobject *ServicesEc2ModelSpotFleetMonitoring) Clone() *ServicesEc2ModelSpotFleetMonitoring {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/SpotFleetMonitoring")
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
	unique_x := &ServicesEc2ModelSpotFleetMonitoring{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelSpotFleetMonitoring) Clone2() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}


