package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelInstanceMonitoringInterface interface {
	JavaLangObjectInterface

	// public void setInstanceId(java.lang.String)
	SetInstanceId(a string) 

	// public java.lang.String getInstanceId()
	GetInstanceId() string

	// public com.amazonaws.services.ec2.model.InstanceMonitoring withInstanceId(java.lang.String)
	WithInstanceId(a string) *ServicesEc2ModelInstanceMonitoring

	// public void setMonitoring(com.amazonaws.services.ec2.model.Monitoring)
	SetMonitoring(a ServicesEc2ModelMonitoringInterface) 

	// public com.amazonaws.services.ec2.model.Monitoring getMonitoring()
	GetMonitoring() *ServicesEc2ModelMonitoring

	// public com.amazonaws.services.ec2.model.InstanceMonitoring withMonitoring(com.amazonaws.services.ec2.model.Monitoring)
	WithMonitoring(a ServicesEc2ModelMonitoringInterface) *ServicesEc2ModelInstanceMonitoring

	// public com.amazonaws.services.ec2.model.InstanceMonitoring clone()
	Clone() *ServicesEc2ModelInstanceMonitoring
}

type ServicesEc2ModelInstanceMonitoring struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.InstanceMonitoring()
func NewServicesEc2ModelInstanceMonitoring() (*ServicesEc2ModelInstanceMonitoring) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/InstanceMonitoring")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelInstanceMonitoring{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceMonitoring) SetInstanceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getInstanceId()
func (jbobject *ServicesEc2ModelInstanceMonitoring) GetInstanceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.InstanceMonitoring withInstanceId(java.lang.String)
func (jbobject *ServicesEc2ModelInstanceMonitoring) WithInstanceId(a string) *ServicesEc2ModelInstanceMonitoring {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceId", "com/amazonaws/services/ec2/model/InstanceMonitoring", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelInstanceMonitoring{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMonitoring(com.amazonaws.services.ec2.model.Monitoring)
func (jbobject *ServicesEc2ModelInstanceMonitoring) SetMonitoring(a ServicesEc2ModelMonitoringInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMonitoring", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Monitoring"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Monitoring getMonitoring()
func (jbobject *ServicesEc2ModelInstanceMonitoring) GetMonitoring() *ServicesEc2ModelMonitoring {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMonitoring", "com/amazonaws/services/ec2/model/Monitoring")
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
	unique_x := &ServicesEc2ModelMonitoring{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.InstanceMonitoring withMonitoring(com.amazonaws.services.ec2.model.Monitoring)
func (jbobject *ServicesEc2ModelInstanceMonitoring) WithMonitoring(a ServicesEc2ModelMonitoringInterface) *ServicesEc2ModelInstanceMonitoring {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMonitoring", "com/amazonaws/services/ec2/model/InstanceMonitoring", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Monitoring"))
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
	unique_x := &ServicesEc2ModelInstanceMonitoring{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelInstanceMonitoring) ToString() string {
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
func (jbobject *ServicesEc2ModelInstanceMonitoring) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelInstanceMonitoring) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.InstanceMonitoring clone()
func (jbobject *ServicesEc2ModelInstanceMonitoring) Clone() *ServicesEc2ModelInstanceMonitoring {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/InstanceMonitoring")
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
	unique_x := &ServicesEc2ModelInstanceMonitoring{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelInstanceMonitoring) Clone2() (*JavaLangObject, error) {
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


