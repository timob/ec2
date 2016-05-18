package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelIamInstanceProfileSpecificationInterface interface {
	JavaLangObjectInterface

	// public void setArn(java.lang.String)
	SetArn(a string) 

	// public java.lang.String getArn()
	GetArn() string

	// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification withArn(java.lang.String)
	WithArn(a string) *ServicesEc2ModelIamInstanceProfileSpecification

	// public void setName(java.lang.String)
	SetName(a string) 

	// public java.lang.String getName()
	GetName() string

	// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification withName(java.lang.String)
	WithName(a string) *ServicesEc2ModelIamInstanceProfileSpecification

	// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification clone()
	Clone() *ServicesEc2ModelIamInstanceProfileSpecification
}

type ServicesEc2ModelIamInstanceProfileSpecification struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification()
func NewServicesEc2ModelIamInstanceProfileSpecification() (*ServicesEc2ModelIamInstanceProfileSpecification) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/IamInstanceProfileSpecification")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelIamInstanceProfileSpecification{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setArn(java.lang.String)
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) SetArn(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArn", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getArn()
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) GetArn() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArn", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification withArn(java.lang.String)
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) WithArn(a string) *ServicesEc2ModelIamInstanceProfileSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withArn", "com/amazonaws/services/ec2/model/IamInstanceProfileSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelIamInstanceProfileSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(java.lang.String)
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) SetName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getName()
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification withName(java.lang.String)
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) WithName(a string) *ServicesEc2ModelIamInstanceProfileSpecification {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withName", "com/amazonaws/services/ec2/model/IamInstanceProfileSpecification", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelIamInstanceProfileSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) ToString() string {
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
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.IamInstanceProfileSpecification clone()
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) Clone() *ServicesEc2ModelIamInstanceProfileSpecification {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/IamInstanceProfileSpecification")
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
	unique_x := &ServicesEc2ModelIamInstanceProfileSpecification{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelIamInstanceProfileSpecification) Clone2() (*JavaLangObject, error) {
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


