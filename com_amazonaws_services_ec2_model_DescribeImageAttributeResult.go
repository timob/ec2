package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeImageAttributeResultInterface interface {
	JavaLangObjectInterface

	// public void setImageAttribute(com.amazonaws.services.ec2.model.ImageAttribute)
	SetImageAttribute(a ServicesEc2ModelImageAttributeInterface) 

	// public com.amazonaws.services.ec2.model.ImageAttribute getImageAttribute()
	GetImageAttribute() *ServicesEc2ModelImageAttribute

	// public com.amazonaws.services.ec2.model.DescribeImageAttributeResult withImageAttribute(com.amazonaws.services.ec2.model.ImageAttribute)
	WithImageAttribute(a ServicesEc2ModelImageAttributeInterface) *ServicesEc2ModelDescribeImageAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeImageAttributeResult clone()
	Clone() *ServicesEc2ModelDescribeImageAttributeResult
}

type ServicesEc2ModelDescribeImageAttributeResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeImageAttributeResult()
func NewServicesEc2ModelDescribeImageAttributeResult() (*ServicesEc2ModelDescribeImageAttributeResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeImageAttributeResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeImageAttributeResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setImageAttribute(com.amazonaws.services.ec2.model.ImageAttribute)
func (jbobject *ServicesEc2ModelDescribeImageAttributeResult) SetImageAttribute(a ServicesEc2ModelImageAttributeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImageAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImageAttribute"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.ImageAttribute getImageAttribute()
func (jbobject *ServicesEc2ModelDescribeImageAttributeResult) GetImageAttribute() *ServicesEc2ModelImageAttribute {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImageAttribute", "com/amazonaws/services/ec2/model/ImageAttribute")
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
	unique_x := &ServicesEc2ModelImageAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeImageAttributeResult withImageAttribute(com.amazonaws.services.ec2.model.ImageAttribute)
func (jbobject *ServicesEc2ModelDescribeImageAttributeResult) WithImageAttribute(a ServicesEc2ModelImageAttributeInterface) *ServicesEc2ModelDescribeImageAttributeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withImageAttribute", "com/amazonaws/services/ec2/model/DescribeImageAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/ImageAttribute"))
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
	unique_x := &ServicesEc2ModelDescribeImageAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeImageAttributeResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeImageAttributeResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeImageAttributeResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeImageAttributeResult clone()
func (jbobject *ServicesEc2ModelDescribeImageAttributeResult) Clone() *ServicesEc2ModelDescribeImageAttributeResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeImageAttributeResult")
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
	unique_x := &ServicesEc2ModelDescribeImageAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeImageAttributeResult) Clone2() (*JavaLangObject, error) {
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


