package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVolumeAttributeResultInterface interface {
	JavaLangObjectInterface

	// public void setVolumeId(java.lang.String)
	SetVolumeId(a string) 

	// public java.lang.String getVolumeId()
	GetVolumeId() string

	// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult withVolumeId(java.lang.String)
	WithVolumeId(a string) *ServicesEc2ModelDescribeVolumeAttributeResult

	// public void setAutoEnableIO(java.lang.Boolean)
	SetAutoEnableIO(a bool) 

	// public java.lang.Boolean getAutoEnableIO()
	GetAutoEnableIO() bool

	// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult withAutoEnableIO(java.lang.Boolean)
	WithAutoEnableIO(a bool) *ServicesEc2ModelDescribeVolumeAttributeResult

	// public java.lang.Boolean isAutoEnableIO()
	IsAutoEnableIO() bool

	// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
	GetProductCodes() []*ServicesEc2ModelProductCode

	// public void setProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	SetProductCodes(a []*ServicesEc2ModelProductCode) 

	// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
	WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelDescribeVolumeAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
	WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelDescribeVolumeAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult clone()
	Clone() *ServicesEc2ModelDescribeVolumeAttributeResult
}

type ServicesEc2ModelDescribeVolumeAttributeResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult()
func NewServicesEc2ModelDescribeVolumeAttributeResult() (*ServicesEc2ModelDescribeVolumeAttributeResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVolumeAttributeResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVolumeAttributeResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVolumeId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) SetVolumeId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVolumeId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVolumeId()
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) GetVolumeId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVolumeId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult withVolumeId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) WithVolumeId(a string) *ServicesEc2ModelDescribeVolumeAttributeResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVolumeId", "com/amazonaws/services/ec2/model/DescribeVolumeAttributeResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVolumeAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setAutoEnableIO(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) SetAutoEnableIO(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAutoEnableIO", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getAutoEnableIO()
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) GetAutoEnableIO() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAutoEnableIO", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult withAutoEnableIO(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) WithAutoEnableIO(a bool) *ServicesEc2ModelDescribeVolumeAttributeResult {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAutoEnableIO", "com/amazonaws/services/ec2/model/DescribeVolumeAttributeResult", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelDescribeVolumeAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isAutoEnableIO()
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) IsAutoEnableIO() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isAutoEnableIO", "java/lang/Boolean")
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

// public java.util.List<com.amazonaws.services.ec2.model.ProductCode> getProductCodes()
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) GetProductCodes() []*ServicesEc2ModelProductCode {
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
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) SetProductCodes(a []*ServicesEc2ModelProductCode)  {
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

// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult withProductCodes(com.amazonaws.services.ec2.model.ProductCode...)
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) WithProductCodes(a ...*ServicesEc2ModelProductCode) *ServicesEc2ModelDescribeVolumeAttributeResult {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/ProductCode")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/DescribeVolumeAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ProductCode"))
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
	unique_x := &ServicesEc2ModelDescribeVolumeAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult withProductCodes(java.util.Collection<com.amazonaws.services.ec2.model.ProductCode>)
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) WithProductCodes2(a []*ServicesEc2ModelProductCode) *ServicesEc2ModelDescribeVolumeAttributeResult {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withProductCodes", "com/amazonaws/services/ec2/model/DescribeVolumeAttributeResult", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelDescribeVolumeAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVolumeAttributeResult clone()
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) Clone() *ServicesEc2ModelDescribeVolumeAttributeResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVolumeAttributeResult")
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
	unique_x := &ServicesEc2ModelDescribeVolumeAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVolumeAttributeResult) Clone2() (*JavaLangObject, error) {
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


