package ec2

import "github.com/timob/javabind"

type ServicesEc2ModelPrefixListIdInterface interface {
	JavaLangObjectInterface

	// public void setPrefixListId(java.lang.String)
	SetPrefixListId(a string) 

	// public java.lang.String getPrefixListId()
	GetPrefixListId() string

	// public com.amazonaws.services.ec2.model.PrefixListId withPrefixListId(java.lang.String)
	WithPrefixListId(a string) *ServicesEc2ModelPrefixListId

	// public com.amazonaws.services.ec2.model.PrefixListId clone()
	Clone() *ServicesEc2ModelPrefixListId
}

type ServicesEc2ModelPrefixListId struct {
	JavaLangObject
}

// public com.amazonaws.services.ec2.model.PrefixListId()
func NewServicesEc2ModelPrefixListId() (*ServicesEc2ModelPrefixListId) {

	obj, err := javabind.GetEnv().NewObject("com/amazonaws/services/ec2/model/PrefixListId")
	if err != nil {
		panic(err)
	}
	x := &ServicesEc2ModelPrefixListId{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setPrefixListId(java.lang.String)
func (jbobject *ServicesEc2ModelPrefixListId) SetPrefixListId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPrefixListId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getPrefixListId()
func (jbobject *ServicesEc2ModelPrefixListId) GetPrefixListId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPrefixListId", "java/lang/String")
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

// public com.amazonaws.services.ec2.model.PrefixListId withPrefixListId(java.lang.String)
func (jbobject *ServicesEc2ModelPrefixListId) WithPrefixListId(a string) *ServicesEc2ModelPrefixListId {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "withPrefixListId", "com/amazonaws/services/ec2/model/PrefixListId", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &ServicesEc2ModelPrefixListId{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *ServicesEc2ModelPrefixListId) ToString() string {
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
func (jbobject *ServicesEc2ModelPrefixListId) Equals(a interface{}) bool {
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
func (jbobject *ServicesEc2ModelPrefixListId) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.amazonaws.services.ec2.model.PrefixListId clone()
func (jbobject *ServicesEc2ModelPrefixListId) Clone() *ServicesEc2ModelPrefixListId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/amazonaws/services/ec2/model/PrefixListId")
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
	unique_x := &ServicesEc2ModelPrefixListId{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *ServicesEc2ModelPrefixListId) Clone2() (*JavaLangObject, error) {
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


