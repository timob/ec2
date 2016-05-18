package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelBundleTaskErrorInterface interface {
	JavaLangObjectInterface

	// public void setCode(java.lang.String)
	SetCode(a string) 

	// public java.lang.String getCode()
	GetCode() string

	// public com.amazonaws.services.ec2.model.BundleTaskError withCode(java.lang.String)
	WithCode(a string) *ServicesEc2ModelBundleTaskError

	// public void setMessage(java.lang.String)
	SetMessage(a string) 

	// public java.lang.String getMessage()
	GetMessage() string

	// public com.amazonaws.services.ec2.model.BundleTaskError withMessage(java.lang.String)
	WithMessage(a string) *ServicesEc2ModelBundleTaskError

	// public com.amazonaws.services.ec2.model.BundleTaskError clone()
	Clone() *ServicesEc2ModelBundleTaskError
}

type ServicesEc2ModelBundleTaskError struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.BundleTaskError()
func NewServicesEc2ModelBundleTaskError() (*ServicesEc2ModelBundleTaskError) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/BundleTaskError")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelBundleTaskError{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setCode(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTaskError) SetCode(a string)  {
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
func (jbobject *ServicesEc2ModelBundleTaskError) GetCode() string {
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

// public com.amazonaws.services.ec2.model.BundleTaskError withCode(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTaskError) WithCode(a string) *ServicesEc2ModelBundleTaskError {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withCode", "com/amazonaws/services/ec2/model/BundleTaskError", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBundleTaskError{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMessage(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTaskError) SetMessage(a string)  {
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
func (jbobject *ServicesEc2ModelBundleTaskError) GetMessage() string {
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

// public com.amazonaws.services.ec2.model.BundleTaskError withMessage(java.lang.String)
func (jbobject *ServicesEc2ModelBundleTaskError) WithMessage(a string) *ServicesEc2ModelBundleTaskError {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withMessage", "com/amazonaws/services/ec2/model/BundleTaskError", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelBundleTaskError{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelBundleTaskError) ToString() string {
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
func (jbobject *ServicesEc2ModelBundleTaskError) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelBundleTaskError) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.BundleTaskError clone()
func (jbobject *ServicesEc2ModelBundleTaskError) Clone() *ServicesEc2ModelBundleTaskError {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/BundleTaskError")
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
	unique_x := &ServicesEc2ModelBundleTaskError{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelBundleTaskError) Clone2() (*JavaLangObject, error) {
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


