package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSnapshotsResultInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.amazonaws.services.ec2.model.Snapshot> getSnapshots()
	GetSnapshots() []*ServicesEc2ModelSnapshot

	// public void setSnapshots(java.util.Collection<com.amazonaws.services.ec2.model.Snapshot>)
	SetSnapshots(a []*ServicesEc2ModelSnapshot) 

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult withSnapshots(com.amazonaws.services.ec2.model.Snapshot...)
	WithSnapshots(a ...*ServicesEc2ModelSnapshot) *ServicesEc2ModelDescribeSnapshotsResult

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult withSnapshots(java.util.Collection<com.amazonaws.services.ec2.model.Snapshot>)
	WithSnapshots2(a []*ServicesEc2ModelSnapshot) *ServicesEc2ModelDescribeSnapshotsResult

	// public void setNextToken(java.lang.String)
	SetNextToken(a string) 

	// public java.lang.String getNextToken()
	GetNextToken() string

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult withNextToken(java.lang.String)
	WithNextToken(a string) *ServicesEc2ModelDescribeSnapshotsResult

	// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult clone()
	Clone() *ServicesEc2ModelDescribeSnapshotsResult
}

type ServicesEc2ModelDescribeSnapshotsResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult()
func NewServicesEc2ModelDescribeSnapshotsResult() (*ServicesEc2ModelDescribeSnapshotsResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSnapshotsResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSnapshotsResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.amazonaws.services.ec2.model.Snapshot> getSnapshots()
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) GetSnapshots() []*ServicesEc2ModelSnapshot {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSnapshots", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelSnapshot)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setSnapshots(java.util.Collection<com.amazonaws.services.ec2.model.Snapshot>)
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) SetSnapshots(a []*ServicesEc2ModelSnapshot)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSnapshots", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult withSnapshots(com.amazonaws.services.ec2.model.Snapshot...)
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) WithSnapshots(a ...*ServicesEc2ModelSnapshot) *ServicesEc2ModelDescribeSnapshotsResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/Snapshot")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshots", "com/amazonaws/services/ec2/model/DescribeSnapshotsResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/Snapshot"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult withSnapshots(java.util.Collection<com.amazonaws.services.ec2.model.Snapshot>)
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) WithSnapshots2(a []*ServicesEc2ModelSnapshot) *ServicesEc2ModelDescribeSnapshotsResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshots", "com/amazonaws/services/ec2/model/DescribeSnapshotsResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) SetNextToken(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNextToken", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getNextToken()
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) GetNextToken() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNextToken", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult withNextToken(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) WithNextToken(a string) *ServicesEc2ModelDescribeSnapshotsResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withNextToken", "com/amazonaws/services/ec2/model/DescribeSnapshotsResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotsResult clone()
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) Clone() *ServicesEc2ModelDescribeSnapshotsResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSnapshotsResult")
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
	unique_x := &ServicesEc2ModelDescribeSnapshotsResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSnapshotsResult) Clone2() (*JavaLangObject, error) {
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


