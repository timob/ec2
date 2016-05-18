package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelAssociateAddressResultInterface interface {
	JavaLangObjectInterface

	// public void setAssociationId(java.lang.String)
	SetAssociationId(a string) 

	// public java.lang.String getAssociationId()
	GetAssociationId() string

	// public com.amazonaws.services.ec2.model.AssociateAddressResult withAssociationId(java.lang.String)
	WithAssociationId(a string) *ServicesEc2ModelAssociateAddressResult

	// public com.amazonaws.services.ec2.model.AssociateAddressResult clone()
	Clone() *ServicesEc2ModelAssociateAddressResult
}

type ServicesEc2ModelAssociateAddressResult struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.AssociateAddressResult()
func NewServicesEc2ModelAssociateAddressResult() (*ServicesEc2ModelAssociateAddressResult) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/AssociateAddressResult")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelAssociateAddressResult{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelAssociateAddressResult) SetAssociationId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAssociationId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getAssociationId()
func (jbobject *ServicesEc2ModelAssociateAddressResult) GetAssociationId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAssociationId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.AssociateAddressResult withAssociationId(java.lang.String)
func (jbobject *ServicesEc2ModelAssociateAddressResult) WithAssociationId(a string) *ServicesEc2ModelAssociateAddressResult {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withAssociationId", "com/amazonaws/services/ec2/model/AssociateAddressResult", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelAssociateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelAssociateAddressResult) ToString() string {
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
func (jbobject *ServicesEc2ModelAssociateAddressResult) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelAssociateAddressResult) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.AssociateAddressResult clone()
func (jbobject *ServicesEc2ModelAssociateAddressResult) Clone() *ServicesEc2ModelAssociateAddressResult {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/AssociateAddressResult")
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
	unique_x := &ServicesEc2ModelAssociateAddressResult{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelAssociateAddressResult) Clone2() (*JavaLangObject, error) {
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


