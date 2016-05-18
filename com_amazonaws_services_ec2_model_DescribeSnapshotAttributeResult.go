package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeSnapshotAttributeResultInterface interface {
	JavaLangObjectInterface

	// public void setSnapshotId(java.lang.String)
	SetSnapshotId(a string) 

	// public java.lang.String getSnapshotId()
	GetSnapshotId() string

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withSnapshotId(java.lang.String)
	WithSnapshotId(a string) *ServicesEc2ModelDescribeSnapshotAttributeResult

	// public java.util.List<com.amazonaws.services.ec2.model.CreateVolumePermission> getCreateVolumePermissions()
	GetCreateVolumePermissions() []*ServicesEc2ModelCreateVolumePermission

	// public void setCreateVolumePermissions(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
	SetCreateVolumePermissions(a []*ServicesEc2ModelCreateVolumePermission) 

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withCreateVolumePermissions(com.amazonaws.services.ec2.model.CreateVolumePermission...)
	WithCreateVolumePermissions(a ...*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelDescribeSnapshotAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withCreateVolumePermissions(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
	WithCreateVolumePermissions2(a []*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelDescribeSnapshotAttributeResult

	// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
	GetProductCodes() []*ServicesEc2ModelProductCode

	// public void setProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	SetProductCodes(a []*ServicesEc2ModelProductCode) 

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
	WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelDescribeSnapshotAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelDescribeSnapshotAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult clone()
	Clone() *ServicesEc2ModelDescribeSnapshotAttributeResult
}

type ServicesEc2ModelDescribeSnapshotAttributeResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult()
func NewServicesEc2ModelDescribeSnapshotAttributeResult() (*ServicesEc2ModelDescribeSnapshotAttributeResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeSnapshotAttributeResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeSnapshotAttributeResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) SetSnapshotId(a string)  {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) GetSnapshotId() string {
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

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withSnapshotId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) WithSnapshotId(a string) *ServicesEc2ModelDescribeSnapshotAttributeResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withSnapshotId", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.CreateVolumePermission> getCreateVolumePermissions()
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) GetCreateVolumePermissions() []*ServicesEc2ModelCreateVolumePermission {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCreateVolumePermissions", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelCreateVolumePermission)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setCreateVolumePermissions(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) SetCreateVolumePermissions(a []*ServicesEc2ModelCreateVolumePermission)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCreateVolumePermissions", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withCreateVolumePermissions(com.amazonaws.services.ec2.model.CreateVolumePermission...)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) WithCreateVolumePermissions(a ...*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelDescribeSnapshotAttributeResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/CreateVolumePermission")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateVolumePermissions", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/CreateVolumePermission"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withCreateVolumePermissions(java.util.Collection<com.amazonaws.services.ec2.model.CreateVolumePermission>)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) WithCreateVolumePermissions2(a []*ServicesEc2ModelCreateVolumePermission) *ServicesEc2ModelDescribeSnapshotAttributeResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCreateVolumePermissions", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) GetProductCodes() []*ServicesEc2ModelProductCode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getProductCodes", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelProductCode)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) SetProductCodes(a []*ServicesEc2ModelProductCode)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setProductCodes", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelDescribeSnapshotAttributeResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ProductCode")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ProductCode"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelDescribeSnapshotAttributeResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeSnapshotAttributeResult clone()
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) Clone() *ServicesEc2ModelDescribeSnapshotAttributeResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeSnapshotAttributeResult")
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
	unique_x := &ServicesEc2ModelDescribeSnapshotAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeSnapshotAttributeResult) Clone2() (*JavaLangObject, error) {
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


