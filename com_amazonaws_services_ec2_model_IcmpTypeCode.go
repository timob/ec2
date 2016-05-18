package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelIcmpTypeCodeInterface interface {
	JavaLangObjectInterface

	// public void setType(java.lang.Integer)
	SetType(a int) 

	// public java.lang.Integer getType()
	GetType() int

	// public com.amazonaws.services.ec2.model.IcmpTypeCode withType(java.lang.Integer)
	WithType(a int) *ServicesEc2ModelIcmpTypeCode

	// public void setCode(java.lang.Integer)
	SetCode(a int) 

	// public java.lang.Integer getCode()
	GetCode() int

	// public com.amazonaws.services.ec2.model.IcmpTypeCode withCode(java.lang.Integer)
	WithCode(a int) *ServicesEc2ModelIcmpTypeCode

	// public com.amazonaws.services.ec2.model.IcmpTypeCode clone()
	Clone() *ServicesEc2ModelIcmpTypeCode
}

type ServicesEc2ModelIcmpTypeCode struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.IcmpTypeCode()
func NewServicesEc2ModelIcmpTypeCode() (*ServicesEc2ModelIcmpTypeCode) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/IcmpTypeCode")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelIcmpTypeCode{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setType(java.lang.Integer)
func (jbobject *ServicesEc2ModelIcmpTypeCode) SetType(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getType()
func (jbobject *ServicesEc2ModelIcmpTypeCode) GetType() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.IcmpTypeCode withType(java.lang.Integer)
func (jbobject *ServicesEc2ModelIcmpTypeCode) WithType(a int) *ServicesEc2ModelIcmpTypeCode {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withType", "com/amazonaws/services/ec2/model/IcmpTypeCode", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelIcmpTypeCode{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCode(java.lang.Integer)
func (jbobject *ServicesEc2ModelIcmpTypeCode) SetCode(a int)  {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCode", javabind.Void, conv_a.Value().Cast("java/lang/Integer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.Integer getCode()
func (jbobject *ServicesEc2ModelIcmpTypeCode) GetCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCode", "java/lang/Integer")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoInteger()
	dst := new(int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.amazonaws.services.ec2.model.IcmpTypeCode withCode(java.lang.Integer)
func (jbobject *ServicesEc2ModelIcmpTypeCode) WithCode(a int) *ServicesEc2ModelIcmpTypeCode {
	conv_a := javabind.NewGoToJavaInteger()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCode", "com/amazonaws/services/ec2/model/IcmpTypeCode", conv_a.Value().Cast("java/lang/Integer"))
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
	unique_x := &ServicesEc2ModelIcmpTypeCode{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelIcmpTypeCode) ToString() string {
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
func (jbobject *ServicesEc2ModelIcmpTypeCode) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelIcmpTypeCode) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.IcmpTypeCode clone()
func (jbobject *ServicesEc2ModelIcmpTypeCode) Clone() *ServicesEc2ModelIcmpTypeCode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/IcmpTypeCode")
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
	unique_x := &ServicesEc2ModelIcmpTypeCode{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelIcmpTypeCode) Clone2() (*JavaLangObject, error) {
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


