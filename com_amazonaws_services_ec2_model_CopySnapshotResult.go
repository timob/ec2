package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelCopySnapshotResultInterface interface {
	JavaLangObjectInterface

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.CopySnapshotResult withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelCopySnapshotResult

	// public com.amazonaws.services.ec2.model.CopySnapshotResult clone()
	Clone() *ServicesEc2ModelCopySnapshotResult
}

type ServicesEc2ModelCopySnapshotResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.CopySnapshotResult()
func NewServicesEc2ModelCopySnapshotResult() (*ServicesEc2ModelCopySnapshotResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/CopySnapshotResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelCopySnapshotResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotResult) SetSnapshotId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshotId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getSnapshotId()
func (jbobject *ServicesEc2ModelCopySnapshotResult) GetSnapshotId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshotId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.CopySnapshotResult withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelCopySnapshotResult) WithSnapshotId(a string) *ServicesEc2ModelCopySnapshotResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/CopySnapshotResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelCopySnapshotResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelCopySnapshotResult) ToString() string {
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
func (jbobject *ServicesEc2ModelCopySnapshotResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelCopySnapshotResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.CopySnapshotResult clone()
func (jbobject *ServicesEc2ModelCopySnapshotResult) Clone() *ServicesEc2ModelCopySnapshotResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/CopySnapshotResult")
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
	unique_x := &ServicesEc2ModelCopySnapshotResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelCopySnapshotResult) Clone2() (*JavaLangObject, error) {
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


