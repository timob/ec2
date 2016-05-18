package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAccountAttributeInterface interface {
	JavaLangObjectInterface

	// public void setAttributeName(java.lang.String)
	SetAttributeName(a string) 

	// public java.lang.String getAttributeName()
	GetAttributeName() string

	// public com.amazonaws.services.ec2.model.AccountAttribute withAttributeName(java.lang.String)
	WithAttributeName(a string) *ServicesEc2ModelAccountAttribute

	// public java.util.List<com.amazonaws.services.ec2.model.AccountAttributeValue> getAttributeValues()
	GetAttributeValues() []*ServicesEc2ModelAccountAttributeValue

	// public void setAttributeValues(java.util.Collection<com.amazonaws.services.ec2.model.AccountAttributeValue>)
	SetAttributeValues(a []*ServicesEc2ModelAccountAttributeValue) 

	// public com.amazonaws.services.ec2.model.AccountAttribute withAttributeValues(com.amazonaws.services.ec2.model.AccountAttributeValue...)
	WithAttributeValues(a ...*ServicesEc2ModelAccountAttributeValue) *ServicesEc2ModelAccountAttribute

	// public com.amazonaws.services.ec2.model.AccountAttribute withAttributeValues(java.util.Collection<com.amazonaws.services.ec2.model.AccountAttributeValue>)
	WithAttributeValues2(a []*ServicesEc2ModelAccountAttributeValue) *ServicesEc2ModelAccountAttribute

	// public com.amazonaws.services.ec2.model.AccountAttribute clone()
	Clone() *ServicesEc2ModelAccountAttribute
}

type ServicesEc2ModelAccountAttribute struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AccountAttribute()
func NewServicesEc2ModelAccountAttribute() (*ServicesEc2ModelAccountAttribute) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AccountAttribute")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAccountAttribute{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAttributeName(java.lang.String)
func (jbobject *ServicesEc2ModelAccountAttribute) SetAttributeName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttributeName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAttributeName()
func (jbobject *ServicesEc2ModelAccountAttribute) GetAttributeName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttributeName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AccountAttribute withAttributeName(java.lang.String)
func (jbobject *ServicesEc2ModelAccountAttribute) WithAttributeName(a string) *ServicesEc2ModelAccountAttribute {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttributeName", "com/amazonaws/services/ec2/model/AccountAttribute", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAccountAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.amazonaws.services.ec2.model.AccountAttributeValue> getAttributeValues()
func (jbobject *ServicesEc2ModelAccountAttribute) GetAttributeValues() []*ServicesEc2ModelAccountAttributeValue {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAttributeValues", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*ServicesEc2ModelAccountAttributeValue)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAttributeValues(java.util.Collection<com.amazonaws.services.ec2.model.AccountAttributeValue>)
func (jbobject *ServicesEc2ModelAccountAttribute) SetAttributeValues(a []*ServicesEc2ModelAccountAttributeValue)  {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAttributeValues", javabind.Void, conv_a.Value().Cast("java/util/Collection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.AccountAttribute withAttributeValues(com.amazonaws.services.ec2.model.AccountAttributeValue...)
func (jbobject *ServicesEc2ModelAccountAttribute) WithAttributeValues(a ...*ServicesEc2ModelAccountAttributeValue) *ServicesEc2ModelAccountAttribute {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "com/amazonaws/services/ec2/model/AccountAttributeValue")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttributeValues", "com/amazonaws/services/ec2/model/AccountAttribute", conv_a.Value().Cast("com/amazonaws/services/ec2/model/AccountAttributeValue"))
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
	unique_x := &ServicesEc2ModelAccountAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public com.amazonaws.services.ec2.model.AccountAttribute withAttributeValues(java.util.Collection<com.amazonaws.services.ec2.model.AccountAttributeValue>)
func (jbobject *ServicesEc2ModelAccountAttribute) WithAttributeValues2(a []*ServicesEc2ModelAccountAttributeValue) *ServicesEc2ModelAccountAttribute {
	conv_a := javabind.NewGoToJavaCollection(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAttributeValues", "com/amazonaws/services/ec2/model/AccountAttribute", conv_a.Value().Cast("java/util/Collection"))
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
	unique_x := &ServicesEc2ModelAccountAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAccountAttribute) ToString() string {
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
func (jbobject *ServicesEc2ModelAccountAttribute) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAccountAttribute) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AccountAttribute clone()
func (jbobject *ServicesEc2ModelAccountAttribute) Clone() *ServicesEc2ModelAccountAttribute {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AccountAttribute")
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
	unique_x := &ServicesEc2ModelAccountAttribute{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAccountAttribute) Clone2() (*JavaLangObject, error) {
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


