package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCreateSnapshotResultInterface interface {
	JavaLangObjectInterface

	// public void setSnapshot(com.amazonaws.services.ec2.model.Snapshot)
	SetSnapshot(a ServicesEc2ModelSnapshotInterface) 

	// public com.amazonaws.services.ec2.model.Snapshot getSnapshot()
	GetSnapshot() *ServicesEc2ModelSnapshot

	// public com.amazonaws.services.ec2.model.CreateSnapshotResult withSnapshot(com.amazonaws.services.ec2.model.Snapshot)
	WithSnapshot(a ServicesEc2ModelSnapshotInterface) *ServicesEc2ModelCreateSnapshotResult

	// public com.amazonaws.services.ec2.model.CreateSnapshotResult clone()
	Clone() *ServicesEc2ModelCreateSnapshotResult
}

type ServicesEc2ModelCreateSnapshotResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CreateSnapshotResult()
func NewServicesEc2ModelCreateSnapshotResult() (*ServicesEc2ModelCreateSnapshotResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CreateSnapshotResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCreateSnapshotResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSnapshot(com.amazonaws.services.ec2.model.Snapshot)
func (jbobject *ServicesEc2ModelCreateSnapshotResult) SetSnapshot(a ServicesEc2ModelSnapshotInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshot", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/Snapshot"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.Snapshot getSnapshot()
func (jbobject *ServicesEc2ModelCreateSnapshotResult) GetSnapshot() *ServicesEc2ModelSnapshot {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshot", "com/amazonaws/services/ec2/model/Snapshot")
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
	unique_x := &ServicesEc2ModelSnapshot{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.CreateSnapshotResult withSnapshot(com.amazonaws.services.ec2.model.Snapshot)
func (jbobject *ServicesEc2ModelCreateSnapshotResult) WithSnapshot(a ServicesEc2ModelSnapshotInterface) *ServicesEc2ModelCreateSnapshotResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshot", "com/amazonaws/services/ec2/model/CreateSnapshotResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Snapshot"))
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
	unique_x := &ServicesEc2ModelCreateSnapshotResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCreateSnapshotResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCreateSnapshotResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCreateSnapshotResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CreateSnapshotResult clone()
func (jbobject *ServicesEc2ModelCreateSnapshotResult) Clone() *ServicesEc2ModelCreateSnapshotResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CreateSnapshotResult")
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
	unique_x := &ServicesEc2ModelCreateSnapshotResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCreateSnapshotResult) Clone2() (*JavaLangObject, error) {
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


