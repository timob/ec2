package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeVpcAttributeResultInterface interface {
	JavaLangObjectInterface

	// public void setVpcId(java.lang.String)
	SetVpcId(a string) 

	// public java.lang.String getVpcId()
	GetVpcId() string

	// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult withVpcId(java.lang.String)
	WithVpcId(a string) *ServicesEc2ModelDescribeVpcAttributeResult

	// public void setEnableDnsSupport(java.lang.Boolean)
	SetEnableDnsSupport(a bool) 

	// public java.lang.Boolean getEnableDnsSupport()
	GetEnableDnsSupport() bool

	// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult withEnableDnsSupport(java.lang.Boolean)
	WithEnableDnsSupport(a bool) *ServicesEc2ModelDescribeVpcAttributeResult

	// public java.lang.Boolean isEnableDnsSupport()
	IsEnableDnsSupport() bool

	// public void setEnableDnsHostnames(java.lang.Boolean)
	SetEnableDnsHostnames(a bool) 

	// public java.lang.Boolean getEnableDnsHostnames()
	GetEnableDnsHostnames() bool

	// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult withEnableDnsHostnames(java.lang.Boolean)
	WithEnableDnsHostnames(a bool) *ServicesEc2ModelDescribeVpcAttributeResult

	// public java.lang.Boolean isEnableDnsHostnames()
	IsEnableDnsHostnames() bool

	// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult clone()
	Clone() *ServicesEc2ModelDescribeVpcAttributeResult
}

type ServicesEc2ModelDescribeVpcAttributeResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult()
func NewServicesEc2ModelDescribeVpcAttributeResult() (*ServicesEc2ModelDescribeVpcAttributeResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeVpcAttributeResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeVpcAttributeResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) SetVpcId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVpcId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getVpcId()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) GetVpcId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVpcId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult withVpcId(java.lang.String)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) WithVpcId(a string) *ServicesEc2ModelDescribeVpcAttributeResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withVpcId", "com/amazonaws/services/ec2/model/DescribeVpcAttributeResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelDescribeVpcAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public void setEnableDnsSupport(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) SetEnableDnsSupport(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnableDnsSupport", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEnableDnsSupport()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) GetEnableDnsSupport() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnableDnsSupport", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult withEnableDnsSupport(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) WithEnableDnsSupport(a bool) *ServicesEc2ModelDescribeVpcAttributeResult {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEnableDnsSupport", "com/amazonaws/services/ec2/model/DescribeVpcAttributeResult", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelDescribeVpcAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEnableDnsSupport()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) IsEnableDnsSupport() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnableDnsSupport", "java/lang/Boolean")
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

// public void setEnableDnsHostnames(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) SetEnableDnsHostnames(a bool)  {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEnableDnsHostnames", javabind.Void, conv_a.Value().Cast("java/lang/Boolean"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Boolean getEnableDnsHostnames()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) GetEnableDnsHostnames() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEnableDnsHostnames", "java/lang/Boolean")
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

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult withEnableDnsHostnames(java.lang.Boolean)
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) WithEnableDnsHostnames(a bool) *ServicesEc2ModelDescribeVpcAttributeResult {
	conv_a := javabind.NewGoToJavaBoolean()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withEnableDnsHostnames", "com/amazonaws/services/ec2/model/DescribeVpcAttributeResult", conv_a.Value().Cast("java/lang/Boolean"))
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
	unique_x := &ServicesEc2ModelDescribeVpcAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Boolean isEnableDnsHostnames()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) IsEnableDnsHostnames() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEnableDnsHostnames", "java/lang/Boolean")
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
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeVpcAttributeResult clone()
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) Clone() *ServicesEc2ModelDescribeVpcAttributeResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeVpcAttributeResult")
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
	unique_x := &ServicesEc2ModelDescribeVpcAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeVpcAttributeResult) Clone2() (*JavaLangObject, error) {
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


