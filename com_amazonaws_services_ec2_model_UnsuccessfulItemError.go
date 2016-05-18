package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelUnsuccessfulItemErrorInterface interface {
	JavaLangObjectInterface

	// public void setCode(java.lang.String)
	SetCode(a string) 

	// public java.lang.String getCode()
	GetCode() string

	// public com.amazonaws.services.ec2.model.UnsuccessfulItemError withCode(java.lang.String)
	WithCode(a string) *ServicesEc2ModelUnsuccessfulItemError

	// public void setMessage(java.lang.String)
	SetMessage(a string) 

	// public java.lang.String getMessage()
	GetMessage() string

	// public com.amazonaws.services.ec2.model.UnsuccessfulItemError withMessage(java.lang.String)
	WithMessage(a string) *ServicesEc2ModelUnsuccessfulItemError

	// public com.amazonaws.services.ec2.model.UnsuccessfulItemError clone()
	Clone() *ServicesEc2ModelUnsuccessfulItemError
}

type ServicesEc2ModelUnsuccessfulItemError struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.UnsuccessfulItemError()
func NewServicesEc2ModelUnsuccessfulItemError() (*ServicesEc2ModelUnsuccessfulItemError) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/UnsuccessfulItemError")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelUnsuccessfulItemError{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setCode(java.lang.String)
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) SetCode(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCode", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getCode()
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) GetCode() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCode", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UnsuccessfulItemError withCode(java.lang.String)
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) WithCode(a string) *ServicesEc2ModelUnsuccessfulItemError {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCode", "com/amazonaws/services/ec2/model/UnsuccessfulItemError", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUnsuccessfulItemError{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMessage(java.lang.String)
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) SetMessage(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMessage", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getMessage()
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) GetMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMessage", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UnsuccessfulItemError withMessage(java.lang.String)
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) WithMessage(a string) *ServicesEc2ModelUnsuccessfulItemError {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMessage", "com/amazonaws/services/ec2/model/UnsuccessfulItemError", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUnsuccessfulItemError{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) ToString() string {
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
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.UnsuccessfulItemError clone()
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) Clone() *ServicesEc2ModelUnsuccessfulItemError {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/UnsuccessfulItemError")
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
	unique_x := &ServicesEc2ModelUnsuccessfulItemError{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelUnsuccessfulItemError) Clone2() (*JavaLangObject, error) {
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


