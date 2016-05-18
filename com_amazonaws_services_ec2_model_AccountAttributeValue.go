package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAccountAttributeValueInterface interface {
	JavaLangObjectInterface

	// public void setAttributeValue(java.lang.String)
	SetAttributeValue(a string) 

	// public java.lang.String getAttributeValue()
	GetAttributeValue() string

	// public com.amazonaws.services.ec2.model.AccountAttributeValue withAttributeValue(java.lang.String)
	WithAttributeValue(a string) *ServicesEc2ModelAccountAttributeValue

	// public com.amazonaws.services.ec2.model.AccountAttributeValue clone()
	Clone() *ServicesEc2ModelAccountAttributeValue
}

type ServicesEc2ModelAccountAttributeValue struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AccountAttributeValue()
func NewServicesEc2ModelAccountAttributeValue() (*ServicesEc2ModelAccountAttributeValue) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AccountAttributeValue")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAccountAttributeValue{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAttributeValue(java.lang.String)
func (jbobject *ServicesEc2ModelAccountAttributeValue) SetAttributeValue(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttributeValue", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAttributeValue()
func (jbobject *ServicesEc2ModelAccountAttributeValue) GetAttributeValue() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttributeValue", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AccountAttributeValue withAttributeValue(java.lang.String)
func (jbobject *ServicesEc2ModelAccountAttributeValue) WithAttributeValue(a string) *ServicesEc2ModelAccountAttributeValue {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttributeValue", "com/amazonaws/services/ec2/model/AccountAttributeValue", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAccountAttributeValue{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAccountAttributeValue) ToString() string {
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
func (jbobject *ServicesEc2ModelAccountAttributeValue) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAccountAttributeValue) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AccountAttributeValue clone()
func (jbobject *ServicesEc2ModelAccountAttributeValue) Clone() *ServicesEc2ModelAccountAttributeValue {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AccountAttributeValue")
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
	unique_x := &ServicesEc2ModelAccountAttributeValue{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAccountAttributeValue) Clone2() (*JavaLangObject, error) {
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


