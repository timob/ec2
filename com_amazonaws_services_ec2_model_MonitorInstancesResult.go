package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelMonitorInstancesResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.InstanceMonitoring> getInstanceMonitorings()
	GetInstanceMonitorings() []*ServicesEc2ModelInstanceMonitoring

	// public void setInstanceMonitorings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceMonitoring>)
	SetInstanceMonitorings(a []*ServicesEc2ModelInstanceMonitoring) 

	// public com.amazonaws.services.ec2.model.MonitorInstancesResult withInstanceMonitorings(com.amazonaws.services.ec2.model.InstanceMonitoring...)
	WithInstanceMonitorings(a ...*ServicesEc2ModelInstanceMonitoring) *ServicesEc2ModelMonitorInstancesResult

	// public com.amazonaws.services.ec2.model.MonitorInstancesResult withInstanceMonitorings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceMonitoring>)
	WithInstanceMonitorings2(a []*ServicesEc2ModelInstanceMonitoring) *ServicesEc2ModelMonitorInstancesResult

	// public com.amazonaws.services.ec2.model.MonitorInstancesResult clone()
	Clone() *ServicesEc2ModelMonitorInstancesResult
}

type ServicesEc2ModelMonitorInstancesResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.MonitorInstancesResult()
func NewServicesEc2ModelMonitorInstancesResult() (*ServicesEc2ModelMonitorInstancesResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/MonitorInstancesResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelMonitorInstancesResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.InstanceMonitoring> getInstanceMonitorings()
func (jbobject *ServicesEc2ModelMonitorInstancesResult) GetInstanceMonitorings() []*ServicesEc2ModelInstanceMonitoring {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceMonitorings", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelInstanceMonitoring)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setInstanceMonitorings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceMonitoring>)
func (jbobject *ServicesEc2ModelMonitorInstancesResult) SetInstanceMonitorings(a []*ServicesEc2ModelInstanceMonitoring)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceMonitorings", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.MonitorInstancesResult withInstanceMonitorings(com.amazonaws.services.ec2.model.InstanceMonitoring...)
func (jbobject *ServicesEc2ModelMonitorInstancesResult) WithInstanceMonitorings(a ...*ServicesEc2ModelInstanceMonitoring) *ServicesEc2ModelMonitorInstancesResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/InstanceMonitoring")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceMonitorings", "com/amazonaws/services/ec2/model/MonitorInstancesResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceMonitoring"))
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
	unique_x := &ServicesEc2ModelMonitorInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.MonitorInstancesResult withInstanceMonitorings(java.util.Collection<com.amazonaws.services.ec2.model.InstanceMonitoring>)
func (jbobject *ServicesEc2ModelMonitorInstancesResult) WithInstanceMonitorings2(a []*ServicesEc2ModelInstanceMonitoring) *ServicesEc2ModelMonitorInstancesResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceMonitorings", "com/amazonaws/services/ec2/model/MonitorInstancesResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelMonitorInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelMonitorInstancesResult) ToString() string {
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
func (jbobject *ServicesEc2ModelMonitorInstancesResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelMonitorInstancesResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.MonitorInstancesResult clone()
func (jbobject *ServicesEc2ModelMonitorInstancesResult) Clone() *ServicesEc2ModelMonitorInstancesResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/MonitorInstancesResult")
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
	unique_x := &ServicesEc2ModelMonitorInstancesResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelMonitorInstancesResult) Clone2() (*JavaLangObject, error) {
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


