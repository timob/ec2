package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelDescribeInstanceAttributeResultInterface interface {
	JavaLangObjectInterface

	// public void setInstanceAttribute(com.amazonaws.services.ec2.model.InstanceAttribute)
	SetInstanceAttribute(a ServicesEc2ModelInstanceAttributeInterface) 

	// public com.amazonaws.services.ec2.model.InstanceAttribute getInstanceAttribute()
	GetInstanceAttribute() *ServicesEc2ModelInstanceAttribute

	// public com.amazonaws.services.ec2.model.DescribeInstanceAttributeResult withInstanceAttribute(com.amazonaws.services.ec2.model.InstanceAttribute)
	WithInstanceAttribute(a ServicesEc2ModelInstanceAttributeInterface) *ServicesEc2ModelDescribeInstanceAttributeResult

	// public com.amazonaws.services.ec2.model.DescribeInstanceAttributeResult clone()
	Clone() *ServicesEc2ModelDescribeInstanceAttributeResult
}

type ServicesEc2ModelDescribeInstanceAttributeResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.DescribeInstanceAttributeResult()
func NewServicesEc2ModelDescribeInstanceAttributeResult() (*ServicesEc2ModelDescribeInstanceAttributeResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/DescribeInstanceAttributeResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelDescribeInstanceAttributeResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setInstanceAttribute(com.amazonaws.services.ec2.model.InstanceAttribute)
func (jbobject *ServicesEc2ModelDescribeInstanceAttributeResult) SetInstanceAttribute(a ServicesEc2ModelInstanceAttributeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInstanceAttribute", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceAttribute"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.InstanceAttribute getInstanceAttribute()
func (jbobject *ServicesEc2ModelDescribeInstanceAttributeResult) GetInstanceAttribute() *ServicesEc2ModelInstanceAttribute {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInstanceAttribute", "com/amazonaws/services/ec2/model/InstanceAttribute")
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
	unique_x := &ServicesEc2ModelInstanceAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.DescribeInstanceAttributeResult withInstanceAttribute(com.amazonaws.services.ec2.model.InstanceAttribute)
func (jbobject *ServicesEc2ModelDescribeInstanceAttributeResult) WithInstanceAttribute(a ServicesEc2ModelInstanceAttributeInterface) *ServicesEc2ModelDescribeInstanceAttributeResult {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withInstanceAttribute", "com/amazonaws/services/ec2/model/DescribeInstanceAttributeResult", conv_a.Value().Cast("com/amazonaws/services/ec2/model/InstanceAttribute"))
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
	unique_x := &ServicesEc2ModelDescribeInstanceAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelDescribeInstanceAttributeResult) ToString() string {
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
func (jbobject *ServicesEc2ModelDescribeInstanceAttributeResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelDescribeInstanceAttributeResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.DescribeInstanceAttributeResult clone()
func (jbobject *ServicesEc2ModelDescribeInstanceAttributeResult) Clone() *ServicesEc2ModelDescribeInstanceAttributeResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/DescribeInstanceAttributeResult")
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
	unique_x := &ServicesEc2ModelDescribeInstanceAttributeResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelDescribeInstanceAttributeResult) Clone2() (*JavaLangObject, error) {
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


