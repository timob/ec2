package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelUnsuccessfulItemInterface interface {
	JavaLangObjectInterface

	// public void setResourceId(java.lang.String)
	SetResourceId(a string) 

	// public java.lang.String getResourceId()
	GetResourceId() string

	// public com.amazonaws.services.ec2.model.UnsuccessfulItem withResourceId(java.lang.String)
	WithResourceId(a string) *ServicesEc2ModelUnsuccessfulItem

	// public void setError(com.amazonaws.services.ec2.model.UnsuccessfulItemError)
	SetError(a ServicesEc2ModelUnsuccessfulItemErrorInterface) 

	// public com.amazonaws.services.ec2.model.UnsuccessfulItemError getError()
	GetError() *ServicesEc2ModelUnsuccessfulItemError

	// public com.amazonaws.services.ec2.model.UnsuccessfulItem withError(com.amazonaws.services.ec2.model.UnsuccessfulItemError)
	WithError(a ServicesEc2ModelUnsuccessfulItemErrorInterface) *ServicesEc2ModelUnsuccessfulItem

	// public com.amazonaws.services.ec2.model.UnsuccessfulItem clone()
	Clone() *ServicesEc2ModelUnsuccessfulItem
}

type ServicesEc2ModelUnsuccessfulItem struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.UnsuccessfulItem()
func NewServicesEc2ModelUnsuccessfulItem() (*ServicesEc2ModelUnsuccessfulItem) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/UnsuccessfulItem")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelUnsuccessfulItem{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setResourceId(java.lang.String)
func (jbobject *ServicesEc2ModelUnsuccessfulItem) SetResourceId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResourceId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getResourceId()
func (jbobject *ServicesEc2ModelUnsuccessfulItem) GetResourceId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResourceId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.UnsuccessfulItem withResourceId(java.lang.String)
func (jbobject *ServicesEc2ModelUnsuccessfulItem) WithResourceId(a string) *ServicesEc2ModelUnsuccessfulItem {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withResourceId", "com/amazonaws/services/ec2/model/UnsuccessfulItem", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelUnsuccessfulItem{}
	unique_x.Callable = dst
	return unique_x
}

// public void setError(com.amazonaws.services.ec2.model.UnsuccessfulItemError)
func (jbobject *ServicesEc2ModelUnsuccessfulItem) SetError(a ServicesEc2ModelUnsuccessfulItemErrorInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setError", javabind.Void, conv_a.Value().Cast("com/amazonaws/services/ec2/model/UnsuccessfulItemError"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.amazonaws.services.ec2.model.UnsuccessfulItemError getError()
func (jbobject *ServicesEc2ModelUnsuccessfulItem) GetError() *ServicesEc2ModelUnsuccessfulItemError {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getError", "com/amazonaws/services/ec2/model/UnsuccessfulItemError")
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

// public com.amazonaws.services.ec2.model.UnsuccessfulItem withError(com.amazonaws.services.ec2.model.UnsuccessfulItemError)
func (jbobject *ServicesEc2ModelUnsuccessfulItem) WithError(a ServicesEc2ModelUnsuccessfulItemErrorInterface) *ServicesEc2ModelUnsuccessfulItem {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withError", "com/amazonaws/services/ec2/model/UnsuccessfulItem", conv_a.Value().Cast("com/amazonaws/services/ec2/model/UnsuccessfulItemError"))
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
	unique_x := &ServicesEc2ModelUnsuccessfulItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelUnsuccessfulItem) ToString() string {
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
func (jbobject *ServicesEc2ModelUnsuccessfulItem) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelUnsuccessfulItem) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.UnsuccessfulItem clone()
func (jbobject *ServicesEc2ModelUnsuccessfulItem) Clone() *ServicesEc2ModelUnsuccessfulItem {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/UnsuccessfulItem")
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
	unique_x := &ServicesEc2ModelUnsuccessfulItem{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelUnsuccessfulItem) Clone2() (*JavaLangObject, error) {
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


